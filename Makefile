.DEFAULT_GOAL := help

# Go backend lives in backend/; all targets run there.
BACKEND := backend

# Resolve air from PATH, else fall back to the Go install dir (GOPATH/bin).
AIR := $(shell command -v air 2>/dev/null || echo "$(shell go env GOPATH)/bin/air")

.PHONY: help dev run build clean tidy tools seed \
        migrate-up migrate-down migrate-status migrate-reset migrate-create

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-16s\033[0m %s\n", $$1, $$2}'

dev: ## Run server with hot reload (requires air; `make tools`)
	cd $(BACKEND) && $(AIR)

run: ## Run the server once
	cd $(BACKEND) && go run ./switchboard/server

build: ## Build the server binary into backend/bin/
	cd $(BACKEND) && go build -o bin/server ./switchboard/server

clean: ## Remove build artifacts and the local db
	cd $(BACKEND) && rm -rf bin tmp *.db

tidy: ## Sync go.mod / go.sum
	cd $(BACKEND) && go mod tidy

tools: ## Install dev tools (air)
	cd $(BACKEND) && go install github.com/air-verse/air@latest

seed: ## Load mock people + the demo trip into the database
	cd $(BACKEND) && go run ./switchboard/cli seed

migrate-up: ## Apply all pending migrations
	cd $(BACKEND) && go run ./switchboard/cli up

migrate-down: ## Roll back the most recent migration
	cd $(BACKEND) && go run ./switchboard/cli down

migrate-status: ## Show migration status
	cd $(BACKEND) && go run ./switchboard/cli status

migrate-reset: ## Roll back all migrations
	cd $(BACKEND) && go run ./switchboard/cli reset

migrate-create: ## Create a migration: make migrate-create name=add_users
	cd $(BACKEND) && go run ./switchboard/cli create $(name) sql
