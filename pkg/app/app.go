package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pumphouse-p/peek-go/pkg/env"
	"github.com/pumphouse-p/peek-go/pkg/net"
	"github.com/pumphouse-p/peek-go/pkg/runtime"
)

type App struct {
	config  Config
	env     *env.Env
	net     *net.Net
	runtime *runtime.Runtime
	router  *mux.Router
}

func (p *App) Run() {
	log.Printf("Serving over HTTP on %v", p.config.ListenOn)
	log.Fatal(http.ListenAndServe(p.config.ListenOn, p.router))
}

func NewApp() *App {
	app := &App{}
	app.router = mux.NewRouter()

	app.env = env.New()

	app.router.HandleFunc("/api/env", app.env.APIGet)
	app.router.HandleFunc("/api/net", app.net.APIGet)
	app.router.HandleFunc("/api/cpu", app.runtime.APIGetCPU)
	app.router.HandleFunc("/api/mem", app.runtime.APIGetMem)
	app.router.HandleFunc("/api/storage", app.runtime.APIGetStorage)
	app.router.HandleFunc("/api/runtime", app.runtime.APIGetRuntime)

	return app
}
