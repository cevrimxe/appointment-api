package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type TestConfig struct {
	Host         string
	SpecialistID int
	ServiceID    int
	Date         string
}

func main() {
	fmt.Println("🧪 Appointment API - İnteraktif Test Runner")
	fmt.Println(strings.Repeat("=", 50))

	for {
		showMenu()
		choice := getUserInput("Hangi test senaryosunu çalıştırmak istiyorsunuz? (1-5, q=çıkış): ")

		if choice == "q" || choice == "Q" {
			fmt.Println("👋 Görüşürüz!")
			break
		}

		switch choice {
		case "1":
			runAvailableSlotsTest()
		case "2":
			runAuthTest()
		case "3":
			runAppointmentTest()
		case "4":
			runSettingsTest()
		case "5":
			runAllTests()
		default:
			fmt.Println("❌ Geçersiz seçim! Lütfen 1-5 arası veya 'q' girin.")
		}

		fmt.Println("\n" + strings.Repeat("-", 50))
	}
}

func showMenu() {
	fmt.Println("\n📋 Test Menüsü:")
	fmt.Println("1. 📅 Available Slots Test")
	fmt.Println("2. 🔐 Auth & Registration Test")
	fmt.Println("3. 📝 Appointment Booking Test")
	fmt.Println("4. ⚙️ Settings & Duration Test")
	fmt.Println("5. 🚀 Tüm Testleri Çalıştır")
	fmt.Println("q. Çıkış")
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func getTestConfig() TestConfig {
	fmt.Println("\n⚙️ Test Konfigürasyonu:")

	host := getUserInput("Host (varsayılan: localhost:8080): ")
	if host == "" {
		host = "localhost:8080"
	}

	specialistStr := getUserInput("Specialist ID (varsayılan: 1): ")
	specialistID := 1
	if specialistStr != "" {
		if id, err := strconv.Atoi(specialistStr); err == nil {
			specialistID = id
		}
	}

	serviceStr := getUserInput("Service ID (varsayılan: 1): ")
	serviceID := 1
	if serviceStr != "" {
		if id, err := strconv.Atoi(serviceStr); err == nil {
			serviceID = id
		}
	}

	date := getUserInput("Test tarihi YYYY-MM-DD (varsayılan: 2025-05-26): ")
	if date == "" {
		date = "2025-05-26"
	}

	return TestConfig{
		Host:         host,
		SpecialistID: specialistID,
		ServiceID:    serviceID,
		Date:         date,
	}
}

func runAvailableSlotsTest() {
	fmt.Println("\n📅 Available Slots Test Başlatılıyor...")
	config := getTestConfig()

	fmt.Printf("\n🔍 Test ediliyor: %s, Specialist %d, Tarih %s\n", config.Host, config.SpecialistID, config.Date)

	url := fmt.Sprintf("http://localhost:8080/api/specialists/%d/available-slots?date=%s", config.SpecialistID, config.Date)
	response := makeGetRequest(url, config.Host)

	var result map[string]interface{}
	json.Unmarshal([]byte(response), &result)

	if result["success"] == true && result["data"] != nil {
		slots := result["data"].([]interface{})
		fmt.Printf("✅ %d slot bulundu: %v\n", len(slots), slots)
	} else {
		fmt.Printf("❌ Problem: %s\n", response)
	}

	// Multi-tenant test
	if strings.Contains(config.Host, "localhost") {
		fmt.Println("\n🏢 Multi-tenant test...")
		testHost := "test." + config.Host
		response = makeGetRequest(url, testHost)

		json.Unmarshal([]byte(response), &result)
		if result["success"] == true && result["data"] != nil {
			slots := result["data"].([]interface{})
			fmt.Printf("✅ Test tenant: %d slot\n", len(slots))
		} else {
			fmt.Printf("❌ Test tenant problem: %s\n", response)
		}
	}
}

func runAuthTest() {
	fmt.Println("\n🔐 Auth Test Başlatılıyor...")
	config := getTestConfig()

	// Registration test
	fmt.Println("\n👤 User Registration Test:")
	email := fmt.Sprintf("test_%d@example.com", time.Now().Unix())
	regData := map[string]string{
		"name":     "Test User",
		"email":    email,
		"password": "password123",
		"phone":    "+90555123456",
	}

	jsonData, _ := json.Marshal(regData)
	response := makePostRequest("http://localhost:8080/api/auth/register", jsonData, config.Host)

	var result map[string]interface{}
	json.Unmarshal([]byte(response), &result)

	if result["success"] == true {
		fmt.Printf("✅ Registration başarılı: %s\n", email)

		// Login test with new user
		fmt.Println("\n🔑 Login Test (Yeni user):")
		loginData := map[string]string{
			"email":    email,
			"password": "password123",
		}

		jsonData, _ = json.Marshal(loginData)
		response = makePostRequest("http://localhost:8080/api/auth/login", jsonData, config.Host)

		json.Unmarshal([]byte(response), &result)
		if result["success"] == true && result["token"] != nil {
			token := result["token"].(string)
			fmt.Printf("✅ Login başarılı, token: %s...\n", token[:20])
		} else {
			fmt.Printf("❌ Login başarısız: %s\n", response)
		}
	} else {
		fmt.Printf("❌ Registration başarısız: %s\n", response)
	}
}

func runAppointmentTest() {
	fmt.Println("\n📝 Appointment Test Başlatılıyor...")
	config := getTestConfig()

	// Get auth token first
	fmt.Println("\n🔐 Auth token alınıyor...")
	email := fmt.Sprintf("test_%d@example.com", time.Now().Unix())

	// Quick registration
	regData := map[string]string{
		"name":     "Test User",
		"email":    email,
		"password": "password123",
		"phone":    "+90555123456",
	}

	jsonData, _ := json.Marshal(regData)
	makePostRequest("http://localhost:8080/api/auth/register", jsonData, config.Host)

	// Login
	loginData := map[string]string{
		"email":    email,
		"password": "password123",
	}

	jsonData, _ = json.Marshal(loginData)
	response := makePostRequest("http://localhost:8080/api/auth/login", jsonData, config.Host)

	var result map[string]interface{}
	json.Unmarshal([]byte(response), &result)

	if result["success"] != true || result["token"] == nil {
		fmt.Println("❌ Auth token alınamadı")
		return
	}

	token := result["token"].(string)
	fmt.Println("✅ Token alındı")

	// Available slots before
	fmt.Println("\n📅 Available slots (öncesi):")
	url := fmt.Sprintf("http://localhost:8080/api/specialists/%d/available-slots?date=%s", config.SpecialistID, config.Date)
	response = makeGetRequest(url, config.Host)

	json.Unmarshal([]byte(response), &result)
	var slotsBefore []interface{}
	if result["success"] == true && result["data"] != nil {
		slotsBefore = result["data"].([]interface{})
		fmt.Printf("   %d slot mevcut\n", len(slotsBefore))
	}

	// Create appointment
	fmt.Println("\n💾 Randevu oluşturuluyor...")
	appointmentTime := config.Date + "T09:00:00Z"

	appointmentData := map[string]interface{}{
		"specialist_id":    config.SpecialistID,
		"service_id":       config.ServiceID,
		"appointment_date": config.Date,
		"appointment_time": appointmentTime,
		"notes":            "Test randevusu",
	}

	jsonData, _ = json.Marshal(appointmentData)
	response = makeAuthRequest("POST", "http://localhost:8080/api/appointments", jsonData, token, config.Host)

	json.Unmarshal([]byte(response), &result)
	if result["success"] == true {
		fmt.Println("✅ Randevu oluşturuldu")

		// Available slots after
		fmt.Println("\n📅 Available slots (sonrası):")
		response = makeGetRequest(url, config.Host)

		json.Unmarshal([]byte(response), &result)
		if result["success"] == true && result["data"] != nil {
			slotsAfter := result["data"].([]interface{})
			fmt.Printf("   %d slot kaldı\n", len(slotsAfter))

			if len(slotsAfter) < len(slotsBefore) {
				fmt.Println("✅ Slot başarıyla rezerve edildi!")
			} else {
				fmt.Println("⚠️ Slot sayısı azalmadı")
			}
		}
	} else {
		fmt.Printf("❌ Randevu oluşturulamadı: %s\n", response)
	}
}

func runSettingsTest() {
	fmt.Println("\n⚙️ Settings Test Başlatılıyor...")
	config := getTestConfig()

	fmt.Println("\n📋 Mevcut settings:")
	response := makeGetRequest("http://localhost:8080/api/admin/settings", config.Host)
	fmt.Printf("   %s\n", response)

	fmt.Println("\n⏱️ Duration hesaplamaları:")
	durations := []struct {
		duration string
		desc     string
	}{
		{"30", "30 dakika"},
		{"60", "60 dakika (standart)"},
		{"90", "90 dakika"},
		{"120", "120 dakika"},
	}

	for _, d := range durations {
		expected := calculateExpectedSlots(d.duration)
		fmt.Printf("   %s: %d slot bekleniyor\n", d.desc, expected)
	}

	// Current slots
	fmt.Printf("\n📅 Mevcut slots (%s):\n", config.Date)
	url := fmt.Sprintf("http://localhost:8080/api/specialists/%d/available-slots?date=%s", config.SpecialistID, config.Date)
	response = makeGetRequest(url, config.Host)

	var result map[string]interface{}
	json.Unmarshal([]byte(response), &result)

	if result["success"] == true && result["data"] != nil {
		slots := result["data"].([]interface{})
		fmt.Printf("   Gerçek: %d slot\n", len(slots))
	}
}

func runAllTests() {
	fmt.Println("\n🚀 Tüm Testler Çalıştırılıyor...")

	tests := []struct {
		name string
		fn   func()
	}{
		{"Available Slots", runAvailableSlotsTest},
		{"Auth & Registration", runAuthTest},
		{"Appointment Booking", runAppointmentTest},
		{"Settings", runSettingsTest},
	}

	for i, test := range tests {
		fmt.Printf("\n%d. %s\n", i+1, test.name)
		fmt.Println(strings.Repeat("-", 30))
		test.fn()

		if i < len(tests)-1 {
			fmt.Println("\n" + strings.Repeat("=", 30))
		}
	}
}

func calculateExpectedSlots(duration string) int {
	totalMinutes := 480 // 09:00-17:00 = 8 saat

	switch duration {
	case "30":
		return totalMinutes / 30
	case "60":
		return totalMinutes / 60
	case "90":
		return totalMinutes / 90
	case "120":
		return totalMinutes / 120
	default:
		return 8
	}
}

// HTTP helper functions
func makeGetRequest(url, host string) string {
	client := &http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}

	req.Host = host

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}

	return string(body)
}

func makePostRequest(url string, body []byte, host string) string {
	client := &http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Host = host

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}

	return string(responseBody)
}

func makeAuthRequest(method, url string, body []byte, token, host string) string {
	client := &http.Client{Timeout: 10 * time.Second}

	var req *http.Request
	var err error

	if body != nil {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
	}

	req.Host = host
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}

	return string(responseBody)
}
