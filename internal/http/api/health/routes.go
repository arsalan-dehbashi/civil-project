package health

import (
	"net/http"

	"github.com/arsalan-dehbashi/civil-project.git/internal/utils"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/healthz", func(c *gin.Context) {
		utils.Success(c, http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/readyz", func(c *gin.Context) {
		utils.Success(c, http.StatusOK, gin.H{"ready": true})
	})
}
