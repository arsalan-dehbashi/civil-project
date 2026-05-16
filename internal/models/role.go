package models

import "time"

type Role struct {
	ID          uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"size:64;not null;uniqueIndex"`
	DisplayName string    `json:"display_name" gorm:"size:128;not null"`
	Level       int       `json:"level" gorm:"not null;index"`
	Description string    `json:"description" gorm:"size:255"`
	IsSystem    bool      `json:"is_system" gorm:"not null;default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Permissions []Permission `json:"permissions,omitempty" gorm:"many2many:role_permissions;"`
	Users       []User       `json:"-" gorm:"many2many:user_roles;"`
}

func (Role) TableName() string {
	return "roles"
}
