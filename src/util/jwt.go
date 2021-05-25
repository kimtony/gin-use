package util

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

type Claims struct {
	InPayload map[string]interface{} `json:"inPayload"`
	jwt.StandardClaims
}

const (
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin    = "请重新登陆"
)

const TokenExpireDuration = time.Hour * 2

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// GenToken 生成JWT
func GenToken(inPayload map[string]interface{}) (string, error) {
	println("--------------------------", TokenExpireDuration)

	// 创建一个我们自己的声明
	c := Claims{
		inPayload, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "yuezhi",                                   // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(jwtSecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { // 校验token
		println("-------claims---------", claims)

		return claims, nil
	}
	return nil, errors.New("invalid token")
}

//TODO:能返回解析token的内容 还需要看看怎么优化
func VerifyAction(strToken string) (g.Map, error) {

	// //用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	// tokenClaims, err := jwt.ParseWithClaims(strToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
	// 	return jwtSecret, nil
	// })

	// if tokenClaims != nil {
	// 	// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
	// 	// 要传入指针，项目中结构体都是用指针传递，节省空间。
	// 	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
	// 		return claims, nil
	// 	}
	// }
	// return nil, err

	parseAuth, err := jwt.Parse(strToken, func(*jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		// return params, errors.New("TOKEN ERROR!")
		fmt.Println("TOKEN ERROR")
	}
	//将token中的内容存入parmMap
	claim := parseAuth.Claims.(jwt.MapClaims)
	var parmMap map[string]interface{}
	parmMap = make(map[string]interface{})
	for key, val := range claim {
		parmMap[key] = val
	}
	fmt.Println(gconv.Map(parmMap))
	
	return gconv.Map(parmMap), nil

}
