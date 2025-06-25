# Appointment API - Frontend Documentation

## Base URL
```
http://localhost:8080
```

## Authentication
Bearer token kullanılır:
```
Authorization: Bearer <jwt_token>
```

---

## 🔐 Authentication Endpoints

### POST /api/auth/register
Yeni kullanıcı kaydı
```json
Request:
{
  "name": "John Doe",
  "email": "user@example.com", 
  "password": "password123",
  "phone": "+90555123456"
}

Response:
{
  "success": true,
  "message": "User registered successfully"
}
```

### POST /api/auth/login
Kullanıcı girişi
```json
Request:
{
  "email": "user@example.com",
  "password": "password123"
}

Response:
{
  "success": true,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "user@example.com"
  }
}
```

### POST /api/auth/forgot-password
Şifre sıfırlama (placeholder)
```json
Request:
{
  "email": "user@example.com"
}
```

### POST /api/auth/reset-password
Şifre değiştirme (placeholder)
```json
Request:
{
  "token": "reset_token",
  "new_password": "newpassword123"
}
```

---

## 👤 User Endpoints (AUTH Required)

### GET /api/user/profile
Kullanıcı profili görüntüleme
```json
Response:
{
  "success": true,
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "user@example.com",
    "phone": "+90555123456",
    "birth_date": "1990-05-15",
    "ust_bel": 85.5,
    "orta_bel": 70.0,
    "alt_bel": 95.2,
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

### PUT /api/user/profile
Kullanıcı profili güncelleme
```json
Request:
{
  "name": "Updated Name",
  "phone": "+90555999888",
  "birth_date": "1990-05-15",
  "ust_bel": 85.5,
  "orta_bel": 70.0,
  "alt_bel": 95.2
}

Response:
{
  "success": true,
  "message": "Profile updated successfully"
}
```

### PUT /api/user/change-password
Şifre değiştirme
```json
Request:
{
  "current_password": "oldpassword",
  "new_password": "newpassword123"
}

Response:
{
  "success": true,
  "message": "Password changed successfully"
}
```

---

## 🏷️ Category Endpoints

### GET /api/categories
Aktif kategorileri listeleme
```json
Response:
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Masaj Hizmetleri",
      "description": "Çeşitli masaj hizmetleri",
      "image_url": "https://example.com/category1.jpg",
      "active": true
    }
  ]
}
```

### GET /api/categories/:id
Belirli kategori detayları
```json
Response:
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Masaj Hizmetleri",
    "description": "Çeşitli masaj hizmetleri",
    "image_url": "https://example.com/category1.jpg",
    "active": true
  }
}
```

---

## 🛠️ Service Endpoints

### GET /api/services
Aktif hizmetleri listeleme
```json
Response:
{
  "success": true,
  "data": [
    {
      "id": 1,
      "category_id": 1,
      "name": "İsveç Masajı",
      "description": "Rahatlatıcı İsveç masajı",
      "price": 250.00,
      "image_url": "https://example.com/service1.jpg",
      "active": true
    }
  ]
}
```

### GET /api/services/:id
Belirli hizmet detayları
```json
Response:
{
  "success": true,
  "data": {
    "id": 1,
    "category_id": 1,
    "name": "İsveç Masajı",
    "description": "Rahatlatıcı İsveç masajı",
    "price": 250.00,
    "image_url": "https://example.com/service1.jpg",
    "active": true
  }
}
```

---

## 👨‍⚕️ Specialist Endpoints

### GET /api/specialists
Aktif uzmanları listeleme
```json
Response:
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Dr. Ahmet Yılmaz",
      "email": "ahmet@example.com",
      "phone": "+90555123456",
      "active": true
    }
  ]
}
```

### GET /api/specialists/:id
Belirli uzman detayları
```json
Response:
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Dr. Ahmet Yılmaz",
    "email": "ahmet@example.com",
    "phone": "+90555123456",
    "active": true
  }
}
```

### GET /api/specialists/:id/working-hours
Uzmanın çalışma saatleri
```json
Response:
{
  "success": true,
  "data": [
    {
      "id": 1,
      "specialist_id": 1,
      "day_of_week": 1,
      "start_time": "09:00",
      "end_time": "17:00", 
      "active": true
    }
  ]
}
```

### GET /api/specialists/:id/available-slots
Uzmanın uygun randevu saatleri
```
Query Parameters:
- date (required): YYYY-MM-DD format (örn: 2025-05-26)

Example: /api/specialists/1/available-slots?date=2025-05-26
```

```json
Response:
{
  "success": true,
  "data": [
    "09:00",
    "10:00", 
    "11:00",
    "14:00",
    "15:00",
    "16:00"
  ]
}
```

---

## 📅 Appointment Endpoints (AUTH Required)

### POST /api/appointments
Randevu oluşturma
```json
Request:
{
  "specialist_id": 1,
  "service_id": 1,
  "appointment_date": "2025-05-26",
  "appointment_time": "2025-05-26T14:00:00Z",
  "notes": "Sırt ağrısı için"
}

Response:
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 1,
    "specialist_id": 1,
    "service_id": 1,
    "appointment_date": "2025-05-26T00:00:00Z",
    "appointment_time": "2025-05-26T14:00:00Z",
    "status": "pending",
    "payment_status": "pending",
    "total_amount": 250.00,
    "notes": "Sırt ağrısı için"
  }
}
```

### GET /api/appointments
Kullanıcının randevularını listeleme
```json
Response:
{
  "success": true,
  "data": [
    {
      "id": 1,
      "user_id": 1,
      "specialist_id": 1,
      "service_id": 1,
      "appointment_date": "2025-05-26T00:00:00Z",
      "appointment_time": "2025-05-26T14:00:00Z",
      "status": "confirmed",
      "payment_status": "completed",
      "total_amount": 250.00,
      "notes": "Sırt ağrısı için"
    }
  ]
}
```

### GET /api/appointments/:id
Belirli randevu detayları
```json
Response:
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 1,
    "specialist_id": 1,
    "service_id": 1,
    "appointment_date": "2025-05-26T00:00:00Z",
    "appointment_time": "2025-05-26T14:00:00Z",
    "status": "confirmed",
    "payment_status": "completed",
    "total_amount": 250.00,
    "notes": "Sırt ağrısı için"
  }
}
```

### PUT /api/appointments/:id
Randevu güncelleme (Sadece notlar güncellenebilir)
```json
Request:
{
  "notes": "Güncellenen not"
}

Response:
{
  "success": true,
  "message": "Appointment updated successfully"
}
```

### DELETE /api/appointments/:id
Randevu iptal etme
```json
Response:
{
  "success": true,
  "message": "Appointment cancelled successfully"
}
```

### POST /api/appointments/:id/payment
Randevu ödemesi (Demo)
```json
Request:
{
  "payment_method": "credit_card",
  "card_token": "demo_token_123",
  "device_id": 1
}

Response:
{
  "success": true,
  "message": "Payment processed successfully",
  "data": {
    "payment_id": 1,
    "amount": 250.00,
    "status": "completed"
  }
}
```

---

## 📞 Contact Endpoints

### POST /api/contact
İletişim formu gönderme
```json
Request:
{
  "name": "John Doe",
  "email": "john@example.com",
  "subject": "Randevu Sorunu",
  "message": "Randevumu iptal etmek istiyorum"
}

Response:
{
  "success": true,
  "message": "Contact message sent successfully"
}
```

---

## 🏥 System Endpoints

### GET /api/health
Sistem durumu kontrolü
```json
Response:
{
  "success": true,
  "message": "API is healthy",
  "timestamp": "2025-01-01T12:00:00Z"
}
```

---

## 🔧 Admin Endpoints (ADMIN AUTH Required)

### Kullanıcı Yönetimi
- `GET /api/admin/users` - Kullanıcıları listeleme (pagination: ?limit=10&offset=0)
- `POST /api/admin/users` - Kullanıcı oluşturma
- `PUT /api/admin/users/:id` - Kullanıcı güncelleme
- `DELETE /api/admin/users/:id` - Kullanıcı silme
- `PUT /api/admin/users/:id/role` - Kullanıcı rolü güncelleme

### Kategori Yönetimi
- `GET /api/admin/categories` - Kategorileri listeleme
- `POST /api/admin/categories` - Kategori oluşturma
- `PUT /api/admin/categories/:id` - Kategori güncelleme
- `DELETE /api/admin/categories/:id` - Kategori silme

### Hizmet Yönetimi
- `GET /api/admin/services` - Tüm hizmetleri listeleme
- `POST /api/admin/services` - Hizmet oluşturma
- `PUT /api/admin/services/:id` - Hizmet güncelleme
- `DELETE /api/admin/services/:id` - Hizmet silme

### Uzman Yönetimi
- `GET /api/admin/specialists` - Tüm uzmanları listeleme
- `POST /api/admin/specialists` - Uzman oluşturma
- `PUT /api/admin/specialists/:id` - Uzman güncelleme
- `DELETE /api/admin/specialists/:id` - Uzman silme

### Çalışma Saatleri Yönetimi
- `GET /api/admin/specialists/:id/working-hours` - Uzman çalışma saatleri
- `PUT /api/admin/specialists/:id/working-hours` - Çalışma saatleri güncelleme

### Randevu Yönetimi
- `GET /api/admin/appointments` - Tüm randevuları listeleme
- `POST /api/admin/appointments` - Admin randevu oluşturma
- `PUT /api/admin/appointments/:id` - Randevu güncelleme
- `DELETE /api/admin/appointments/:id` - Randevu silme
- `PUT /api/admin/appointments/:id/status` - Randevu durumu güncelleme

### Cihaz Yönetimi
- `GET /api/admin/devices` - Cihazları listeleme
- `POST /api/admin/devices` - Cihaz oluşturma
- `PUT /api/admin/devices/:id` - Cihaz güncelleme
- `DELETE /api/admin/devices/:id` - Cihaz silme

### Ödeme Yönetimi
- `GET /api/admin/payments` - Ödemeleri listeleme
- `POST /api/admin/payments` - Ödeme kaydı oluşturma
- `PUT /api/admin/payments/:id` - Ödeme güncelleme
- `DELETE /api/admin/payments/:id` - Ödeme silme

### Ayarlar Yönetimi
- `GET /api/admin/settings` - Sistem ayarlarını listeleme
- `PUT /api/admin/settings/:key` - Ayar güncelleme
- `PUT /api/admin/settings/appointment-duration` - Randevu süresi güncelleme (dakika)

### İletişim Mesajları
- `GET /api/admin/contact-messages` - İletişim mesajlarını listeleme
- `PUT /api/admin/contact-messages/:id/read` - Mesajı okundu olarak işaretleme
- `DELETE /api/admin/contact-messages/:id` - Mesaj silme

### Raporlar
- `GET /api/admin/reports/sales` - Satış raporları
- `GET /api/admin/reports/payments` - Ödeme raporları
- `GET /api/admin/reports/appointments` - Randevu raporları

---

## 📊 Response Formats

### Success Response
```json
{
  "success": true,
  "data": {},
  "message": "Operation successful"
}
```

### Error Response
```json
{
  "success": false,
  "error": "Error message description",
  "code": "ERROR_CODE"
}
```

### Pagination Response
```json
{
  "success": true,
  "data": [],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 100,
    "pages": 10
  }
}
```

---

## 🔑 Status Codes

### Appointment Status
- `pending` - Onay bekliyor
- `confirmed` - Onaylandı
- `completed` - Tamamlandı
- `cancelled` - İptal edildi

### Payment Status
- `pending` - Ödeme bekliyor
- `completed` - Ödeme tamamlandı
- `failed` - Ödeme başarısız
- `refunded` - İade edildi

### Day of Week (working_hours)
- `0` - Pazar
- `1` - Pazartesi
- `2` - Salı
- `3` - Çarşamba
- `4` - Perşembe
- `5` - Cuma
- `6` - Cumartesi

---

## 💡 Important Notes

### Available Slots
- Appointment duration global olarak settings'den alınır (varsayılan: 60 dakika)
- Mevcut randevular otomatik olarak çıkarılır
- Working hours ve active status kontrol edilir
- Multi-tenant destekli

### Authentication
- JWT token kullanılır
- Token süre sınırı vardır
- Admin endpoints için admin yetkisi gerekir

### Date Formats
- Date: `YYYY-MM-DD` (örn: 2025-05-26)
- DateTime: `YYYY-MM-DDTHH:MM:SSZ` (örn: 2025-05-26T14:00:00Z)
- Time: `HH:MM` (örn: 14:00)

### Multi-Tenant
- Host header ile tenant belirlenir
- `localhost:8080` - Ana tenant
- `test.localhost:8080` - Test tenant 