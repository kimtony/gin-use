package util

import (
	"errors"
	"gin-use/configs"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	InPayload map[string]interface{} `json:"inPayload"`
	jwt.StandardClaims
}

var (
	//jwt
	TokenExpireDuration = time.Hour * configs.Get().JWT.ExpireDuration //单位小时
	jwtSecret           = []byte(configs.Get().JWT.Secret)

	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

// GenToken 生成JWT
func GenToken(inPayload map[string]interface{}) (string, error) {

	// 创建一个我们自己的声明
	c := Claims{
		inPayload, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "linkai",                                   // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(jwtSecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (map[string]interface{}, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, TokenInvalid
		}
		return jwtSecret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.InPayload, nil
	}

	return nil, TokenExpired
}
