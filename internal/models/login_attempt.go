package models

import "time"

type LoginAttempt struct {
	ID          uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	PhoneNumber string    `json:"phone_number" gorm:"size:32;not null;index"`
	UserID      *uint64   `json:"user_id" gorm:"index"`
	IPAddress   string    `json:"ip_address" gorm:"size:64;index"`
	UserAgent   string    `json:"user_agent" gorm:"size:255"`
	Success     bool      `json:"success" gorm:"not null;default:false;index"`
	Reason      string    `json:"reason" gorm:"size:128"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null;index"`

	User *User `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (LoginAttempt) TableName() string {
	return "login_attempts"
}
