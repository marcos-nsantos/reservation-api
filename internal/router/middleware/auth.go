package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/marcos-nsantos/reservation-api/internal/token"
)

func AuthMiddleware(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			return
		}

		tokenString, err := token.Extract(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header format"})
			return
		}

		parsedToken, err := jwt.ParseWithClaims(tokenString, &token.CustomClaims{}, func(token *jwt.Token) (any, error) {
			return []byte(key), nil
		})

		if err != nil || !parsedToken.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid parsedToken"})
			return
		}

		claims, ok := parsedToken.Claims.(*token.CustomClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid parsedToken claims"})
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
