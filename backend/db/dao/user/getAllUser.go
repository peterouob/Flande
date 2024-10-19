package user

import (
	"context"
	"database/sql"
	"ecomm/db"
	user2 "ecomm/protocol/user"
	"errors"
	"fmt"
	"log"
)

func GetAllUser(ctx context.Context) ([]user2.GetAllUserResp, error) {
	var resp []user2.GetAllUserResp
	query := "SELECT * FROM users"

	if err := db.DB.SelectContext(ctx, &resp, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("no users found")
		}
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}

	log.Println(resp)
	return resp, nil
}

//func GetAllUserById(ctx context.Context, req user2.GetAllUserReq) (*user2.GetAllUserResp, error) {
//	resp := &user2.GetAllUserResp{}
//	query := "SELECT * FROM users WHERE id = ?"
//	if err := db.DB.Select(&resp, query); !errors.Is(err, sql.ErrNoRows) {
//		return nil, errors.New("error to find the user from id:" + string(req.Id))
//	}
//	return resp, nil
//}
