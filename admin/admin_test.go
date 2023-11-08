package admin

import (
	"context"
	"errors"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/metrics"
)

func setup(t *testing.T) context.Context {
	t.Cleanup(metrics.Deinit)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	t.Cleanup(cancel)

	return ctx
}

func TestAdminStartStop(t *testing.T) {

	ctx := setup(t)
	a := AdminApi{}
	err := a.Init("admin_test")
	if err != nil {
		t.Fatal(err)
	}

	errs := make(chan error)
	// Omit a port and the library will pick one for us
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		err := a.Serve(listener)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errs <- err
		}
		close(errs)
	}()

	// Hit the metrics handler to make sure it's replying
	resp, err := http.Get("http://" + listener.Addr().String() + "/metrics")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatal("Non-200 status code received", resp.StatusCode)
	}

	err = a.Shutdown(ctx)
	if err != nil {
		t.Fatal(err)
	}

	err = <-errs
	if err != nil {
		t.Fatal(err)
	}
}
