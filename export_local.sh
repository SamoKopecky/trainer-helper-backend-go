#/bin/bash

docker run --rm --name keycloak_export -v ./keycloak_data:/opt/keycloak/data/h2 -v ./keycloak_export:/tmp/keycloak-export quay.io/keycloak/keycloak:latest export --file /tmp/keycloak-export/export.json --verbose --realm $1
