package store

import (
	"database/sql"

	"github.com/david-clare/guide-me/backend/internal/domain"
)

// CreateDocument inserts a document metadata row, minting an id if absent.
func (s *Store) CreateDocument(d *domain.Document) error {
	if d.ID == "" {
		d.ID = newID()
	}
	_, err := s.db.Exec(
		`INSERT INTO documents (id, account_id, kind, person_id, stage_id, file_name, storage_key, content_type)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		d.ID, d.AccountID, string(d.Kind), d.PersonID, nullStr(d.StageID),
		d.FileName, d.StorageKey, nullStr(d.ContentType),
	)
	return err
}

// ListDocumentsByPerson returns a person's documents (passports, GHIC, etc).
func (s *Store) ListDocumentsByPerson(personID string) ([]domain.Document, error) {
	rows, err := s.db.Query(
		`SELECT id, account_id, kind, person_id, stage_id, file_name, storage_key, content_type, uploaded_at
		 FROM documents WHERE person_id = ? ORDER BY uploaded_at`, personID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanDocuments(rows)
}

// ListDocumentsByStage returns the boarding passes attached to a stage.
func (s *Store) ListDocumentsByStage(stageID string) ([]domain.Document, error) {
	rows, err := s.db.Query(
		`SELECT id, account_id, kind, person_id, stage_id, file_name, storage_key, content_type, uploaded_at
		 FROM documents WHERE stage_id = ? ORDER BY uploaded_at`, stageID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanDocuments(rows)
}

// DeleteDocument removes a document metadata row by id.
func (s *Store) DeleteDocument(id string) error {
	_, err := s.db.Exec(`DELETE FROM documents WHERE id = ?`, id)
	return err
}

func scanDocuments(rows *sql.Rows) ([]domain.Document, error) {
	var out []domain.Document
	for rows.Next() {
		var d domain.Document
		var stageID, contentType sql.NullString
		if err := rows.Scan(
			&d.ID, &d.AccountID, &d.Kind, &d.PersonID, &stageID,
			&d.FileName, &d.StorageKey, &contentType, &d.UploadedAt,
		); err != nil {
			return nil, err
		}
		d.StageID = stageID.String
		d.ContentType = contentType.String
		out = append(out, d)
	}
	return out, rows.Err()
}
