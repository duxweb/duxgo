package models

import (
	"fmt"
	"github.com/duxweb/go-fast/database"
	"github.com/duxweb/go-fast/helper"
	"github.com/duxweb/go-fast/models"
	"github.com/gookit/goutil/jsonutil"
	"github.com/samber/lo"
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
	models.Fields
	Username string       `gorm:"uniqueIndex;size:250" json:"username"`
	Nickname string       `gorm:"size:250" json:"nickname"`
	Avatar   string       `gorm:"size:250" json:"avatar"`
	Password string       `gorm:"size:250" json:"-"`
	Status   bool         `gorm:"default:true" json:"status"`
	Roles    []SystemRole `gorm:"many2many:system_user_role"`
	//Operates []LogOperate `gorm:"polymorphic:User;polymorphicValue:Admin"`
}

func (u SystemUser) GetPermission() map[string]bool {
	permissionData := map[string]bool{}
	permissions := []string{}
	for _, role := range u.Roles {
		items := map[string]bool{}
		_ = jsonutil.DecodeString(role.Permission.String(), &items)
		if len(items) == 0 {
			continue
		}
		for k, v := range items {
			permissionData[k] = v
			if v {
				permissions = append(permissions, k)
			}
		}
	}

	for k, v := range permissionData {
		if lo.IndexOf[string](permissions, k) == -1 {
			continue
		}
		permissionData[k] = v
	}
	return permissionData
}
