package service

import (
	"backend/dao"
	model "backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMessage(c *gin.Context) {
	var (
		msg  string
		err  error
		resp []model.MessageInfo
	)
	// message := model.MessageInfo{}
	userID, _ := c.Get("userID")
	message := model.MessageInfo{MsgUserId: int(userID.(int64))}
	resp, err = dao.GetMessageByID(&message)
	if err != nil {
		msg = "数据库查询失败"
		// z.Error(fmt.Sprintf("%s:%s", msg, err))
		response.Failed(c, http.StatusOK, msg, nil)
		return
	}
	response.Success(c, resp)
}
