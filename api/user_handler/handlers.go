package user_handler

import (
	"net/http"
	"trainer-helper/api"
	"trainer-helper/fetcher"
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

	params, err := api.BindParams[userPostRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	err = cc.UserService.RegisterUser(params.Email, params.Username)
	if err == fetcher.ErrUserAlreadyExists {
		return cc.NoContent(http.StatusConflict)
	}
	if err != nil {
		return err
	}
	return cc.NoContent(http.StatusCreated)
}
