package model

import (
	"gorm.io/gorm"
)

type ApplicatUser struct {
	gorm.Model
	AppUserName string `gorm:"type:string;comment:报名人用户名" json:"appusername"`
	TrueName    string `gorm:"type:string;comment:真实姓名" json:"truename"`
	Class       string `gorm:"type:string;comment:班级" json:"class"`
}
