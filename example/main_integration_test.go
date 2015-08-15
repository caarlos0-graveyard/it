package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ContaAzul/ab/datastores"
	"github.com/caarlos0/it"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

var testServer *httptest.Server

func TestMain(m *testing.M) {
	serverUp := func(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
		return server(db).ServeHTTP
	}
	it := it.New()
	handler := it.Init(serverUp, datastores.NewDBConnectionPool)
	defer it.Shutdown()
	testServer = httptest.NewServer(http.HandlerFunc(handler))
	defer testServer.Close()
	m.Run()
}

func TestCreatesAndListsBook(t *testing.T) {
	res, err := http.Post(testServer.URL+"/books/Integration Test book", "", nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	res, err = http.Get(testServer.URL + "/books")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var books []Book
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&books)
	assert.Len(t, books, 1)
}