package person_handler

import (
	"log"
	"net/http"
	"trainer-helper/api"
	"trainer-helper/crud"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)

	crud := crud.CRUDPerson{Db: cc.Db}
	persons, err := crud.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	return cc.JSON(http.StatusOK, persons)
}
