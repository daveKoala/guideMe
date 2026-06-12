-- +goose Up
-- +goose StatementBegin
CREATE TABLE accounts (
    id         TEXT PRIMARY KEY,
    email      TEXT UNIQUE,
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at TEXT NOT NULL DEFAULT (datetime('now'))
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE people (
    id         TEXT PRIMARY KEY,
    account_id TEXT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    name       TEXT NOT NULL,
    type       TEXT NOT NULL CHECK (type IN ('adult', 'child', 'infant')),
    ghic_id    TEXT,
    dob        TEXT,
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at TEXT NOT NULL DEFAULT (datetime('now'))
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE trips (
    id               TEXT PRIMARY KEY,
    account_id       TEXT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    name             TEXT NOT NULL,
    type             TEXT NOT NULL CHECK (type IN ('outbound', 'return')),
    status           TEXT NOT NULL CHECK (status IN ('planned', 'booked', 'completed', 'cancelled')),
    timezone         TEXT NOT NULL,
    needs_insurance  INTEGER NOT NULL DEFAULT 0,
    needs_ghic       INTEGER NOT NULL DEFAULT 0,
    lead_passenger_id TEXT REFERENCES people(id) ON DELETE SET NULL,
    edit_token       TEXT NOT NULL,
    read_token       TEXT NOT NULL,
    offline_enabled  INTEGER NOT NULL DEFAULT 1,
    created_at       TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at       TEXT NOT NULL DEFAULT (datetime('now'))
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE trip_passengers (
    trip_id   TEXT NOT NULL REFERENCES trips(id) ON DELETE CASCADE,
    person_id TEXT NOT NULL REFERENCES people(id) ON DELETE CASCADE,
    PRIMARY KEY (trip_id, person_id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE stages (
    id          TEXT PRIMARY KEY,
    trip_id     TEXT NOT NULL REFERENCES trips(id) ON DELETE CASCADE,
    kind        TEXT NOT NULL CHECK (kind IN ('flight', 'travel_to_airport', 'accommodation', 'note')),
    subkind     TEXT,
    start       TEXT NOT NULL DEFAULT '',
    sort_order  INTEGER NOT NULL DEFAULT 0,
    values_json TEXT NOT NULL DEFAULT '{}',
    created_at  TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at  TEXT NOT NULL DEFAULT (datetime('now'))
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE insurance (
    id                TEXT PRIMARY KEY,
    trip_id           TEXT NOT NULL REFERENCES trips(id) ON DELETE CASCADE,
    policy_number     TEXT,
    emergency_contact TEXT,
    account_url       TEXT,
    medical_json      TEXT NOT NULL DEFAULT '{}',
    created_at        TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at        TEXT NOT NULL DEFAULT (datetime('now'))
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE insurance_covers (
    insurance_id TEXT NOT NULL REFERENCES insurance(id) ON DELETE CASCADE,
    person_id    TEXT NOT NULL REFERENCES people(id) ON DELETE CASCADE,
    PRIMARY KEY (insurance_id, person_id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE documents (
    id           TEXT PRIMARY KEY,
    account_id   TEXT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    kind         TEXT NOT NULL CHECK (kind IN ('passport', 'ghic_card', 'boarding_pass')),
    person_id    TEXT NOT NULL REFERENCES people(id) ON DELETE CASCADE,
    stage_id     TEXT REFERENCES stages(id) ON DELETE CASCADE,
    file_name    TEXT NOT NULL,
    storage_key  TEXT NOT NULL,
    content_type TEXT,
    uploaded_at  TEXT NOT NULL DEFAULT (datetime('now'))
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_people_account     ON people(account_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX idx_trips_account      ON trips(account_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX idx_stages_trip        ON stages(trip_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX idx_insurance_trip     ON insurance(trip_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX idx_documents_person   ON documents(person_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX idx_documents_stage    ON documents(stage_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS documents;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS insurance_covers;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS insurance;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS stages;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS trip_passengers;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS trips;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS people;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS accounts;
-- +goose StatementEnd
