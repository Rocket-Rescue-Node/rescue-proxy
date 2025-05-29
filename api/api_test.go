package api

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/Rocket-Rescue-Node/rescue-proxy/metrics"
	"github.com/Rocket-Rescue-Node/rescue-proxy/pb"
	"github.com/Rocket-Rescue-Node/rescue-proxy/test"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type apiTest struct {
	listener *bufconn.Listener
	logger   *zap.Logger
	client   pb.ApiClient
	ctx      context.Context
}

func setup(t *testing.T) apiTest {
	_, err := metrics.Init("api_test_" + t.Name())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(metrics.Deinit)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	t.Cleanup(cancel)

	out := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		_ = out.Close()
	})

	conn, err := grpc.NewClient(
		"passthrough:bufconn",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return out.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatal("couldn't dial mock api listener", err)
	}

	client := pb.NewApiClient(conn)

	return apiTest{
		listener: out,
		logger:   zaptest.NewLogger(t),
		client:   client,
		ctx:      ctx,
	}
}

func TestApiStartStop(t *testing.T) {

	at := setup(t)
	el := test.NewMockExecutionLayer(50, 5, 200, 0, t.Name())
	cl := test.NewMockConsensusLayer(400, t.Name())
	a := API{
		EL:     el,
		CL:     cl,
		Logger: at.logger,
	}
	err := a.Init(at.listener)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(a.Deinit)

	// Force a ticker tick
	a.ticker.t.Reset(100 * time.Millisecond)
	time.Sleep(110 * time.Millisecond)
	a.ticker.t.Reset(1 * time.Second)

	resp, err := at.client.GetOdaoNodes(at.ctx, &pb.OdaoNodesRequest{})
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.GetNodeIds()) <= 0 {
		t.Fatal("api didn't return any odao nodes")
	}

	resp2, err := at.client.GetRocketPoolNodes(at.ctx, &pb.RocketPoolNodesRequest{})
	if err != nil {
		t.Fatal(err)
	}

	if len(resp2.GetNodeIds()) <= 0 {
		t.Fatal("api didn't return any odao nodes")
	}

	resp3, err := at.client.GetSoloValidators(at.ctx, &pb.SoloValidatorsRequest{})
	if err != nil {
		t.Fatal(err)
	}

	if len(resp3.GetWithdrawalAddresses()) <= 0 {
		t.Fatal("api didn't return any odao nodes")
	}
}
