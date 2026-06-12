// Package domain holds the core entities the backend persists. The shapes
// mirror the Vue frontend's TypeScript types (src/types/*.ts) but are
// Go-idiomatic and annotated for JSON responses.
package domain

// PersonType is one of adult | child | infant.
type PersonType string

// TripType is one of outbound | return.
type TripType string

// TripStatus is one of planned | booked | completed | cancelled.
type TripStatus string

// DocumentKind is one of passport | ghic_card | boarding_pass.
type DocumentKind string

// Account is the ownership root. Auth attaches here later.
type Account struct {
	ID        string `json:"id"`
	Email     string `json:"email,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Person is an account-level traveller. Trips reference people by id.
type Person struct {
	ID        string       `json:"id"`
	AccountID string       `json:"account_id"`
	Name      string       `json:"name"`
	Type      PersonType   `json:"type"`
	GhicID    string       `json:"ghic_id,omitempty"`
	Dob       string       `json:"dob,omitempty"`
	Documents []Document   `json:"documents,omitempty"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
}

// Document is metadata for a stored blob (passport, GHIC card, boarding pass).
// The bytes live in the blob store under StorageKey; the DB holds only this row.
type Document struct {
	ID          string       `json:"id"`
	AccountID   string       `json:"account_id"`
	Kind        DocumentKind `json:"kind"`
	PersonID    string       `json:"person_id"`
	StageID     string       `json:"stage_id,omitempty"` // set only for boarding_pass
	FileName    string       `json:"file_name"`
	StorageKey  string       `json:"storage_key"`
	ContentType string       `json:"content_type,omitempty"`
	UploadedAt  string       `json:"uploaded_at"`
}

// Party is the set of travellers on a trip.
type Party struct {
	LeadPassenger string   `json:"lead_passenger"`
	Passengers    []string `json:"passengers"`
}

// Sharing exposes the trip via tokenised links.
type Sharing struct {
	EditToken      string `json:"-"`
	ReadToken      string `json:"-"`
	OwnerEditURL   string `json:"owner_edit_url"`
	ReadOnlyURL    string `json:"read_only_url"`
	OfflineEnabled bool   `json:"offline_enabled"`
}

// Medical is the schemaless 3rd-party assistance block on a policy.
// Persisted as the insurance.medical_json column.
type Medical struct {
	AssistID string `json:"assist_id,omitempty"`
	Phone    string `json:"phone,omitempty"`
	URL      string `json:"url,omitempty"`
}

// Insurance is a policy attached to a trip, covering one or more people.
type Insurance struct {
	ID               string   `json:"id"`
	TripID           string   `json:"-"`
	PolicyNumber     string   `json:"policy_number,omitempty"`
	EmergencyContact string   `json:"emergency_contact,omitempty"`
	AccountURL       string   `json:"account_url,omitempty"`
	Covers           []string `json:"covers"`
	Medical          Medical  `json:"medical"`
}

// Stage is one timeline item. Values is a schemaless bag whose keys depend on
// kind/subkind (see the frontend stageRegistry); persisted as values_json.
type Stage struct {
	ID        string            `json:"id"`
	TripID    string            `json:"-"`
	Kind      string            `json:"kind"`
	Subkind   string            `json:"subkind,omitempty"`
	Start     string            `json:"start"`
	SortOrder int               `json:"-"`
	Values    map[string]string `json:"values"`
}

// Trip is the full aggregate: meta + party + sharing + stages + insurance.
type Trip struct {
	ID             string      `json:"id"`
	AccountID      string      `json:"account_id"`
	Name           string      `json:"name"`
	Type           TripType    `json:"type"`
	Status         TripStatus  `json:"status"`
	Timezone       string      `json:"timezone"`
	NeedsInsurance bool        `json:"needs_insurance"`
	NeedsGhic      bool        `json:"needs_ghic"`
	Party          Party       `json:"party"`
	Sharing        Sharing     `json:"sharing"`
	Stages         []Stage     `json:"stages"`
	Insurance      []Insurance `json:"insurance,omitempty"`
	CreatedAt      string      `json:"created_at"`
	UpdatedAt      string      `json:"updated_at"`
}
