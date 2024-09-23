package router

import (
	"ecomm/logger"
	"ecomm/service/user"
	"ecomm/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	l := logger.NewLogger()
	r.Use(logger.GinLogger(l), logger.GinRecovery(l, true))
	r.Use(utils.Cors)
	r.POST("/login", user.LoginUser)
	r.POST("/create", user.CreateUser)
	r.POST("/getAll", utils.AuthByJWT(), user.GetAllUser)
}
