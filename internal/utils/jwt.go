package utils

import (
	constErrors "expression-backend/internal/errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const hmacSecret = "abcdefghijklmnopqrstuvwxyz"

func GenerateToken(login string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"nbf":   now.Unix(),
		"exp":   now.Add(time.Hour * 24).Unix(),
		"iat":   now.Unix(),
	})

	tokenString, err := token.SignedString([]byte(hmacSecret))
	if err != nil {
		return "", constErrors.FailedTokenGeneration
	}

	return tokenString, nil
}

func ValidateToken(token string) (string, error) {
	tokenFromString, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, constErrors.UnexpectedTokenMethod
		}

		return []byte(hmacSecret), nil
	})

	if err != nil {
		return "", constErrors.InvalidToken
	}

	if claims, ok := tokenFromString.Claims.(jwt.MapClaims); ok {
		return (claims["login"]).(string), nil
	} else {
		return "", constErrors.InvalidToken
	}
}
