package transport

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wolkodaf/todo/backend/internal/config"
)

func NewRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()
	router.RedirectTrailingSlash = false

	corsCfg := cors.DefaultConfig()
	if len(cfg.HTTPServer.AllowedOrigins) > 0 && cfg.HTTPServer.AllowedOrigins[0] != "" {
		corsCfg.AllowOrigins = cfg.HTTPServer.AllowedOrigins
	} else {
		corsCfg.AllowOrigins = []string{"http://localhost:5173"}
	}
	corsCfg.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsCfg.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	corsCfg.ExposeHeaders = []string{"Content-Length"}
	corsCfg.AllowCredentials = true
	corsCfg.MaxAge = 12 * time.Hour

	router.Use(cors.New(corsCfg))

	// Простой health-check.
	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return router
}
