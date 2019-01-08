package main

import (
	"net/http"
	"time"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
)

func routes(e *echo.Echo) {
	e.GET("/api/sponsors", routeSponsors)
	e.GET("/api/posts", routePosts)
	e.POST("/api/sponsors/add", routeAddSponsor, adminTokenValidator())
	e.POST("/api/posts/add", routeAddPost, adminTokenValidator())
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

// SponsorViewModel sponsor view model
type SponsorViewModel struct {
	ID   uint
	Name string
}

func routeSponsors(c echo.Context) (err error) {
	var sponsors []Sponsor
	var ret []*SponsorViewModel
	if err = DB.Order("id DESC").Limit(500).Find(&sponsors).Error; err != nil {
		return
	}
	copier.Copy(&ret, &sponsors)
	return c.JSON(http.StatusOK, ret)
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

// PostViewModel view model for post
type PostViewModel struct {
	ID        uint
	CreatedAt time.Time
	ImageURL  string
	Content   string
	Link      string
}

func routePosts(c echo.Context) (err error) {
	var posts []Post
	var ret []*PostViewModel
	var db = DB
	if len(c.QueryParam("lastId")) > 0 {
		db = db.Where("id < ?", c.QueryParam("lastId"))
	}
	if err = db.Order("id DESC").Limit(50).Find(&posts).Error; err != nil {
		return
	}
	copier.Copy(&ret, &posts)
	return c.JSON(http.StatusOK, ret)
}

// AddPostForm add sponsor form
type AddPostForm struct {
	Content  string `json:"content"`
	ImageURL string `json:"imageUrl"`
	Link     string `json:"link"`
}

func routeAddPost(c echo.Context) (err error) {
	form := new(AddPostForm)
	if err = c.Bind(form); err != nil {
		return
	}
	if err = DB.Create(&Post{Content: form.Content, ImageURL: form.ImageURL, Link: form.Link}).Error; err != nil {
		return
	}
	return c.String(http.StatusOK, "OK")
}
