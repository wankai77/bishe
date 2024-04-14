package service

import (
	"backend/dao"
	model "backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetActInfo(c *gin.Context) {
	var (
		msg  string
		err  error
		resp model.ReleaseInfo
		ID   int
	)
	act := model.ReleaseInfo{}

	resp, err = dao.GetActInfoByID(ID, &act)
	if err != nil {
		msg = "数据库查询失败"
		// z.Error(fmt.Sprintf("%s:%s", msg, err))
		response.Failed(c, http.StatusOK, msg, nil)
		return
	}
	response.Success(c, resp)
}
