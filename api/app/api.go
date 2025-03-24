package app

import (
	"fmt"
	"net/http"
	"time"
	"trainer-helper/api"
	"trainer-helper/api/exercise_handler"
	"trainer-helper/api/exercise_handler/count_handler"
	"trainer-helper/api/person_handler"
	"trainer-helper/api/timeslot_handler"
	"trainer-helper/api/timeslot_handler/revert_handler"
	"trainer-helper/api/work_set_handler"
	"trainer-helper/crud"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func CustomLogger(next echo.HandlerFunc) echo.HandlerFunc {
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

		// Determine the color for the method
		var methodColor string
		switch method {
		case "GET":
			methodColor = ColorBlue // Blue for GET
		case "POST":
			methodColor = ColorGreen // Green for POST
		case "PUT":
			methodColor = ColorYellow // Yellow for PUT
		case "DELETE":
			methodColor = ColorRed // Red for DELETE
		default:
			methodColor = ColorCyan // Default color for other methods
		}

		// Log the details
		fmt.Printf("[%s] %s%-6s%s %s %s(%d) [%v]\n",
			stop.Format("2006-01-02 15:04:05"), // Timestamp
			methodColor, method, ColorReset,    // HTTP method with color and reset
			path,    // URL path
			ip,      // Client IP
			status,  // Status code
			latency, // Latency
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
	e.Use(CustomLogger)
	e.Use(middleware.CORS())

	e.GET("/-/ping", pong)
	e.GET("/timeslot", timeslot_handler.Get)
	e.POST("/timeslot", timeslot_handler.Post)
	e.DELETE("/timeslot", timeslot_handler.Delete)
	e.PUT("/timeslot", timeslot_handler.Put)
	e.PUT("/timeslot/revert", timeslot_revert_handler.Put)
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
