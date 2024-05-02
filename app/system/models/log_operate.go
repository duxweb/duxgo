package models

import (
	"github.com/duxweb/go-fast/models"
)

type LogOperateUser struct {
	models.LogOperate
	User SystemUser
}

func (LogOperateUser) TableName() string {
	return "app_log_operate"
}
