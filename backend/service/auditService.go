package service

import (
	"backend/dao"
	model "backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRelInfo(c *gin.Context) {
	var (
		msg  string
		err  error
		resp []model.ReleaseInfo
	)
	act := model.ReleaseInfo{}
	resp, err = dao.GetActInfo(&act)
	if err != nil {
		msg = "数据库查询失败"
		// z.Error(fmt.Sprintf("%s:%s", msg, err))
		response.Failed(c, http.StatusOK, msg, nil)
		return
	}
	response.Success(c, resp)
}

// 审核活动
type Audit struct {
	State string `form:"state" binding:"required"`
	ActID int    `form:"actid" binding:"required"`
}

func AuditRelInfo(c *gin.Context) {
	var (
		reqForm Audit
		msg     string
		err     error
	)
	if err = c.ShouldBind(&reqForm); err != nil {
		msg = "请求不合法"
		fmt.Println(err)
		response.Failed(c, http.StatusOK, msg, nil)
		return
	}
	act := model.ReleaseInfo{
		State: reqForm.State,
		ActID: reqForm.ActID,
	}
	err = dao.UpdateState(&act)
	if err != nil {

		msg = "审核已通过"

		response.Failed(c, http.StatusOK, msg, nil)
		return
	}
	response.Success(c, nil)
}
