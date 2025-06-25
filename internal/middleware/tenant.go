package middleware

import (
	"appointment-api/internal/models"
	"appointment-api/internal/services"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TenantMiddleware(tenantService services.TenantService, mainDB *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Request.Host

		// Get tenant config by host
		tenant, err := tenantService.GetTenantByHost(host)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "Tenant not found",
			})
			c.Abort()
			return
		}

		// Check if schema exists, create if not
		var schemaExists bool
		err = mainDB.QueryRow("SELECT EXISTS(SELECT 1 FROM information_schema.schemata WHERE schema_name = $1)", tenant.Schema).Scan(&schemaExists)
		if err == nil && !schemaExists {
			// Create schema if it doesn't exist
			err = tenantService.CreateTenantSchema(tenant.Schema)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"error":   "Failed to create tenant schema: " + err.Error(),
				})
				c.Abort()
				return
			}
		}

		// Set search_path to tenant schema
		_, err = mainDB.Exec(fmt.Sprintf("SET search_path TO %s, public", tenant.Schema))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to set tenant context",
			})
			c.Abort()
			return
		}

		// Store tenant in context
		c.Set("tenant", tenant)
		c.Set("tenant_schema", tenant.Schema)
		c.Next()

		// Reset search_path after request
		mainDB.Exec("SET search_path TO public")
	}
}

// Helper function to get current tenant from context
func GetCurrentTenant(c *gin.Context) (*models.TenantConfig, bool) {
	tenant, exists := c.Get("tenant")
	if !exists {
		return nil, false
	}
	return tenant.(*models.TenantConfig), true
}
