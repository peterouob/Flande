package router

import (
	"ecomm/logger"
	"ecomm/service/user"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	l := logger.NewLogger()
	r.Use(logger.GinLogger(l), logger.GinRecovery(l, true))
	r.POST("/login", user.LoginUser)
	r.POST("/create", user.CreateUser)
}
