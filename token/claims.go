package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type UserClaims struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func NewUserClaims(id int64, name string, duration time.Duration) (*UserClaims, error) {
	tokeId, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.New("error occur in generator uuid")
	}
	return &UserClaims{
		Id:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokeId.String(),
			Subject:   name,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		}}, nil
}
