.PHONY: deps dev pgcli migrate-up migrate-down migrate-status

APP_NAME="auth-service"
DB_MIGRATIONS_DIR="./migrations"

ifneq (,$(wildcard ./.env))
	include .env
	export
endif

deps:
	@cargo install cargo-watch

pgcli:
	pgcli $(DATABASE_URL)

dev:
	cargo watch -x 'run --bin $(APP_NAME)'

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
