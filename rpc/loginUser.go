package rpc

import (
	"context"
	"ecomm/db"
	"ecomm/etcd"
	user2 "ecomm/protocol/user"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func (s *Server) StartLoginRpc() error {
	const addr = "127.0.0.1:50050"
	mysql := db.ConnectMysql()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.New("error to listen port :50050")
	}
	grpcServer := grpc.NewServer()
	user2.RegisterUserServiceServer(grpcServer, &Server{db: mysql})

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

func (s *Server) LoginUserRpc(ctx context.Context, req *user2.LoginUserReq) (*user2.LoginUserResp, error) {
	if err := s.db.QueryRow("select `name` from `users` where name = ? and password = ?", req.Name, req.Password).Err(); err != nil {
		return nil, errors.New("db error occur in select database")
	}
	return &user2.LoginUserResp{Name: req.Name}, nil
}
