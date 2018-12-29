package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/virtcanhead/health"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Static("/", "public")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(health.New())
	e.Logger.Fatal(e.Start(":" + Port))
}
