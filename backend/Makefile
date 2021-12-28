.PHONY: migrate-up migrate-down migrate-status

DB_MIGRATIONS_DIR="./migrations"

# load .env
ifneq (,$(wildcard ./.env))
	include .env
	export
endif

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
