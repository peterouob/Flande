package token

import (
	"context"
	"ecomm/db"
	"time"
)

func SaveTokenAuth(td *Token) error {
	at := time.Unix(td.AtExp, 0)
	rt := time.Unix(td.ReExp, 0)
	now := time.Now()
	if err := db.Rdb.Set(context.Background(), td.AccessUUid, td.AccessToken, at.Sub(now)).Err(); err != nil {
		return err
	}
	if err := db.Rdb.Set(context.Background(), td.RefreshUUid, td.RefreshToken, rt.Sub(now)).Err(); err != nil {
		return err
	}
	return nil
}
