// Package blob abstracts document storage behind a small interface so the dev
// build can write to the local filesystem while production swaps in cloud blob
// storage without touching callers.
package blob

import (
	"context"
	"fmt"
	"io"
)

// Store reads and writes opaque objects keyed by a storage key.
type Store interface {
	// Put writes r under key, overwriting any existing object.
	Put(ctx context.Context, key string, r io.Reader, contentType string) error
	// Get opens the object at key for reading. Caller closes the reader.
	Get(ctx context.Context, key string) (io.ReadCloser, error)
	// Delete removes the object at key. Missing keys are not an error.
	Delete(ctx context.Context, key string) error
}

// New returns a Store for the given backend. Only "local" is wired today.
func New(backend, dir string) (Store, error) {
	switch backend {
	case "local", "":
		return NewLocalFS(dir)
	default:
		return nil, fmt.Errorf("unknown blob backend %q", backend)
	}
}
