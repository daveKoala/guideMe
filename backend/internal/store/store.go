// Package store is a thin SQLite repository over the domain entities. It hides
// the hybrid layout (real rows + a couple of JSON columns) behind plain Go
// calls; the headline helpers are LoadTrip / SaveTrip which assemble and persist
// the full Trip aggregate in one transaction.
package store

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// Store wraps a *sql.DB connection to the SQLite database.
type Store struct {
	db *sql.DB
}

// execer is satisfied by both *sql.DB and *sql.Tx, letting the row helpers run
// inside or outside a transaction.
type execer interface {
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

// Open dials the SQLite file, enabling foreign keys and a busy timeout, then
// returns a ready Store.
func Open(dbPath string) (*Store, error) {
	// foreign_keys must be on per-connection; busy_timeout avoids SQLITE_BUSY
	// under the brief write contention of the CLI/tests.
	dsn := fmt.Sprintf("file:%s?_pragma=foreign_keys(1)&_pragma=busy_timeout(5000)", dbPath)
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}
	return &Store{db: db}, nil
}

// New wraps an existing *sql.DB (used by tests).
func New(db *sql.DB) *Store { return &Store{db: db} }

// DB exposes the underlying handle for callers that need raw access.
func (s *Store) DB() *sql.DB { return s.db }

// Close closes the database connection.
func (s *Store) Close() error { return s.db.Close() }

// newID mints a random UUID string for a new row.
func newID() string { return uuid.NewString() }

// newToken mints a random sharing token.
func newToken() string { return uuid.NewString() }

// marshalJSON serializes v for a *_json column. Never returns an error for the
// map/struct types we store; falls back to "{}" defensively.
func marshalJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(b)
}

// scanJSON unmarshals a *_json column into out. Empty strings are treated as
// an empty object so zero-value columns don't error.
func scanJSON(raw string, out any) error {
	if raw == "" {
		raw = "{}"
	}
	return json.Unmarshal([]byte(raw), out)
}

// nullStr converts a possibly-empty string to a SQL NULL when empty.
func nullStr(s string) any {
	if s == "" {
		return nil
	}
	return s
}

// boolToInt maps a bool to SQLite's 0/1 integer representation.
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
