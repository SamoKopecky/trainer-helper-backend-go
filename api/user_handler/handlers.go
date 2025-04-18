package user_handler

import (
	"net/http"
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
