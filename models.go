package main

import (
	"github.com/jinzhu/gorm"
)

// Sponsor sponsor
type Sponsor struct {
	gorm.Model

	Name string `gorm:"NOT NULL;INDEX"`
}
