package service

import (
	"backend/config/result"
	"backend/dao"
	model "backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	// z        = middleware.GetLogger()
	response *result.Result
)

func GetInfo(c *gin.Context) {
	var (
		msg string
		err error
	)
	userID, _ := c.Get("userID")

	user := model.UserInfo{ID: int(userID.(int64))} //根据 userID 的值创建一个 model.UserInfo 的实例，并将其 ID 字段设置为 userID 的值
	err = dao.GetUserInfoByID(&user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			msg = "不存在该用户"
			// z.Error(fmt.Sprintf("%s:%s", msg, err))
			response.Failed(c, http.StatusOK, msg, nil)
			return
		}
		msg = "数据库查询失败"
		// z.Error(fmt.Sprintf("%s:%s", msg, err))
		response.Failed(c, http.StatusOK, msg, nil)
		return
	}

	resp := struct {
		ID          int    `gorm:"type:bigint;comment:id" json:"id"`
		Email       string `gorm:"type:varchar(30);comment:邮箱" json:"email"`
		Name        string `gorm:"type:varchar(30);not null;comment:用户名" json:"Name"`
		Password    string `gorm:"type:varchar(255);not null;comment:密码" json:"Password"`
		Avatar      string `gorm:"type:varchar(1024);not null;comment:头像地址" json:"avatar"`
		Intro       string `gorm:"type:varchar(255);comment:个人简介" json:"intro"`
		Permissions int8   `gorm:"type:tinyint(1);default:0;comment:是否有权限(0-否 1-是) ;column:Permissions" json:"Permissions"`
	}{
		ID:          user.ID,
		Email:       user.Email,
		Name:        user.Name,
		Password:    user.Password,
		Avatar:      user.Avatar,
		Intro:       user.Intro,
		Permissions: user.Permissions,
	}

	response.Success(c, resp)
}
