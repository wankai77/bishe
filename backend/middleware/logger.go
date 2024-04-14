package middleware

import (
	"backend/core/log"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

var z *zap.Logger

func Logger() gin.HandlerFunc {
	logger := log.Log()
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime) / time.Microsecond
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		header := c.Request.Header
		proto := c.Request.Proto
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		err := c.Err()
		body, _ := ioutil.ReadAll(c.Request.Body)
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIp,
			"req_method":   reqMethod,
			"req_url":      reqUri,
			"header":       header,
			"proto":        proto,
			"err":          err,
			"body":         body,
		}).Info()
	}
}

//	func init() {
//		// 在 init 函数中初始化 z 变量
//		logger, err := zap.NewProduction()
//		if err != nil {
//			panic("failed to initialize logger: " + err.Error())
//		}
//		z = logger
//	}
func GetLogger() *zap.Logger {
	return z
}
