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
	weekday "trainer-helper/api/week_day"
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
			crudWeek := crud.NewWeek(db)
			crudWeekDay := crud.NewWeekDay(db)
			crudExercise := crud.NewExercise(db)
			crudWorkSet := crud.NewWorkSet(db)

			iam := fetcher.IAM{
				AppConfig:  cfg,
				AuthConfig: fetcher.CreateAuthConfig(cfg)}

			cc := &api.DbContext{Context: c,
				ExerciseCrud:        crudExercise,
				TimeslotCrud:        crudTimeslot,
				WorkSetCrud:         crudWorkSet,
				ExerciseTypeCrud:    crudExerciseType,
				BlockCrud:           crudBlock,
				WeekCrud:            crudWeek,
				WeekDayCrud:         crudWeekDay,
				IAMFetcher:          iam,
				TimeslotService:     service.Timeslot{TimeslotCrud: crudTimeslot, WeekDayCrud: crudWeekDay, Fetcher: iam},
				UserService:         service.User{Fetcher: iam},
				ExerciseTypeService: service.ExerciseType{Store: crudExerciseType},
				BlockService:        service.Block{Store: crudBlock},
				WeekService:         service.Week{WeekStore: crudWeek, WeekDayStore: crudWeekDay, ExerciseStore: crudExercise, WorkSetStore: crudWorkSet},
				ExerciseService:     service.Exercise{Store: crudExercise},
				AIService: service.AI{
					Fetcher:           fetcher.AI{AppConfig: cfg},
					ExerciseTypeStore: crudExerciseType,
					ExerciseStore:     crudExercise,
					WorkSetStore:      crudWorkSet},

				Config: cfg,
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
	timeslots.GET("", timeslot.GetMany)
	timeslots.POST("", timeslot.Post, trainerOnlyMiddleware)
	timeslots.DELETE("/:id", timeslot.Delete, trainerOnlyMiddleware)
	timeslots.PUT("", timeslot.Put, trainerOnlyMiddleware)
	// TODO: Don't use action, use action field in request param
	timeslots.POST("/undelete/:id", timeslot.PostUndelete, trainerOnlyMiddleware)
	timeslots.GET("/detailed", timeslot.GetManyEnhanced)

	exercises := jg.Group("/exercises")
	exercises.GET("", exercise.GetMany)
	exercises.POST("", exercise.Post)
	exercises.PUT("", exercise.Put)
	exercises.DELETE("/:id", exercise.Delete)
	exercises.PUT("/count", exercise.PutCount)
	exercises.DELETE("/count", exercise.DeleteCount)
	exercises.POST("/undelete/:id", exercise.PostUndelete)

	workSets := jg.Group("/work-sets")
	workSets.PUT("", work_set.Put)
	workSets.POST("/undelete", work_set.PostUndelete)

	exerciseTypes := jg.Group("/exercise-types")
	exerciseTypes.GET("", exercise_type.GetMany)
	exerciseTypes.POST("", exercise_type.Post, trainerOnlyMiddleware)
	exerciseTypes.PUT("", exercise_type.Put, trainerOnlyMiddleware)
	exerciseTypes.POST("/duplicate", exercise_type.PostDuplicate, trainerOnlyMiddleware)
	exerciseTypes.GET("/:id/files", exercise_type.GetMedia)
	exerciseTypes.POST("/:id/files", exercise_type.PostMedia, trainerOnlyMiddleware)

	users := jg.Group("/users")
	users.GET("", user.GetMany)
	users.POST("", user.Post, trainerOnlyMiddleware)
	users.DELETE("/:id", user.Delete, trainerOnlyMiddleware)
	users.PUT("", user.Put, trainerOnlyMiddleware)

	blocks := jg.Group("/blocks")
	blocks.GET("", block.GetMany)
	blocks.POST("", block.Post, trainerOnlyMiddleware)
	blocks.DELETE("/:id", block.Delete, trainerOnlyMiddleware)
	blocks.POST("/undelete/:id", block.PostUndelete, trainerOnlyMiddleware)

	weeks := jg.Group("/weeks")
	weeks.GET("", week.GetFiltered)
	weeks.GET("/:id", week.Get)
	weeks.POST("", week.Post, trainerOnlyMiddleware)
	weeks.PUT("", week.Put, trainerOnlyMiddleware)
	weeks.DELETE("/:id", week.Delete, trainerOnlyMiddleware)
	weeks.POST("/undelete/:id", week.PostUndelete, trainerOnlyMiddleware)
	weeks.POST("/duplicate", week.PostDuplicate, trainerOnlyMiddleware)

	week_days := jg.Group("/week-days")
	week_days.GET("", weekday.GetMany)
	week_days.GET("/:id", weekday.Get)
	week_days.POST("", weekday.Post, trainerOnlyMiddleware)
	week_days.PUT("", weekday.Put, trainerOnlyMiddleware)
	week_days.DELETE("/:id", weekday.Delete, trainerOnlyMiddleware)
	week_days.POST("/from-raw", weekday.PostFromRaw, trainerOnlyMiddleware)
	week_days.POST("/undelete/:id", weekday.PostUndelete, trainerOnlyMiddleware)
	week_days.DELETE("/timeslots/:id", weekday.DeleteTimeslot, trainerOnlyMiddleware)

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
