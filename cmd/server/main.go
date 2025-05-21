package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/varshaprasad96/mcp-kserve/pkg/handlers"
	"github.com/varshaprasad96/mcp-kserve/pkg/kserve"
)

func main() {
	// Initialize configuration
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Could not read config file: %v", err)
	}

	// Set default values
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("k8s.namespace", "default")

	// Initialize KServe client
	kserveClient, err := kserve.NewClient(viper.GetString("k8s.namespace"))
	if err != nil {
		log.Fatalf("Failed to create KServe client: %v", err)
	}

	// Initialize model handler
	modelHandler := handlers.NewModelHandler(kserveClient)

	// Initialize router
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	// Model endpoints
	router.POST("/api/v1/models", modelHandler.DeployModel)
	router.GET("/api/v1/models/:name", modelHandler.GetModelStatus)
	router.GET("/api/v1/models", modelHandler.ListModels)
	router.DELETE("/api/v1/models/:name", modelHandler.DeleteModel)

	// Start server
	port := viper.GetString("server.port")
	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
