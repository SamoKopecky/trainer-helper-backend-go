package user

// TODO: Convert struct requests for ids to path params

import (
	"net/http"
	"trainer-helper/api"
	"trainer-helper/model"
	"trainer-helper/schemas"

	"github.com/labstack/echo/v4"
	"slices"
)

func Get(c echo.Context) (err error) {
	cc := c.(*schemas.DbContext)

	users, err := cc.UserService.GetUsers(cc.Claims)
	if err != nil {
		return err
	}
	models := make([]model.User, len(users))
	var index int
	for i, user := range users {
		if user.Id == cc.Claims.Subject {
			index = i
			continue
		}
		models[i] = user.ToUserModel()
	}

	// Delete trainer user
	return cc.JSON(http.StatusOK, slices.Delete(models, index, index+1))
}

func Post(c echo.Context) (err error) {
	cc := c.(*schemas.DbContext)
	traineeRole, _ := cc.Claims.AppTraineeRole()

	params, err := api.BindParams[userPostRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	userId, err := cc.UserService.RegisterUser(params.Email, params.Username, traineeRole, cc.Claims.Subject)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusCreated, struct {
		UserId string `json:"user_id"`
	}{UserId: userId})
}

func Delete(c echo.Context) (err error) {
	cc := c.(*schemas.DbContext)

	params, err := api.BindParams[userDeleteRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}
	traineeRole, _ := cc.Claims.AppTraineeRole()

	err = cc.UserService.UnregisterUser(params.Id, traineeRole)
	if err != nil {
		return
	}
	return cc.NoContent(http.StatusOK)
}

func Put(c echo.Context) (err error) {
	cc := c.(*schemas.DbContext)

	params, err := api.BindParams[userPutRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	err = cc.UserService.UpdateNickname(params.Id, params.Nickname)
	if err != nil {
		return
	}
	return cc.NoContent(http.StatusOK)
}
