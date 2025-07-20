package model

import "time"

type OTP struct {
	ID           uint      `gorm:"primarykey"`
	MobileNumber string    `gorm:"not null;size:20"`
	CodeHash     string    `gorm:"not null"` // Store a HASH of the OTP, not the OTP itself.
	ExpiresAt    time.Time `gorm:"not null"`
	CreatedAt    time.Time // Used for rate-limiting checks.
}
