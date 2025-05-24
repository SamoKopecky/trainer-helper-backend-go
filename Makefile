.PHONY: help

TA ?= ""
MSG ?= "replace_me"

# Help
help:
	@awk '/^#/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print substr($$1,1,index($$1,":")),c}1{c=0}' $(MAKEFILE_LIST) | column -s: -t

# Connect to db
pgcli:
	pgcli "postgresql://trainer_helper:alpharius@localhost/trainer_helper"

# Start services in the backgound
up-d:
	docker compose up -d

up:
	docker compose up

# Stop and delete all docker containers
down:
	docker compose down

# Stop db only
db-refresh:
	docker compose down db
	docker compose up db -d

# Run tests
test:
	docker compose up db -d
	gotest $(TA) ./...

reset-db:
	docker compose down
	docker compose up db -d

# Purge all keycloak configuration & data
kc-purge:
	rm ./keycloak_data/*

# Import dev keycloak configuration & data
kc-import:
	./keycloak_export/manage_realms.sh import

# Export current keycloak configuration & data to dev
kc-export:
	./keycloak_export/manage_realms.sh export


# Run app in dev mode
run:
	APP_ENV="dev" APP_KC_ADMIN_CLIENT_SECRET="0F32CR8bzQAMgLCYAR6pa2HbksVViCMc" air -- --debug

# Add a new migration
add-migration:
	migrate create -dir migrations/ -ext sql $(MSG)

# Generate swagger docs
swagger-docs:
	swag init -g api/app/api.go
