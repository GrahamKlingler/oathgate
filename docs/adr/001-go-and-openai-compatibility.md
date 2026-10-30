# ADR-001: Go, and OpenAI-compatible API

- Status: Accepted
- Date: 2026-10-29

## Context

Applications that use LLMs call a provider directly, so each one handles its
own auth, rate limiting, caching, and backend choice. Oathgate is a gateway that
sits between those applications and LLM backends (Ollama locally, other
providers later) and centralizes those concerns.

Constraints:

- **Low added latency.** The gateway is on every request's path; its overhead
  should be negligible next to model inference time.
- **Many concurrent, long-lived connections.** Completions take seconds, and
  streaming responses hold connections open for their whole duration.
- **Simple deployment.** It ships as a container to Docker Compose and Kubernetes (Helm).
- **No client changes.** Existing applications should adopt the gateway without
  rewriting their LLM calls.

## Decision

1. **Implement the gateway in Go.**
2. **Expose an OpenAI-compatible HTTP API** (`/v1/chat/completions`,
   `/v1/models`), so the gateway is a drop-in `base_url` for OpenAI SDKs.

## Alternatives considered

### Language

- **Rust:** faster and no garbage collector, but slower to build and iterate
  on. The gateway is I/O-bound, so Rust's raw speed buys little.
- **Python (FastAPI/asyncio):** best LLM ecosystem, but higher per-request
  overhead, the GIL limits CPU-bound work, and deployment needs a runtime and
  dependencies rather than one binary.
- **Node.js:** good async I/O, but CPU work on the request path (e.g. token
  counting) blocks the single event loop.

Go fits the constraints: `net/http` and `httputil.ReverseProxy` cover proxying
in the standard library, goroutines make many concurrent connections cheap,
it compiles to a single static binary for small container images, and it
includes a race detector, which CI uses.

### API

- **Custom API:** full design freedom, but every client needs a new SDK or
  integration code, which breaks the "no client changes" constraint.
- **Provider-native APIs** (e.g. Anthropic, Gemini): ties clients to one
  provider's schema and defeats the gateway's ability to swap backends.

The OpenAI API is the de facto standard: official and third-party SDKs exist
in most languages, and many backends (Ollama, vLLM, LM Studio, and others)
already serve it, so compatible backends need no translation at all.

## Consequences

**Easier**

- Any OpenAI SDK works by changing only `base_url`, which is the Phase 1 acceptance test.
- Compatible backends such as Ollama can be passed through without any request translation.
- One static binary keeps the container image and Helm chart simple.

**Harder**

- The gateway follows OpenAI's schema as it changes (for example, the newer
  Responses API), and we must decide which endpoints to support.
- Backends that don't speak the OpenAI format need an adapter (planned for
  `internal/backend`), and provider-specific features that don't fit the schema
  are lost or need extensions.
- If the Python LLM ecosystem is needed (e.g. tokenizers), we must find a Go
  equivalent or call out to a separate service.
