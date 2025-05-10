package service

import (
	"errors"
	"trainer-helper/fetcher"
	"trainer-helper/schema"
)

type User struct {
	Fetcher fetcher.IAM
}

func (u User) GetUsers(claims *schema.JwtClaims) (users []fetcher.KeycloakUser, err error) {
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

func (u User) RegisterUser(email, username, traineeRole, trainerId string) (userId string, err error) {
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

	kcRole, err := u.Fetcher.GetRole(traineeRole)
	if err != nil {
		return
	}

	err = u.Fetcher.AddUserRoles(userLocation, kcRole)
	if err != nil {
		return
	}

	err = u.UpdateTrainerId(userLocation.UserId(), trainerId)
	if err != nil {
		return
	}

	return userLocation.UserId(), nil
}

func (u User) UnregisterUser(userId, traineeRole string) error {
	userLocation := u.Fetcher.GetUserLocation(userId)
	kcRole, err := u.Fetcher.GetRole(traineeRole)
	if err != nil {
		return err
	}

	err = u.Fetcher.RemoveUserRoles(userLocation, kcRole)
	if err != nil {
		return err
	}
	return nil
}

func (u User) updateAttributes(userId string, attributes fetcher.KeycloakAttributes) error {
	user, err := u.Fetcher.GetUserById(userId)
	if err != nil {
		return err
	}

	if len(attributes.Nickname) > 0 {
		user.Attributes.Nickname = attributes.Nickname
	}
	if len(attributes.TrainerId) > 0 {
		user.Attributes.TrainerId = attributes.TrainerId
	}

	err = u.Fetcher.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u User) UpdateNickname(userId, nickname string) error {
	attributes := fetcher.KeycloakAttributes{Nickname: []string{nickname}}
	err := u.updateAttributes(userId, attributes)

	if err != nil {
		return err
	}

	return nil
}

func (u User) UpdateTrainerId(userId, trainerId string) error {
	attributes := fetcher.KeycloakAttributes{TrainerId: []string{trainerId}}
	err := u.updateAttributes(userId, attributes)

	if err != nil {
		return err
	}

	return nil
}
