package models

import "time"

type UserRole struct {
	UserID    uint64    `json:"user_id" gorm:"primaryKey;autoIncrement:false"`
	RoleID    uint64    `json:"role_id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time `json:"created_at"`
}

func (UserRole) TableName() string {
	return "user_roles"
}

type RolePermission struct {
	RoleID       uint64    `json:"role_id" gorm:"primaryKey;autoIncrement:false"`
	PermissionID uint64    `json:"permission_id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt    time.Time `json:"created_at"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
