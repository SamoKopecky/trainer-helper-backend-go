package service

import (
	"trainer-helper/api"
	"trainer-helper/fetcher"
)

type Person struct {
	Fetcher fetcher.IAM
}

func (p Person) GetUsers(claims *api.JwtClaims) (users []fetcher.KeycloakUser, err error) {
	role, isTrainer := claims.AppTraineeRole()
	if isTrainer {
		users, err = p.Fetcher.GetUsersByRole(role)
		if err != nil {
			return
		}
	} else {
		user, err := p.Fetcher.GetUserById(claims.Subject)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
		return users, nil
	}
	return
}
