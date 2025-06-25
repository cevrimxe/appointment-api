-- Public Schema - Tenant Management
-- Basit tenant tablosu - domain'e göre schema seçimi

CREATE TABLE IF NOT EXISTS public.tenants (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    domain VARCHAR(255) NOT NULL UNIQUE,  -- tam domain (localhost:3000, benimsitem.com)
    schema_name VARCHAR(100) NOT NULL UNIQUE,
    active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index for performance
CREATE INDEX IF NOT EXISTS idx_tenants_domain ON public.tenants(domain);
CREATE INDEX IF NOT EXISTS idx_tenants_active ON public.tenants(domain, active);

-- Sample data
INSERT INTO public.tenants (name, domain, schema_name) VALUES 
('Test Site', 'localhost:3000', 'test_schema'),
('Main Site', 'localhost:8080', 'main_schema'),
('Production Site', 'benimsitem.com', 'benimsitem_schema'); 