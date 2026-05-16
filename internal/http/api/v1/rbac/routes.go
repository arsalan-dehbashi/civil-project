package rbac

import (
	rbachandler "github.com/arsalan-dehbashi/civil-project.git/internal/http/handlers/rbac"
	rbacrepo "github.com/arsalan-dehbashi/civil-project.git/internal/repository/rbac"
	rbacservice "github.com/arsalan-dehbashi/civil-project.git/internal/service/rbac"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(api *gin.RouterGroup, db *gorm.DB) {
	repo := rbacrepo.NewGormRepository(db)
	service := rbacservice.NewService(repo)
	handler := rbachandler.NewHandler(service)

	roles := api.Group("/roles")
	roles.GET("", handler.ListRoles)
	roles.POST("", handler.CreateRole)
	roles.GET("/:id", handler.GetRole)
	roles.PATCH("/:id", handler.UpdateRole)
	roles.DELETE("/:id", handler.DeleteRole)
	roles.PUT("/:id/permissions", handler.SyncRolePermissions)

	permissions := api.Group("/permissions")
	permissions.GET("", handler.ListPermissions)
	permissions.POST("", handler.CreatePermission)
	permissions.GET("/:id", handler.GetPermission)
	permissions.PATCH("/:id", handler.UpdatePermission)
	permissions.DELETE("/:id", handler.DeletePermission)

	users := api.Group("/users")
	users.PUT("/:id/roles", handler.SyncUserRoles)
}
