package models

import (
	"github.com/duxweb/go-fast/models"
)

// ToolsMagicGroup @AutoMigrate()
type ToolsMagicGroup struct {
	models.Fields
	ParentID uint   `gorm:"comment:上级id" json:"parent_id"`
	Label    string `gorm:"size:50;comment:标签" json:"label"`
	Name     string `gorm:"size:50;comment:名称" json:"name"`
	Icon     string `gorm:"size:20;comment:图标" json:"icon"`
	Parent   *ToolsMagicGroup
	Children []ToolsMagicGroup `gorm:"foreignkey:ParentID;default:0" json:"children"`
}
