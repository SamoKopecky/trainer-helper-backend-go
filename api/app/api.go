package app

import (
	"fmt"
	"net"
	"net/http"
	"trainer-helper/api"
	"trainer-helper/api/block"
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

	_ "trainer-helper/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
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

func contextMiddleware(db *bun.DB, cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			crudTimeslot := crud.NewTimeslot(db)
			crudExerciseType := crud.NewExerciseType(db)
			crudBlock := crud.NewBlock(db)
			iam := fetcher.IAM{
				AppConfig:  cfg,
				AuthConfig: fetcher.CreateAuthConfig(cfg)}

			cc := &api.DbContext{Context: c,
				ExerciseCrud:        crud.NewExercise(db),
				TimeslotCrud:        crudTimeslot,
				WorkSetCrud:         crud.NewWorkSet(db),
				ExerciseTypeCrud:    crudExerciseType,
				BlockCrud:           crudBlock,
				WeekCrud:            crud.NewWeek(db),
				WeekDayCrud:         crud.NewWeekDay(db),
				IAMFetcher:          iam,
				TimeslotService:     service.Timeslot{Crud: crudTimeslot, Fetcher: iam},
				UserService:         service.User{Fetcher: iam},
				ExerciseTypeService: service.ExerciseType{Store: crudExerciseType},
				BlockService:        service.Block{Store: crudBlock},
			}

			return next(cc)
		}

	}
}

//	@title			Trainer Helper
//	@version		0.0.1
//	@description	Trainer helper application backend API

//	@contact.name	SamuelKopecky
//	@contact.email	samo.kopecky@protonmail.com

// @host		localhost:2001
func RunApi(db *bun.DB, appConfig *config.Config) {
	e := echo.New()
	e.HTTPErrorHandler = logError
	e.Use(contextMiddleware(db, appConfig))
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/-/ping", pong)
	e.GET("/-/metrics", pong, localhostOnlyMiddleware)
	e.GET("/-/swagger/*", echoSwagger.WrapHandler)

	jg := e.Group("")
	jg.Use(jwtMiddleware(appConfig))
	jg.Use(claimContextMiddleware)

	timeslots := jg.Group("/timeslots")
	timeslots.GET("", timeslot.Get)
	timeslots.POST("", timeslot.Post, trainerOnlyMiddleware)
	timeslots.DELETE("/:id", timeslot.Delete, trainerOnlyMiddleware)
	timeslots.PUT("", timeslot.Put, trainerOnlyMiddleware)
	// TODO: Don't use action, use action field in request param
	timeslots.POST("/undelete", timeslot.PostUndelete, trainerOnlyMiddleware)

	exercises := jg.Group("/exercises")
	// TODO: When to use :id param
	exercises.GET("/:id", exercise.Get)
	exercises.POST("", exercise.Post)
	exercises.PUT("", exercise.Put)
	exercises.DELETE("", exercise.Delete)
	exercises.PUT("/count", exercise.PutCount)
	exercises.DELETE("/count", exercise.DeleteCount)
	exercises.POST("/undelete", exercise.PostUndelete)
	exercises.POST("/duplicate", exercise.PostDuplicate, trainerOnlyMiddleware)

	workSets := jg.Group("/work-sets")
	workSets.PUT("", work_set.Put)
	workSets.POST("/undelete", work_set.PostUndelete)

	exerciseTypes := jg.Group("/exercise-types")
	exerciseTypes.GET("", exercise_type.Get)
	exerciseTypes.POST("", exercise_type.Post, trainerOnlyMiddleware)
	exerciseTypes.PUT("", exercise_type.Put, trainerOnlyMiddleware)
	exerciseTypes.POST("/duplicate", exercise_type.PostDuplicate, trainerOnlyMiddleware)

	users := jg.Group("/users")
	users.GET("", user.Get)
	users.POST("", user.Post, trainerOnlyMiddleware)
	users.DELETE("", user.Delete, trainerOnlyMiddleware)
	users.PUT("", user.Put, trainerOnlyMiddleware)

	blocks := jg.Group("/blocks")
	blocks.GET("", block.Get)

	weeks := jg.Group("/weeks")
	weeks.POST("", week.Post, trainerOnlyMiddleware)
	weeks.PUT("", week.Put, trainerOnlyMiddleware)
	weeks.DELETE("/:id", week.Delete, trainerOnlyMiddleware)

	e.Logger.Fatal(e.Start(":2001"))
}

// @Summary      Ping endpoint
// @Description  Checks if the service is responsive.
// @Produce      json
// @Success      200  {string}  string  "pong"
// @Router       /-/ping [get]
func pong(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
