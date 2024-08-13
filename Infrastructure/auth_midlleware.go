package infrastructure

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(isAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Header("Content-Type", "application/json")
			c.IndentedJSON(401, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.Header("Content-Type", "application/json")
			c.IndentedJSON(401, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		t := Token{}
		token, e := t.ValidateToken(authParts[1])

		if e != nil {
			c.Header("Content-Type", "application/json")
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": e.Error()})
			c.Abort()
			return
		}

		if isAdmin {
			allow := t.ValidateAdmin(token)
			if !allow {
				c.Header("Content-Type", "application/json")
				c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Forbidden"})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
