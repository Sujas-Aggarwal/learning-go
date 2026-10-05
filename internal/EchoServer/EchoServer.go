package EchoServer

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

var PORT string = ":8000"

func StartServer() {
	app := echo.New()

	app.Use(middleware.RequestLogger())
	app.Use(middleware.Recover())

	app.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})

	if err := app.Start(PORT); err != nil {
		app.Logger.Error("failed to start server", "error", err)
	}
}
