package main

import (
	"github.com/jinzhu/gorm"
)

// Sponsor sponsor
type Sponsor struct {
	gorm.Model

	Name string `gorm:"NOT NULL;INDEX"`
}

// Post post
type Post struct {
	gorm.Model

	ImageURL string `gorm:"NOT NULL"`
	Content  string `gorm:"NOT NULL"`
	Link     string `gorm:"NOT NULL"`
}
