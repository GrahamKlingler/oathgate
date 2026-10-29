// Package proxy forwards OpenAI-compatible requests to a backend.
package proxy

import (
	// "errors"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// New returns a handler that forwards requests unchanged to backendURL.
// Ollama already serves POST /v1/chat/completions and GET /v1/models, so a
// plain passthrough covers both endpoints this week.
//
// TODO:
//  1. url.Parse(backendURL); return the error if it fails.
//  2. return httputil.NewSingleHostReverseProxy(u), nil
//     (stdlib; it copies method, path, headers, and body, and streams the response back).
//  3. Later phases (auth, limiter, cache) wrap this handler instead of replacing it.
func New(backendURL string) (http.Handler, error) {
	url, err := url.Parse(backendURL)
	if err != nil {
		return nil, err
	}
	return httputil.NewSingleHostReverseProxy(url), nil
}
