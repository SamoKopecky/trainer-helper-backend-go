package app

import (
	"fmt"
	"net/http"
	"time"
	"trainer-helper/api"
	"trainer-helper/api/exercise_count_handler"
	"trainer-helper/api/exercise_handler"
	"trainer-helper/api/person_handler"
	"trainer-helper/api/timeslot_handler"
	"trainer-helper/api/work_set_handler"
	"trainer-helper/crud"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
)

// Custom logger middleware
func ColorfulLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Start timer
		start := time.Now()

		// Process request
		err := next(c)

		// Stop timer
		stop := time.Now()
		latency := stop.Sub(start)

		// Get request details
		method := c.Request().Method
		path := c.Request().URL.Path
		status := c.Response().Status
		ip := c.RealIP()

		// Determine the color based on the status code
		color := ColorGreen
		if status >= 400 && status < 500 {
			color = ColorYellow
		} else if status >= 500 {
			color = ColorRed
		}

		// Log the details in a colorful format
		fmt.Printf("%s[%s] %s%-10s %s%s %s(%d) %s[%v]%s\n",
			ColorCyan, stop.Format("2006-01-02 15:04:05"), // Timestamp
			color, method, // HTTP method
			path,                  // URL path
			ColorBlue, ip, status, // IP and status
			ColorYellow, latency, // Latency
			ColorReset, // Reset color
		)

		return err
	}
}

func RunApi(db *bun.DB) {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &api.DbContext{Context: c,
				CRUDExercise: crud.NewCRUDExercise(db),
				CRUDTimeslot: crud.NewCRUDTimeslot(db),
				CRUDWorkSet:  crud.NewCRUDWorkSet(db),
				CRUDPerson:   crud.CRUDPerson{Db: db}}
			return next(cc)
		}
	})
	e.Use(ColorfulLogger)
	e.Use(middleware.CORS())

	e.GET("/-/ping", pong)
	e.GET("/timeslot", timeslot_handler.Get)
	e.POST("/timeslot", timeslot_handler.Post)
	e.DELETE("/timeslot", timeslot_handler.Delete)
	e.PUT("/timeslot", timeslot_handler.Put)
	e.GET("/exercise/:id", exercise_handler.Get)
	e.PUT("/exercise", exercise_handler.Put)
	e.DELETE("/exercise", exercise_handler.Delete)
	e.POST("/exercise", exercise_handler.Post)
	e.PUT("/exercise/count", exercise_count_handler.Put)
	e.DELETE("/exercise/count", exercise_count_handler.Delete)
	e.PUT("/workset", work_set_handler.Put)
	e.GET("/person", person_handler.Get)

	e.Logger.Fatal(e.Start(":1323"))
}

func pong(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
