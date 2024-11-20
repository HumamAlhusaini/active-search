package main

import (
	"github.com/HumamAlhusaini/active-search/internal/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/", "public")

	e.GET("/active-search-table", handler.GetActiveSearchExampleTable)
	e.GET("/active-search", handler.GetActiveSearchExample)

	e.Start(":7000")
}
