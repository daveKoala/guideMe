// Package config loads application configuration following the 12-factor
// approach: hard-coded defaults are overridden by a .env file, which is in turn
// overridden by real environment variables. Pure standard library.
package config

import (
	"bufio"
	"os"
	"strings"
)

// Config holds all runtime configuration. Every field has a sane default so the
// server and CLI run with zero setup.
type Config struct {
	// Port the HTTP server listens on.
	Port string
	// Env names the runtime environment (development, production, ...).
	Env string
	// DBPath is the SQLite database file path.
	DBPath string
	// LogLevel controls log verbosity (debug, info, warn, error).
	LogLevel string
	// BlobBackend selects the document storage backend (local, ...).
	BlobBackend string
	// BlobDir is the local filesystem directory for the local blob backend.
	BlobDir string
}

// defaults defines the baseline configuration. Keys match environment variable
// names.
var defaults = map[string]string{
	"PORT":         "8080",
	"ENV":          "development",
	"DB_PATH":      "./guide-me.db",
	"LOG_LEVEL":    "info",
	"BLOB_BACKEND": "local",
	"BLOB_DIR":     "./blobs",
}

// Load builds a Config using the precedence: defaults < .env file < real env.
// The .env file is optional; a missing file is not an error.
func Load() Config {
	loadDotEnv(".env")
	return Config{
		Port:     get("PORT"),
		Env:      get("ENV"),
		DBPath:      get("DB_PATH"),
		LogLevel:    get("LOG_LEVEL"),
		BlobBackend: get("BLOB_BACKEND"),
		BlobDir:     get("BLOB_DIR"),
	}
}

// get returns the real env var if set, otherwise the compiled-in default.
func get(key string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return defaults[key]
}

// loadDotEnv parses a simple KEY=VALUE file and sets any keys that are not
// already present in the real environment. Blank lines and lines beginning with
// '#' are ignored. Surrounding quotes on values are stripped. Real environment
// variables always win, so this never clobbers an explicit override.
func loadDotEnv(path string) {
	f, err := os.Open(path)
	if err != nil {
		return // optional file
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, val, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		val = strings.Trim(strings.TrimSpace(val), `"'`)
		if key == "" {
			continue
		}
		if _, exists := os.LookupEnv(key); !exists {
			_ = os.Setenv(key, val)
		}
	}
}
