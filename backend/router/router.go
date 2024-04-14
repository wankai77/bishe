package router

import (
	"backend/middleware"
	"backend/service"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	//跌机恢复
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())

	// router.Use(middleware.Logger())
	register(router)
	return router
}

// 路由注册
func register(router *gin.Engine) {
	//todo 后续接口url
	userRouters1 := router.Group("/api/user")
	userRouters1.POST("/register", service.Register)
	userRouters1.POST("/login", service.Login)
	userRouters2 := router.Group("/api/home").Use(middleware.AuthUser())
	userRouters2.GET("/getInfo", service.GetInfo)
	userRouters3 := router.Group("/api/personal").Use(middleware.AuthUser())
	userRouters3.POST("/resetPassword", service.ResetPassword)
	userRouters3.POST("/release", service.ReleaseActivity)
	userRouters4 := router.Group("/api/audit").Use(middleware.AuthUser())
	userRouters4.GET("/getrelinfo", service.GetRelInfo)
	userRouters4.POST("/auditrelinfo", service.AuditRelInfo)
	userRouters5 := router.Group("/api/msg").Use(middleware.AuthUser())
	userRouters5.GET("/getmsg", service.GetMessage)
	userRouters6 := router.Group("/api/detail").Use(middleware.AuthUser())
	userRouters6.GET("/getactinfo", service.GetActInfo)
}
