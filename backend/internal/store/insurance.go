package store

import (
	"database/sql"

	"github.com/david-clare/guide-me/backend/internal/domain"
)

// CreateInsurance inserts a policy and its covers join rows.
func (s *Store) CreateInsurance(ins *domain.Insurance) error {
	return insertInsurance(s.db, ins)
}

// ListInsuranceByTrip returns a trip's policies with their covers populated.
func (s *Store) ListInsuranceByTrip(tripID string) ([]domain.Insurance, error) {
	return listInsuranceByTrip(s.db, tripID)
}

// DeleteInsurance removes a policy (cascades to its covers).
func (s *Store) DeleteInsurance(id string) error {
	_, err := s.db.Exec(`DELETE FROM insurance WHERE id = ?`, id)
	return err
}

// insertInsurance writes a policy + covers using the given execer.
func insertInsurance(e execer, ins *domain.Insurance) error {
	if ins.ID == "" {
		ins.ID = newID()
	}
	if _, err := e.Exec(
		`INSERT INTO insurance (id, trip_id, policy_number, emergency_contact, account_url, medical_json)
		 VALUES (?, ?, ?, ?, ?, ?)`,
		ins.ID, ins.TripID, nullStr(ins.PolicyNumber), nullStr(ins.EmergencyContact),
		nullStr(ins.AccountURL), marshalJSON(ins.Medical),
	); err != nil {
		return err
	}
	for _, personID := range ins.Covers {
		if _, err := e.Exec(
			`INSERT INTO insurance_covers (insurance_id, person_id) VALUES (?, ?)`,
			ins.ID, personID,
		); err != nil {
			return err
		}
	}
	return nil
}

// listInsuranceByTrip reads policies + their covers using the given execer.
func listInsuranceByTrip(e execer, tripID string) ([]domain.Insurance, error) {
	rows, err := e.Query(
		`SELECT id, trip_id, policy_number, emergency_contact, account_url, medical_json
		 FROM insurance WHERE trip_id = ? ORDER BY created_at`, tripID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []domain.Insurance{}
	for rows.Next() {
		var ins domain.Insurance
		var policy, emergency, accountURL sql.NullString
		var medicalJSON string
		if err := rows.Scan(&ins.ID, &ins.TripID, &policy, &emergency, &accountURL, &medicalJSON); err != nil {
			return nil, err
		}
		ins.PolicyNumber = policy.String
		ins.EmergencyContact = emergency.String
		ins.AccountURL = accountURL.String
		if err := scanJSON(medicalJSON, &ins.Medical); err != nil {
			return nil, err
		}
		out = append(out, ins)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Populate covers per policy.
	for i := range out {
		covers, err := listCovers(e, out[i].ID)
		if err != nil {
			return nil, err
		}
		out[i].Covers = covers
	}
	return out, nil
}

// listCovers reads the person ids covered by a policy.
func listCovers(e execer, insuranceID string) ([]string, error) {
	rows, err := e.Query(
		`SELECT person_id FROM insurance_covers WHERE insurance_id = ? ORDER BY person_id`,
		insuranceID,
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
