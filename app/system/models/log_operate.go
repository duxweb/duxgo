package models

import (
	"github.com/duxweb/go-fast/database"
	"gorm.io/datatypes"
)

type LogOperate struct {
	database.Fields
	UserType      string         `gorm:"size:250;comment:关联类型" json:"user_type"`
	UserID        uint           `gorm:"size:20;comment:关联 id" json:"user_id"`
	RequestMethod string         `gorm:"size:20;comment:请求方法" json:"request_method"`
	RequestUrl    string         `gorm:"size:20;comment:请求链接" json:"request_url"`
	RequestParams datatypes.JSON `gorm:"size:20;comment:请求链接" json:"request_params"`
	RequestTime   float64        `gorm:"precision:3;comment:访客量" json:"request_time"`
	RouteName     string         `gorm:"size:50;comment:关联类型" json:"route_name"`
	RouteTitle    string         `gorm:"size:50;comment:关联标题" json:"route_title"`
	ClientUa      string         `gorm:"size:250;comment:UA" json:"client_ua"`
	ClientIp      string         `gorm:"size:100;comment:IP" json:"client_ip"`
	ClientBrowser string         `gorm:"size:250;comment:浏览器" json:"client_browser"`
	ClientDevice  string         `gorm:"size:250;comment:设备" json:"client_device"`
	User          any            `gorm:"-"`
}
