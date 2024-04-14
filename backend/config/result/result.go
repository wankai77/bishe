package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `josn:"data"`
}

func (r *Result) Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := Result{}
	res.Code = int(ApiCode.SUCCESS)
	res.Message = ApiCode.GetMessage(ApiCode.SUCCESS)
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func (r *Result) Failed(c *gin.Context, code int, message string, data interface{}) {
	res := Result{}
	res.Code = code
	res.Message = message
	res.Data = gin.H{}
	c.JSON(http.StatusOK, res)
}
