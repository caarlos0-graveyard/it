package client_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/caarlos0/it/client"
	"github.com/stretchr/testify/assert"
)

var server *httptest.Server

func TestMain(m *testing.M) {
	server = httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "PUT" {
					w.WriteHeader(http.StatusMethodNotAllowed)
				}
				w.Write([]byte("ok"))
			},
		),
	)
	defer server.Close()
	m.Run()
}

func TestPut(t *testing.T) {
	resp, err := client.Put(server.URL, "", nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPutForm(t *testing.T) {
	resp, err := client.PutForm(server.URL, url.Values{})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestFailsWhenNotPut(t *testing.T) {
	resp, err := http.Post(server.URL, "", nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)
}
