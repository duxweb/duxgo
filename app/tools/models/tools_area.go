package models

// ToolsArea @AutoMigrate()
type ToolsArea struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	ParentCode string `gorm:"size:50;comment:上级编码" json:"parent_code"`
	Code       string `gorm:"size:50;comment:地区编码" json:"code"`
	Name       string `gorm:"comment:名称" json:"name"`
	Level      uint   `gorm:"comment:层级" json:"level"`
	Leaf       *bool  `gorm:"size:50;comment:终节点" json:"leaf"`
}
