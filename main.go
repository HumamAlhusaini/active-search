package main

import (
	"net/http"

	"project/internal/views/components"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/", "public")

	e.GET("/", func(c echo.Context) error {
		buf := templ.GetBuffer()
		defer templ.ReleaseBuffer(buf)

		accordion := components.AccordionExample()
		if err := accordion.Render(c.Request().Context(), buf); err != nil {
			return err
		}

		return c.HTML(http.StatusOK, buf.String())
	})

	e.Start(":8080")
}
