package rbac

import (
	"context"

	"github.com/arsalan-dehbashi/civil-project.git/internal/models"
)

type Repository interface {
	ListRoles(ctx context.Context) ([]models.Role, error)
	CreateRole(ctx context.Context, role *models.Role) error
	GetRole(ctx context.Context, id uint64) (*models.Role, error)
	UpdateRole(ctx context.Context, id uint64, updates map[string]any) (*models.Role, error)
	DeleteRole(ctx context.Context, id uint64) error
	SyncRolePermissions(ctx context.Context, roleID uint64, permissionIDs []uint64) (*models.Role, error)

	ListPermissions(ctx context.Context) ([]models.Permission, error)
	CreatePermission(ctx context.Context, permission *models.Permission) error
	GetPermission(ctx context.Context, id uint64) (*models.Permission, error)
	UpdatePermission(ctx context.Context, id uint64, updates map[string]any) (*models.Permission, error)
	DeletePermission(ctx context.Context, id uint64) error

	SyncUserRoles(ctx context.Context, userID uint64, roleIDs []uint64) (*models.User, error)
}
