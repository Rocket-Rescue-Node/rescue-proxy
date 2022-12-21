package admin

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

type AdminApi struct {
	http.Server
}

func (a *AdminApi) Init(listenAddr string) {

	a.Addr = listenAddr
	a.Handler = mux.NewRouter()
}

func (a *AdminApi) Handle(path string, handler http.Handler) {
	a.Handler.(*mux.Router).Path(path).Handler(handler)
}

func (a *AdminApi) Start() error {
	listener, err := net.Listen("tcp", a.Addr)
	if err != nil {
		return err
	}

	go func() {
		_ = a.Serve(listener)
	}()
	return nil
}

func (a *AdminApi) Stop() {
	a.Close()
}
