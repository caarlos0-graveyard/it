package main_test

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caarlos0/it"
	"github.com/caarlos0/it/example"
	"github.com/stretchr/testify/assert"
)

var testServer *httptest.Server

func TestMain(m *testing.M) {
	serverFn := func(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
		return main.Server(db).ServeHTTP
	}
	it := it.New()
	handler := it.Init(serverFn, main.NewConnectionPool)
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

	var books []main.Book
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&books)
	assert.Len(t, books, 1)
}
