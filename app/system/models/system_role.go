package models

import (
	"github.com/duxweb/go-fast/database"
	"gorm.io/datatypes"
)

// SystemRole @AutoMigrate()
type SystemRole struct {
	database.Fields
	Name       string         `gorm:"size:250" json:"username"`
	Permission datatypes.JSON `json:"permission"`
	Users      []SystemUser   `gorm:"many2many:system_user_role"`
}
