// Package proxy forwards OpenAI-compatible requests to a backend.
package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func New(backendURL string) (http.Handler, error) {
	u, err := url.Parse(backendURL)
	if err != nil {
		return nil, err
	}
	return httputil.NewSingleHostReverseProxy(u), nil
}
