package app

import (
	"demo/config"
	"demo/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data":    "Server is up and running",
		"release": config.ConfigSingleton.CommitHash,
	}
	c.JSON(http.StatusOK, res)
}

func SetupApp() *gin.Engine {
	// Initialize config and logger
	config.New()
	log.New()

	// Set gin to release mode to disable debug logs
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.GET("/", HealthCheck)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return r
}
