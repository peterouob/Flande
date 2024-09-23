package user

import (
	"database/sql"
	"ecomm/protocol/user"
)

type Server struct {
	db *sql.DB
	user.UnimplementedUserServiceServer
}

func (s *Server) StartRpcService() {
	go s.StartCreateRpc()
	go s.StartLoginRpc()
	go s.StartGetAllUserRpc()
}
