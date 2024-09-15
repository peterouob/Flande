package rpc

import (
	"database/sql"
	"ecomm/userRpc/user"
)

type Server struct {
	db *sql.DB
	user.UnimplementedUserServiceServer
}
