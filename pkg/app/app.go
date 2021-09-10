package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pumphouse-p/peek-go/pkg/env"
)

type App struct {
	config Config
	env    *env.Env
	router *mux.Router
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

	return app
}
