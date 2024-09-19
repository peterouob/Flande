package utils

import (
	"ecomm/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Forbidden(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusForbidden, gin.H{
		"code": -1,
		"msg":  msg,
	})
	c.Abort()
}

func SetCookie(c *gin.Context, name, value string) {
	c.SetCookie(name, value, 365*3600, "/", config.Config.GetString("server.host"), false, true)
}

func RemoveCookie(c *gin.Context, key string) {
	c.SetCookie(key, "", -1, "", config.Config.GetString("server.host"), false, true)
}

func Cors(c *gin.Context) {
	method := c.Request.Method
	c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
	fmt.Println(c.GetHeader("Origin"))
	c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token")
	c.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
	c.Header("Access-Control-Allow-Credentials", "true")
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.Next()
}
