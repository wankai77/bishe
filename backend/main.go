package main

import (
	"backend/config/config"
	"backend/core/db"
	"backend/core/log"
	"backend/router"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	db.SetupDBLink()
	log.Init()
	log := log.Log()
	gin.SetMode(config.Config.Server.Model)
	router := router.InitRouter()
	srv := &http.Server{
		Addr:    config.Config.Server.Address,
		Handler: router,
	}
	//启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("listen:%s \n", err)
		}
		log.Printf("listen:%s \n", config.Config.Server.Address)
	}()
	quit := make(chan os.Signal)
	//监听消息
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Info("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}
