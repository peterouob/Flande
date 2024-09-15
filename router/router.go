package router

import (
	"ecomm/logger"
	"ecomm/service"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	l := logger.NewLogger()
	r.Use(logger.GinLogger(l), logger.GinRecovery(l, true))
	r.POST("/login", service.LoginUser)
	r.POST("/create", service.CreateUser)
}
