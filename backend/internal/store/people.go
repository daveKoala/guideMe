package store

import (
	"database/sql"
	"errors"

	"github.com/david-clare/guide-me/backend/internal/domain"
)

// CreatePerson inserts a person, minting an id if absent.
func (s *Store) CreatePerson(p *domain.Person) error {
	if p.ID == "" {
		p.ID = newID()
	}
	_, err := s.db.Exec(
		`INSERT INTO people (id, account_id, name, type, ghic_id, dob)
		 VALUES (?, ?, ?, ?, ?, ?)`,
		p.ID, p.AccountID, p.Name, string(p.Type), nullStr(p.GhicID), nullStr(p.Dob),
	)
	return err
}

// GetPerson loads a person (without documents) by id.
func (s *Store) GetPerson(id string) (*domain.Person, error) {
	return scanPerson(s.db.QueryRow(
		`SELECT id, account_id, name, type, ghic_id, dob, created_at, updated_at
		 FROM people WHERE id = ?`, id,
	))
}

// ListPeopleByAccount returns every person under an account.
func (s *Store) ListPeopleByAccount(accountID string) ([]domain.Person, error) {
	rows, err := s.db.Query(
		`SELECT id, account_id, name, type, ghic_id, dob, created_at, updated_at
		 FROM people WHERE account_id = ? ORDER BY created_at`, accountID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []domain.Person
	for rows.Next() {
		p, err := scanPerson(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, *p)
	}
	return out, rows.Err()
}

// UpdatePerson writes the mutable fields of an existing person.
func (s *Store) UpdatePerson(p *domain.Person) error {
	_, err := s.db.Exec(
		`UPDATE people
		 SET name = ?, type = ?, ghic_id = ?, dob = ?, updated_at = datetime('now')
		 WHERE id = ?`,
		p.Name, string(p.Type), nullStr(p.GhicID), nullStr(p.Dob), p.ID,
	)
	return err
}

// DeletePerson removes a person (cascades to their documents).
func (s *Store) DeletePerson(id string) error {
	_, err := s.db.Exec(`DELETE FROM people WHERE id = ?`, id)
	return err
}

// rowScanner is satisfied by both *sql.Row and *sql.Rows.
type rowScanner interface {
	Scan(dest ...any) error
}

func scanPerson(row rowScanner) (*domain.Person, error) {
	var p domain.Person
	var ghic, dob sql.NullString
	err := row.Scan(&p.ID, &p.AccountID, &p.Name, &p.Type, &ghic, &dob, &p.CreatedAt, &p.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	p.GhicID = ghic.String
	p.Dob = dob.String
	return &p, nil
}
