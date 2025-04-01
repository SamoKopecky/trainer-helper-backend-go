.PHONY: help

# Help
help:
	@awk '/^#/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print substr($$1,1,index($$1,":")),c}1{c=0}' $(MAKEFILE_LIST) | column -s: -t

# Connect to db
pgcli:
	pgcli "postgresql://trainer_helper:alpharius@localhost/trainer_helper"

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
