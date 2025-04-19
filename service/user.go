package service

import (
	"errors"
	"fmt"
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

func (u User) RegisterUser(email, username, traineeRole string) (userId string, err error) {
	userLocation, err := u.Fetcher.CreateUser(email, username)
	if errors.Is(err, fetcher.ErrUserAlreadyExists) {
		// If use is already created, get user id by email
		userLocation, err = u.Fetcher.GetUserLocationByEmail(email)
		if err != nil {
			return
		}
	} else if err == nil {
		err = u.Fetcher.InvokeUserUpdate(userLocation)
		if err != nil {
			return
		}
	} else {
		return
	}

	fmt.Println(userLocation)
	kcRole, err := u.Fetcher.GetRole(traineeRole)
	if err != nil {
		return
	}

	err = u.Fetcher.AddUserRoles(userLocation, kcRole)
	if err != nil {
		return
	}

	return userLocation.UserId(), nil
}
