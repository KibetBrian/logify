package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)


func Stringer(i interface{}) string {
	return fmt.Sprintf("%v", i)
}

var SecretKey = Stringer(viper.Get("JWT_SECRET_KEY"))

func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
	})

	return claims.SignedString([]byte(SecretKey))
}

func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims.Issuer, nil
}
