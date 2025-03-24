PHONY: db

db:
	pgcli "postgresql://root:alpharius@localhost/trainer_helper"

up-db:
	docker compose up db -d

down:
	docker compose down

test:
	go test ./...

reset-db:
	docker compose down
	docker compose up db -d

up:
	docker compose up
