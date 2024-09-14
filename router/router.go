package router

import (
	"ecomm/service"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.POST("/create", service.CreateUser)
}
