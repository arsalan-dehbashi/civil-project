package router

import (
	"fmt"
	"time"

	"github.com/arsalan-dehbashi/civil-project.git/internal/app"
	"github.com/arsalan-dehbashi/civil-project.git/internal/http/middleware"
	"github.com/arsalan-dehbashi/civil-project.git/internal/utils"
	"github.com/gin-gonic/gin"
)

func New(a *app.App) *gin.Engine {
	r := gin.New()
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(p gin.LogFormatterParams) string {

			if p.StatusCode < 400 && p.Latency < 500*time.Millisecond {
				return ""
			}

			return fmt.Sprintf("[GIN] %s | %d | %v | %s | %s %s\n",
				p.TimeStamp.Format("2006/01/02 - 15:04;05"),
				p.StatusCode,
				p.Latency,
				p.ClientIP,
				p.Method,
				p.Path,
			)
		},
	}))
	r.Use(gin.Recovery())

	r.Use(middleware.CORS(middleware.CORSOptions{
		Origins:          utils.SplitAndTrimCSV(a.Cfg.CORSAllowedOrigins),
		Methods:          utils.SplitAndTrimCSV(a.Cfg.CORSAllowedMethods),
		Headers:          utils.SplitAndTrimCSV(a.Cfg.CORSAllowedHeaders),
		MaxAge:           time.Duration(a.Cfg.CORSMaxAgeSeconds) * time.Second,
		AllowCredentials: false,
		AllowWildcard:    true,
	}))

	r.GET("/healthz", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })
	r.GET("/readyz", func(c *gin.Context) { c.JSON(200, gin.H{"ready": true}) })

	return r
}
