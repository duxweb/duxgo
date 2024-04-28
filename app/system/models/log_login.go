package models

import (
	"github.com/duxweb/go-fast/database"
)

// LogLogin @AutoMigrate()
type LogLogin struct {
	database.Fields
	UserType string `gorm:"size:250;comment:关联类型" json:"user_type"`
	UserId   uint   `gorm:"size:20;comment:关联 id" json:"user_id"`
	Browser  string `gorm:"size:250;comment:浏览器" json:"browser"`
	Ip       string `gorm:"size:100;comment:IP" json:"ip"`
	Platform string `gorm:"size:100;comment:平台" json:"platform"`
	Status   bool   `gorm:"default:true" json:"status"`
}
