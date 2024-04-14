package model

import (
	"time"

	"gorm.io/gorm"
)

// 活动信息
type ReleaseInfo struct {
	gorm.Model
	ActID         int       `gorm:"type:bigint;comment:actid" json:"actid"`
	ReleaseUser   UserInfo  `gorm:"foreignkey:ID"`
	ReleaseUserId string    `gorm:"type:varchar(30);not null;comment:发布人id" json:"releaseUserId"`
	ActName       string    `gorm:"type:varchar(30);not null;comment:活动名" json:"actName"`
	Type          string    `gorm:"type:varchar(30);comment:类型" json:"type"`
	EndTime       time.Time `gorm:"type:datetime;comment:截止时间" json:"endTime"`
	Detail        string    `gorm:"type:varchar(1024);comment:详情" json:"detail"`
	Actimg        string    `gorm:"type:varchar(1024);not null;comment:图片地址" json:"actimg"`
	State         string    `gorm:"type:varchar(30);default:待审核;comment:状态" json:"state"`
	Applicattime  string    `gorm:"type:datetime;comment:申请时间" json:"applicattime"`
}
