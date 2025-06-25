# Test Senaryoları

Bu klasör, Appointment API'sinin çeşitli özelliklerini test etmek için hazırlanmış interaktif test runner'ını içerir.

## 📁 Dosyalar

- **`test_runner.go`** - Ana interaktif test runner (tek dosya)
- **`README.md`** - Bu dokümantasyon

## 🧪 Test Senaryoları

### 1. Available Slots Test
- **Amaç**: Available slots sistemini test eder
- **Testler**:
  - Belirtilen tarih için slot kontrolü
  - Multi-tenant slot kontrolü (test.localhost:8080)
  - Slot sayısı analizi

### 2. Auth & Registration Test
- **Amaç**: Authentication ve user registration sistemini test eder
- **Testler**:
  - Yeni user registration
  - Login testi (yeni oluşturulan user ile)
  - Token alma işlemi

### 3. Appointment Booking Test
- **Amaç**: Randevu oluşturma ve slot management'ı test eder
- **Testler**:
  - User registration + login
  - Randevu öncesi available slots kontrolü
  - Randevu oluşturma
  - Randevu sonrası slot azalması kontrolü

### 4. Settings & Duration Test
- **Amaç**: Settings ve appointment duration sistemini test eder
- **Testler**:
  - Mevcut settings kontrolü
  - Farklı duration senaryoları (30dk, 60dk, 90dk, 120dk)
  - Beklenen vs gerçek slot sayısı karşılaştırması

### 5. Tüm Testleri Çalıştır
- Yukarıdaki tüm testleri sırayla çalıştırır

## 🚀 Kullanım

### Interaktif Test Runner
```bash
# Ana dizinden
go run tests/test_runner.go
```

### Menü Seçenekleri
```
📋 Test Menüsü:
1. 📅 Available Slots Test
2. 🔐 Auth & Registration Test  
3. 📝 Appointment Booking Test
4. ⚙️ Settings & Duration Test
5. 🚀 Tüm Testleri Çalıştır
q. Çıkış
```

### Parametreler
Her test için şu parametreleri sorulur:
- **Host**: API host adresi (varsayılan: localhost:8080)
- **Specialist ID**: Test edilecek specialist (varsayılan: 1)
- **Service ID**: Test edilecek service (varsayılan: 1)
- **Test Tarihi**: YYYY-MM-DD formatında (varsayılan: 2025-05-26)

### Örnek Kullanım
```bash
go run tests/test_runner.go
# Test 1 seç
# Host: Enter (localhost:8080)
# Specialist ID: Enter (1) 
# Service ID: Enter (1)
# Test tarihi: 2025-05-27
# Test çalışır
# Menüye geri döner
# q ile çık
```

## ⚙️ Ön Koşullar

1. **API Server**: Ana API server'ın çalışıyor olması gerekir
   ```bash
   go run cmd/server/main.go
   ```

2. **Database**: PostgreSQL ve tenant'ların hazır olması gerekir

3. **Test Data**: 
   - En az 1 specialist olmalı
   - En az 1 service olmalı
   - Working hours tanımlı olmalı

## 📊 Beklenen Sonuçlar

### Available Slots (Pazartesi-Cuma)
- **60dk duration**: 8 slot (09:00-16:00)
- **30dk duration**: 16 slot
- **90dk duration**: 5 slot
- **120dk duration**: 4 slot

### Cumartesi-Pazar
- Boş array (working hours yok)

### Auth
- Registration: Başarılı
- Login: Token döner

### Appointment Booking
- Randevu oluşturma: Başarılı
- Slot azalması: Öncesi → Sonrası slot sayısı azalır

## 🔧 Özellikler

- ✅ **Interaktif**: Menü tabanlı seçim
- ✅ **Parametrik**: Her test için özelleştirilebilir parametreler
- ✅ **Multi-tenant**: Otomatik test.localhost:8080 testi
- ✅ **Döngü**: Test bittikten sonra menüye dönüş
- ✅ **Varsayılan değerler**: Enter'a basınca varsayılan kullanımı
- ✅ **Temiz çıktı**: Renkli emoji'li sonuçlar

## 💡 İpuçları

- Varsayılan değerleri kullanmak için Enter'a basın
- Multi-tenant test için host'u localhost:8080 bırakın
- Farklı tarihler denemek için YYYY-MM-DD formatını kullanın
- Appointment test'i öncesinde user registration otomatik yapılır
- q ile çıkana kadar testleri tekrar çalıştırabilirsiniz 