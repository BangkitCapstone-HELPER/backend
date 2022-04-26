package lib

import (
	"fmt"
	"time"

	"github.com/BangkitCapstone-HELPER/backend/internal/app/config"
	"github.com/golang-jwt/jwt"
	"go.uber.org/fx"
)

type JWT interface {
	Encode(payload map[string]interface{}) (string, error)
	Decode(encrypted string) (map[string]interface{}, error)
}

type JwtParams struct {
	fx.In

	config.JWTConfig
}

func NewJWT(params JwtParams) JWT {
	return &params
}

func (j JwtParams) Encode(payload map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{}
	for k, v := range payload {
		claims[k] = v
	}
	claims["exp"] = time.Now().Add(j.Exp()).Unix()
	fmt.Println(claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.Secret()))
	return tokenString, err
}

func (j JwtParams) Decode(encrypted string) (map[string]interface{}, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.Secret()), nil
	}

	token, err := jwt.ParseWithClaims(encrypted, &jwt.MapClaims{}, keyFunc)
	if err != nil {
		return nil, err
	}

	claims, _ := token.Claims.(*jwt.MapClaims)
	return *claims, nil
}
