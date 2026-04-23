package middlewares

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// TODO: replace with JWT validation
		userID := c.GetHeader("X-User-ID")
		if userID == "" {
			userID = "dev-user"
		}

		c.Set("user_id", "01952865-c71c-7708-8e6c-7e6d0800c149")
		c.Next()
	}
}
