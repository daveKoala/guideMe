package store_test

import (
	"path/filepath"
	"testing"

	"github.com/david-clare/guide-me/backend/internal/domain"
	"github.com/david-clare/guide-me/backend/internal/store"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

// newTestStore opens a fresh temp-file database and applies the goose
// migrations so each test runs against the real schema.
func newTestStore(t *testing.T) *store.Store {
	t.Helper()
	dbPath := filepath.Join(t.TempDir(), "test.db")
	s, err := store.Open(dbPath)
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	t.Cleanup(func() { s.Close() })

	if err := goose.SetDialect("sqlite3"); err != nil {
		t.Fatalf("dialect: %v", err)
	}
	if err := goose.Up(s.DB(), "../../migrations"); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return s
}

func TestSaveAndLoadTripAggregate(t *testing.T) {
	s := newTestStore(t)

	acct := &domain.Account{Email: "test@guide-me.app"}
	if err := s.CreateAccount(acct); err != nil {
		t.Fatalf("account: %v", err)
	}

	dave := &domain.Person{AccountID: acct.ID, Name: "Dave", Type: "adult"}
	kid := &domain.Person{AccountID: acct.ID, Name: "Kid", Type: "child"}
	for _, p := range []*domain.Person{dave, kid} {
		if err := s.CreatePerson(p); err != nil {
			t.Fatalf("person: %v", err)
		}
	}

	trip := &domain.Trip{
		AccountID: acct.ID,
		Name:      "Test trip",
		Type:      "outbound",
		Status:    "planned",
		Timezone:  "Europe/London",
		Party: domain.Party{
			LeadPassenger: dave.ID,
			Passengers:    []string{dave.ID, kid.ID},
		},
		Sharing: domain.Sharing{OfflineEnabled: true},
		Stages: []domain.Stage{
			// Deliberately out of sort_order to prove LoadTrip orders them.
			{Kind: "flight", Start: "2026-07-01T07:30:00", SortOrder: 2,
				Values: map[string]string{"airline": "Vueling", "flight_no": "VY1"}},
			{Kind: "note", Start: "2026-07-01T05:00:00", SortOrder: 1,
				Values: map[string]string{"title": "Pack passports"}},
		},
		Insurance: []domain.Insurance{
			{PolicyNumber: "AXA-1", Covers: []string{dave.ID, kid.ID},
				Medical: domain.Medical{AssistID: "MA-1", Phone: "+44"}},
		},
	}

	if err := s.SaveTrip(trip); err != nil {
		t.Fatalf("SaveTrip: %v", err)
	}
	if trip.ID == "" || trip.Sharing.EditToken == "" || trip.Sharing.ReadToken == "" {
		t.Fatalf("SaveTrip did not mint id/tokens: %+v", trip.Sharing)
	}

	got, err := s.LoadTrip(trip.ID)
	if err != nil {
		t.Fatalf("LoadTrip: %v", err)
	}

	if got.Name != "Test trip" || got.AccountID != acct.ID {
		t.Errorf("meta mismatch: %+v", got)
	}
	if !got.Sharing.OfflineEnabled {
		t.Errorf("offline_enabled lost")
	}
	wantEditURL := "/trips/" + trip.ID + "/edit/" + trip.Sharing.EditToken
	if got.Sharing.OwnerEditURL != wantEditURL {
		t.Errorf("edit url = %q, want %q", got.Sharing.OwnerEditURL, wantEditURL)
	}

	// Party.
	if got.Party.LeadPassenger != dave.ID || len(got.Party.Passengers) != 2 {
		t.Errorf("party mismatch: %+v", got.Party)
	}

	// Stages ordered by sort_order: note (1) before flight (2).
	if len(got.Stages) != 2 {
		t.Fatalf("want 2 stages, got %d", len(got.Stages))
	}
	if got.Stages[0].Kind != "note" || got.Stages[1].Kind != "flight" {
		t.Errorf("stage order wrong: %s, %s", got.Stages[0].Kind, got.Stages[1].Kind)
	}
	if got.Stages[1].Values["airline"] != "Vueling" {
		t.Errorf("stage values json round-trip failed: %+v", got.Stages[1].Values)
	}

	// Insurance + covers + medical json.
	if len(got.Insurance) != 1 {
		t.Fatalf("want 1 policy, got %d", len(got.Insurance))
	}
	if len(got.Insurance[0].Covers) != 2 {
		t.Errorf("want 2 covers, got %d", len(got.Insurance[0].Covers))
	}
	if got.Insurance[0].Medical.AssistID != "MA-1" {
		t.Errorf("medical json round-trip failed: %+v", got.Insurance[0].Medical)
	}
}

func TestDeleteTripCascades(t *testing.T) {
	s := newTestStore(t)

	acct := &domain.Account{}
	if err := s.CreateAccount(acct); err != nil {
		t.Fatalf("account: %v", err)
	}
	p := &domain.Person{AccountID: acct.ID, Name: "Solo", Type: "adult"}
	if err := s.CreatePerson(p); err != nil {
		t.Fatalf("person: %v", err)
	}
	trip := &domain.Trip{
		AccountID: acct.ID, Name: "Doomed", Type: "outbound", Status: "planned", Timezone: "UTC",
		Party:  domain.Party{LeadPassenger: p.ID, Passengers: []string{p.ID}},
		Stages: []domain.Stage{{Kind: "note", SortOrder: 1, Values: map[string]string{"title": "x"}}},
	}
	if err := s.SaveTrip(trip); err != nil {
		t.Fatalf("SaveTrip: %v", err)
	}

	if err := s.DeleteTrip(trip.ID); err != nil {
		t.Fatalf("DeleteTrip: %v", err)
	}

	stages, err := s.ListStagesByTrip(trip.ID)
	if err != nil {
		t.Fatalf("ListStagesByTrip: %v", err)
	}
	if len(stages) != 0 {
		t.Errorf("stages not cascaded: %d remain", len(stages))
	}
	if _, err := s.LoadTrip(trip.ID); err != store.ErrNotFound {
		t.Errorf("want ErrNotFound after delete, got %v", err)
	}
}
