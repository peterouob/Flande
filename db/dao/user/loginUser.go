package user

import (
	"context"
	"ecomm/db"
	user2 "ecomm/protocol/user"
	"errors"
)

func LoginUser(ctx context.Context, req user2.LoginUserReq) error {
	query := "select `name` from `users` where name = ? and password = ?"
	if err := db.DB.Get(&req, query, req.Name, req.Password); err != nil {
		return errors.New("error occur in login user")
	}
	return nil
}
