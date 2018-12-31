package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func routes(e *echo.Echo) {
	e.GET("/api/sponsors", routeSponsors)
	e.POST("/api/sponsors/add", routeAddSponsor, adminTokenValidator())
}

func adminTokenValidator() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Header.Get("X-Admin-Token") != AdminToken {
				return c.String(http.StatusForbidden, "INVALID ADMIN TOKEN")
			}

			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}

func routeSponsors(c echo.Context) (err error) {
	var sponsors []Sponsor
	if err = DB.Order("id DESC").Limit(500).Find(&sponsors).Error; err != nil {
		return
	}
	return c.JSON(http.StatusOK, sponsors)
}

// AddSponsorForm add sponsor form
type AddSponsorForm struct {
	Name string `json:"name"`
}

func routeAddSponsor(c echo.Context) (err error) {
	form := new(AddSponsorForm)
	if err = c.Bind(form); err != nil {
		return
	}
	if err = DB.Unscoped().Delete(&Sponsor{}, map[string]interface{}{"name": form.Name}).Error; err != nil {
		return
	}
	if err = DB.Create(&Sponsor{Name: form.Name}).Error; err != nil {
		return
	}
	return c.String(http.StatusOK, "OK")
}
