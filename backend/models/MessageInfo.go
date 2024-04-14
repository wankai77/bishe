package model

import (
	"time"

	"gorm.io/gorm"
)

// 用户个人信息
type MessageInfo struct {
	gorm.Model
	MsgID     int       `gorm:"type:bigint;comment:msgid" json:"msgid"`
	MsgUserId int       `gorm:"type:int;not null;comment:消息人id" json:"msguserId"`
	MsgUser   UserInfo  `gorm:"foreignkey:ID"`
	Type      string    `gorm:"type:varchar(30);comment:类型" json:"type"`
	Detail    string    `gorm:"type:varchar(1024);comment:详情" json:"detail"`
	State     string    `gorm:"type:varchar(30);default:待审核;comment:状态" json:"state"`
	Time      time.Time `gorm:"type:datetime;comment:接收消息时间" json:"time"`
}
