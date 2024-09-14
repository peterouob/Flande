package rpc

import (
	"context"
	"database/sql"
	"ecomm/db"
	"ecomm/rpc/user"
	"errors"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	db *sql.DB
	user.UnimplementedUserServiceServer
}

func (s *Server) GetAllUserRpc(ctx context.Context, req *user.GetAllUserReq) (*user.GetAllUserResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) CreateUserRpc(ctx context.Context, req *user.CreateUserReq) (*user.CreateUserResp, error) {
	query := "INSERT INTO `users` (uid,name,password,email,phone,sex) VALUES  (?,?,?,?,?,?)"
	if _, err := s.db.Exec(query, req.Uid, req.Name, req.Password, req.Email, req.Phone, req.Sex); err != nil {
		return nil, errors.New("db error occur in insert user statement:" + err.Error())
	}
	return &user.CreateUserResp{Name: req.Name}, nil
}

func (s *Server) StartCreateRpc() error {
	mysql := db.ConnectMysql()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return errors.New("error to listen port :50051")
	}
	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &Server{db: mysql})
	if err := grpcServer.Serve(lis); err != nil {
		return errors.New("error occur in grpc server when serve")
	}
	log.Println("Start create rpc ...")
	return nil
}
