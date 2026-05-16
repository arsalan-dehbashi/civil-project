package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CORSOptions struct {
	Origins          []string
	Methods          []string
	Headers          []string
	MaxAge           time.Duration
	AllowCredentials bool
	AllowWildcard    bool
}

func CORS(opt CORSOptions) gin.HandlerFunc {
	cfg := cors.Config{
		AllowOrigins:     opt.Origins,
		AllowMethods:     opt.Methods,
		AllowHeaders:     opt.Headers,
		ExposeHeaders:    []string{},
		AllowCredentials: opt.AllowCredentials,
		MaxAge:           opt.MaxAge,
		AllowWildcard:    opt.AllowWildcard,
	}

	return cors.New(cfg)
}
