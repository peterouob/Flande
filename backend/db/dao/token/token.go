package token

import (
	"context"
	"ecomm/db"
	"errors"
	"log"
	"time"
)

func SaveTokenAuth(td *Token) error {
	at := time.Unix(td.AtExp, 0)
	now := time.Now()
	if err := db.Rdb.Set(context.Background(), td.AccessUUid, td.AccessToken, at.Sub(now)).Err(); err != nil {
		return err
	}
	return nil
}

func SaveRefreshToken(td *RefreshToken) error {
	rt := time.Unix(td.ReExp, 0)
	now := time.Now()
	if err := db.Rdb.Set(context.Background(), td.RefreshUUid, td.RefreshToken, rt.Sub(now)).Err(); err != nil {
		return err
	}
	return nil
}

func DeleteOldToken(auuid, ruuid interface{}) error {
	if err := db.Rdb.Del(context.Background(), auuid.(string)).Err(); err != nil {
		return errors.New("error occur in delete origin uuid :" + err.Error())
	}

	if err := db.Rdb.Del(context.Background(), ruuid.(string)).Err(); err != nil {
		return errors.New("error occur in delete refresh uuid :" + err.Error())
	}

	log.Println("Success delete old token set")
	return nil
}

func GetValueById(uuid, ruuid interface{}) (string, string, error) {
	if uuid.(string) == " " || ruuid.(string) == " " {
		return "", "", errors.New("uuid and ruuid are empty")
	}
	uid := db.Rdb.Get(context.Background(), uuid.(string)).String()
	rid := db.Rdb.Get(context.Background(), ruuid.(string)).String()
	return uid, rid, nil
}
