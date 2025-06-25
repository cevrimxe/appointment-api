package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 1. Login and get token
	loginData := map[string]string{
		"email":    "test2@test.com",
		"password": "test123",
	}

	jsonData, _ := json.Marshal(loginData)
	resp, err := http.Post("http://localhost:8080/api/auth/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Login error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Login Response: %s\n", string(body))

	// Parse response to get token
	var loginResp map[string]interface{}
	json.Unmarshal(body, &loginResp)

	if data, ok := loginResp["data"].(map[string]interface{}); ok {
		if token, ok := data["token"].(string); ok {
			fmt.Printf("Token: %s\n", token)

			// 2. Test profile endpoint with token
			req, _ := http.NewRequest("GET", "http://localhost:8080/api/user/profile", nil)
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("tenant-id", "localhost:8080") // Add tenant header

			client := &http.Client{}
			resp2, err := client.Do(req)
			if err != nil {
				fmt.Printf("Profile error: %v\n", err)
				return
			}
			defer resp2.Body.Close()

			body2, _ := io.ReadAll(resp2.Body)
			fmt.Printf("Profile Response (%d): %s\n", resp2.StatusCode, string(body2))
		}
	}
}
