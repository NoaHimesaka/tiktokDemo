package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

//todo
//jwt的秘钥
var Mykey = []byte("idontknow")

type MyStandardClaims struct {
	UserId int64 `json:"userid"`
	jwt.StandardClaims
}

//token生成
func GenerateToken(userid int64) (string, error) {
	claims := MyStandardClaims{
		UserId: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60*5, //过期时间设置为5个小时
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, err := token.SignedString(Mykey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

//获取token中的信息
func GetIdInToken(token string) int64 {
	claims, _ := jwt.ParseWithClaims(token, &MyStandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return Mykey, nil
		})
	return claims.Claims.(*MyStandardClaims).UserId

}

//token校验
func TokenCheck(token string) bool {
	//如果校验成功，返回true
	_, err := jwt.ParseWithClaims(token, &MyStandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return Mykey, nil
		})
	if err != nil {
		return false
	}
	return true
}
