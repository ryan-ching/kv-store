# Go KV Store — 4 Week Learning Roadmap

## Goal

Build an in-memory key-value store with a REST API. By the end, you should be able to reason about Go's concurrency model well enough to start MIT 6.824. The project is the vehicle — the learning is the point.

## Rules of engagement

- Vibecode the boilerplate. Router wiring, JSON marshaling, flag parsing — fine to generate.
- Write the concurrency logic yourself. Every goroutine, channel, mutex, and select statement must come from your hands. If you can't explain why it's there, it doesn't go in.
- Keep a `questions.md` next to the code. Every time something feels magical or surprising, write it down. Answer them before the week is out.
- No tutorials past week 1. After that, read the standard library source and the Go memory model spec when you're stuck.

## Week 1 — Types, packages, and a single-threaded store

Build a working KV store with no concurrency. Single goroutine, single map, REST endpoints for GET/PUT/DELETE.

Things you should figure out, not look up:
- Why Go has both `[]byte` and `string`. When do you choose which?
- What does it mean that a map is a reference type? What happens when you pass it to a function?
- Errors as values — what's the actual ergonomic difference vs exceptions? When does it bite?
- Package layout: `cmd/`, `internal/`, `pkg/`. What's the convention and why does it exist?

Deliverable: `curl` can PUT a key, GET it back, DELETE it. No tests yet. Ugly is fine.

Open question to sit with: your handlers probably look repetitive. Is that idiomatic Go, or are you missing an abstraction? Don't refactor yet — just notice.

## Week 2 — Concurrency, the hard way

Make the store safe for concurrent access. Spin up a load generator (another small Go program) that hammers it with parallel reads and writes.

Do this in stages, and benchmark each:
1. Naive: wrap the map in `sync.Mutex`.
2. Smarter: `sync.RWMutex`. Measure the difference under read-heavy vs write-heavy loads. Does it match your intuition?
3. Stranger: replace the mutex with a single goroutine that owns the map, communicating via channels. This is the "share memory by communicating" pattern. Slower or faster? Why?

Things to figure out:
- What does `go test -race` actually detect? Trigger it on purpose at least once.
- What's the difference between buffered and unbuffered channels in terms of synchronization guarantees?
- When you close a channel, what happens to readers? Writers? Why is closing from the reader side a bug?

Open question: one of those three approaches is the "right" one for a KV store. Form an opinion before week 3 ends. You'll revisit it.

## Week 3 — Lifecycle, cancellation, and the things that go wrong

Add: TTL on keys (entries expire), graceful shutdown (SIGTERM drains in-flight requests), and a `/stats` endpoint that reports without blocking writes.

This is where `context.Context` and `select` earn their keep.

Things to figure out:
- Why does idiomatic Go pass `ctx` as the first argument to everything? What problem is it solving that a global cancel channel doesn't?
- How does the `net/http` server use context to signal client disconnect? Read the source.
- For TTL expiry: do you scan periodically, use a heap, use per-key timers? Three valid answers, each with different tradeoffs. Pick one and be able to defend it.
- What's a goroutine leak? Cause one accidentally, then fix it. Use `runtime.NumGoroutine()` to confirm.

Open question: your `/stats` endpoint needs to read counters that other goroutines are updating. `sync/atomic`, `sync.Mutex`, or a channel-based approach? The answer changes depending on how often you read vs write. Think about it.

## Week 4 — Testing, profiling, and something distributed-adjacent

Three threads, in parallel:

**Testing.** Table-driven tests for the handlers. Race-detector clean (`go test -race ./...`). At least one test that uses `httptest.Server`. Figure out what `t.Parallel()` does and when it's safe.

**Profiling.** Run a load test, capture a CPU profile with `pprof`, find the hot spot. Then capture a goroutine profile under load — what does it tell you about your design?

**The distributed teaser.** Add a `/replicate` endpoint and run two instances. Make writes to instance A asynchronously forward to instance B. Don't worry about consistency yet — just get bytes flowing.

Then sit with this: what could go wrong? Network partition between A and B. A goes down mid-write. B is slow and you're queueing writes faster than it accepts them. Write these failure modes in `questions.md`. You don't need to solve them — but if any of them feel mysterious, that's exactly what 6.824 is going to teach you.

## What "done" looks like

- You can explain, without notes, the difference between a mutex and a channel-based approach and when you'd use each.
- `go test -race ./...` passes.
- You've read at least three files from the Go standard library source (suggested: `net/http/server.go`, `sync/rwmutex.go`, `context/context.go`).
- Your `questions.md` has 15+ entries, most of them answered.

## What this won't teach you

This project gives you Go syntax, concurrency primitives, and the standard library. It does *not* teach you distributed systems — that's 6.824's job. The replication exercise in week 4 is a teaser, not a lesson. Don't go down the rabbit hole of implementing Raft here. Save that energy.

## A note on ambiguity

This roadmap deliberately under-specifies. The choices you make (RWMutex vs channels, periodic scan vs heap for TTL, where to put the package boundary) are the actual learning. If you find yourself wanting a "correct" answer, that's the signal to dig deeper instead of asking.
