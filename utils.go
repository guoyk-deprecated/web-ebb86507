package main

import (
	"github.com/labstack/echo"
)

// ExtractGA extract GA user track id from request
func ExtractGA(c echo.Context) (ga string, ok bool) {
	cookie, _ := c.Request().Cookie("_ga")
	if cookie == nil {
		return
	}
	ga = cookie.Value
	ok = len(ga) > 0
	return
}
