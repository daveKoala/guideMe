package store

import (
	"database/sql"
	"errors"

	"github.com/david-clare/guide-me/backend/internal/domain"
)

// ErrNotFound is returned when a lookup matches no row.
var ErrNotFound = errors.New("store: not found")

// CreateAccount inserts a new account, minting an id if absent.
func (s *Store) CreateAccount(a *domain.Account) error {
	if a.ID == "" {
		a.ID = newID()
	}
	_, err := s.db.Exec(
		`INSERT INTO accounts (id, email) VALUES (?, ?)`,
		a.ID, nullStr(a.Email),
	)
	return err
}

// GetAccount loads an account by id.
func (s *Store) GetAccount(id string) (*domain.Account, error) {
	var a domain.Account
	var email sql.NullString
	err := s.db.QueryRow(
		`SELECT id, email, created_at, updated_at FROM accounts WHERE id = ?`, id,
	).Scan(&a.ID, &email, &a.CreatedAt, &a.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	a.Email = email.String
	return &a, nil
}
