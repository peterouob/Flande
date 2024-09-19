package user

import (
	"context"
	"ecomm/db/dao/user"
	"ecomm/etcd"
	"ecomm/kafka/producer"
	user2 "ecomm/protocol/user"
	"ecomm/service"
	"github.com/gin-gonic/gin"
	eclient "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"time"
)

func LoginUser(c *gin.Context) {
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
	var req user2.LoginUserReq
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
	conn, err := grpc.NewClient("etcd:///service/login", grpc.WithResolvers(builder), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"err:": err.Error(),
		})
		return
	}
	defer conn.Close()

	resp, err := user.LoginUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"err:": err.Error(),
		})
		return
	}
	err = producer.ProducerMessage(service.Producer, "user", "login user name="+resp.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":          -1,
			"producer err:": err.Error(),
		})
		return
	}

	accessToken, _, err := service.Token.CreateToken(resp.Id, resp.Name, 15*time.Minute)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":       -1,
			"token err:": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg:":  "login success",
		"data":  resp,
		"token": accessToken,
	})
}
