package models

import "time"

type Permission struct {
	ID          uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"size:128;not null;uniqueIndex"`
	Resource    string    `json:"resource" gorm:"size:64;not null;index"`
	Action      string    `json:"action" gorm:"size:64;not null;index"`
	Description string    `json:"description" gorm:"size:255"`
	IsSystem    bool      `json:"is_system" gorm:"not null;default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Roles []Role `json:"-" gorm:"many2many:role_permissions;"`
}

func (Permission) TableName() string {
	return "permissions"
}
