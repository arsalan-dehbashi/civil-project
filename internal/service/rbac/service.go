package rbac

import (
	"context"
	"strings"

	rbacdto "github.com/arsalan-dehbashi/civil-project.git/internal/dto/rbac"
	"github.com/arsalan-dehbashi/civil-project.git/internal/models"
	rbacrepo "github.com/arsalan-dehbashi/civil-project.git/internal/repository/rbac"
)

type service struct {
	repo rbacrepo.Repository
}

func NewService(repo rbacrepo.Repository) Service {
	return &service{repo: repo}
}

func (s *service) ListRoles(ctx context.Context) ([]rbacdto.RoleResponse, error) {
	roles, err := s.repo.ListRoles(ctx)
	if err != nil {
		return nil, err
	}
	return rbacdto.RolesFromModels(roles), nil
}

func (s *service) CreateRole(ctx context.Context, input rbacdto.CreateRoleRequest) (*rbacdto.RoleResponse, error) {
	role := &models.Role{
		Name:        strings.TrimSpace(input.Name),
		DisplayName: strings.TrimSpace(input.DisplayName),
		Level:       input.Level,
		Description: strings.TrimSpace(input.Description),
		IsSystem:    input.IsSystem,
	}

	if role.Name == "" || role.DisplayName == "" {
		return nil, ValidationError("name and display_name are required")
	}

	if err := s.repo.CreateRole(ctx, role); err != nil {
		return nil, err
	}

	response := rbacdto.RoleFromModel(*role)
	return &response, nil
}

func (s *service) GetRole(ctx context.Context, id uint64) (*rbacdto.RoleResponse, error) {
	role, err := s.repo.GetRole(ctx, id)
	if err != nil {
		return nil, err
	}

	response := rbacdto.RoleFromModel(*role)
	return &response, nil
}

func (s *service) UpdateRole(ctx context.Context, id uint64, input rbacdto.UpdateRoleRequest) (*rbacdto.RoleResponse, error) {
	updates := map[string]any{}
	if input.Name != nil {
		value := strings.TrimSpace(*input.Name)
		if value == "" {
			return nil, ValidationError("name cannot be empty")
		}
		updates["name"] = value
	}
	if input.DisplayName != nil {
		value := strings.TrimSpace(*input.DisplayName)
		if value == "" {
			return nil, ValidationError("display_name cannot be empty")
		}
		updates["display_name"] = value
	}
	if input.Level != nil {
		updates["level"] = *input.Level
	}
	if input.Description != nil {
		updates["description"] = strings.TrimSpace(*input.Description)
	}
	if input.IsSystem != nil {
		updates["is_system"] = *input.IsSystem
	}
	if len(updates) == 0 {
		return nil, ValidationError("no role fields provided")
	}

	role, err := s.repo.UpdateRole(ctx, id, updates)
	if err != nil {
		return nil, err
	}

	response := rbacdto.RoleFromModel(*role)
	return &response, nil
}

func (s *service) DeleteRole(ctx context.Context, id uint64) error {
	return s.repo.DeleteRole(ctx, id)
}

func (s *service) SyncRolePermissions(ctx context.Context, roleID uint64, permissionIDs []uint64) (*rbacdto.RoleResponse, error) {
	role, err := s.repo.SyncRolePermissions(ctx, roleID, uniqueUint64(permissionIDs))
	if err != nil {
		return nil, err
	}

	response := rbacdto.RoleFromModel(*role)
	return &response, nil
}

func (s *service) ListPermissions(ctx context.Context) ([]rbacdto.PermissionResponse, error) {
	permissions, err := s.repo.ListPermissions(ctx)
	if err != nil {
		return nil, err
	}
	return rbacdto.PermissionsFromModels(permissions), nil
}

func (s *service) CreatePermission(ctx context.Context, input rbacdto.CreatePermissionRequest) (*rbacdto.PermissionResponse, error) {
	permission := &models.Permission{
		Name:        strings.TrimSpace(input.Name),
		Resource:    strings.TrimSpace(input.Resource),
		Action:      strings.TrimSpace(input.Action),
		Description: strings.TrimSpace(input.Description),
		IsSystem:    input.IsSystem,
	}

	if permission.Name == "" || permission.Resource == "" || permission.Action == "" {
		return nil, ValidationError("name, resource, and action are required")
	}

	if err := s.repo.CreatePermission(ctx, permission); err != nil {
		return nil, err
	}

	response := rbacdto.PermissionFromModel(*permission)
	return &response, nil
}

func (s *service) GetPermission(ctx context.Context, id uint64) (*rbacdto.PermissionResponse, error) {
	permission, err := s.repo.GetPermission(ctx, id)
	if err != nil {
		return nil, err
	}

	response := rbacdto.PermissionFromModel(*permission)
	return &response, nil
}

func (s *service) UpdatePermission(ctx context.Context, id uint64, input rbacdto.UpdatePermissionRequest) (*rbacdto.PermissionResponse, error) {
	updates := map[string]any{}
	if input.Name != nil {
		value := strings.TrimSpace(*input.Name)
		if value == "" {
			return nil, ValidationError("name cannot be empty")
		}
		updates["name"] = value
	}
	if input.Resource != nil {
		value := strings.TrimSpace(*input.Resource)
		if value == "" {
			return nil, ValidationError("resource cannot be empty")
		}
		updates["resource"] = value
	}
	if input.Action != nil {
		value := strings.TrimSpace(*input.Action)
		if value == "" {
			return nil, ValidationError("action cannot be empty")
		}
		updates["action"] = value
	}
	if input.Description != nil {
		updates["description"] = strings.TrimSpace(*input.Description)
	}
	if input.IsSystem != nil {
		updates["is_system"] = *input.IsSystem
	}
	if len(updates) == 0 {
		return nil, ValidationError("no permission fields provided")
	}

	permission, err := s.repo.UpdatePermission(ctx, id, updates)
	if err != nil {
		return nil, err
	}

	response := rbacdto.PermissionFromModel(*permission)
	return &response, nil
}

func (s *service) DeletePermission(ctx context.Context, id uint64) error {
	return s.repo.DeletePermission(ctx, id)
}

func (s *service) SyncUserRoles(ctx context.Context, userID uint64, roleIDs []uint64) (*rbacdto.UserRolesResponse, error) {
	user, err := s.repo.SyncUserRoles(ctx, userID, uniqueUint64(roleIDs))
	if err != nil {
		return nil, err
	}

	response := rbacdto.UserRolesFromModel(*user)
	return &response, nil
}

func uniqueUint64(values []uint64) []uint64 {
	seen := make(map[uint64]struct{}, len(values))
	out := make([]uint64, 0, len(values))

	for _, value := range values {
		if value == 0 {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}

		seen[value] = struct{}{}
		out = append(out, value)
	}

	return out
}
