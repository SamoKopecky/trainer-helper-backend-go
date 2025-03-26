package api

import (
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

const (
	trainerPostfix = "trainer_app"
	traineePostfix = "trainee_app"
	rolesKey       = "roles"
)

func (jcc JwtClaims) GetAppRole() (string, bool) {
	var trainerRole string
	var traineeRole string
	for k, v := range jcc.RealmAccess {
		if k == rolesKey {
			for _, role := range v {
				if strings.Contains(role, trainerPostfix) {
					trainerRole = role
				} else if strings.Contains(role, traineePostfix) {
					traineeRole = role
				}
			}
			break
		}
	}

	if trainerRole != "" {
		return strings.Replace(trainerRole, "trainer", "trainee", -1), true
	}
	return traineeRole, false
}

type JwtClaims struct {
	RealmAccess map[string][]string `json:"realm_access"`
	jwt.RegisteredClaims
}
