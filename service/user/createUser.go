package user

import (
	"context"
	"ecomm/db/dao/user"
	"ecomm/etcd"
	"ecomm/kafka/producer"
	user2 "ecomm/protocol/user"
	"ecomm/service"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	eclient "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

func CreateUser(c *gin.Context) {

	cli, err := eclient.NewFromURL(etcd.EtcdAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg:": "error occur in etcd new from url",
			"err:": err.Error(),
		})
		return
	}
	builder, err := resolver.NewBuilder(cli)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg:": "error occur in etcd new builder",
			"err:": err.Error(),
		})
		return
	}

	var req user2.CreateUserReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"err:": err.Error(),
		})
		return
	}
	if req.Name == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"err:": "please enter full information",
		})
		return
	}
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Uid = node.Generate().Int64()
	conn, err := grpc.NewClient("etcd:///service/create", grpc.WithResolvers(builder), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"err:": err.Error(),
		})
		return
	}
	defer conn.Close()

	err = user.CreateUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"err:": err.Error(),
		})
		return
	}
	err = producer.ProducerMessage(service.Producer, "user", "register user")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":          -1,
			"producer err:": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"id":   req.Uid,
	})
}
