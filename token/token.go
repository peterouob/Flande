package token

import (
	"ecomm/config"
	token2 "ecomm/db/dao/token"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	"time"
)

var err error
var (
	auuid chan interface{}
	ruuid chan interface{}
)

func init() {
	auuid = make(chan interface{}, 1024)
	ruuid = make(chan interface{}, 1024)
}

func CreateToken(id int64) (*token2.Token, *token2.RefreshToken, error) {
	tokenVal := config.Config.GetString("token.key")
	rtokenVal := config.Config.GetString("token.rkey")
	t, err := createToken(id, tokenVal)
	if err != nil {
		log.Println("Error int create token :" + err.Error())
		return nil, nil, err
	}
	rt, err := createRefreshToken(id, rtokenVal)
	if err != nil {
		log.Println("Error int create refresh token :" + err.Error())
		return nil, nil, err
	}
	go func() {
		err := refreshTokenRoutine(rt.RefreshToken, id, rtokenVal)
		if err != nil {
			log.Println("error in refresh token routine :" + err.Error())
		}
	}()
	return t, rt, nil
}
func refreshTokenRoutine(refreshToken string, userId int64, refreshTokenSecret string) error {
	for {
		token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(refreshTokenSecret), nil
		})
		if err != nil {
			return errors.New("error parsing refresh token :" + err.Error())
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp := int64(claims["exp"].(float64))
			sleepDuration := time.Duration(exp-time.Now().Unix()-60) * time.Second
			if sleepDuration > 0 {
				time.Sleep(sleepDuration)
			}
			uid := <-auuid
			rid := <-ruuid
			if err := token2.DeleteOldToken(uid, rid); err != nil {
				return err
			}
			newToken, newRtoken, err := CreateToken(userId)
			if err != nil {
				return errors.New("create new token error:" + err.Error())
			}
			refreshToken = newRtoken.RefreshToken
			log.Println("refresh token refreshed successfully")

			saveErr := token2.SaveTokenAuth(newToken)
			if saveErr != nil {
				return errors.New("error save token error:" + err.Error())
			}
			log.Println("new tokens saved to redis successfully")
		} else {
			return errors.New("error parsing refresh token :" + err.Error())
		}
	}
}
func createToken(id int64, value string) (*token2.Token, error) {
	t := &token2.Token{}
	t.AccessUUid = uuid.NewString()
	t.AtExp = time.Now().Add(time.Minute * 15).Unix()
	claim := jwt.MapClaims{}
	claim["authorized"] = true
	claim["access_uuid"] = t.AccessUUid
	claim["user_id"] = id
	claim["exp"] = t.AtExp
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t.AccessToken, err = tk.SignedString([]byte(value))
	if err != nil {
		fmt.Println("sign token error: ", err)
		return nil, err
	}
	auuid <- claim["access_uuid"]
	err := token2.SaveTokenAuth(t)
	if err != nil {
		return nil, errors.New("error in save token :" + err.Error())
	}
	return t, nil
}
func createRefreshToken(id int64, value string) (*token2.RefreshToken, error) {
	t := &token2.RefreshToken{}
	t.RefreshUUid = uuid.NewString()
	t.ReExp = time.Now().Add(time.Minute * 30).Unix()
	rclaim := jwt.MapClaims{}
	rclaim["authorized"] = true
	rclaim["refresh_uuid"] = t.RefreshUUid
	rclaim["user_id"] = id
	rclaim["exp"] = time.Now().Add(time.Minute * 30).Unix()
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rclaim)
	t.RefreshToken, err = rt.SignedString([]byte(value))
	if err != nil {
		return nil, errors.New("sign refresh token error:" + err.Error())
	}
	ruuid <- rclaim["refresh_uuid"]
	err := token2.SaveRefreshToken(t)
	if err != nil {
		return nil, errors.New("error in save token :" + err.Error())
	}
	return t, nil
}

// VerifyToken 檢查token簽名方法
func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %v", tk.Header["alg"])
		}
		return []byte(config.Config.GetString("token.rkey")), nil
	})
	if err != nil {
		return nil, errors.New("error in parse token")
	}
	return token, nil
}
