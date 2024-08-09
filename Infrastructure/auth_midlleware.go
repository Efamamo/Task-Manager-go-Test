package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(isAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token, e := ValidateToken(authHeader)

		if e != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": e})
			c.Abort()
			return
		}
		if isAdmin {
			allow := ValidateAdmin(token)

			if !allow {
				c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": e})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
