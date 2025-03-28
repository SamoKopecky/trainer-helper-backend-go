package person_handler

import (
	"net/http"
	"trainer-helper/model"
	"trainer-helper/schemas"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) (err error) {
	cc := c.(*schemas.DbContext)

	users, err := cc.PersonService.GetUsers(cc.Claims)
	if err != nil {
		return err
	}
	models := make([]model.Person, len(users))
	for i, user := range users {
		models[i] = user.ToPersonModel()
	}

	return cc.JSON(http.StatusOK, models)
}
