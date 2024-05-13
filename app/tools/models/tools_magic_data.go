package models

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/datatypes"
)

// ToolsMagicData @AutoMigrate()
type ToolsMagicData struct {
	ID        uint            `json:"id" gorm:"primaryKey"`
	ParentID  uint            `gorm:"comment:上级id" json:"parent_id"`
	MagicID   uint            `gorm:"comment:关联id" json:"magic_id"`
	Data      datatypes.JSON  `gorm:"comment:数据权限" json:"external"`
	CreatedAt carbon.DateTime `json:"created_at"`
	UpdatedAt carbon.DateTime `json:"updated_at"`
	Parent    *ToolsMagicData
	Children  []ToolsMagicData `gorm:"foreignkey:ParentID;default:0" json:"children"`
}
