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

	app.GET("/constructor", handler.NewConstructor().Handler)

	app.Start("0.0.0.0:8080")
}
