package services

import (
	"appointment-api/internal/models"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
)

// TenantInfo cache için optimize edilmiş tenant bilgisi
type TenantInfo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Schema string `json:"schema"`
	// DB connection bilgileri gerekirse buraya eklenebilir
}

// TenantCache thread-safe tenant cache yapısı
type TenantCache struct {
	mu       sync.RWMutex
	tenants  map[string]*TenantInfo // domain -> TenantInfo
	db       *sql.DB
	stopCh   chan struct{}
	interval time.Duration
}

// TenantCacheService interface
type TenantCacheService interface {
	Start() error
	Stop()
	GetTenantByDomain(domain string) (*TenantInfo, error)
	RefreshCache() error
	GetCacheStats() (int, []string)
}

// NewTenantCache yeni bir tenant cache oluşturur
func NewTenantCache(db *sql.DB, refreshInterval time.Duration) TenantCacheService {
	if refreshInterval <= 0 {
		refreshInterval = 5 * time.Minute // Default 5 dakika
	}

	return &TenantCache{
		tenants:  make(map[string]*TenantInfo),
		db:       db,
		stopCh:   make(chan struct{}),
		interval: refreshInterval,
	}
}

// Start cache'i başlatır ve tüm tenantları yükler
func (tc *TenantCache) Start() error {
	log.Println("🚀 Starting tenant cache...")

	// İlk yükleme
	if err := tc.RefreshCache(); err != nil {
		return fmt.Errorf("failed to load initial tenant cache: %w", err)
	}

	// Goroutine ile periodik yenileme
	go tc.periodicRefresh()

	log.Printf("✅ Tenant cache started with %d tenants, refresh interval: %v", tc.getCacheSize(), tc.interval)
	return nil
}

// Stop cache'i durdurur
func (tc *TenantCache) Stop() {
	log.Println("🛑 Stopping tenant cache...")
	close(tc.stopCh)
}

// GetTenantByDomain domain'e göre tenant bilgisi döner (O(1) erişim)
func (tc *TenantCache) GetTenantByDomain(domain string) (*TenantInfo, error) {
	// Önce cache'ten bak
	tc.mu.RLock()
	if tenant, exists := tc.tenants[domain]; exists {
		tc.mu.RUnlock()
		return tenant, nil
	}
	tc.mu.RUnlock()

	// Cache'te yoksa DB'den çek
	log.Printf("🔍 Cache miss for domain: %s, querying database...", domain)
	tenant, err := tc.fetchTenantFromDB(domain)
	if err != nil {
		return nil, err
	}

	if tenant == nil {
		return nil, fmt.Errorf("tenant not found for domain: %s", domain)
	}

	// Cache'e ekle
	tc.addToCache(tenant)
	log.Printf("✅ Added tenant %s to cache", domain)

	return tenant, nil
}

// RefreshCache tüm cache'i yeniler
func (tc *TenantCache) RefreshCache() error {
	log.Println("🔄 Refreshing tenant cache...")

	tenants, err := tc.fetchAllTenantsFromDB()
	if err != nil {
		return fmt.Errorf("failed to fetch tenants from database: %w", err)
	}

	tc.mu.Lock()
	defer tc.mu.Unlock()

	// Eski cache'i temizle
	tc.tenants = make(map[string]*TenantInfo)

	// Yeni tenantları ekle
	for _, tenant := range tenants {
		tc.tenants[tenant.Domain] = tenant
	}

	log.Printf("✅ Cache refreshed with %d tenants", len(tenants))
	return nil
}

// GetCacheStats cache istatistiklerini döner
func (tc *TenantCache) GetCacheStats() (int, []string) {
	tc.mu.RLock()
	defer tc.mu.RUnlock()

	count := len(tc.tenants)
	domains := make([]string, 0, count)

	for domain := range tc.tenants {
		domains = append(domains, domain)
	}

	return count, domains
}

// periodicRefresh belirli aralıklarla cache'i yeniler
func (tc *TenantCache) periodicRefresh() {
	ticker := time.NewTicker(tc.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := tc.RefreshCache(); err != nil {
				log.Printf("❌ Failed to refresh tenant cache: %v", err)
			}
		case <-tc.stopCh:
			log.Println("🔄 Stopped periodic tenant cache refresh")
			return
		}
	}
}

// addToCache thread-safe olarak cache'e tenant ekler
func (tc *TenantCache) addToCache(tenant *TenantInfo) {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.tenants[tenant.Domain] = tenant
}

// getCacheSize thread-safe olarak cache boyutunu döner
func (tc *TenantCache) getCacheSize() int {
	tc.mu.RLock()
	defer tc.mu.RUnlock()
	return len(tc.tenants)
}

// fetchTenantFromDB DB'den tek bir tenant çeker
func (tc *TenantCache) fetchTenantFromDB(domain string) (*TenantInfo, error) {
	query := `
		SELECT id, name, domain, schema_name 
		FROM public.tenants 
		WHERE domain = $1 AND active = true`

	var tenant TenantInfo
	err := tc.db.QueryRow(query, domain).Scan(
		&tenant.ID,
		&tenant.Name,
		&tenant.Domain,
		&tenant.Schema,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &tenant, nil
}

// fetchAllTenantsFromDB DB'den tüm active tenantları çeker
func (tc *TenantCache) fetchAllTenantsFromDB() ([]*TenantInfo, error) {
	query := `
		SELECT id, name, domain, schema_name 
		FROM public.tenants 
		WHERE active = true 
		ORDER BY domain`

	rows, err := tc.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenants []*TenantInfo
	for rows.Next() {
		var tenant TenantInfo
		err := rows.Scan(
			&tenant.ID,
			&tenant.Name,
			&tenant.Domain,
			&tenant.Schema,
		)
		if err != nil {
			return nil, err
		}
		tenants = append(tenants, &tenant)
	}

	return tenants, rows.Err()
}

// ConvertToTenantConfig TenantInfo'yu models.TenantConfig'e çevirir
func (ti *TenantInfo) ConvertToTenantConfig() *models.TenantConfig {
	return &models.TenantConfig{
		ID:     ti.ID,
		Name:   ti.Name,
		Host:   ti.Domain,
		Schema: ti.Schema,
	}
}
