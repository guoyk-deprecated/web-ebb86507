package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func routes(e *echo.Echo) {
	e.GET("/api/sponsors", routeSponsors)
}

func routeSponsors(c echo.Context) (err error) {
	var sponsors []Sponsor
	if err = DB.Find(&sponsors).Error; err != nil {
		return
	}
	return c.JSON(http.StatusOK, sponsors)
}
