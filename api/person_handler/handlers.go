package person_handler

import (
	"net/http"
	"trainer-helper/api"
	"trainer-helper/model"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)

	persons, err := cc.CRUDPerson.GetAll()
	if err != nil {
		return err
	}

	if len(persons) == 0 {
		persons = []model.Person{}
	}

	return cc.JSON(http.StatusOK, persons)
}
