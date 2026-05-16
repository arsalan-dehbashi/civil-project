package v1

import (
	"github.com/arsalan-dehbashi/civil-project.git/internal/app"
	"github.com/arsalan-dehbashi/civil-project.git/internal/http/api/v1/rbac"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, a *app.App) {
	api := r.Group("/api/v1")

	rbac.RegisterRoutes(api, a.DB)
}
