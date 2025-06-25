# Appointment API

Multi-tenant randevu sistemi backend API'si.

## Kurulum

### 1. Environment Ayarları

`config.env` dosyasını kendi database bilgilerinizle güncelleyin:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=YOUR_POSTGRES_PASSWORD
DB_NAME=appointment_db
DB_SSLMODE=disable

# Server Configuration
SERVER_PORT=8080

# JWT Configuration
JWT_SECRET=my-super-secret-jwt-key-for-appointment-api-2024

# App Configuration
APP_ENV=development
APP_DEBUG=true
```

### 2. Database Oluşturma

PostgreSQL'de database oluşturun:
```sql
CREATE DATABASE appointment_db;
```

### 3. Migration Çalıştırma

```bash
# Database tablolarını oluştur
psql -U postgres -d appointment_db -f migrations/001_create_tables.sql
psql -U postgres -d appointment_db -f migrations/002_create_main_schema.sql
```

### 4. Uygulamayı Çalıştırma

```bash
go run cmd/server/main.go
```

## API Endpoints

Detaylar için `endpoints.md` dosyasına bakın.

### Public Endpoints
- `GET /api/categories` - Kategorileri listele
- `GET /api/services` - Hizmetleri listele
- `POST /api/auth/register` - Kullanıcı kaydı
- `POST /api/auth/login` - Giriş

### Admin Endpoints (JWT + Admin gerekli)
- `POST /api/admin/categories` - Kategori oluştur
- `PUT /api/admin/categories/:id` - Kategori güncelle
- `DELETE /api/admin/categories/:id` - Kategori sil
- `POST /api/admin/services` - Hizmet oluştur
- `PUT /api/admin/services/:id` - Hizmet güncelle
- `DELETE /api/admin/services/:id` - Hizmet sil

## Test

Server çalışınca: `http://localhost:8080/api/health` 