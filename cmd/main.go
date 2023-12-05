package main

import (
	"context"

	"github.com/JontyCoertzen/echo-templ-htmx/pkg/views"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Static("/assets", "dist")

	e.GET("/", func(c echo.Context) error {
		views.Layout(views.Hello("John")).Render(context.Background(), c.Response().Writer)

		return nil
	})

	e.Logger.Fatal(e.Start(":42069"))
}
