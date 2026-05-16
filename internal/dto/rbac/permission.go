package rbac

import (
	"time"

	"github.com/arsalan-dehbashi/civil-project.git/internal/models"
)

type CreatePermissionRequest struct {
	Name        string `json:"name" binding:"required"`
	Resource    string `json:"resource" binding:"required"`
	Action      string `json:"action" binding:"required"`
	Description string `json:"description"`
	IsSystem    bool   `json:"is_system"`
}

type UpdatePermissionRequest struct {
	Name        *string `json:"name"`
	Resource    *string `json:"resource"`
	Action      *string `json:"action"`
	Description *string `json:"description"`
	IsSystem    *bool   `json:"is_system"`
}

type PermissionResponse struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Resource    string    `json:"resource"`
	Action      string    `json:"action"`
	Description string    `json:"description"`
	IsSystem    bool      `json:"is_system"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func PermissionFromModel(permission models.Permission) PermissionResponse {
	return PermissionResponse{
		ID:          permission.ID,
		Name:        permission.Name,
		Resource:    permission.Resource,
		Action:      permission.Action,
		Description: permission.Description,
		IsSystem:    permission.IsSystem,
		CreatedAt:   permission.CreatedAt,
		UpdatedAt:   permission.UpdatedAt,
	}
}

func PermissionsFromModels(permissions []models.Permission) []PermissionResponse {
	responses := make([]PermissionResponse, 0, len(permissions))
	for _, permission := range permissions {
		responses = append(responses, PermissionFromModel(permission))
	}
	return responses
}
