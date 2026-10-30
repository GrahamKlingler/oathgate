# Oathgate

An OpenAI-compatible LLM gateway written in Go. Clients use any OpenAI SDK and
point `base_url` at the gateway; the gateway forwards requests to a backend
(Ollama locally).

## Quickstart

Requires Go 1.25+ and Docker.

```bash
make up                                     # Redis, Postgres, Ollama
docker compose -f deploy/compose/docker-compose.yml exec ollama ollama pull llama3.2:1b
make run                                    # gateway on :8080
```

Check it:

```bash
curl -i localhost:8080/readyz               # 200 once Ollama is reachable
pip install openai && python scripts/phase1-1.py
```

## Endpoints

| Method | Path                   | Description                                  |
|--------|------------------------|----------------------------------------------|
| POST   | `/v1/chat/completions` | Passed through to the backend (non-streaming) |
| GET    | `/v1/models`           | Passed through to the backend                 |
| GET    | `/healthz`             | Liveness: 200 while the process is up         |
| GET    | `/readyz`              | Readiness: 200 only if the backend responds   |

## Configuration

YAML file, default `deploy/compose/gateway.yaml` (override with `-config <path>`).
Environment variables take precedence over the file.

| YAML key      | Env var               | Example                  |
|---------------|-----------------------|--------------------------|
| `listen_addr` | `GATEWAY_LISTEN_ADDR` | `:8080`                  |
| `backend_url` | `GATEWAY_BACKEND_URL` | `http://localhost:11434` |

`backend_url` is required; the gateway refuses to start without it.

## Development

```bash
make build   # binaries in bin/
make test    # unit tests with the race detector
make lint    # golangci-lint
make down    # stop the local stack
```

CI runs lint and `go test -race ./...` on every push.

## Layout

```
cmd/gateway/      gateway main package
cmd/sim/          backend simulator
cmd/gwctl/        admin CLI
internal/         auth, limiter, cache, router, backend, proxy, config, telemetry
deploy/helm/      Helm chart
deploy/compose/   local stack
bench/            load generators, workloads, results, plotting scripts
docs/adr/         architecture decision records
scripts/          smoke tests
```
