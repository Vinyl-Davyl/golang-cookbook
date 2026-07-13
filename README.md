# Go Cookbook

A **personal, well-documented Go reference** written entirely in Go. Every topic is a runnable recipe with godoc comments, external references to official docs and books, and a CLI to browse and execute examples.

> Inspired by *The Go Programming Language* (Donovan & Kernighan), *Effective Go*, the [Go Tour](https://go.dev/tour/), and modern idioms from the Go 1.21–1.23 releases (`slog`, improved `net/http` routing, `slices`, generics).

## Requirements

- Go 1.23 or later
- Internet (optional — only needed for the HTTP client recipe)

## Quick Start

```bash
# Install dependencies
go mod tidy

# List all recipes by category
go run ./cmd/cookbook list

# Run a specific recipe
go run ./cmd/cookbook run errors-wrap

# Show documentation and references for a recipe
go run ./cmd/cookbook doc concurrency-errgroup

# Run all tests
go test ./...

# Browse godoc locally
go doc ./recipes/errors
```

## Recipe Index

### basics
| ID | Title |
|----|-------|
| `basics-variables` | Variables, Types, and Zero Values |
| `basics-control-flow` | Control Flow: if, for, switch, range |
| `basics-defer-panic-recover` | defer, panic, and recover |

### functions
| ID | Title |
|----|-------|
| `functions-multiple-returns` | Multiple Return Values |
| `functions-variadic` | Variadic Functions |
| `functions-closures` | Closures and Function Values |

### structs
| ID | Title |
|----|-------|
| `structs-embedding` | Struct Embedding (Composition) |
| `structs-method-receivers` | Value vs Pointer Receivers |

### pointers
| ID | Title |
|----|-------|
| `pointers-when-to-use` | When to Use Pointers |

### errors
| ID | Title |
|----|-------|
| `errors-wrap` | Wrapping Errors with Context |
| `errors-is-as` | errors.Is and errors.As |
| `errors-join` | Joining Multiple Errors |
| `errors-sentinel` | Sentinel Errors |
| `errors-custom-type` | Custom Error Types |

### interfaces
| ID | Title |
|----|-------|
| `interfaces-small` | Small, Focused Interfaces |
| `interfaces-type-assertions` | Type Assertions and Type Switches |
| `interfaces-accept-return` | Accept Interfaces, Return Structs |

### slices
| ID | Title |
|----|-------|
| `slices-append-capacity` | Append, Length, and Capacity |
| `slices-filter-map` | Filter and Map Patterns |
| `maps-patterns` | Map Idioms |

### strings
| ID | Title |
|----|-------|
| `strings-builder` | strings.Builder |
| `strings-json` | JSON Marshal and Unmarshal |

### io
| ID | Title |
|----|-------|
| `io-read-write-file` | Reading and Writing Files |
| `io-json-encode-decode` | Streaming JSON with Encoder/Decoder |

### concurrency
| ID | Title |
|----|-------|
| `concurrency-goroutines` | Goroutines |
| `concurrency-channels` | Channels |
| `concurrency-select` | select Statement |
| `concurrency-semaphore` | Bounded Concurrency (Semaphore) |
| `concurrency-producer-consumer` | Producer-Consumer Pattern |
| `concurrency-errgroup` | errgroup for Concurrent Tasks |

### context
| ID | Title |
|----|-------|
| `context-timeout-cancel` | Timeouts and Cancellation |
| `context-propagation` | Context Propagation |

### http
| ID | Title |
|----|-------|
| `http-server-routing` | HTTP Server with Go 1.22+ Routing |
| `http-client` | HTTP Client with Context |
| `http-middleware` | HTTP Middleware Pattern |

### testing
| ID | Title |
|----|-------|
| `testing-table-driven` | Table-Driven Tests |
| `testing-subtests-parallel` | Subtests and Parallel Execution |

### generics
| ID | Title |
|----|-------|
| `generics-functions` | Generic Functions |
| `generics-constraints` | Type Constraints |

### logging
| ID | Title |
|----|-------|
| `logging-slog` | Structured Logging with slog |

### sync
| ID | Title |
|----|-------|
| `sync-mutex-waitgroup` | sync.Mutex and sync.WaitGroup |
| `sync-once` | sync.Once for Lazy Initialization |

### time
| ID | Title |
|----|-------|
| `time-tickers-timers` | Timers and Tickers |

### project
| ID | Title |
|----|-------|
| `project-module-layout` | Modules and Project Layout |

### embed
| ID | Title |
|----|-------|
| `embed-files` | Embedding Files with embed |

## Project Structure

```
golang-cookbook/
├── cmd/cookbook/       # CLI to list, run, and document recipes
├── cookbook/           # Recipe model and registry
├── recipes/            # One package per topic area
│   ├── basics/
│   ├── errors/
│   ├── concurrency/
│   └── ...
├── go.mod
└── README.md
```

## How to Use as a Learning Reference

1. **Read top-down** — start with `basics`, then `functions`, `structs`, `errors`.
2. **Run every recipe** — `go run ./cmd/cookbook run <id>` prints live output.
3. **Read the source** — each recipe's `Run` function is the annotated example.
4. **Follow references** — every recipe links to official Go docs and blog posts.
5. **Browse godoc** — `go doc github.com/golang-cookbook/cookbook/recipes/errors`.

## Core Principles (2025–2026 Idioms)

| Principle | Recipe |
|-----------|--------|
| Errors are values, wrap with `%w` | `errors-wrap` |
| Pass `context.Context` as first param | `context-propagation` |
| Stdlib-first (`slog`, `net/http` routing) | `logging-slog`, `http-server-routing` |
| Small interfaces at the consumer | `interfaces-small` |
| Bounded concurrency over worker pools | `concurrency-semaphore`, `concurrency-errgroup` |
| Table-driven tests | `testing-table-driven` |
| Flat project layout, grow when needed | `project-module-layout` |

## Recommended Books & Official Resources

- [The Go Programming Language](https://www.gopl.io/) — Donovan & Kernighan
- [Effective Go](https://go.dev/doc/effective_go)
- [A Tour of Go](https://go.dev/tour/)
- [Go Blog](https://go.dev/blog/)
- [Go by Example](https://gobyexample.com/)
- [Go Wiki — CodeReviewComments](https://go.dev/wiki/CodeReviewComments)

## License

MIT — use freely for personal learning and reference.
