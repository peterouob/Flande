package main

import (
	"ecomm/router"
	"ecomm/rpc"
	"github.com/gin-gonic/gin"
)

func main() {
	s := rpc.Server{}

	s.StartRpcService()
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8082")
}