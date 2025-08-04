package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redplakad/goapi/handlers"
	"github.com/redplakad/goapi/middleware"
	"github.com/redplakad/goapi/models"
)

func main() {
	// Load configuration
	config := LoadConfig()

	// Initialize database
	database := NewDatabase(config)
	defer database.Close()

	// Initialize repositories
	nominatifKreditRepo := models.NewNominatifKreditRepository(database.DB)

	// Initialize handlers
	nominatifKreditHandler := handlers.NewNominatifKreditHandler(nominatifKreditRepo)

	// Initialize Gin router
	r := gin.Default()

	// Add CORS middleware
	r.Use(middleware.CORS())

	// Health check endpoints
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/health", func(c *gin.Context) {
		if err := database.DB.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Database connection failed",
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Database connection successful",
		})
	})

	// API routes
	api := r.Group("/api")
	{
		// Nominatif Kredit routes
		nominatifKredit := api.Group("/nominatif-kredit")
		{
			nominatifKredit.GET("", nominatifKreditHandler.GetAll)
			nominatifKredit.GET("/:id", nominatifKreditHandler.GetByID)
			nominatifKredit.GET("/rekening/:nomor_rekening", nominatifKreditHandler.GetByNomorRekening)
		}
	}

	// Start server
	r.Run(":" + config.Port)
}
