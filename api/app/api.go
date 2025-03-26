package app

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"net/http"
	"trainer-helper/api"
	"trainer-helper/api/exercise_handler"
	exercise_count_handler "trainer-helper/api/exercise_handler/count_handler"
	exercise_duplicate_handler "trainer-helper/api/exercise_handler/duplicate_handler"
	"trainer-helper/api/person_handler"
	"trainer-helper/api/timeslot_handler"
	timeslot_revert_handler "trainer-helper/api/timeslot_handler/revert_handler"
	"trainer-helper/api/work_set_handler"
	"trainer-helper/config"
	"trainer-helper/crud"
	"trainer-helper/fetcher"
	"trainer-helper/service"
	"trainer-helper/utils"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/uptrace/bun"
)

const (
	ColorReset   = "\033[0m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
)

func logError(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	code := http.StatusInternalServerError
	errMsg := "Internal server error"
	he, ok := err.(*echo.HTTPError)
	if ok {
		code = he.Code
		errMsg = fmt.Sprint(he.Message)
	}

	if err := c.JSON(code, map[string]string{"message": errMsg}); err != nil {
		c.Logger().Error(err)
	}
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
}

func getKey(token *jwt.Token) (any, error) {
	// TODO: Use cache, use config
	keySet, err := jwk.Fetch(context.Background(), "http://localhost:8080/realms/trainer-helper/protocol/openid-connect/certs")
	if err != nil {
		return rsa.PublicKey{}, err
	}

	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return rsa.PublicKey{}, errors.New("expecting JWT header to have a key ID in the kid field")
	}

	key, found := keySet.LookupKeyID(keyID)
	if !found {
		return rsa.PublicKey{}, fmt.Errorf("unable to find key %q", keyID)
	}
	var pubkey rsa.PublicKey
	if err := key.Raw(&pubkey); err != nil {
		return rsa.PublicKey{}, fmt.Errorf("Unable to get the public key. Error: %s", err.Error())
	}
	utils.PrettyPrint(pubkey)
	return &pubkey, nil
}

func RunApi(db *bun.DB, appConfig config.Config) {
	e := echo.New()
	e.HTTPErrorHandler = logError
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			crudTimeslot := crud.NewTimeslot(db)
			iam := fetcher.IAM{
				AppConfig:  appConfig,
				AuthConfig: fetcher.CreateAuthConfig(appConfig)}

			cc := &api.DbContext{Context: c,
				ExerciseCrud:    crud.NewExercise(db),
				TimeslotCrud:    crudTimeslot,
				WorkSetCrud:     crud.NewWorkSet(db),
				IAMFetcher:      iam,
				TimeslotService: service.Timeslot{Crud: crudTimeslot, Fetcher: iam},
			}

			return next(cc)
		}
	})
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/-/ping", pong)
	g := e.Group("")
	g.Use(echojwt.WithConfig(echojwt.Config{
		KeyFunc:       getKey,
		SigningMethod: "RS256",
	}))
	g.GET("/timeslot", timeslot_handler.Get)
	g.POST("/timeslot", timeslot_handler.Post)
	g.DELETE("/timeslot", timeslot_handler.Delete)
	g.PUT("/timeslot", timeslot_handler.Put)
	g.PUT("/timeslot/revert", timeslot_revert_handler.Put)
	g.GET("/exercise/:id", exercise_handler.Get)
	g.PUT("/exercise", exercise_handler.Put)
	g.DELETE("/exercise", exercise_handler.Delete)
	g.POST("/exercise", exercise_handler.Post)
	g.PUT("/exercise/count", exercise_count_handler.Put)
	g.DELETE("/exercise/count", exercise_count_handler.Delete)
	g.POST("/exercise/duplicate", exercise_duplicate_handler.Post)
	g.PUT("/workset", work_set_handler.Put)
	g.GET("/person", person_handler.Get)

	e.Logger.Fatal(e.Start(":2001"))
}

func pong(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
