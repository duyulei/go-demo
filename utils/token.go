package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	key []byte = []byte("romp.com")
)

// 校验token是否有效
func CheckToken(token string) bool {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return false
	}
	return true
}

// 产生json web token
func GenToken() string {
	claims := &jwt.StandardClaims{
		ExpiresAt: int64(time.Now().Unix() + 1000),
		Issuer:    "dulei",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		fmt.Println("生成失败", err)
		return ""
	}
	return ss
}
// Audience string json:"aud,omitempty"  aud 标识token的接收者.
// ExpiresAt int64 json:"exp,omitempty"  exp 过期时间.通常与Unix UTC时间做对比过期后token无效
// Id string json:"jti,omitempty" jti 是自定义的id号
// IssuedAt int64 json:"iat,omitempty"  iat 签名发行时间.
// Issuer string json:"iss,omitempty"  iss 是签名的发行者.
// NotBefore int64 json:"nbf,omitempty"  nbf 这条token信息生效时间.这个值可以不设置,但是设定后,一定要大于当前Unix UTC,否则token将会延迟生效.
// Subject string json:"sub,omitempty" sub 签名面向的用户

