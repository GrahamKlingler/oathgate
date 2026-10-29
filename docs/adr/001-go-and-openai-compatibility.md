# ADR-001: Go, and OpenAI-compatible API

- Status: Proposed
- Date: 2026-10-29

## Context

<!-- What problem is the gateway solving? Who calls it, what does it sit in front of?
     What constraints matter: latency overhead, concurrency, deployment, team skills? -->
The gateway is solving an LLM usage. It sits in front of an LLM provider. 

## Decision

<!-- 1. Implement the gateway in Go.
     2. Expose an OpenAI-compatible HTTP API (/v1/chat/completions, /v1/models). -->
Gateway implemented in Go.

## Alternatives considered

<!-- Language: e.g. Python, Rust, Node. One or two lines each on why not.
     API: e.g. a custom API, or provider-native APIs. Why not? -->
Rust is faster, but it is slower to build.

## Consequences

<!-- What gets easier: existing SDKs work unchanged, ...
     What gets harder / what you're giving up: tied to OpenAI's schema evolving, ... -->

