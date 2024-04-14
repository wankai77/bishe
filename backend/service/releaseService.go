package service

import (
	"backend/core/db"
	model "backend/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// type Activity struct {
// 	ReleaseUserId string `gorm:"type:varchar(30);not null;comment:发布人id" json:"releaseUserId"`
// 	ActName       string `gorm:"type:varchar(30);not null;comment:活动名" json:"actName"`
// 	Type          string `gorm:"type:varchar(30);comment:类型" json:"type"`
// 	EndTime       string `gorm:"type:Datetime;comment:截止时间" json:"endTime"`
// 	Detail        string `gorm:"type:varchar(1024);comment:详情" json:"detail"`
// }

func ReleaseActivity(c *gin.Context) {
	idStr := c.PostForm("ID")
	id, _ := strconv.Atoi(idStr)
	actName := c.PostForm("actName")
	type_ := c.PostForm("type")
	endtimetamp := c.PostForm("endTime")
	detail := c.PostForm("detail")
	actimg := c.PostForm("actimg")
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(endtimetamp)
	endtimetampint, _ := strconv.ParseInt(endtimetamp, 10, 64)
	endtime := StringtoTime(endtimetampint)
	fmt.Println(endtime)
	newAct := model.ReleaseInfo{
		ActID:        id,
		ActName:      actName,
		Type:         type_,
		EndTime:      endtime,
		Detail:       detail,
		Actimg:       actimg,
		Applicattime: currentTime,
	}
	// 在此处执行将活动数据存储到数据库或其他操作
	if err := db.Db.Create(&newAct).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    800,
			"message": "发布失败",
		})
		return
	}
	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "发布成功",
	})
}

func StringtoTime(timestamp int64) time.Time {
	// 指定时区为东八区（北京时间）
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("无法加载时区：", err)
	}
	// 使用 time.Unix() 将时间戳转换为 time.Time 类型
	t := time.Unix(timestamp, 0).In(loc)

	// 输出转换后的时间
	fmt.Println("时间戳转换后的北京时间：", t)

	return t
}
