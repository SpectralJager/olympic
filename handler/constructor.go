package handler

import (
	"fmt"
	"olympics/view"

	"github.com/labstack/echo/v4"
)

type Constructor struct {
}

func NewConstructor() *Constructor {
	return &Constructor{}
}

func (h *Constructor) GET(c echo.Context) error {
	switch c.Request().Header.Get("Page-Action") {
	case "get_edit_exercise_modal":
		return h.getEditExerciseModal(c)
	default:
		return h.view(c)
	}
}

func (h *Constructor) POST(c echo.Context) error {
	action := c.Request().Header.Get("Page-Action")
	switch action {
	case "append":
		return h.append(c)
	case "save_edit_exercise":
		return h.saveEditExercise(c)
	default:
		return fmt.Errorf("unexpected POST Page-Action: %s", action)
	}
}

func (h *Constructor) DELETE(c echo.Context) error {
	action := c.Request().Header.Get("Page-Action")
	switch action {
	case "remove":
		return Empty(c)
	default:
		return fmt.Errorf("unexpected DELETE Page-Action: %s", action)
	}
}

func (h *Constructor) view(c echo.Context) error {
	return render(c, view.ConstructorPage())
}

func (h *Constructor) append(c echo.Context) error {
	value := c.FormValue("value")
	return render(c, view.TrainingExercise(value))
}

func (h *Constructor) getEditExerciseModal(c echo.Context) error {
	value := c.FormValue("value")
	return render(c, view.EditExercise(value))
}

func (h *Constructor) saveEditExercise(c echo.Context) error {
	value := c.FormValue("value")
	return render(c, view.TrainingExercise(value))
}
