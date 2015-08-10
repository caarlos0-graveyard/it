package client

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

func PutForm(url string, data url.Values) (resp *http.Response, err error) {
	return Put(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

func Put(url string, bodyType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", bodyType)
	client := http.Client{}
	return client.Do(req)
}
