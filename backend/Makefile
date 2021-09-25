.PHONY: deps dev

APP_NAME="backend"

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
