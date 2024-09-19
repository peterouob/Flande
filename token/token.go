package token

import (
	"ecomm/config"
	token2 "ecomm/db/dao/token"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	"time"
)

var err error

func CreateToken(id int64) (*token2.Token, error) {
	t := &token2.Token{}
	t.AccessUUid = uuid.NewString()
	t.RefreshUUid = uuid.NewString()
	t.AtExp = time.Now().Add(time.Minute * 15).Unix()
	t.ReExp = time.Now().Add(time.Hour * 24).Unix()
	tokenVal := config.Config.GetString("token.key")
	claim := jwt.MapClaims{}
	claim["authorized"] = true
	claim["access_uuid"] = t.AccessUUid
	claim["user_id"] = id
	claim["exp"] = t.AtExp
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t.AccessToken, err = tk.SignedString([]byte(tokenVal))
	if err != nil {
		fmt.Println("sign token error: ", err)
		return nil, err
	}

	rtokenVal := config.Config.GetString("token.rkey")
	rclaim := jwt.MapClaims{}
	rclaim["authorized"] = true
	rclaim["refresh_uuid"] = t.RefreshUUid
	rclaim["user_id"] = id
	rclaim["exp"] = time.Now().Add(time.Minute * 15).Unix()
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t.RefreshToken, err = rt.SignedString([]byte(rtokenVal))
	if err != nil {
		fmt.Println("sign token error: ", err)
		return nil, err
	}

	go refreshTokenRoutine(t.RefreshToken, id, rtokenVal)

	return t, nil
}

func refreshTokenRoutine(refreshToken string, userId int64, refreshTokenSecret string) {
	for {
		token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(refreshTokenSecret), nil
		})
		if err != nil {
			fmt.Println("error parsing refresh token:", err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp := int64(claims["exp"].(float64))
			sleepDuration := time.Duration(exp-time.Now().Unix()-60) * time.Second
			if sleepDuration > 0 {
				time.Sleep(sleepDuration)
			}
			newToken, err := CreateToken(userId)
			if err != nil {
				fmt.Println("error creating new token:", err)
				return
			}
			refreshToken = newToken.RefreshToken
			log.Println("refresh token refreshed successfully")

			saveErr := token2.SaveTokenAuth(newToken)
			if saveErr != nil {
				fmt.Println("error saving token to redis:", saveErr)
				return
			}
			log.Println("new tokens saved to redis successfully")
		} else {
			log.Println("invalid refresh token")
			return
		}
	}
}
