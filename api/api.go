package api

import (
	"context"
	"net"

	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/executionlayer"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/pb"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type API struct {
	pb.UnimplementedApiServer
	EL         *executionlayer.ExecutionLayer
	Logger     *zap.Logger
	ListenAddr string
	listener   net.Listener
	server     *grpc.Server
}

func NewAPI(listenAddr string, el *executionlayer.ExecutionLayer, logger *zap.Logger) *API {
	out := &API{
		EL:         el,
		Logger:     logger,
		ListenAddr: listenAddr,
	}

	return out
}

func (a *API) GetRocketPoolNodes(ctx context.Context, request *pb.RocketPoolNodesRequest) (*pb.RocketPoolNodes, error) {
	out := &pb.RocketPoolNodes{}
	out.NodeIds = make([][]byte, 0, 1024)

	err := a.EL.ForEachNode(func(addr common.Address) bool {
		out.NodeIds = append(out.NodeIds, addr.Bytes())
		return true
	})

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (a *API) Init() error {
	var err error

	a.listener, err = net.Listen("tcp", a.ListenAddr)
	if err != nil {
		return err
	}

	a.server = grpc.NewServer()

	pb.RegisterApiServer(a.server, a)

	a.Logger.Info("Starting grpc server", zap.String("url", a.ListenAddr))
	go func() {
		if err := a.server.Serve(a.listener); err != nil {
			a.Logger.Panic("gRPC server stopped", zap.Error(err))
		}
	}()

	return nil
}

func (a *API) Deinit() {
	a.server.GracefulStop()
	a.listener.Close()
}
