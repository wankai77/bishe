package middleware

import (
	"backend/config/result"
	"backend/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	response *result.Result
	j        *utils.JWT
)

// AuthUser 用户合法身份校验
func AuthUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.Abort()
			msg := "获取 token 失败"
			z.Error(msg)
			response.Failed(ctx, http.StatusUnauthorized, msg, nil)
			return
		}
		payload, err := j.TokenParse(token)
		if err != nil {
			ctx.Abort()
			msg := "token 不合法"
			// z.Error(fmt.Sprintf("%s:%s", msg, err))
			response.Failed(ctx, http.StatusUnauthorized, msg, nil)
			return
		}
		// z.Info(fmt.Sprintf("请求接口的用户ID：%d", payload.UserID))

		ctx.Set("userID", payload.UserID)
		token, err = j.TokenNew(payload.UserID, payload.Username) //更新token
		if err != nil {
			ctx.Abort()
			msg := "更新 token 失败"
			z.Error(fmt.Sprintf("%s:%s", msg, err))
			response.Failed(ctx, http.StatusInternalServerError, msg, nil)
			return
		}
		ctx.Writer.Header().Set("Authorization", token)
		ctx.Next()

	}
}
