package service

import (
	"fmt"
	"net/http"

	"backend/dao"
	model "backend/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type reqResetPassword struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
	Desc     string `form:"desc"`
}

func ResetPassword(c *gin.Context) {
	fmt.Println("load....")
	var (
		reqForm reqResetPassword
		msg     string
		err     error
	)
	if err = c.ShouldBind(&reqForm); err != nil {
		msg = "请求不合法"
		fmt.Println(err)
		response.Failed(c, http.StatusOK, msg, nil)
		return
	}
	hasedPassword, _ := bcrypt.GenerateFromPassword([]byte(reqForm.Password), bcrypt.DefaultCost)
	user := model.UserInfo{
		Name:     reqForm.Name,
		Password: string(hasedPassword),
		Intro:    reqForm.Desc,
	}
	err = dao.UpdatePassword(&user)
	if err != nil {

		msg = "更新密码失败"

		response.Failed(c, http.StatusOK, msg, nil)
		return
	}
	response.Success(c, nil)
}
