package models

import "github.com/duxweb/go-fast/models"

// LogVisitSpider @AutoMigrate()
type LogVisitSpider struct {
	models.Fields
	HasType string `gorm:"size:250;comment:关联类型" json:"has_type"`
	HasId   uint   `gorm:"size:20;comment:关联 id" json:"has_id"`
	Name    string `gorm:"size:250;comment:蜘蛛名" json:"name"`
	Path    string `gorm:"size:250;comment:页面路径" json:"path"`
	Num     uint   `gorm:"default:0;comment:访客量" json:"num"`
}
