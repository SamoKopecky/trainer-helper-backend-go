PHONY: db

db:
	pgcli "postgresql://root:alpharius@localhost/trainer_helper"

up-db:
	docker compose up db -d

up:
	docker compose up

down:
	docker compose down

test:
	go test ./...

reset-db:
	docker compose down
	docker compose up db -d

kc-purge:
	rm ./keycloak_data/*

kc-import:
	./keycloak_export/manage_realms.sh import

kc-export:
	./keycloak_export/manage_realms.sh export

run:
	APP_KC_ADMIN_CLIENT_SECRET="0F32CR8bzQAMgLCYAR6pa2HbksVViCMc" air -- --debug
