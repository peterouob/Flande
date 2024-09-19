package user

import (
	"context"
	"database/sql"
	"ecomm/db"
	"ecomm/protocol/user"
	"errors"
	"fmt"
)

func CreateUser(ctx context.Context, req user.CreateUserReq) error {
	query := "SELECT `name` FROM users where `name` = ?"
	err := db.DB.Get(&user.LoginUserReq{}, query, req.Name)
	if errors.Is(err, sql.ErrNoRows) {
		query = "INSERT INTO `users` (uid,name,password,email,phone,sex) VALUES  (?,?,?,?,?,?)"
		_, err := db.DB.Exec(query, req.Uid, req.Name, req.Password, req.Email, req.Phone, req.Sex)
		if err != nil {
			return errors.New("error occur in create")
		}
	} else {
		return fmt.Errorf("have the same username")
	}
	return nil
}
