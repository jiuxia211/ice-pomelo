package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jiuxia211/ice-pomelo/pkg/constants"
)

type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

type ApiClaims struct {
	jwt.StandardClaims
}

func CreateToken(userId int64) (string, error) {
	expireTime := time.Now().Add(24 * 7 * time.Hour) // 过期时间为7天
	nowTime := time.Now()                            // 当前时间
	claims := Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间戳
			IssuedAt:  nowTime.Unix(),    // 当前时间戳
			Issuer:    "ice_pomelo",      // 颁发者签名
		},
	}
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenStruct.SignedString([]byte(constants.JWTValue))
}

func CheckToken(token string) (*Claims, error) {
	response, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JWTValue), nil
	})

	if err != nil {
		return nil, err
	}

	if resp, ok := response.Claims.(*Claims); ok && response.Valid {
		return resp, nil
	}

	return nil, err
}

func CreateApiToken() (string, error) {
	expireTime := time.Now().Add(time.Minute) // 过期时间为1分钟
	nowTime := time.Now()                     // 当前时间
	apiClaims := ApiClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间戳
			IssuedAt:  nowTime.Unix(),    // 当前时间戳
			Issuer:    "ice_pomelo",      // 颁发者签名
		},
	}
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, apiClaims)
	return tokenStruct.SignedString([]byte(constants.JWTValue))
}

func CheckApiToken(token string) (*Claims, error) {
	response, err := jwt.ParseWithClaims(token, &ApiClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JWTValue), nil
	})

	if err != nil {
		return nil, err
	}

	if resp, ok := response.Claims.(*Claims); ok && response.Valid {
		return resp, nil
	}

	return nil, err
}
