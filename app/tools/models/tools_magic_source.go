package models

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/datatypes"
)

// ToolsMagicSource @AutoMigrate()
type ToolsMagicSource struct {
	ID        uint            `json:"id" gorm:"primaryKey"`
	Name      string          `gorm:"size:255;comment:名称" json:"name"`
	Type      uint            `gorm:"size:1;comment:类型;default:0" json:"type"`
	Data      datatypes.JSON  `gorm:"comment:数据" json:"data"`
	Url       string          `gorm:"size:255;comment:地址" json:"url"`
	CreatedAt carbon.DateTime `json:"created_at"`
	UpdatedAt carbon.DateTime `json:"updated_at"`
}
