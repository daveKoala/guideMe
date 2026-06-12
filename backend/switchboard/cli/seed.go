package main

import (
	"fmt"

	"github.com/david-clare/guide-me/backend/internal/domain"
	"github.com/david-clare/guide-me/backend/internal/store"
)

// seedPerson mirrors one entry from the frontend mock_users.ts, keyed by its
// readable mock id so the trip's references resolve after id remapping.
type seedPerson struct {
	mockID      string
	name        string
	ptype       domain.PersonType
	ghicID      string
	dob         string
	hasPassport bool
}

var seedPeople = []seedPerson{
	{"person_dave", "Dave", "adult", "GHIC-90011223344", "1985-04-12", true},
	{"person_adult2", "Adult 2", "adult", "GHIC-90011223355", "1987-09-30", true},
	{"person_child1", "Child 1", "child", "", "2015-02-20", false},
	{"person_child2", "Child 2", "child", "", "2017-06-08", false},
	{"person_child3", "Child 3", "child", "", "2019-11-15", false},
	{"person_sue", "Grandma Sue", "adult", "GHIC-90011229999", "1956-01-03", true},
	{"person_tom", "Friend Tom", "adult", "", "1990-07-22", false},
}

// runSeed clears the database and loads the mock account, people, demo trip and
// document metadata. Idempotent: clearing first means re-running is safe.
// Server-generated UUIDs replace the readable mock ids; an in-memory map keeps
// the trip's references (party, covers, boarding passes) pointing at the right
// rows.
func runSeed(dbPath string) error {
	s, err := store.Open(dbPath)
	if err != nil {
		return err
	}
	defer s.Close()

	if err := clearAll(s); err != nil {
		return fmt.Errorf("clear: %w", err)
	}

	// Account.
	account := &domain.Account{Email: "demo@guide-me.app"}
	if err := s.CreateAccount(account); err != nil {
		return fmt.Errorf("account: %w", err)
	}

	// People — record mock id -> new uuid.
	idOf := map[string]string{}
	for _, sp := range seedPeople {
		p := &domain.Person{
			AccountID: account.ID,
			Name:      sp.name,
			Type:      sp.ptype,
			GhicID:    sp.ghicID,
			Dob:       sp.dob,
		}
		if err := s.CreatePerson(p); err != nil {
			return fmt.Errorf("person %s: %w", sp.mockID, err)
		}
		idOf[sp.mockID] = p.ID

		if sp.hasPassport {
			if err := s.CreateDocument(&domain.Document{
				AccountID:   account.ID,
				Kind:        "passport",
				PersonID:    p.ID,
				FileName:    "passport.png",
				StorageKey:  fmt.Sprintf("seed/%s/passport.png", sp.mockID),
				ContentType: "image/png",
			}); err != nil {
				return fmt.Errorf("passport %s: %w", sp.mockID, err)
			}
		}
	}

	// Demo trip aggregate (mirrors mock_trip.ts).
	trip := buildDemoTrip(account.ID, idOf)
	if err := s.SaveTrip(trip); err != nil {
		return fmt.Errorf("trip: %w", err)
	}

	// Boarding passes hang off the flight stage (now it has an id).
	flightID := findStageID(trip, "flight")
	for _, mockID := range []string{"person_dave", "person_adult2"} {
		if err := s.CreateDocument(&domain.Document{
			AccountID:   account.ID,
			Kind:        "boarding_pass",
			PersonID:    idOf[mockID],
			StageID:     flightID,
			FileName:    "boardingpass.png",
			StorageKey:  fmt.Sprintf("seed/%s/boardingpass.png", mockID),
			ContentType: "image/png",
		}); err != nil {
			return fmt.Errorf("boarding pass %s: %w", mockID, err)
		}
	}

	fmt.Printf("seeded: account=%s people=%d trip=%s\n", account.ID, len(seedPeople), trip.ID)
	return nil
}

// buildDemoTrip constructs the Barcelona trip, mapping mock person ids to uuids.
func buildDemoTrip(accountID string, idOf map[string]string) *domain.Trip {
	party := func(ids ...string) []string {
		out := make([]string, len(ids))
		for i, id := range ids {
			out[i] = idOf[id]
		}
		return out
	}

	return &domain.Trip{
		AccountID:      accountID,
		Name:           "Family trip to Barcelona",
		Type:           "outbound",
		Status:         "planned",
		Timezone:       "Europe/London",
		NeedsInsurance: false,
		NeedsGhic:      false,
		Party: domain.Party{
			LeadPassenger: idOf["person_dave"],
			Passengers:    party("person_dave", "person_adult2", "person_child1", "person_child2", "person_child3"),
		},
		Sharing: domain.Sharing{OfflineEnabled: true},
		Stages: []domain.Stage{
			{
				Kind: "travel_to_airport", Subkind: "taxi", Start: "2026-07-01T05:00:00", SortOrder: 0,
				Values: map[string]string{
					"company":         "Addison Lee",
					"pickup_time":     "2026-07-01T05:00:00",
					"pickup_location": "12 Elm Road, Manchester",
				},
			},
			{
				Kind: "note", Start: "2026-07-02T09:00:00", SortOrder: 1,
				Values: map[string]string{
					"title": "Sagrada Família tickets",
					"when":  "2026-07-02T09:00:00",
					"notes": "Pre-booked timed entry at 09:15. Bring passports for ID check.",
				},
			},
			{
				Kind: "flight", Start: "2026-07-01T07:30:00", SortOrder: 2,
				Values: map[string]string{
					"airline": "Vueling", "flight_no": "VY7821",
					"from": "MAN", "to": "BCN",
					"depart": "2026-07-01T07:30:00", "arrive": "2026-07-01T10:50:00",
				},
			},
			{
				Kind: "accommodation", Start: "2026-07-01T14:00:00", SortOrder: 3,
				Values: map[string]string{
					"name":      "Hotel Arts Barcelona",
					"address":   "Marina, 19-21, 08005 Barcelona",
					"check_in":  "2026-07-01T14:00:00",
					"check_out": "2026-07-08T11:00:00",
				},
			},
			{
				Kind: "note", Start: "2026-07-02T09:00:00", SortOrder: 4,
				Values: map[string]string{
					"title": "Sagrada Família tickets",
					"when":  "2026-07-02T09:00:00",
					"notes": "Pre-booked timed entry at 09:15. Bring passports for ID check.",
				},
			},
		},
		Insurance: []domain.Insurance{
			{
				PolicyNumber:     "AXA-99481726",
				EmergencyContact: "+44 20 7946 0000",
				AccountURL:       "https://example-insurer.com/account",
				Covers:           party("person_dave", "person_adult2", "person_child1", "person_child2", "person_child3"),
				Medical:          domain.Medical{AssistID: "MA-55012", Phone: "+44 20 7946 1111", URL: "https://example-assist.com/members"},
			},
			{
				PolicyNumber:     "SKI-22310",
				EmergencyContact: "+44 20 7946 2222",
				AccountURL:       "https://example-ski-insurer.com",
				Covers:           party("person_dave"),
				Medical:          domain.Medical{AssistID: "SKI-ASSIST-77", Phone: "+44 20 7946 3333", URL: "https://example-ski-assist.com"},
			},
		},
	}
}

// findStageID returns the id of the first stage of the given kind after SaveTrip
// has populated ids.
func findStageID(t *domain.Trip, kind string) string {
	for _, st := range t.Stages {
		if st.Kind == kind {
			return st.ID
		}
	}
	return ""
}

// clearAll empties every table. Deleting accounts cascades to everything else.
func clearAll(s *store.Store) error {
	_, err := s.DB().Exec(`DELETE FROM accounts`)
	return err
}
