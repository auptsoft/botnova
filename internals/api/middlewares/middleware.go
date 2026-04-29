package middlewares

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev-secret-change-me"
	}

	return func(c *gin.Context) {
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"isSuccessful": false, "message": "Missing Authorization header"})
			return
		}

		headerParts := strings.SplitN(authHeader, " ", 2)
		if len(headerParts) != 2 || !strings.EqualFold(headerParts[0], "Bearer") {
			c.AbortWithStatusJSON(401, gin.H{"isSuccessful": false, "message": "Invalid Authorization header format"})
			return
		}

		tokenString := strings.TrimSpace(headerParts[1])
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"isSuccessful": false, "message": "Missing token"})
			return
		}

		claims := &jwt.RegisteredClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"isSuccessful": false, "message": "Invalid or expired token"})
			return
		}

		if claims.Subject == "" {
			c.AbortWithStatusJSON(401, gin.H{"isSuccessful": false, "message": "Invalid token subject"})
			return
		}

		c.Set("user_id", claims.Subject)
		c.Next()
	}
}
