package token

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(key string, ID uint64) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": ID,
		"exp":    time.Now().Add(time.Hour * 6).Unix(),
		"iat":    time.Now().Unix(),
	})

	return claims.SignedString([]byte(key))
}

func Extract(header string) (string, error) {
	fields := strings.Fields(header)
	if len(fields) != 2 {
		return "", fmt.Errorf("invalid Authorization header format")
	}

	prefix, token := fields[0], fields[1]
	if strings.ToLower(prefix) != "bearer" {
		return "", fmt.Errorf("authorization scheme not supported")
	}

	return token, nil
}

func GetUserID(tokenString, secretKey string) (uint64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, fmt.Errorf("could not parse provided token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	claimUserID, ok := claims["userID"]
	if !ok {
		return 0, fmt.Errorf("could not parse user ID from token")
	}

	ID, err := strconv.ParseUint(fmt.Sprintf("%v", claimUserID), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse user ID from token")
	}

	return ID, nil
}
