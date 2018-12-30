package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/virtcanhead/health"
)

var (
	// DB the database
	DB *gorm.DB
)

func main() {
	var err error
	if DB, err = gorm.Open("mysql", DatabaseDSN); err != nil {
		panic(err)
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Static("/", "public")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(health.New(&GORMResource{DB}))
	e.Logger.Fatal(e.Start(":" + Port))
}
