package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redplakad/goapi/handlers"
	"github.com/redplakad/goapi/middleware"
	"github.com/redplakad/goapi/models"
)

func main() {
	// Load configuration
	config := LoadConfig()

	// Connect to database
	database := NewDatabase(config)
	defer database.Close()

	// Initialize repositories
	nominatifRepo := models.NewNominatifKreditRepository(database.DB)

	// Initialize handlers
	nominatifHandler := handlers.NewNominatifKreditHandler(nominatifRepo)
	apiKeyHandler := handlers.NewAPIKeyHandler()

	// Setup Gin router
	r := gin.Default()

	// Add CORS middleware
	r.Use(middleware.CORS())

	// Public routes (tidak perlu API key)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/health", func(c *gin.Context) {
		if err := database.DB.Ping(); err != nil {
			c.JSON(500, gin.H{"status": "Database connection failed"})
			return
		}
		c.JSON(200, gin.H{"status": "OK", "database": "Connected"})
	})

	// Protected routes dengan API Key
	api := r.Group("/api")
	api.Use(middleware.APIKeyAuth()) // Semua route di grup ini perlu API key
	{
		// API Key validation endpoint
		api.GET("/validate", apiKeyHandler.ValidateAPIKey)
		api.GET("/keys", apiKeyHandler.ListAPIKeys) // untuk admin

		// Nominatif Kredit routes - semua perlu API key
		api.GET("/nominatif-kredit", nominatifHandler.GetAll)
		api.GET("/nominatif-kredit/:id", nominatifHandler.GetByID)
		api.GET("/nominatif-kredit/rekening/:nomor_rekening", nominatifHandler.GetByNomorRekening)
	}

	// Start server
	port := getEnv("PORT", "8080")
	log.Printf("Server starting on port %s", port)
	log.Printf("API Keys configured for authentication")
	r.Run(":" + port)
}
