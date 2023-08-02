package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"math"
)

// 根据token获取id，校验失败返回error，校验成功返回该登录用户的id
func GetIdFromToken(authToken string) (id int, err error) {
	token, err := ParseToken(authToken)
	if err != nil {
		return -1, err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	tokenId := claims["id"]
	if i, ok := tokenId.(float64); ok {
		return int(i), nil
	} else {
		return -1, err
	}
}

// 校验token的方法，校验成功返回nil，失败返回err
// 根据用户的id进行校验，失败返回err
func TokenVerify(authToken string, id int) error {
	token, err := ParseToken(authToken)
	if err != nil {
		return err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	tokenId := claims["id"]
	isEq := false
	if i, ok := tokenId.(float64); ok {
		if int(math.Floor(i)) == id {
			isEq = true
		} else {
			return errors.New("用户登陆错误")
		}
	}
	if isEq {
		return nil
	} else {
		return errors.New("用户校验失败")
	}
}
