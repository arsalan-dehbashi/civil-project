package rbac

import (
	"time"

	"github.com/arsalan-dehbashi/civil-project.git/internal/models"
)

type SyncUserRolesRequest struct {
	RoleIDs []uint64 `json:"role_ids" binding:"required"`
}

type UserRolesResponse struct {
	ID              uint64         `json:"id"`
	PhoneNumber     string         `json:"phone_number"`
	IsPhoneVerified bool           `json:"is_phone_verified"`
	IsActive        bool           `json:"is_active"`
	LastLoginAt     *time.Time     `json:"last_login_at"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	Roles           []RoleResponse `json:"roles,omitempty"`
}

func UserRolesFromModel(user models.User) UserRolesResponse {
	return UserRolesResponse{
		ID:              user.ID,
		PhoneNumber:     user.PhoneNumber,
		IsPhoneVerified: user.IsPhoneVerified,
		IsActive:        user.IsActive,
		LastLoginAt:     user.LastLoginAt,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		Roles:           RolesFromModels(user.Roles),
	}
}
