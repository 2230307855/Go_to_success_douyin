package utils

import "github.com/dgrijalva/jwt-go"

func ParseToken(authtoken string) (*jwt.Token, error) {
	token, err := jwt.Parse(authtoken, func(token *jwt.Token) (interface{}, error) {
		return []byte(SingingKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
