package store

import (
	"database/sql"

	"github.com/david-clare/guide-me/backend/internal/domain"
)

// CreateStage inserts a stage on a trip.
func (s *Store) CreateStage(st *domain.Stage) error {
	return insertStage(s.db, st)
}

// ListStagesByTrip returns a trip's stages in timeline order.
func (s *Store) ListStagesByTrip(tripID string) ([]domain.Stage, error) {
	return listStagesByTrip(s.db, tripID)
}

// UpdateStage replaces the mutable fields of an existing stage.
func (s *Store) UpdateStage(st *domain.Stage) error {
	_, err := s.db.Exec(
		`UPDATE stages
		 SET kind = ?, subkind = ?, start = ?, sort_order = ?, values_json = ?, updated_at = datetime('now')
		 WHERE id = ?`,
		st.Kind, nullStr(st.Subkind), st.Start, st.SortOrder, marshalJSON(st.Values), st.ID,
	)
	return err
}

// DeleteStage removes a stage (cascades to its boarding-pass documents).
func (s *Store) DeleteStage(id string) error {
	_, err := s.db.Exec(`DELETE FROM stages WHERE id = ?`, id)
	return err
}

// insertStage writes a stage row using the given execer (db or tx).
func insertStage(e execer, st *domain.Stage) error {
	if st.ID == "" {
		st.ID = newID()
	}
	if st.Values == nil {
		st.Values = map[string]string{}
	}
	_, err := e.Exec(
		`INSERT INTO stages (id, trip_id, kind, subkind, start, sort_order, values_json)
		 VALUES (?, ?, ?, ?, ?, ?, ?)`,
		st.ID, st.TripID, st.Kind, nullStr(st.Subkind), st.Start, st.SortOrder, marshalJSON(st.Values),
	)
	return err
}

// listStagesByTrip reads stages ordered by sort_order using the given execer.
func listStagesByTrip(e execer, tripID string) ([]domain.Stage, error) {
	rows, err := e.Query(
		`SELECT id, trip_id, kind, subkind, start, sort_order, values_json
		 FROM stages WHERE trip_id = ? ORDER BY sort_order, start`, tripID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []domain.Stage{}
	for rows.Next() {
		var st domain.Stage
		var subkind sql.NullString
		var valuesJSON string
		if err := rows.Scan(&st.ID, &st.TripID, &st.Kind, &subkind, &st.Start, &st.SortOrder, &valuesJSON); err != nil {
			return nil, err
		}
		st.Subkind = subkind.String
		st.Values = map[string]string{}
		if err := scanJSON(valuesJSON, &st.Values); err != nil {
			return nil, err
		}
		out = append(out, st)
	}
	return out, rows.Err()
}
