package service

import (
	"trainer-helper/api"
	"trainer-helper/fetcher"
)

type User struct {
	Fetcher fetcher.IAM
}

func (u User) GetUsers(claims *api.JwtClaims) (users []fetcher.KeycloakUser, err error) {
	role, isTrainer := claims.AppTraineeRole()
	if isTrainer {
		users, err = u.Fetcher.GetUsersByRole(role)
		if err != nil {
			return
		}
	} else {
		user, err := u.Fetcher.GetUserById(claims.Subject)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
		return users, nil
	}
	return
}

func (u User) RegisterUser(email, username string) error {
	userIdUrl, err := u.Fetcher.CreateUser(email, username)
	if err != nil {
		return err
	}
	err = u.Fetcher.InvokeUserUpdate(userIdUrl)
	if err != nil {
		return err
	}

	return nil
}
