// Package io covers file I/O and JSON streaming patterns.
package io

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/golang-cookbook/cookbook"
)

// ReadWriteFile demonstrates os.ReadFile, os.WriteFile, and path handling.
func ReadWriteFile() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "io-read-write-file",
		Category: "io",
		Title:    "Reading and Writing Files",
		Summary:  "Use os.ReadFile/WriteFile for small files; io.Copy for streams; always handle errors.",
		References: []cookbook.Reference{
			{Title: "os package", URL: "https://pkg.go.dev/os"},
			{Title: "io package", URL: "https://pkg.go.dev/io"},
		},
		Run: func() error {
			dir := os.TempDir()
			path := filepath.Join(dir, "cookbook-demo.txt")

			content := []byte("hello from the Go cookbook\n")
			if err := os.WriteFile(path, content, 0o644); err != nil {
				return fmt.Errorf("write: %w", err)
			}

			read, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("read: %w", err)
			}
			fmt.Printf("read %d bytes: %q", len(read), read)

			_ = os.Remove(path)
			return nil
		},
	}
}

// JSONEncodeDecode demonstrates streaming JSON with Encoder/Decoder.
func JSONEncodeDecode() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "io-json-encode-decode",
		Category: "io",
		Title:    "Streaming JSON with Encoder/Decoder",
		Summary:  "Use json.NewEncoder/NewDecoder for HTTP bodies and large streams instead of loading all at once.",
		References: []cookbook.Reference{
			{Title: "encoding/json — Encoder", URL: "https://pkg.go.dev/encoding/json#Encoder"},
		},
		Run: func() error {
			type Event struct {
				Type string `json:"type"`
				ID   int    `json:"id"`
			}

			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)
			for _, e := range []Event{{"click", 1}, {"view", 2}} {
				if err := enc.Encode(e); err != nil {
					return err
				}
			}
			fmt.Printf("encoded stream:\n%s", buf.String())

			dec := json.NewDecoder(&buf)
			for {
				var e Event
				if err := dec.Decode(&e); err != nil {
					if err != io.EOF {
						return err
					}
					break
				}
				fmt.Printf("decoded: %+v\n", e)
			}
			return nil
		},
	}
}
