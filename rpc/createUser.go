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

func (s *Server) StartCreateRpc() error {
	const addr = "127.0.0.1:50051"
	mysql := db.ConnectMysql()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.New("error to listen port :50051")
	}
	grpcServer := grpc.NewServer()
	user2.RegisterUserServiceServer(grpcServer, &Server{db: mysql})

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

func (s *Server) CreateUserRpc(ctx context.Context, req *user2.CreateUserReq) (*user2.CreateUserResp, error) {
	query := "INSERT INTO `users` (uid,name,password,email,phone,sex) VALUES  (?,?,?,?,?,?)"
	if _, err := s.db.Exec(query, req.Uid, req.Name, req.Password, req.Email, req.Phone, req.Sex); err != nil {
		return nil, errors.New("db error occur in insert user statement:" + err.Error())
	}
	return &user2.CreateUserResp{Name: req.Name}, nil
}
