package it

import (
	"log"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/jmoiron/sqlx"
)

// ServerHandlerFn blah
type ServerHandlerFn func(*sqlx.DB) func(http.ResponseWriter, *http.Request)

// Framework blah
type Framework struct {
	db *DB
}

// NewFramework blah
func NewFramework() Framework {
	return Framework{}
}

// Init blah
func (f *Framework) Init(server ServerHandlerFn, connectToDatabase DBPoolFn) func(http.ResponseWriter, *http.Request) {
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
