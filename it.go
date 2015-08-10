package it

import (
	"log"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/jmoiron/sqlx"
)

// HTTPHandler blah
type HTTPHandler func(w http.ResponseWriter, r *http.Request)

// ServerHandlerFn blah
type ServerHandlerFn func(db *sqlx.DB) HTTPHandler

// Framework blah
type Framework struct {
	db *DB
}

// Init blah
func (f *Framework) Init(server ServerHandlerFn, connectToDatabase DBPoolFn) HTTPHandler {
	var cfg Config
	env.Parse(&cfg)
	f.db = &DB{
		cfg:     cfg,
		connect: connectToDatabase,
	}
	return server(f.db.Init())
}

// Shutdown blah
func (f *Framework) Shutdown() {
	f.db.Shutdown()
	log.Println("Shutdown IT Framework")
}
