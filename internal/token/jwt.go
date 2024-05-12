package token

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	UserID uint64 `json:"userID"`
}

func GenerateJWT(key string, ID uint64) (string, error) {
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 6)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		UserID: ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
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
