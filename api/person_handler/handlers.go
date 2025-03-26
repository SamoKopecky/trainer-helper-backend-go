package person_handler

import (
	"net/http"
	"trainer-helper/api"
	"trainer-helper/model"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) (err error) {
	cc := c.(*api.DbContext)

	users, err := cc.IAMFetcher.GetUsers()
	if err != nil {
		return err
	}

	models := make([]model.Person, len(users))
	for i, user := range users {
		models[i] = user.ToPersonModel()
	}

	return cc.JSON(http.StatusOK, models)
}
