package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func JwtGeneration(name string, id int) string {
	//获取jwt对象并设置加密算法
	Jwt := jwt.New(jwt.SigningMethodHS256)
	//设置映射对象 键为string值任意
	var claims map[string]interface{}
	//配置成jwt的内置格式
	claims = Jwt.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["username"] = name
	//设置过期时间为一天
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	//设置签发时间
	claims["iat"] = time.Now().Unix()
	//设置签发主题
	claims["sub"] = "jwtTest"
	//设置密匙,字符串以二进制方式进入
	authCode, _ := Jwt.SignedString([]byte(SingingKey))
	return authCode
}
