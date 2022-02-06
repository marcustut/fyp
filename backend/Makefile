.PHONY: deps migrate-up migrate-down migrate-status

DB_MIGRATIONS_DIR="./migrations"

# load .env
ifneq (,$(wildcard ./.env))
	include .env
	export
endif

deps:
	brew install yq
	brew install pre-commit
	pre-commit --version
	pre-commit install
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/segmentio/golines@latest
	go install github.com/cosmtrek/air@latest

up:
	docker-compose up -d

# Create a new migration
migrate-new:
	dbmate -u $(DATABASE_URL) -d $(DB_MIGRATIONS_DIR) --no-dump-schema new $(NAME)

# Apply latest migration
migrate-up:
	dbmate -u $(DATABASE_URL) -d $(DB_MIGRATIONS_DIR) --no-dump-schema up

# Rollback latest migration
migrate-down:
	dbmate -u $(DATABASE_URL) -d $(DB_MIGRATIONS_DIR) --no-dump-schema down

# Check migration status
migrate-status:
	dbmate -u $(DATABASE_URL) -d $(DB_MIGRATIONS_DIR) --no-dump-schema status

# Migrate scheme in ent to database
migrate_schema:
	go run ./cmd/migration/main.go

# Generate Ent and Gqlgen
generate:
	go generate ./...

# Start dev server
start:
	air