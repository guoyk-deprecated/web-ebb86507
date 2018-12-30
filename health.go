package main

import "github.com/jinzhu/gorm"

// GORMResource gorm resource
type GORMResource struct {
	DB *gorm.DB
}

// HealthCheck implements health check
func (r *GORMResource) HealthCheck() error {
	return r.DB.Exec("SELECT version();").Error
}
