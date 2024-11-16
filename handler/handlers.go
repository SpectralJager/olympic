package handler

import (
	"net/http"
	"olympics/view"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func NotFound(c echo.Context) error {
	return c.String(http.StatusNotFound, "not found")
}

func Index(c echo.Context) error {
	return render(c, view.IndexPage())
}
