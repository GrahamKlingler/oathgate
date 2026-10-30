// Command gateway is the OpenAI-compatible LLM gateway.
package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/grahamklingler/oathgate/internal/config"
	"github.com/grahamklingler/oathgate/internal/proxy"
)

func main() {
	configPath := flag.String("config", "deploy/compose/gateway.yaml", "path to YAML config")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatal(err)
	}
	backend, err := proxy.New(cfg.BackendURL)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("POST /v1/chat/completions", backend)
	mux.Handle("GET /v1/models", backend)

	// Liveness: the process is up. Always 200, never checks dependencies.
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Readiness: 200 only if the backend answers GET /v1/models with 200.
	mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{
			Timeout: 2 * time.Second,
		}
		resp, err := client.Get(cfg.BackendURL + "/v1/models")
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		_ = resp.Body.Close() // nothing useful to do if close fails; never crash on a probe
		if resp.StatusCode != http.StatusOK {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)

	})

	log.Printf("gateway listening on %s, backend %s", cfg.ListenAddr, cfg.BackendURL)
	log.Fatal(http.ListenAndServe(cfg.ListenAddr, mux))
}
