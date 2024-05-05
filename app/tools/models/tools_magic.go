package models

import (
	"encoding/json"
	"github.com/duxweb/go-fast/cache"
	"github.com/duxweb/go-fast/database"
	"github.com/golang-module/carbon/v2"
	"gorm.io/datatypes"
)

// ToolsMagic @AutoMigrate()
type ToolsMagic struct {
	ID        uint            `json:"id" gorm:"primaryKey"  nestedset:"id"`
	GroupID   uint            `gorm:"comment:分组id" json:"group_id"`
	Label     string          `gorm:"size:255;comment:标签" json:"label"`
	Name      string          `gorm:"size:255;comment:名称" json:"name"`
	Type      string          `gorm:"size:50;comment:数据类型;default:common" json:"type"`
	TreeLabel string          `gorm:"size:50;comment:树形标签" json:"tree_label"`
	Inline    *bool           `gorm:"comment:附属模型" json:"inline"`
	Page      *bool           `gorm:"comment:页面操作" json:"page"`
	External  datatypes.JSON  `gorm:"comment:数据权限" json:"external"`
	Fields    datatypes.JSON  `gorm:"comment:数据字段" json:"fields"`
	CreatedAt carbon.DateTime `json:"created_at"`
	UpdatedAt carbon.DateTime `json:"updated_at"`
	Group     ToolsMagicGroup
}

func GetMagicMenu() []map[string]any {
	key := []byte("magic.menus")
	data, err := cache.Injector().Get(key)
	groupData := []map[string]any{}
	if err == nil {
		_ = json.Unmarshal(data, &groupData)
		return groupData
	}
	groups := []ToolsMagicGroup{}
	database.Gorm().Model(ToolsMagicGroup{}).Find(&groups)

	for _, group := range groups {
		magics := []ToolsMagic{}
		database.Gorm().Model(ToolsMagic{}).Where("group_id = ?", group.ID).Where("inline = ?", false).Find(&magics)

		items := []map[string]any{}
		for _, magic := range magics {
			items = append(items, map[string]any{
				"name":  magic.Name,
				"label": magic.Label,
			})
		}
		groupData = append(groupData, map[string]any{
			"name":     group.Name,
			"label":    group.Label,
			"icon":     group.Icon,
			"children": items,
		})
	}

	marshal, _ := json.Marshal(groupData)
	_ = cache.Injector().Set(key, marshal, 0)

	return groupData
}
