package it

import (
	"log"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/caarlos0/it/base"
	"github.com/caarlos0/it/db"
	"github.com/jmoiron/sqlx"
)

// ServerHandlerFn blah
type ServerHandlerFn func(*sqlx.DB) func(http.ResponseWriter, *http.Request)

// IT blah
type IT struct {
	db *db.DB
}

// New blah
func New() IT {
	return IT{}
}

// Init blah
func (i *IT) Init(server ServerHandlerFn, connectToDatabase db.PoolFn) func(http.ResponseWriter, *http.Request) {
	var cfg base.Config
	env.Parse(&cfg)
	i.db = db.New(connectToDatabase, &cfg)
	return server(i.db.Init())
}

// Shutdown blah
func (i *IT) Shutdown() {
	i.db.Shutdown()
	log.Println("Shutdown IT IT")
}
