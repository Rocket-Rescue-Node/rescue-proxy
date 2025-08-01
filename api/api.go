package api

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/Rocket-Rescue-Node/rescue-proxy/consensuslayer"
	"github.com/Rocket-Rescue-Node/rescue-proxy/executionlayer"
	"github.com/Rocket-Rescue-Node/rescue-proxy/metrics"
	"github.com/Rocket-Rescue-Node/rescue-proxy/pb"
	apiv1 "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type API struct {
	pb.UnimplementedApiServer
	EL     executionlayer.ExecutionLayer
	CL     consensuslayer.ConsensusLayer
	Logger *zap.Logger
	server *grpc.Server
	m      *metrics.MetricsRegistry

	soloValidatorCache     map[common.Address]interface{}
	soloValidatorCacheLock sync.RWMutex
	ticker                 struct {
		t    *time.Ticker
		done chan bool
	}
	ctx context.Context
}

func (a *API) GetRocketPoolNodes(ctx context.Context, request *pb.RocketPoolNodesRequest) (*pb.RocketPoolNodes, error) {
	out := &pb.RocketPoolNodes{}
	out.NodeIds = make([][]byte, 0, 1024)

	err := a.EL.ForEachNode(func(addr common.Address) bool {
		out.NodeIds = append(out.NodeIds, addr.Bytes())
		return true
	})

	if err != nil {
		a.m.Counter("get_rocket_pool_nodes_error").Inc()
		return nil, err
	}

	a.m.Counter("get_rocket_pool_nodes_ok").Inc()
	return out, nil
}

func (a *API) GetOdaoNodes(ctx context.Context, request *pb.OdaoNodesRequest) (*pb.OdaoNodes, error) {
	out := &pb.OdaoNodes{}
	out.NodeIds = make([][]byte, 0, 8)

	err := a.EL.ForEachOdaoNode(func(addr common.Address) bool {
		out.NodeIds = append(out.NodeIds, addr.Bytes())
		return true
	})

	if err != nil {
		a.m.Counter("get_odao_nodes_error").Inc()
		return nil, err
	}

	a.m.Counter("get_odao_nodes_ok").Inc()
	return out, nil
}

func (a *API) GetSoloValidators(ctx context.Context, request *pb.SoloValidatorsRequest) (*pb.SoloValidators, error) {
	a.soloValidatorCacheLock.RLock()
	defer a.soloValidatorCacheLock.RUnlock()

	out := &pb.SoloValidators{}
	out.WithdrawalAddresses = make([][]byte, 0, len(a.soloValidatorCache))

	for addr := range a.soloValidatorCache {
		out.WithdrawalAddresses = append(out.WithdrawalAddresses, addr.Bytes())
	}

	a.m.Counter("get_solo_validators_ok").Inc()
	return out, nil
}

func (a *API) ValidateEIP1271(ctx context.Context, request *pb.ValidateEIP1271Request) (*pb.ValidateEIP1271Response, error) {
	if len(request.DataHash) != 32 {
		return &pb.ValidateEIP1271Response{Error: fmt.Sprintf("invalid DataHash length: expected 32 bytes, got %d", len(request.DataHash))}, nil
	}
	if len(request.Address) != 20 {
		return &pb.ValidateEIP1271Response{Error: fmt.Sprintf("invalid Address length: expected 20 bytes, got %d", len(request.Address))}, nil
	}
	dataHash := common.BytesToHash(request.DataHash)
	address := common.BytesToAddress(request.Address)

	valid, err := a.EL.ValidateEIP1271(ctx, dataHash, request.Signature, address)
	if err != nil {
		a.m.Counter("validate_eip1271_error").Inc()
		return &pb.ValidateEIP1271Response{Error: err.Error()}, nil
	}

	a.m.Counter("validate_eip1271_ok").Inc()
	return &pb.ValidateEIP1271Response{Valid: valid}, nil
}

func (a *API) updateCache() error {
	a.soloValidatorCacheLock.Lock()
	defer a.soloValidatorCacheLock.Unlock()

	validators, err := a.CL.GetValidators()
	if err != nil {
		a.m.Counter("solo_validator_cache_update_failed").Inc()
		return err
	}

	newMap := make(map[common.Address]interface{})
	for _, validator := range validators {
		if validator.Status == apiv1.ValidatorStateUnknown ||
			validator.Status == apiv1.ValidatorStateExitedUnslashed ||
			validator.Status == apiv1.ValidatorStateExitedSlashed ||
			validator.Status == apiv1.ValidatorStateWithdrawalPossible ||
			validator.Status == apiv1.ValidatorStateWithdrawalDone {
			continue
		}

		creds := validator.Validator.WithdrawalCredentials

		// Check that the creds are an EL address
		credsAreValid := bytes.HasPrefix(creds, []byte{0x01}) || bytes.HasPrefix(creds, []byte{0x02})
		if !credsAreValid {
			continue
		}

		// The address is the last 20 bytes of the credential.
		address := common.BytesToAddress(creds[len(creds)-common.AddressLength:])

		// Add to the new map
		newMap[address] = struct{}{}
	}

	// swap out the cache
	a.m.Gauge("solo_validator_cache_update_size").Set(float64(len(newMap)))
	a.m.Counter("solo_validator_cache_update_success").Inc()
	a.soloValidatorCache = newMap

	return nil
}

func (a *API) Init(listener net.Listener) error {

	a.m = metrics.NewMetricsRegistry("api")

	a.server = grpc.NewServer()

	pb.RegisterApiServer(a.server, a)

	a.Logger.Info("Starting grpc server", zap.String("url", listener.Addr().String()))
	go func() {
		if err := a.server.Serve(listener); err != nil {
			a.Logger.Panic("gRPC server stopped", zap.Error(err))
		}
	}()

	a.Logger.Info("Seeding the solo validator cache")
	err := a.updateCache()
	if err != nil {
		return fmt.Errorf("unable to seed the solo validator cache: %w", err)
	}

	// Updates to the validator set are taxing on the bn to query, so do it every 16 epochs (about once every hour and a half)
	a.ticker.t = time.NewTicker(time.Second * 32 * 12 * 16)
	a.ticker.done = make(chan bool)
	a.Logger.Info("Starting solo validator background worker")

	ctx, cancel := context.WithCancel(context.Background())
	a.ctx = ctx
	go func() {
		for {
			select {
			case <-a.ticker.t.C:
				a.Logger.Debug("Updating solo validator worker")
				err := a.updateCache()
				if err != nil {
					a.Logger.Warn("Error updating solo validator cache", zap.Error(err))
				}
			case <-a.ticker.done:
				a.Logger.Info("Stopped background solo validator worker")
				cancel()
				return
			}
		}
	}()

	return nil
}

func (a *API) Deinit() {
	a.server.Stop()
	a.ticker.t.Stop()
	close(a.ticker.done)
	<-a.ctx.Done()
}
