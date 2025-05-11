package user

// TODO: Convert struct requests for ids to path params

import (
	"net/http"
	"trainer-helper/api"
	"trainer-helper/model"

	"slices"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) (err error) {
	cc := c.(*api.DbContext)

	isTrainer := cc.Claims.IsTrainer()
	users, err := cc.UserService.GetUsers(cc.Claims)
	if err != nil {
		return err
	}
	userModels := make([]model.User, len(users))
	var index int
	for i, user := range users {
		if user.Id == cc.Claims.Subject && isTrainer {
			index = i
			continue
		}
		userModels[i] = user.ToUserModel()
	}

	if isTrainer {
		// Delete trainer user
		userModels = slices.Delete(userModels, index, index+1)
	}

	return cc.JSON(http.StatusOK, userModels)
}

func Post(c echo.Context) (err error) {
	cc := c.(*api.DbContext)
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
	cc := c.(*api.DbContext)

	id := cc.Param("id")
	traineeRole, _ := cc.Claims.AppTraineeRole()

	err = cc.UserService.UnregisterUser(id, traineeRole)
	if err != nil {
		return
	}
	return cc.NoContent(http.StatusOK)
}

func Put(c echo.Context) (err error) {
	cc := c.(*api.DbContext)

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
