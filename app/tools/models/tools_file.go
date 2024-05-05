package models

import (
	"github.com/duxweb/go-fast/models"
)

// ToolsFile @AutoMigrate()
type ToolsFile struct {
	models.Fields
	DirID   uint   `gorm:"comment:目录id" json:"dir_id"`
	HasType string `gorm:"size:50;comment:关联类型" json:"has_type"`
	Driver  string `gorm:"size:50;comment:驱动类型" json:"driver"`
	Url     string `gorm:"comment:链接" json:"url"`
	Path    string `gorm:"comment:路径" json:"path"`
	Name    string `gorm:"comment:文件名" json:"name"`
	Ext     string `gorm:"size:20;comment:后缀" json:"ext"`
	Size    uint   `gorm:"comment:大小" json:"size"`
	Mime    string `gorm:"comment:MIME" json:"mime"`
	Dir     ToolsFileDir
}
