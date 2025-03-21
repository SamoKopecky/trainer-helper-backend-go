PHONY: run

db:
	pgcli "postgresql://root:alpharius@localhost/trainer_helper"

up-db:
	docker compose up db -d


test:
	go test ./...
