package swarm

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

func setup() (*Client, *http.ServeMux, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client := NewClient("TESTCUSTOMER", "TESTAPITOKEN")
	url, _ := url.Parse(server.URL)

	client.BaseURL = url
	client.CustomerURL = url

	return client, mux, server.Close
}
