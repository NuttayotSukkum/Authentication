package Configs

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"os"
	"time"
)

var hmacSampleSecret []byte

func Token(userId uint) (string, error) {
	InitEnv()
	hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Minute * 1).Unix(),
	})
	signedToken, err := token.SignedString(hmacSampleSecret)
	log.Println("secret Key:", string(hmacSampleSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func PareToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token")
	}
	id := uint(claims["userId"].(float64))
	return id, nil
}
