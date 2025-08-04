package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Daftar API keys yang valid
// Dalam production, ini bisa disimpan di database atau config file
var validAPIKeys = map[string]APIKeyInfo{
	"api-key-laravel-frontend-2025": {
		Name:        "Laravel Frontend",
		Description: "Main Laravel frontend application",
		Active:      true,
	},
	"api-key-mobile-app-2025": {
		Name:        "Mobile App",
		Description: "Mobile application",
		Active:      true,
	},
	"api-key-dashboard-admin-2025": {
		Name:        "Admin Dashboard",
		Description: "Administrative dashboard",
		Active:      true,
	},
}

// APIKeyInfo stores information about API key
type APIKeyInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

// APIKeyAuth middleware untuk validasi API key
func APIKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cek API key dari header X-API-Key
		apiKey := c.GetHeader("X-API-Key")
		
		// Jika tidak ada, cek dari Authorization header dengan format "ApiKey <key>"
		if apiKey == "" {
			authHeader := c.GetHeader("Authorization")
			if authHeader != "" {
				parts := strings.SplitN(authHeader, " ", 2)
				if len(parts) == 2 && strings.ToLower(parts[0]) == "apikey" {
					apiKey = parts[1]
				}
			}
		}

		// Jika masih tidak ada API key
		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "API key is required. Please provide X-API-Key header or Authorization: ApiKey <key>",
				"error":   "missing_api_key",
			})
			c.Abort()
			return
		}

		// Validasi API key
		keyInfo, exists := validAPIKeys[apiKey]
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid API key",
				"error":   "invalid_api_key",
			})
			c.Abort()
			return
		}

		// Cek apakah API key masih aktif
		if !keyInfo.Active {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "API key is disabled",
				"error":   "disabled_api_key",
			})
			c.Abort()
			return
		}

		// Set informasi API key di context untuk logging atau tracking
		c.Set("api_key", apiKey)
		c.Set("api_key_name", keyInfo.Name)
		c.Set("api_key_description", keyInfo.Description)

		c.Next()
	}
}

// GetValidAPIKeys returns list of valid API keys (for admin purposes)
func GetValidAPIKeys() map[string]APIKeyInfo {
	return validAPIKeys
}

// AddAPIKey adds a new API key
func AddAPIKey(key string, info APIKeyInfo) {
	validAPIKeys[key] = info
}

// RemoveAPIKey removes an API key
func RemoveAPIKey(key string) {
	delete(validAPIKeys, key)
}

// DeactivateAPIKey deactivates an API key
func DeactivateAPIKey(key string) {
	if info, exists := validAPIKeys[key]; exists {
		info.Active = false
		validAPIKeys[key] = info
	}
}

// ActivateAPIKey activates an API key
func ActivateAPIKey(key string) {
	if info, exists := validAPIKeys[key]; exists {
		info.Active = true
		validAPIKeys[key] = info
	}
}
