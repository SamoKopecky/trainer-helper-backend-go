package person_handler

import (
	"log"
	"net/http"
	"trainer-helper/api"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)

	persons, err := cc.CRUDPerson.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	return cc.JSON(http.StatusOK, persons)
}
