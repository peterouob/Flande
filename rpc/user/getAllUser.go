package user

import (
	"context"
	"ecomm/etcd"
	user2 "ecomm/protocol/user"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func (s *Server) StartGetAllUserRpc() error {
	const addr = "127.0.0.1:50052"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.New("error to listen port :50052")
	}
	grpcServer := grpc.NewServer()
	user2.RegisterUserServiceServer(grpcServer, &Server{})

	ctx := context.Background()
	if err := etcd.RegisterEtcd(ctx, "create", addr); err != nil {
		log.Fatalf("register %s failed %v", "create", err)
	}
	fmt.Printf("start grpc server:%s\n", addr)

	if err := grpcServer.Serve(lis); err != nil {
		return errors.New("error occur in grpc server when serve")
	}
	return nil
}
