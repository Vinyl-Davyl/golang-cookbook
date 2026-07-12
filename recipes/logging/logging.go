// Package logging covers structured logging with log/slog (Go 1.21+).
package logging

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/golang-cookbook/cookbook"
)

// StructuredSlog demonstrates slog for structured, leveled logging.
func StructuredSlog() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "logging-slog",
		Category: "logging",
		Title:    "Structured Logging with slog",
		Summary:  "Use log/slog instead of log.Printf; attach key-value attributes for machine-readable logs.",
		References: []cookbook.Reference{
			{Title: "Go 1.21 Release Notes — log/slog", URL: "https://go.dev/doc/go1.21#slog"},
			{Title: "log/slog package", URL: "https://pkg.go.dev/log/slog"},
		},
		Run: func() error {
			logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}))

			logger.Info("server started",
				slog.String("host", "localhost"),
				slog.Int("port", 8080),
			)

			logger.Error("request failed",
				slog.String("method", "GET"),
				slog.String("path", "/api/users"),
				slog.Int("status", 500),
			)

			// With helper for common attributes
			reqLogger := logger.With(slog.String("request_id", "abc-123"))
			reqLogger.Debug("processing request")

			fmt.Println("(see structured log lines above)")
			return nil
		},
	}
}
