package service

import (
	"context"
	"ecomm/etcd"
	"ecomm/kafka/producer"
	user2 "ecomm/protocol/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	req.Uid = int64(uuid.New().ID()) >> 27
	conn, err := grpc.NewClient("etcd:///service/create", grpc.WithResolvers(builder), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"err:": err.Error(),
		})
		return
	}
	defer conn.Close()
	grpcClient := user2.NewUserServiceClient(conn)
	resp, err := grpcClient.CreateUserRpc(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"err:": err.Error(),
		})
		return
	}

	err = producer.ProducerMessage(p, "user", "register user")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":          -1,
			"producer err:": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"resp": resp,
	})
}
