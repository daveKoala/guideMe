package store

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/david-clare/guide-me/backend/internal/domain"
)

// ListTripsByAccount returns trip aggregates for an account.
func (s *Store) ListTripsByAccount(accountID string) ([]domain.Trip, error) {
	rows, err := s.db.Query(`SELECT id FROM trips WHERE account_id = ? ORDER BY created_at`, accountID)
	if err != nil {
		return nil, err
	}
	ids := []string{}
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			rows.Close()
			return nil, err
		}
		ids = append(ids, id)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}

	out := make([]domain.Trip, 0, len(ids))
	for _, id := range ids {
		t, err := s.LoadTrip(id)
		if err != nil {
			return nil, err
		}
		out = append(out, *t)
	}
	return out, nil
}

// LoadTrip assembles the full Trip aggregate: meta + party + sharing + stages
// (ordered) + insurance (with covers).
func (s *Store) LoadTrip(id string) (*domain.Trip, error) {
	var t domain.Trip
	var lead sql.NullString
	var needsIns, needsGhic, offline int
	err := s.db.QueryRow(
		`SELECT id, account_id, name, type, status, timezone,
		        needs_insurance, needs_ghic, lead_passenger_id,
		        edit_token, read_token, offline_enabled, created_at, updated_at
		 FROM trips WHERE id = ?`, id,
	).Scan(
		&t.ID, &t.AccountID, &t.Name, &t.Type, &t.Status, &t.Timezone,
		&needsIns, &needsGhic, &lead,
		&t.Sharing.EditToken, &t.Sharing.ReadToken, &offline, &t.CreatedAt, &t.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	t.NeedsInsurance = needsIns != 0
	t.NeedsGhic = needsGhic != 0
	t.Sharing.OfflineEnabled = offline != 0
	t.Sharing.OwnerEditURL = fmt.Sprintf("/trips/%s/edit/%s", t.ID, t.Sharing.EditToken)
	t.Sharing.ReadOnlyURL = fmt.Sprintf("/trips/%s/share/%s", t.ID, t.Sharing.ReadToken)

	passengers, err := s.listPassengers(t.ID)
	if err != nil {
		return nil, err
	}
	t.Party = domain.Party{LeadPassenger: lead.String, Passengers: passengers}

	if t.Stages, err = listStagesByTrip(s.db, t.ID); err != nil {
		return nil, err
	}
	if t.Insurance, err = listInsuranceByTrip(s.db, t.ID); err != nil {
		return nil, err
	}
	return &t, nil
}

// SaveTrip upserts the entire aggregate in one transaction: the trip row, its
// party, stages and insurance are fully replaced to match the given struct.
// Tokens are minted on first save when absent.
func (s *Store) SaveTrip(t *domain.Trip) (err error) {
	if t.ID == "" {
		t.ID = newID()
	}
	if t.Sharing.EditToken == "" {
		t.Sharing.EditToken = newToken()
	}
	if t.Sharing.ReadToken == "" {
		t.Sharing.ReadToken = newToken()
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	if _, err = tx.Exec(
		`INSERT INTO trips
		   (id, account_id, name, type, status, timezone, needs_insurance, needs_ghic,
		    lead_passenger_id, edit_token, read_token, offline_enabled)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		 ON CONFLICT(id) DO UPDATE SET
		    account_id = excluded.account_id, name = excluded.name, type = excluded.type,
		    status = excluded.status, timezone = excluded.timezone,
		    needs_insurance = excluded.needs_insurance, needs_ghic = excluded.needs_ghic,
		    lead_passenger_id = excluded.lead_passenger_id,
		    edit_token = excluded.edit_token, read_token = excluded.read_token,
		    offline_enabled = excluded.offline_enabled, updated_at = datetime('now')`,
		t.ID, t.AccountID, t.Name, string(t.Type), string(t.Status), t.Timezone,
		boolToInt(t.NeedsInsurance), boolToInt(t.NeedsGhic), nullStr(t.Party.LeadPassenger),
		t.Sharing.EditToken, t.Sharing.ReadToken, boolToInt(t.Sharing.OfflineEnabled),
	); err != nil {
		return err
	}

	// Replace party.
	if _, err = tx.Exec(`DELETE FROM trip_passengers WHERE trip_id = ?`, t.ID); err != nil {
		return err
	}
	for _, personID := range t.Party.Passengers {
		if _, err = tx.Exec(
			`INSERT INTO trip_passengers (trip_id, person_id) VALUES (?, ?)`, t.ID, personID,
		); err != nil {
			return err
		}
	}

	// Replace stages.
	if _, err = tx.Exec(`DELETE FROM stages WHERE trip_id = ?`, t.ID); err != nil {
		return err
	}
	for i := range t.Stages {
		t.Stages[i].TripID = t.ID
		if t.Stages[i].SortOrder == 0 {
			t.Stages[i].SortOrder = i
		}
		if err = insertStage(tx, &t.Stages[i]); err != nil {
			return err
		}
	}

	// Replace insurance.
	if _, err = tx.Exec(`DELETE FROM insurance WHERE trip_id = ?`, t.ID); err != nil {
		return err
	}
	for i := range t.Insurance {
		t.Insurance[i].TripID = t.ID
		if err = insertInsurance(tx, &t.Insurance[i]); err != nil {
			return err
		}
	}

	return tx.Commit()
}

// DeleteTrip removes a trip and all its children via cascade.
func (s *Store) DeleteTrip(id string) error {
	_, err := s.db.Exec(`DELETE FROM trips WHERE id = ?`, id)
	return err
}

// listPassengers reads the ordered passenger person ids for a trip.
func (s *Store) listPassengers(tripID string) ([]string, error) {
	rows, err := s.db.Query(
		`SELECT person_id FROM trip_passengers WHERE trip_id = ? ORDER BY person_id`, tripID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []string{}
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		out = append(out, id)
	}
	return out, rows.Err()
}
