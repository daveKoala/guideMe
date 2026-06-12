// Command cli is the project's database-migration tool. It wraps Goose as a
// library so no global goose binary is required.
//
// Usage:
//
//	go run ./switchboard/cli <command> [args]
//
// Commands: up, down, status, reset, create <name> [sql|go], version
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/david-clare/guide-me/backend/config"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

// migrationsDir holds the goose migration files, relative to the backend root.
const migrationsDir = "migrations"

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	command := os.Args[1]
	args := os.Args[2:]

	cfg := config.Load()

	// seed is our own command, not a goose migration.
	if command == "seed" {
		if err := runSeed(cfg.DBPath); err != nil {
			log.Fatalf("seed: %v", err)
		}
		return
	}

	db, err := sql.Open("sqlite", cfg.DBPath)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatalf("set dialect: %v", err)
	}

	ctx := context.Background()
	if err := goose.RunContext(ctx, command, db, migrationsDir, args...); err != nil {
		log.Fatalf("goose %s: %v", command, err)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, `migration cli — wraps goose

usage: go run ./switchboard/cli <command> [args]

commands:
  up                   migrate the database to the most recent version
  up-by-one            migrate up a single version
  down                 roll back the version by 1
  reset                roll back all migrations
  status               dump the migration status for the database
  version              print the current version of the database
  create <name> sql    create a new timestamped sql migration
  seed                 load mock people + the demo trip into the database`)
}
