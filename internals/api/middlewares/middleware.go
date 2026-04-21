package middlewares

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// TODO: replace with JWT validation
		userID := c.GetHeader("X-User-ID")
		if userID == "" {
			userID = "dev-user"
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
