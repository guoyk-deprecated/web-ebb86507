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

	ImageURL string `gorm:"Type:TEXT;NOT NULL"`
	Content  string `gorm:"Type:TEXT;NOT NULL"`
	Link     string `gorm:"Type:TEXT;NOT NULL"`
}
