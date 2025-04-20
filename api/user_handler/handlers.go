package user_handler

import (
	"net/http"
	"trainer-helper/api"
	"trainer-helper/model"
	"trainer-helper/schemas"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) (err error) {
	cc := c.(*schemas.DbContext)

	users, err := cc.UserService.GetUsers(cc.Claims)
	if err != nil {
		return err
	}
	models := make([]model.User, len(users))
	for i, user := range users {
		models[i] = user.ToUserModel()
	}

	return cc.JSON(http.StatusOK, models)
}

func Post(c echo.Context) (err error) {
	cc := c.(*schemas.DbContext)
	traineeRole, _ := cc.Claims.AppTraineeRole()

	params, err := api.BindParams[userPostRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	userId, err := cc.UserService.RegisterUser(params.Email, params.Username, traineeRole)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusCreated, struct {
		UserId string `json:"user_id"`
	}{UserId: userId})
}

// func Put(c echo.Context) (err error) {
//
// }
