package rbac

import (
	"context"

	rbacdto "github.com/arsalan-dehbashi/civil-project.git/internal/dto/rbac"
)

type Service interface {
	ListRoles(ctx context.Context) ([]rbacdto.RoleResponse, error)
	CreateRole(ctx context.Context, input rbacdto.CreateRoleRequest) (*rbacdto.RoleResponse, error)
	GetRole(ctx context.Context, id uint64) (*rbacdto.RoleResponse, error)
	UpdateRole(ctx context.Context, id uint64, input rbacdto.UpdateRoleRequest) (*rbacdto.RoleResponse, error)
	DeleteRole(ctx context.Context, id uint64) error
	SyncRolePermissions(ctx context.Context, roleID uint64, permissionIDs []uint64) (*rbacdto.RoleResponse, error)

	ListPermissions(ctx context.Context) ([]rbacdto.PermissionResponse, error)
	CreatePermission(ctx context.Context, input rbacdto.CreatePermissionRequest) (*rbacdto.PermissionResponse, error)
	GetPermission(ctx context.Context, id uint64) (*rbacdto.PermissionResponse, error)
	UpdatePermission(ctx context.Context, id uint64, input rbacdto.UpdatePermissionRequest) (*rbacdto.PermissionResponse, error)
	DeletePermission(ctx context.Context, id uint64) error

	SyncUserRoles(ctx context.Context, userID uint64, roleIDs []uint64) (*rbacdto.UserRolesResponse, error)
}
