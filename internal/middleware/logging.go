package middleware

import (
	"appointment-api/internal/models"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLogger middleware that logs all incoming requests with tenant info
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Get request info
		method := c.Request.Method
		path := c.Request.URL.Path
		rawQuery := c.Request.URL.RawQuery
		host := c.Request.Host
		clientIP := c.ClientIP()

		// Build full path with query
		fullPath := path
		if rawQuery != "" {
			fullPath = path + "?" + rawQuery
		}

		// Log request start (without tenant/user info yet)
		fmt.Printf("🌐 [%s] %s %s | Host: %s",
			start.Format("15:04:05"),
			method,
			fullPath,
			host)

		// Process request
		c.Next()

		// Get tenant info after middleware processing
		tenant := "unknown"
		if tenantValue, exists := c.Get("tenant"); exists {
			if tenantConfig, ok := tenantValue.(*models.TenantConfig); ok {
				tenant = tenantConfig.Name
			}
		}

		// Get user info after middleware processing
		userInfo := "anonymous"
		if user, exists := c.Get("current_user"); exists {
			if userMap, ok := user.(map[string]interface{}); ok {
				if email, ok := userMap["email"].(string); ok {
					userInfo = email
				}
			}
		}

		// Calculate latency
		latency := time.Since(start)
		statusCode := c.Writer.Status()

		// Choose emoji based on status code
		var statusEmoji string
		switch {
		case statusCode >= 200 && statusCode < 300:
			statusEmoji = "✅"
		case statusCode >= 300 && statusCode < 400:
			statusEmoji = "🔄"
		case statusCode >= 400 && statusCode < 500:
			statusEmoji = "❌"
		case statusCode >= 500:
			statusEmoji = "💥"
		default:
			statusEmoji = "❓"
		}

		// Log response with tenant and user info
		fmt.Printf(" | Tenant: %s | User: %s | IP: %s\n%s [%s] %d | %v | %s %s\n",
			tenant,
			userInfo,
			clientIP,
			statusEmoji,
			time.Now().Format("15:04:05"),
			statusCode,
			latency,
			method,
			fullPath)

		// Add separator for errors or slow requests
		if statusCode >= 400 || latency > 1*time.Second {
			fmt.Println("   " + strings.Repeat("─", 50))
		}
	}
}

// TenantLogger specifically logs tenant resolution
func TenantLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Request.Host

		// Log tenant resolution attempt
		fmt.Printf("🏢 Resolving tenant for host: %s\n", host)

		c.Next()

		// Log resolved tenant info
		if tenantValue, exists := c.Get("tenant"); exists {
			if tenantMap, ok := tenantValue.(map[string]interface{}); ok {
				name := tenantMap["name"].(string)
				schema := tenantMap["schema"].(string)
				fmt.Printf("   ✅ Resolved to: %s (schema: %s)\n", name, schema)
			}
		} else {
			fmt.Printf("   ❌ Failed to resolve tenant\n")
		}
	}
}
