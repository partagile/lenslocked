package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/partagile/lenslocked/rand"
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
	// Verify we have a valid email address for user and get assoc. ID
	strings.ToLower(email)
	var userID int
	row := service.DB.QueryRow(`
		SELECT id
		FROM users 
		WHERE email = $1;`, email)
	err := row.Scan(&userID)
	if err != nil {
		// TODO: potentiall y return specific error if user does not exist
		return nil, fmt.Errorf("pw reset create %w", err)
	}
	// Build the PasswordReset
	bytesPerToken := service.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("pw reset create: %w", err)
	}
	duration := service.Duration
	if duration == 0 {
		duration = DefaultResetDuration
	}
	pwReset := PasswordReset{
		UserID:    userID,
		Token:     token,
		TokenHash: service.hash(token),
		ExpiresAt: time.Now().Add(duration),
	}

	// INSERT the PasswordReset into the DB
	row = service.DB.QueryRow(`
    INSERT INTO password_resets (user_id, token_hash, expires_at)
    VALUES ($1, $2, $3)
    ON CONFLICT (user_id) DO
    UPDATE
    SET token_hash = $2, expires_at = $3
    RETURNING id;`, pwReset.UserID, pwReset.TokenHash, pwReset.ExpiresAt)
	err = row.Scan(&pwReset.ID)
	if err != nil {
		return nil, fmt.Errorf("pw reset create: %w", err)
	}
	return &pwReset, nil
}

func (service *PasswordResetService) Consume(token string) (*User, error) {
	// 1. validate token; ensure it is not expired
	tokenHash := service.hash(token)
	var user User
	var pwReset PasswordReset
	row := service.DB.QueryRow(`
		SELECT password_resets.id, 
			password_resets.expires_at,
			users.id, 
			users.email, 
			users.password_hash
		FROM password_resets
			JOIN users on users.id = password_resets.user_id
		WHERE password_resets.token_hash = $1;`, tokenHash)
	err := row.Scan(&pwReset.ID, &pwReset.ExpiresAt,
		&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("pw reset consume %w:", err)
	}
	if time.Now().After(pwReset.ExpiresAt) {
		return nil, fmt.Errorf("token expiration %v", token)
	}
	// 2. ensure we have the userid for the token
	// 3. delete the token once consumed
	err = service.delete(pwReset.ID)
	if err != nil {
		return nil, fmt.Errorf("consume can't delete %w", err)
	}
	return &user, nil
}

func (service *PasswordResetService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}

func (service *PasswordResetService) delete(id int) error {
	_, err := service.DB.Exec(`
		DELETE FROM password_resets
		WHERE id = $1;`, id)
	if err != nil {
		return fmt.Errorf("failed to delete reset token %w", err)
	}
	return nil
}
