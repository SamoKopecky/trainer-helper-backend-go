package app

import (
	"fmt"
	"net"
	"net/http"
	"trainer-helper/api"
	"trainer-helper/api/exercise"
	"trainer-helper/api/exercise_type"
	"trainer-helper/api/timeslot"
	"trainer-helper/api/user"
	"trainer-helper/api/week"
	"trainer-helper/api/work_set"
	"trainer-helper/config"
	"trainer-helper/crud"
	"trainer-helper/fetcher"
	"trainer-helper/schema"
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
		cc := c.(*api.DbContext)
		cc.Claims = cc.Get("user").(*jwt.Token).Claims.(*schema.JwtClaims)
		return next(c)
	}
}

func trainerOnlyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*api.DbContext)
		if !cc.Claims.IsTrainer() {
			return cc.NoContent(http.StatusForbidden)
		}

		return next(c)
	}
}

func localhostOnlyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		addr := c.Request().RemoteAddr
		host, _, err := net.SplitHostPort(addr)
		if err != nil {
			return c.NoContent(http.StatusForbidden)
		}

		ip := net.ParseIP(host)
		if ip == nil {
			c.Logger().Warnf("Cant parse ip", host)
			return c.NoContent(http.StatusForbidden)
		}

		if !ip.IsLoopback() {
			c.Logger().Warnf("Forbidden access to metrics from IP", ip)
			return c.NoContent(http.StatusForbidden)
		}

		return next(c)
	}
}

func jwtMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	keyFunc := func(token *jwt.Token) (any, error) {
		return getKey(cfg, token)
	}
	return echojwt.WithConfig(echojwt.Config{
		KeyFunc:       keyFunc,
		SigningMethod: "RS256",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(schema.JwtClaims)
		},
	})
}

func RunApi(db *bun.DB, appConfig *config.Config) {
	e := echo.New()
	e.HTTPErrorHandler = logError
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			crudTimeslot := crud.NewTimeslot(db)
			crudExerciseType := crud.NewExerciseType(db)
			crudWeek := crud.NewWeek(db)
			iam := fetcher.IAM{
				AppConfig:  appConfig,
				AuthConfig: fetcher.CreateAuthConfig(appConfig)}

			cc := &api.DbContext{Context: c,
				ExerciseCrud:        crud.NewExercise(db),
				TimeslotCrud:        crudTimeslot,
				WorkSetCrud:         crud.NewWorkSet(db),
				ExerciseTypeCrud:    crudExerciseType,
				WeekCrud:            crudWeek,
				IAMFetcher:          iam,
				TimeslotService:     service.Timeslot{Crud: crudTimeslot, Fetcher: iam},
				UserService:         service.User{Fetcher: iam},
				ExerciseTypeService: service.ExerciseType{Store: crudExerciseType},
				WeekService:         service.Week{Store: crudWeek},
			}

			return next(cc)
		}
	})
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/-/ping", pong)
	e.GET("/-/metrics", pong, localhostOnlyMiddleware)

	jg := e.Group("")
	jg.Use(jwtMiddleware(appConfig))
	jg.Use(claimContextMiddleware)
	jg.GET("/timeslot", timeslot.Get)
	jg.GET("/exercise/:id", exercise.Get)
	jg.PUT("/exercise", exercise.Put)
	jg.DELETE("/exercise", exercise.Delete)
	jg.POST("/exercise", exercise.Post)
	jg.PUT("/exercise/count", exercise.PutCount)
	jg.DELETE("/exercise/count", exercise.DeleteCount)
	jg.PUT("/workset", work_set.Put)
	jg.GET("/user", user.Get)
	jg.GET("/exerciseType", exercise_type.Get)
	jg.POST("/workset/undelete", work_set.PostUndelete)
	jg.POST("/exercise/undelete", exercise.PostUndelete)
	jg.GET("/week", week.Get)

	to := jg.Group("")
	to.Use(trainerOnlyMiddleware)
	to.DELETE("/timeslot", timeslot.Delete)
	to.POST("/timeslot", timeslot.Post)
	to.PUT("/timeslot", timeslot.Put)
	to.POST("/timeslot/undelete", timeslot.PostUndelete)
	to.POST("/exercise/duplicate", exercise.PostDuplicate)
	to.POST("/exerciseType", exercise_type.Post)
	to.PUT("/exerciseType", exercise_type.Put)
	to.POST("/exerciseType/duplicate", exercise_type.PostDuplicate)
	to.POST("/user", user.Post)
	to.DELETE("/user", user.Delete)
	to.PUT("/user", user.Put)

	e.Logger.Fatal(e.Start(":2001"))
}

func pong(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
