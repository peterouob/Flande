package utils

import (
	"ecomm/config"
	"ecomm/token"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

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

type auth struct {
	id  interface{} `json:"id"`
	rid interface{} `json:"rid"`
}

func AuthByJWT() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"code": "-1",
				"msg:": "not have auth header",
			})
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "Format of Authorization is wrong",
			})
			c.Abort()
			return
		}
		_, err := token.VerifyToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "Verify of Authorization is wrong",
			})
			c.Abort()
			return
		}
		c.Next()
	}

}
