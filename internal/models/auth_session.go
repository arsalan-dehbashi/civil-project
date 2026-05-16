package models

import "time"

type AuthSession struct {
	ID               uint64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID           uint64     `json:"user_id" gorm:"not null;index"`
	JTI              string     `json:"jti" gorm:"size:64;not null;uniqueIndex"`
	RefreshTokenHash string     `json:"-" gorm:"size:255;not null;uniqueIndex"`
	UserAgent        string     `json:"user_agent" gorm:"size:255"`
	IPAddress        string     `json:"ip_address" gorm:"size:64"`
	ExpiresAt        time.Time  `json:"expires_at" gorm:"not null;index"`
	RevokedAt        *time.Time `json:"revoked_at" gorm:"index"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`

	User User `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (AuthSession) TableName() string {
	return "auth_sessions"
}
