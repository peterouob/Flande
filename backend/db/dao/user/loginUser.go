package user

import (
	"context"
	"ecomm/db"
	user2 "ecomm/protocol/user"
	"errors"
)

func LoginUser(ctx context.Context, req user2.LoginUserReq) (*user2.LoginUserResp, error) {
	resp := &user2.LoginUserResp{}
	var id int64
	query := "select `uid` from `users` where name = ? and password = ?"
	err := db.DB.GetContext(ctx, &id, query, req.Name, req.Password)
	if err != nil {
		return nil, errors.New("error occurred during login :" + err.Error())
	}
	resp.Name = req.Name
	resp.Id = id
	return resp, nil
}
