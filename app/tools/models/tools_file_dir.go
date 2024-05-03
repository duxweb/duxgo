package models

import (
	"github.com/duxweb/go-fast/models"
)

// ToolsFileDir @AutoMigrate()
type ToolsFileDir struct {
	models.Fields
	Name    string `gorm:"size:50;comment:名称" json:"name"`
	HasType string `gorm:"size:20;comment:关联类型" json:"has_type"`
}
