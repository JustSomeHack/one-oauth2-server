package controllers

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//SetupRouter starts up API service
func SetupRouter() (*gin.Engine, error) {
	if os.Getenv("PRODUCTION") == "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	authentication := router.Group("/")
	{
		authentication.POST("/authorize", AuthenticationPOST)
		authentication.POST("/validate", ValidatePOST)
	}

	health := router.Group("/health")
	{
		health.GET("", healthGET)
	}

	return router, nil
}

func healthGET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
