package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redplakad/goapi/middleware"
)

type APIKeyHandler struct{}

func NewAPIKeyHandler() *APIKeyHandler {
	return &APIKeyHandler{}
}

// ListAPIKeys returns list of all API keys (for admin purposes)
func (h *APIKeyHandler) ListAPIKeys(c *gin.Context) {
	apiKeys := middleware.GetValidAPIKeys()
	
	// Convert to response format (hide actual keys for security)
	response := make(map[string]interface{})
	for key, info := range apiKeys {
		// Mask the API key for security (show only first and last 4 characters)
		maskedKey := ""
		if len(key) > 8 {
			maskedKey = key[:4] + "****" + key[len(key)-4:]
		} else {
			maskedKey = "****"
		}
		
		response[maskedKey] = gin.H{
			"name":        info.Name,
			"description": info.Description,
			"active":      info.Active,
		}
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "API keys retrieved successfully",
		Data:    response,
	})
}

// ValidateAPIKey endpoint untuk test API key
func (h *APIKeyHandler) ValidateAPIKey(c *gin.Context) {
	// Jika sampai ke sini, berarti API key valid (sudah dicek di middleware)
	apiKeyName, _ := c.Get("api_key_name")
	apiKeyDescription, _ := c.Get("api_key_description")

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "API key is valid",
		Data: gin.H{
			"api_key_name":        apiKeyName,
			"api_key_description": apiKeyDescription,
			"timestamp":          gin.H{
				"validated_at": "now",
			},
		},
	})
}
