package helper

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

var secret = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateToken(id uint) (string, error) {

	claims := jwt.MapClaims{
		"id": id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(secret)

	if err != nil {
		return "token bermasalah", err
	}

	return tokenStr, nil

}

func ValidateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, errors.New("error teroooos")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return jwt.MapClaims{}, errors.New("failed decode token")
	}

}
