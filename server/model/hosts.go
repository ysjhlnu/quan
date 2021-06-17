// 自动生成模板Hosts
package model

import (
	"quan/global"
)

// 如果含有time.Time 请自行import time包
type Hosts struct {
      global.GVA_MODEL
      Instanceid  string `json:"instanceid" form:"instanceid" gorm:"column:instanceid;comment:实例ID;type:varchar(50);"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:实例名字;type:varchar(50);"`
      Privateip  string `json:"privateip" form:"privateip" gorm:"column:privateip;comment:内网IP;type:varchar(50);"`
      Region  string `json:"region" form:"region" gorm:"column:region;comment:地区;type:varchar(50);"`
      Env  string `json:"env" form:"env" gorm:"column:env;comment:环境;type:varchar(50);"`
      Status  string `json:"status" form:"status" gorm:"column:status;comment:运行状态;type:varchar(50);"`
      Type  string `json:"type" form:"type" gorm:"column:type;comment:实例模板;type:varchar(50);"`
      Ps  string `json:"ps" form:"ps" gorm:"column:ps;comment:其他;type:varchar(50);"`
}


func (Hosts) TableName() string {
  return "hosts"
}

