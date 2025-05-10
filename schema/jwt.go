package schema

import (
	"regexp"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

const (
	trainerPostfix = "trainer_app"
	traineePostfix = "trainee_app"
	trainerRegex   = `.*` + trainerPostfix + `.*`
	traineeRegex   = `.*` + traineePostfix + `.*`
	rolesKey       = "roles"
)

type JwtClaims struct {
	RealmAccess map[string][]string `json:"realm_access"`
	jwt.RegisteredClaims
}

func (jcc JwtClaims) IsTrainer() bool {
	if roles, ok := jcc.RealmAccess[rolesKey]; ok {
		for _, role := range roles {
			if matched, _ := regexp.MatchString(trainerRegex, role); matched {
				return true
			}
		}
	}
	return false
}

func (jcc JwtClaims) AppRole() (role string, isTrainer bool) {
	isTrainer = false
	trainerCompiled, _ := regexp.Compile(trainerRegex)
	traineeCompiled, _ := regexp.Compile(traineeRegex)

	if roles, ok := jcc.RealmAccess[rolesKey]; ok {
		for _, r := range roles {
			if match := trainerCompiled.FindString(r); match != "" {
				role = match
				isTrainer = true
				// If trainer found, return imidiedatly
				return
			}

			if match := traineeCompiled.FindString(r); match != "" {
				role = match
				// If trainee found, wait until possible trainer is found
			}
		}
	}
	return
}

func (jcc JwtClaims) AppTraineeRole() (role string, isTrainer bool) {
	role, isTrainer = jcc.AppRole()
	if isTrainer {
		role = strings.Replace(role, trainerPostfix, traineePostfix, -1)
		return
	}
	return
}
