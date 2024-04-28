package models

import (
	"fmt"
	"github.com/duxweb/go-fast/database"
	"github.com/duxweb/go-fast/helper"
	"gorm.io/gorm"
)

// SystemUserMigrate @AutoMigrate()
var SystemUserMigrate = database.Migrate{
	Model: &SystemUser{},
	Seed: func(db *gorm.DB) {
		fmt.Println("admin")
		db.Model(SystemUser{}).Create(&SystemUser{
			Username: "admin",
			Password: helper.HashEncode("admin"),
		})
	},
}

type SystemUser struct {
	database.Fields
	Username string       `gorm:"uniqueIndex;size:250" json:"username"`
	Nickname string       `gorm:"size:250" json:"nickname"`
	Avatar   string       `gorm:"size:250" json:"avatar"`
	Password string       `gorm:"size:250" json:"-"`
	Status   bool         `gorm:"default:true" json:"status"`
	Roles    []SystemRole `gorm:"many2many:system_user_role"`
}
