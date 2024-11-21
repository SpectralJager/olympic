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

	constructorHandler := handler.NewConstructor()
	app.GET("/constructor", constructorHandler.GET)
	app.POST("/constructor", constructorHandler.POST)
	app.DELETE("/constructor", constructorHandler.DELETE)

	app.Start("0.0.0.0:8080")
}
