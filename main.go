package main

import (
	"flag"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/virtcanhead/health"
)

var (
	verbose bool

	// DB the database
	DB *gorm.DB
)

func main() {
	flag.BoolVar(&verbose, "verbose", false, "enable verbose mode")
	flag.Parse()

	var err error
	if DB, err = gorm.Open("mysql", DatabaseDSN); err != nil {
		panic(err)
	}

	if err = DB.AutoMigrate(&Sponsor{}, &Vote{}, &Post{}).Error; err != nil {
		panic(err)
	}

	DB.LogMode(verbose)

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Static("/", "public")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(health.New(&GORMResource{DB}))
	routes(e)
	e.Logger.Fatal(e.Start(":" + Port))
}
