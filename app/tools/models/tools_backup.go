package models

import "github.com/duxweb/go-fast/models"

// ToolsBackup @AutoMigrate()
type ToolsBackup struct {
	models.Fields
	Name string `gorm:"comment:文件名称" json:"name"`
	Url  string `gorm:"comment:下载链接" json:"url"`
}
