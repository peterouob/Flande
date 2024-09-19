package user

import (
	"context"
	"ecomm/etcd"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func (s *Server) StartLoginRpc() error {
	const addr = "127.0.0.1:50050"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.New("error to listen port :50050")
	}
	grpcServer := grpc.NewServer()
	ctx := context.Background()
	if err := etcd.RegisterEtcd(ctx, "login", addr); err != nil {
		log.Fatalf("register %s failed %v", "login", err)
	}
	fmt.Printf("start grpc server:%s\n", addr)

	if err := grpcServer.Serve(lis); err != nil {
		return errors.New("error occur in grpc server when serve")
	}
	return nil
}
