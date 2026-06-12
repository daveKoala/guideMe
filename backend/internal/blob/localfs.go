package blob

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// LocalFS stores objects as files under a base directory. The storage key is
// used as a relative path, so keys like "accounts/<id>/passport/<uuid>.png"
// nest naturally on disk.
type LocalFS struct {
	base string
}

// NewLocalFS creates the base directory if needed and returns a LocalFS.
func NewLocalFS(dir string) (*LocalFS, error) {
	if dir == "" {
		return nil, fmt.Errorf("blob: local dir is empty")
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, fmt.Errorf("blob: create base dir: %w", err)
	}
	return &LocalFS{base: dir}, nil
}

// path resolves key to an absolute path, guarding against escaping the base.
func (l *LocalFS) path(key string) (string, error) {
	clean := filepath.Clean("/" + key) // force key to be treated as rooted
	full := filepath.Join(l.base, clean)
	if !strings.HasPrefix(full, filepath.Clean(l.base)+string(os.PathSeparator)) {
		return "", fmt.Errorf("blob: invalid key %q", key)
	}
	return full, nil
}

// Put writes r to the file at key, creating parent directories.
func (l *LocalFS) Put(ctx context.Context, key string, r io.Reader, _ string) error {
	full, err := l.path(key)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		return err
	}
	f, err := os.Create(full)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := io.Copy(f, r); err != nil {
		return err
	}
	return nil
}

// Get opens the file at key.
func (l *LocalFS) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	full, err := l.path(key)
	if err != nil {
		return nil, err
	}
	return os.Open(full)
}

// Delete removes the file at key; a missing file is not an error.
func (l *LocalFS) Delete(ctx context.Context, key string) error {
	full, err := l.path(key)
	if err != nil {
		return err
	}
	if err := os.Remove(full); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}
