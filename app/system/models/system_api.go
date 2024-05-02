package models

import "github.com/duxweb/go-fast/models"

// SystemApi @AutoMigrate()
type SystemApi struct {
	models.Fields
	Name      string `gorm:"size:250" json:"name"`
	SecretID  string `gorm:"size:250" json:"secret_id"`
	SecretKey string `gorm:"size:250" json:"secret_key"`
	Status    bool   `gorm:"default:true" json:"status"`
}
