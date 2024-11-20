package handler

import (
	"olympics/view"

	"github.com/labstack/echo/v4"
)

type Constructor struct {
}

func NewConstructor() *Constructor {
	return &Constructor{}
}

func (h *Constructor) Handler(c echo.Context) error {
	switch c.Request().Header.Get("Page-Action") {

	default:
		return h.view(c)
	}
}

func (h *Constructor) view(c echo.Context) error {
	return render(c, view.ConstructorPage())
}
