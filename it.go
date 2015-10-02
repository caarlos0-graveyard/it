package it

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/caarlos0/it/base"
	"github.com/caarlos0/it/db"
)

// ServerHandlerFn blah
type ServerHandlerFn func(*sql.DB) func(http.ResponseWriter, *http.Request)

// IT is the main structure of the integration testing framework.
type IT struct {
	db *db.DB
}

// New creates a new IT instance
func New() IT {
	return IT{}
}

// Init the IT framework: loads config, creates the database and startup the
// server.
func (i *IT) Init(server ServerHandlerFn, connectToDatabase db.PoolFn) func(http.ResponseWriter, *http.Request) {
	var cfg base.Config
	env.Parse(&cfg)
	i.db = db.New(connectToDatabase, &cfg)
	return server(i.db.Init())
}

// Shutdown the IT framework.
func (i *IT) Shutdown() {
	i.db.Shutdown()
	log.Println("Shutdown IT...")
}
