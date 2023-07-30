package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

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
		if int(i) == id {
			isEq = true
		}
	}
	if isEq {
		return nil
	} else {
		return errors.New("用户校验失败")
	}
}
