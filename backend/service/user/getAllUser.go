package user

import (
	"context"
	"ecomm/db/dao/user"
	"ecomm/kafka/producer"
	"ecomm/service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

func GetAllUser(c *gin.Context) {
	builder, err := service.InitBuilder()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"err":  err.Error(),
		})
		return
	}

	conn, err := grpc.NewClient("etcd:///service/login", grpc.WithResolvers(builder), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"err:": err.Error(),
		})
		return
	}
	defer conn.Close()
	resp, err := user.GetAllUser(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"err:": err.Error(),
		})
		return
	}
	err = producer.ProducerMessage(service.Producer, "admin", "get all user")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":          -1,
			"producer err:": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg:": "get data success",
		"data": resp,
	})
}
