package service

import (
	"net/http"

	"backend/core/db"
	model "backend/models"
	"backend/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	//获取参数
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	//数据验证
	if len(username) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    415,
			"message": "用户名不能为空",
		})
		return
	}

	//判断用户名是否存在
	var user model.UserInfo
	db.Db.Where("name = ?", username).First(&user)

	if user.ID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    429,
			"message": "用户已存在",
		})
		return
	}
	//判断email是否存在
	db.Db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    425,
			"message": "邮箱已注册",
		})
		return
	}

	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "密码加密错误",
		})
		return
	}
	newUser := model.UserInfo{
		Name:     username,
		Email:    email,
		Password: string(hasedPassword),
	}
	if err := db.Db.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    800,
			"message": "注册失败",
		})
		return
	}

	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
	})
}

// 登录
var (
	message string
	j       *utils.JWT
)

func Login(c *gin.Context) {
	var user model.UserInfo
	username := c.PostForm("username")
	password := c.PostForm("password")
	//判断用户是否存在
	db.Db.Where("name = ?", username).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    435,
			"message": "用户不存在...",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    431,
			"message": "密码错误",
		})
		return
	}
	//返回结果
	// c.JSON(http.StatusOK, gin.H{
	// 	"code":    200,
	// 	"message": "登录成功",
	// })
	//token
	token, err := j.TokenNew(int64(user.ID), username) //新建token
	if err != nil {
		message = "新建 token 失败"
		// z.Error(fmt.Sprintf("%s:%s", message, err))
		response.Failed(c, http.StatusOK, message, nil)
		return
	}
	resp := struct {
		Token string `json:"token,omitempty"`
	}{Token: token}
	response.Success(c, resp)
}
