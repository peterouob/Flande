package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Token struct {
	secretKey string
}
type TokenInterface interface{}

var _ TokenInterface = (*Token)(nil)

func NewToken(secretKey string) *Token {
	return &Token{
		secretKey: secretKey,
	}
}
func (t *Token) CreateToken(id int64, name string, duration time.Duration) (string, *UserClaims, error) {
	claims, err := NewUserClaims(id, name, duration)
	if err != nil {
		return "", nil, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(t.secretKey))
	if err != nil {
		return "", nil, errors.New("error in signed string: " + err.Error())
	}
	return tokenStr, claims, nil
}

func (t *Token) VerifyToken(tokenStr string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("error occur in jwt signing method")
		}
		return []byte(t.secretKey), nil
	})
	if err != nil {
		return nil, errors.New("error occur in jwt parse with claims:" + err.Error())
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.New("invalid token claims :" + err.Error())
	}
	return claims, nil
}
