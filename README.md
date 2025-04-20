# Trainer Helper Backend

Backend for the Trainer Helper app written in go.

See frontend repo [here](https://github.com/SamoKopecky/trainer-helper-frontend/).

# Running

Use `make` to run all useful commands.

# Keycloak config

- Add `nickname` attribute in `realm settings -> user profile`
- Add `trainerId` attribute in `realm settings -> user profile`, then add a custom mapper in `clients->trainer_helper->client scopes->trainer-helper-dedicated`
- Use service acount auth for admin realm
- configure email in trainer helper realm
