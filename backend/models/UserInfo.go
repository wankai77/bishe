package model

import "gorm.io/gorm"

// 用户个人信息
type UserInfo struct {
	gorm.Model
	ID          int    `gorm:"type:bigint;comment:id" json:"id"`
	Email       string `gorm:"type:varchar(30);comment:邮箱" json:"email"`
	Name        string `gorm:"type:varchar(30);not null;comment:用户名" json:"Name"`
	Password    string `gorm:"type:varchar(255);not null;comment:密码" json:"Password"`
	Avatar      string `gorm:"type:varchar(1024);not null;comment:头像地址" json:"avatar"`
	Intro       string `gorm:"type:varchar(255);comment:个人简介" json:"intro"`
	Permissions int8   `gorm:"type:tinyint(1);default:0;comment:是否有权限(0-否 1-是) ;column:Permissions" json:"Permissions"`
}
