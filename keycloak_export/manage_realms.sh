#/bin/bash

ss -tuln | grep ':8080' >/dev/null
if [ $? -eq 0 ]; then
	echo "Keycloak is running on :8080, turn it off"
	exit 1
fi

realms="master trainer-helper"
action=$1

if [ "$action" = "import" ]; then
	for realm in $realms; do
		echo "----------------"
		echo "Importing $realm"
		echo "----------------"
		docker run --rm \
			--name keycloak_export \
			-v ./keycloak_data:/opt/keycloak/data/h2 \
			-v ./keycloak_export:/tmp/keycloak-export \
			quay.io/keycloak/keycloak:latest \
			import --file "/tmp/keycloak-export/export_$realm.json" \
			--verbose \
			--override true
	done
elif [ "$action" = "export" ]; then
	for realm in $realms; do
		echo "----------------"
		echo "Exporting $realm"
		echo "----------------"
		docker run --rm \
			--name keycloak_export \
			-v ./keycloak_data:/opt/keycloak/data/h2 \
			-v ./keycloak_export:/tmp/keycloak-export \
			quay.io/keycloak/keycloak:latest \
			export --file "/tmp/keycloak-export/export_$realm.json" \
			--verbose \
			--realm $realm
	done

else
	echo "unknown command"
	exit 1
fi
