package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint64         `json:"id" gorm:"primaryKey;autoIncrement"`
	PhoneNumber     string         `json:"phone_number" gorm:"size:32;not null;uniqueIndex"`
	PasswordHash    string         `json:"-" gorm:"size:255"`
	IsPhoneVerified bool           `json:"is_phone_verified" gorm:"not null;default:false"`
	IsActive        bool           `json:"is_active" gorm:"not null;default:true"`
	LastLoginAt     *time.Time     `json:"last_login_at"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`

	Roles []Role `json:"roles,omitempty" gorm:"many2many:user_roles;"`
}

func (User) TableName() string {
	return "users"
}
