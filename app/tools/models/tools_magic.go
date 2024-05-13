package models

import (
	"github.com/duxweb/go-fast/database"
	"github.com/duxweb/go-fast/menu"
	coreModel "github.com/duxweb/go-fast/models"
	"github.com/golang-module/carbon/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm/clause"
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

type ToolsMagicFields struct {
	Key      string             `json:"key"`
	List     bool               `json:"list"`
	Name     string             `json:"name"`
	Type     string             `json:"type"`
	Label    string             `json:"label"`
	Search   bool               `json:"search"`
	Setting  map[string]any     `json:"setting"`
	Required bool               `json:"required"`
	Child    []ToolsMagicFields `json:"child"`
}

func GetMagicMenu(appMenu *menu.MenuData) {
	//key := []byte("magic.menus")
	//data, err := cache.Injector().Get(key)
	apps := []ToolsMagicGroup{}
	database.Gorm().Model(ToolsMagicGroup{}).Preload(clause.Associations, coreModel.ChildrenPreload).Where("parent_id = 0").Find(&apps)

	for i, app := range apps {

		t := appMenu.Add(&menu.MenuData{
			Name:  "tools.data." + app.Name,
			Label: app.Label,
			Icon:  app.Icon,
			Meta: map[string]any{
				"sort": 500 + i,
			},
		})

		// 一级数据
		appMagics := []ToolsMagic{}
		database.Gorm().Model(ToolsMagic{}).Where("group_id = ?", app.ID).Where("inline = ?", false).Find(&appMagics)
		if len(appMagics) > 0 {
			for _, magic := range appMagics {
				t.Item("tools.data."+magic.Name, magic.Label, "data/"+magic.Name, 0)
			}
		}

		// 一级分组
		if len(app.Children) > 0 {
			for _, topGroup := range app.Children {

				// 二级分组
				g := t.Group(topGroup.Name, topGroup.Label, topGroup.Icon)

				// 二级数据
				topMagics := []ToolsMagic{}
				database.Gorm().Model(ToolsMagic{}).Where("group_id = ?", topGroup.ID).Where("inline = ?", false).Find(&topMagics)
				if len(topMagics) > 0 {
					for _, magic := range topMagics {
						g.Item("tools.data."+magic.Name, magic.Label, "data/"+magic.Name, 0)
					}
				}
			}
		}
	}

	//_ = cache.Injector().Set(key, []byte("true"), 0)
}
