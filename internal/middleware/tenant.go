package middleware

import (
	"appointment-api/internal/models"
	"appointment-api/internal/services"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func TenantMiddleware(tenantService services.TenantService, mainDB *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Domain'i al - Origin header'dan önce Host'tan
		domain := getDomainFromRequest(c)
		if domain == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "No valid domain found",
			})
			c.Abort()
			return
		}

		// Tenant'ı domain'e göre al
		tenant, err := tenantService.GetTenantByDomain(domain)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "Tenant not found for domain: " + domain,
			})
			c.Abort()
			return
		}

		// Schema'nın var olup olmadığını kontrol et, yoksa oluştur
		var schemaExists bool
		err = mainDB.QueryRow("SELECT EXISTS(SELECT 1 FROM information_schema.schemata WHERE schema_name = $1)", tenant.Schema).Scan(&schemaExists)
		if err == nil && !schemaExists {
			// Schema'yı oluştur
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

		// Tenant schema'ya geç
		_, err = mainDB.Exec(fmt.Sprintf("SET search_path TO %s, public", tenant.Schema))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to set tenant context",
			})
			c.Abort()
			return
		}

		// Context'e tenant bilgilerini kaydet
		c.Set("tenant", tenant)
		c.Set("tenant_schema", tenant.Schema)
		c.Set("tenant_domain", domain)

		c.Next()

		// İstek bittikten sonra search_path'i resetle
		mainDB.Exec("SET search_path TO public")
	}
}

// Domain'i request'ten al - Host header veya Origin header'dan
func getDomainFromRequest(c *gin.Context) string {
	// 1. Origin header'ından al (frontend CORS istekleri için)
	origin := c.GetHeader("Origin")
	if origin != "" {
		domain := normalizeDomain(origin)
		if domain != "" {
			return domain
		}
	}

	// 2. Host header'ından al
	host := c.Request.Host
	if host != "" {
		return normalizeDomain(host)
	}

	// 3. Referer header'ından al
	referer := c.GetHeader("Referer")
	if referer != "" {
		domain := normalizeDomain(referer)
		if domain != "" {
			return domain
		}
	}

	return ""
}

// Domain'i temizle ve normalize et
func normalizeDomain(domain string) string {
	// Protocol kaldır
	domain = strings.TrimPrefix(domain, "https://")
	domain = strings.TrimPrefix(domain, "http://")

	// www. prefix'i kaldır
	domain = strings.TrimPrefix(domain, "www.")

	// Path'i kaldır
	if idx := strings.Index(domain, "/"); idx != -1 {
		domain = domain[:idx]
	}

	return strings.ToLower(strings.TrimSpace(domain))
}

// Helper function to get current tenant from context
func GetCurrentTenant(c *gin.Context) (*models.TenantConfig, bool) {
	tenant, exists := c.Get("tenant")
	if !exists {
		return nil, false
	}
	return tenant.(*models.TenantConfig), true
}
