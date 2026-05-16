package rbac

import (
	"time"

	"github.com/arsalan-dehbashi/civil-project.git/internal/models"
)

type CreateRoleRequest struct {
	Name        string `json:"name" binding:"required"`
	DisplayName string `json:"display_name" binding:"required"`
	Level       int    `json:"level" binding:"required"`
	Description string `json:"description"`
	IsSystem    bool   `json:"is_system"`
}

type UpdateRoleRequest struct {
	Name        *string `json:"name"`
	DisplayName *string `json:"display_name"`
	Level       *int    `json:"level"`
	Description *string `json:"description"`
	IsSystem    *bool   `json:"is_system"`
}

type SyncRolePermissionsRequest struct {
	PermissionIDs []uint64 `json:"permission_ids" binding:"required"`
}

type RoleResponse struct {
	ID          uint64               `json:"id"`
	Name        string               `json:"name"`
	DisplayName string               `json:"display_name"`
	Level       int                  `json:"level"`
	Description string               `json:"description"`
	IsSystem    bool                 `json:"is_system"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
	Permissions []PermissionResponse `json:"permissions,omitempty"`
}

func RoleFromModel(role models.Role) RoleResponse {
	response := RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		DisplayName: role.DisplayName,
		Level:       role.Level,
		Description: role.Description,
		IsSystem:    role.IsSystem,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}

	if len(role.Permissions) > 0 {
		response.Permissions = PermissionsFromModels(role.Permissions)
	}

	return response
}

func RolesFromModels(roles []models.Role) []RoleResponse {
	responses := make([]RoleResponse, 0, len(roles))
	for _, role := range roles {
		responses = append(responses, RoleFromModel(role))
	}
	return responses
}
