package models

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	DefaultResetDuration = 1 * time.Hour
)

type PasswordReset struct {
	ID     int
	UserID int
	// Token is set when creating a new PasswordReset request.
	Token     string
	TokenHash string
	ExpiresAt time.Time
}

type PasswordResetService struct {
	DB *sql.DB
	//BytesPerToken - used to determine # of bytes to use when generating
	//PasswordReset token. Ignored if not set or < MinBytesPerToken; at which
	//point the value will be MinBytesPerToken
	BytesPerToken int
	// Duration - the amount of time that a PasswordReset is valid for.
	// Default value is DefaultResetDuration
	Duration time.Duration
}

func (service *PasswordResetService) Create(email string) (*PasswordReset, error) {
	return nil, fmt.Errorf("TODO: Implement PasswordResetService.Create")
}

func (service *PasswordResetService) Consume(token string) (*User, error) {
	return nil, fmt.Errorf("TODO: Implement PasswordResetService.Consume")
}
