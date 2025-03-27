package app

import (
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
	"trainer-helper/schemas"
	"trainer-helper/service"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
)

func logError(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	code := http.StatusInternalServerError
	errMsg := "Internal server error"
	httpError, ok := err.(*echo.HTTPError)
	if ok {
		code = httpError.Code
		errMsg = fmt.Sprint(httpError.Message)
	}

	if err := c.JSON(code, map[string]string{"message": errMsg}); err != nil {
		c.Logger().Error(err)
	}
}

func claimContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*schemas.DbContext)
		cc.Claims = cc.Get("user").(*jwt.Token).Claims.(*api.JwtClaims)
		return next(c)
	}
}

func trainerOnlyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*schemas.DbContext)
		if !cc.Claims.IsTrainer() {
			return cc.NoContent(http.StatusForbidden)
		}

		return next(c)
	}

}

func jwtMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		KeyFunc:       getKey,
		SigningMethod: "RS256",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(api.JwtClaims)
		},
	})
}

func RunApi(db *bun.DB, appConfig *config.Config) {
	e := echo.New()
	e.HTTPErrorHandler = logError
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			crudTimeslot := crud.NewTimeslot(db)
			iam := fetcher.IAM{
				AppConfig:  appConfig,
				AuthConfig: fetcher.CreateAuthConfig(appConfig)}

			cc := &schemas.DbContext{Context: c,
				ExerciseCrud:    crud.NewExercise(db),
				TimeslotCrud:    crudTimeslot,
				WorkSetCrud:     crud.NewWorkSet(db),
				IAMFetcher:      iam,
				TimeslotService: service.Timeslot{Crud: crudTimeslot, Fetcher: iam},
				PersonService:   service.Person{Fetcher: iam},
			}

			return next(cc)
		}
	})
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/-/ping", pong)

	jg := e.Group("")
	jg.Use(jwtMiddleware())
	jg.Use(claimContextMiddleware)
	jg.GET("/timeslot", timeslot_handler.Get)
	jg.GET("/exercise/:id", exercise_handler.Get)
	jg.PUT("/exercise", exercise_handler.Put)
	jg.DELETE("/exercise", exercise_handler.Delete)
	jg.POST("/exercise", exercise_handler.Post)
	jg.PUT("/exercise/count", exercise_count_handler.Put)
	jg.DELETE("/exercise/count", exercise_count_handler.Delete)
	jg.PUT("/workset", work_set_handler.Put)
	jg.GET("/person", person_handler.Get)

	to := jg.Group("")
	to.Use(trainerOnlyMiddleware)
	to.DELETE("/timeslot", timeslot_handler.Delete)
	to.POST("/timeslot", timeslot_handler.Post)
	to.PUT("/timeslot", timeslot_handler.Put)
	to.PUT("/timeslot/revert", timeslot_revert_handler.Put)
	to.POST("/exercise/duplicate", exercise_duplicate_handler.Post)

	e.Logger.Fatal(e.Start(":2001"))
}

func pong(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
