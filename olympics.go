package olympics

import (
	"olympics/handler"

	"github.com/labstack/echo/v4"
)

func Run() {
	app := echo.New()

	app.Static("/static", "public/static")

	app.GET("*", handler.NotFound)
	app.GET("/", handler.Index)

	app.Start(":8080")
}
