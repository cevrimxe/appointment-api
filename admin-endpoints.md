# 🔐 Admin Panel API Documentation

Bu dokümantasyon admin paneli için kullanılacak tüm API endpoint'lerini içerir.

## 📋 İçindekiler
- [Authentication](#authentication)
- [Response Format](#response-format)
- [Error Handling](#error-handling)
- [Categories](#categories)
- [Services](#services)
- [Users](#users)
- [Specialists](#specialists)
- [Appointments](#appointments)
- [Devices](#devices)
- [Settings](#settings)
- [Payments](#payments)
- [Contact Messages](#contact-messages)
- [Reports & Analytics](#reports--analytics)

---

## 🔑 Authentication

Tüm admin endpoint'leri JWT token ile korumalıdır ve `admin` rolü gerektirir.

**Header:**
```
Authorization: Bearer <jwt_token>
```

**Admin Login:**
```http
POST /auth/login
Content-Type: application/json

{
  "email": "admin@example.com",
  "password": "password123"
}
```

---

## 📄 Response Format

Tüm API yanıtları aşağıdaki standart formatı kullanır:

**Başarılı Response:**
```json
{
  "success": true,
  "data": { ... },
  "message": "Operation completed successfully"
}
```

**Hatalı Response:**
```json
{
  "success": false,
  "error": "Error message description"
}
```

---

## ❌ Error Handling

| Status Code | Açıklama |
|-------------|-----------|
| 200 | Başarılı |
| 201 | Oluşturuldu |
| 400 | Geçersiz istek |
| 401 | Yetkisiz erişim |
| 403 | Yasak erişim |
| 404 | Bulunamadı |
| 500 | Sunucu hatası |

---

## 📂 Categories

### List Categories
```http
GET /admin/categories
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Genel Muayene",
      "description": "Genel sağlık kontrolü",
      "active": true,
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ]
}
```

### Create Category
```http
POST /admin/categories
Content-Type: application/json

{
  "name": "Yeni Kategori",
  "description": "Kategori açıklaması",
  "active": true
}
```

### Update Category
```http
PUT /admin/categories/{id}
Content-Type: application/json

{
  "name": "Güncellenmiş Kategori",
  "description": "Yeni açıklama",
  "active": true
}
```

### Delete Category
```http
DELETE /admin/categories/{id}
```

---

## 🛠️ Services

### List Services
```http
GET /admin/services
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "category_id": 1,
      "name": "Genel Kontrol",
      "description": "Rutin sağlık kontrolü",
      "price": 150.00,
      "image_url": "https://example.com/image.jpg",
      "active": true,
      "created_at": "2024-01-01T10:00:00Z"
    }
  ]
}
```

### Create Service
```http
POST /admin/services
Content-Type: application/json

{
  "category_id": 1,
  "name": "Yeni Hizmet",
  "description": "Hizmet açıklaması",
  "price": 200.00,
  "image_url": "https://example.com/image.jpg",
  "active": true
}
```

### Update Service
```http
PUT /admin/services/{id}
Content-Type: application/json

{
  "category_id": 1,
  "name": "Güncellenmiş Hizmet",
  "description": "Yeni açıklama",
  "price": 250.00,
  "active": true
}
```

### Delete Service
```http
DELETE /admin/services/{id}
```

---

## 👥 Users

### List Users (Pagination)
```http
GET /admin/users?limit=50&offset=0
```

**Response:**
```json
{
  "success": true,
  "data": {
    "users": [
      {
        "id": 1,
        "email": "user@example.com",
        "name": "Ahmet Yılmaz",
        "phone": "+90555123456",
        "role": "user",
        "birth_date": "1990-01-01T00:00:00Z",
        "ust_bel": 32,
        "orta_bel": 30,
        "alt_bel": 34,
        "created_at": "2024-01-01T10:00:00Z"
      }
    ],
    "total": 156,
    "limit": 50,
    "offset": 0
  }
}
```

### Create User
```http
POST /admin/users
Content-Type: application/json

{
  "email": "newuser@example.com",
  "password": "password123",
  "name": "Yeni Kullanıcı",
  "phone": "+90555123456",
  "role": "user",
  "birth_date": "1990-01-01T00:00:00Z"
}
```

### Update User
```http
PUT /admin/users/{id}
Content-Type: application/json

{
  "email": "updated@example.com",
  "name": "Güncellenmiş İsim",
  "phone": "+90555654321",
  "role": "user"
}
```

### Update User Role
```http
PUT /admin/users/{id}/role
Content-Type: application/json

{
  "role": "admin"
}
```

### Delete User
```http
DELETE /admin/users/{id}
```

---

## 👨‍⚕️ Specialists

### List Specialists
```http
GET /admin/specialists
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Dr. Mehmet Özkan",
      "email": "dr.ozkan@example.com",
      "phone": "+90555987654",
      "active": true,
      "created_at": "2024-01-01T10:00:00Z"
    }
  ]
}
```

### Create Specialist
```http
POST /admin/specialists
Content-Type: application/json

{
  "name": "Dr. Yeni Doktor",
  "email": "dr.yeni@example.com",
  "phone": "+90555111222",
  "active": true
}
```

### Update Specialist
```http
PUT /admin/specialists/{id}
Content-Type: application/json

{
  "name": "Dr. Güncellenmiş İsim",
  "email": "dr.updated@example.com",
  "phone": "+90555333444",
  "active": true
}
```

### Get Specialist Working Hours
```http
GET /admin/specialists/{id}/working-hours
```

**Response:**
```json
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

### Update Specialist Working Hours
```http
PUT /admin/specialists/{id}/working-hours
Content-Type: application/json

[
  {
    "day_of_week": 1,
    "start_time": "09:00",
    "end_time": "17:00",
    "active": true
  },
  {
    "day_of_week": 2,
    "start_time": "10:00",
    "end_time": "18:00",
    "active": true
  }
]
```

### Delete Specialist
```http
DELETE /admin/specialists/{id}
```

---

## 📅 Appointments

### List Appointments (Pagination)
```http
GET /admin/appointments?limit=50&offset=0
```

**Response:**
```json
{
  "success": true,
  "data": {
    "appointments": [
      {
        "id": 1,
        "user_id": 1,
        "specialist_id": 1,
        "service_id": 1,
        "appointment_date": "2024-02-15T00:00:00Z",
        "appointment_time": "0000-01-01T10:30:00Z",
        "status": "pending",
        "payment_status": "pending",
        "total_amount": 150.00,
        "notes": "Randevu notu",
        "created_at": "2024-01-01T10:00:00Z"
      }
    ],
    "total": 245,
    "limit": 50,
    "offset": 0
  }
}
```

### Create Appointment
```http
POST /admin/appointments
Content-Type: application/json

{
  "user_id": 1,
  "specialist_id": 1,
  "service_id": 1,
  "appointment_date": "2024-02-15T00:00:00Z",
  "appointment_time": "0000-01-01T10:30:00Z",
  "notes": "Admin tarafından oluşturulan randevu"
}
```

### Update Appointment
```http
PUT /admin/appointments/{id}
Content-Type: application/json

{
  "user_id": 1,
  "specialist_id": 1,
  "service_id": 1,
  "appointment_date": "2024-02-16T00:00:00Z",
  "appointment_time": "0000-01-01T11:00:00Z",
  "notes": "Güncellenmiş randevu"
}
```

### Update Appointment Status
```http
PUT /admin/appointments/{id}/status
Content-Type: application/json

{
  "status": "confirmed"
}
```

**Status Values:** `pending`, `confirmed`, `completed`, `cancelled`

### Delete Appointment
```http
DELETE /admin/appointments/{id}
```

---

## 📱 Devices

### List Devices
```http
GET /admin/devices
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "brand": "Samsung",
      "name": "Galaxy S23",
      "device_date": "2024-01-01T00:00:00Z",
      "price": 15000.00,
      "active": true,
      "created_at": "2024-01-01T10:00:00Z"
    }
  ]
}
```

### Create Device
```http
POST /admin/devices
Content-Type: application/json

{
  "brand": "Apple",
  "name": "iPhone 15",
  "device_date": "2024-01-01T00:00:00Z",
  "price": 18000.00,
  "active": true
}
```

### Update Device
```http
PUT /admin/devices/{id}
Content-Type: application/json

{
  "brand": "Apple",
  "name": "iPhone 15 Pro",
  "device_date": "2024-01-01T00:00:00Z",
  "price": 20000.00,
  "active": true
}
```

### Delete Device
```http
DELETE /admin/devices/{id}
```

---

## ⚙️ Settings

### List Settings
```http
GET /admin/settings
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "key": "appointment_duration",
      "value": "60",
      "description": "Default appointment duration in minutes",
      "created_at": "2024-01-01T10:00:00Z"
    }
  ]
}
```

### Update Setting
```http
PUT /admin/settings/{key}
Content-Type: application/json

{
  "value": "90",
  "description": "Updated appointment duration"
}
```

### Update Appointment Duration (Special)
```http
PUT /admin/settings/appointment-duration
Content-Type: application/json

{
  "duration": 90
}
```

---

## 💳 Payments

### List Payments (Pagination)
```http
GET /admin/payments?limit=50&offset=0
```

**Response:**
```json
{
  "success": true,
  "data": {
    "payments": [
      {
        "id": 1,
        "appointment_id": 1,
        "device_id": 1,
        "amount": 150.00,
        "payment_method": "card",
        "transaction_id": "demo_1_1672531200",
        "status": "completed",
        "created_at": "2024-01-01T10:00:00Z"
      }
    ],
    "limit": 50,
    "offset": 0
  }
}
```

### Create Payment
```http
POST /admin/payments
Content-Type: application/json

{
  "appointment_id": 1,
  "device_id": 1,
  "amount": 150.00,
  "payment_method": "card",
  "transaction_id": "manual_payment_123"
}
```

### Update Payment
```http
PUT /admin/payments/{id}
Content-Type: application/json

{
  "amount": 200.00,
  "status": "completed"
}
```

**Payment Status Values:** `pending`, `completed`, `failed`, `refunded`
**Payment Method Values:** `card`, `cash`, `transfer`

### Delete Payment
```http
DELETE /admin/payments/{id}
```

---

## 📧 Contact Messages

### List Contact Messages (Pagination)
```http
GET /admin/contact-messages?limit=50&offset=0
```

**Response:**
```json
{
  "success": true,
  "data": {
    "messages": [
      {
        "id": 1,
        "name": "Ahmet Yılmaz",
        "email": "ahmet@example.com",
        "subject": "Randevu Sorunu",
        "message": "Randevu alamıyorum, yardım edebilir misiniz?",
        "is_read": false,
        "created_at": "2024-01-01T10:00:00Z"
      }
    ],
    "total": 45,
    "limit": 50,
    "offset": 0
  }
}
```

### Mark Message as Read
```http
PUT /admin/contact-messages/{id}/read
```

### Delete Contact Message
```http
DELETE /admin/contact-messages/{id}
```

---

## 📊 Reports & Analytics

### Sales Reports
```http
GET /admin/reports/sales?start_date=2024-01-01&end_date=2024-01-31
```

**Response:**
```json
{
  "success": true,
  "data": {
    "start_date": "2024-01-01",
    "end_date": "2024-01-31",
    "total_revenue": 45000.50,
    "currency": "TRY"
  }
}
```

### Payment Reports
```http
GET /admin/reports/payments?limit=50&offset=0
```

**Response:**
```json
{
  "success": true,
  "data": {
    "payments": [...],
    "total_amount": 45000.50,
    "payment_count": 125,
    "currency": "TRY",
    "limit": 50,
    "offset": 0
  }
}
```

### Appointment Reports
```http
GET /admin/reports/appointments?limit=50&offset=0
```

**Response:**
```json
{
  "success": true,
  "data": {
    "appointments": [...],
    "total_count": 245,
    "status_counts": {
      "pending": 12,
      "confirmed": 156,
      "completed": 67,
      "cancelled": 10
    },
    "limit": 50,
    "offset": 0
  }
}
```

### Dashboard Stats (Simple)
```http
GET /admin/stats
```

**Response:**
```json
{
  "success": true,
  "data": {
    "today_appointments": 12,
    "total_revenue": 45000.50,
    "total_users": 156,
    "pending_appointments": 8
  }
}
```

### Dashboard Stats (Comprehensive) 🆕
```http
GET /admin/dashboard/stats
```

**Response:**
```json
{
  "success": true,
  "data": {
    "revenue": {
      "monthly": 15750.0,
      "yearly": 189000.0
    },
    "appointments": {
      "today": 12,
      "monthly": 245
    },
    "customers": {
      "total": 342,
      "new_monthly": 18
    },
    "trends": {
      "revenue_monthly_growth": 15.2,
      "revenue_yearly_growth": 22.0,
      "appointments_today_growth": 3,
      "appointments_monthly_growth": 12,
      "customers_growth": 5.2
    }
  }
}
```

---

## 🔐 Authentication Extras

### Forgot Password
```http
POST /auth/forgot-password
Content-Type: application/json

{
  "email": "admin@example.com"
}
```

### Reset Password
```http
POST /auth/reset-password
Content-Type: application/json

{
  "token": "reset_1_1672531200",
  "new_password": "newpassword123"
}
```

---

## 💡 Important Notes

1. **Pagination:** Tüm liste endpoint'leri `limit` ve `offset` parametrelerini destekler
2. **Authentication:** Tüm admin endpoint'leri JWT token ve admin rolü gerektirir
3. **Date Format:** Tüm tarihler ISO 8601 formatında (`YYYY-MM-DDTHH:mm:ssZ`)
4. **Time Format:** Saat değerleri `HH:MM` formatında
5. **Currency:** Tüm fiyatlar TRY cinsinden
6. **File Uploads:** Image upload'ları için ayrı endpoint gerekebilir
7. **Demo Features:** Payment ve email işlemleri demo modunda çalışır

---

## 🚀 Base URL

```
Production: https://api.yourapp.com
Development: http://localhost:8080
```

---

**Son Güncelleme:** 2024-01-01  
**API Version:** v1 