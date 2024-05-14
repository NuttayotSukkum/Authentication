package Configs

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"os"
	"time"
)

var hmacSampleSecret []byte

func Token(userId string) (string, error) {
	InitEnv()
	hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	})
	signedToken, err := token.SignedString(hmacSampleSecret)
	log.Println("secret Key:", string(hmacSampleSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func PareToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token")
	}
	id, ok := claims["userId"].(string)
	if !ok {
		return "", fmt.Errorf("invalid token")
	}
	return id, nil
}
