// Package httprec covers HTTP server and client patterns using the stdlib (Go 1.22+ routing).
package httprec

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/golang-cookbook/cookbook"
)

// ServerRouting demonstrates Go 1.22+ method-aware routing in net/http.
func ServerRouting() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "http-server-routing",
		Category: "http",
		Title:    "HTTP Server with Go 1.22+ Routing",
		Summary:  "Use net/http.ServeMux with method and path patterns — no third-party router required.",
		References: []cookbook.Reference{
			{Title: "Go 1.22 Release Notes — HTTP routing", URL: "https://go.dev/doc/go1.22#routing"},
			{Title: "net/http package", URL: "https://pkg.go.dev/net/http"},
		},
		Run: func() error {
			mux := http.NewServeMux()
			mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "ok")
			})
			mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "user id: %s", r.PathValue("id"))
			})

			ln, err := net.Listen("tcp", "127.0.0.1:0")
			if err != nil {
				return err
			}
			srv := &http.Server{Handler: mux}
			go func() { _ = srv.Serve(ln) }()
			defer srv.Shutdown(context.Background())

			base := "http://" + ln.Addr().String()
			resp, err := http.Get(base + "/health")
			if err != nil {
				return err
			}
			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			fmt.Printf("GET /health → %s\n", body)

			resp, err = http.Get(base + "/users/42")
			if err != nil {
				return err
			}
			body, _ = io.ReadAll(resp.Body)
			resp.Body.Close()
			fmt.Printf("GET /users/42 → %s\n", body)
			return nil
		},
	}
}

// Client demonstrates creating requests with context and timeouts.
func Client() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "http-client",
		Category: "http",
		Title:    "HTTP Client with Context",
		Summary:  "Use http.NewRequestWithContext; set client timeouts; always close response bodies.",
		References: []cookbook.Reference{
			{Title: "net/http — Client", URL: "https://pkg.go.dev/net/http#Client"},
		},
		Run: func() error {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://httpbin.org/get", nil)
			if err != nil {
				return err
			}

			client := &http.Client{Timeout: 10 * time.Second}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("request failed (network may be unavailable): %v\n", err)
				return nil
			}
			defer resp.Body.Close()

			fmt.Printf("status: %s\n", resp.Status)
			return nil
		},
	}
}

// Middleware demonstrates wrapping handlers for cross-cutting concerns.
func Middleware() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "http-middleware",
		Category: "http",
		Title:    "HTTP Middleware Pattern",
		Summary:  "Middleware is a function that wraps http.Handler to add logging, auth, recovery, etc.",
		References: []cookbook.Reference{
			{Title: "Writing Web Applications", URL: "https://go.dev/doc/articles/wiki/"},
		},
		Run: func() error {
			logging := func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					fmt.Printf("  → %s %s\n", r.Method, r.URL.Path)
					next.ServeHTTP(w, r)
				})
			}

			handler := logging(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "hello")
			}))

			mux := http.NewServeMux()
			mux.Handle("GET /", handler)

			ln, err := net.Listen("tcp", "127.0.0.1:0")
			if err != nil {
				return err
			}
			srv := &http.Server{Handler: mux}
			go func() { _ = srv.Serve(ln) }()
			defer srv.Shutdown(context.Background())

			resp, err := http.Get("http://" + ln.Addr().String() + "/")
			if err != nil {
				return err
			}
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
			fmt.Println("middleware logged the request")
			return nil
		},
	}
}
