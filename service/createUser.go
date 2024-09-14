package service

import (
	"context"
	"ecomm/rpc/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var req user.CreateUserReq
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
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"err:": err.Error(),
		})
		return
	}
	defer conn.Close()
	grpcClient := user.NewUserServiceClient(conn)
	resp, err := grpcClient.CreateUserRpc(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"err:": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"resp": resp,
	})
}
