package middleware

import (
	"appointment-api/internal/models"
	"appointment-api/internal/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Authorization header required",
			})
			c.Abort()
			return
		}

		// Extract token from "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Invalid authorization format",
			})
			c.Abort()
			return
		}

		token := tokenParts[1]
		user, err := authService.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Invalid token",
			})
			c.Abort()
			return
		}

		// Store user in context
		c.Set("user", user)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "User not authenticated",
			})
			c.Abort()
			return
		}

		userModel := user.(*models.User)
		if userModel.Role != models.RoleAdmin {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"error":   "Admin access required",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Helper function to get current user from context
func GetCurrentUser(c *gin.Context) (*models.User, bool) {
	user, exists := c.Get("user")
	if !exists {
		return nil, false
	}
	return user.(*models.User), true
}
