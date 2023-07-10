package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/partagile/lenslocked/rand"
)

const (
	// Minimum # of bytes to be used for each session token
	MinBytesPerToken = 32
)

type Session struct {
	ID     int
	UserID int
	// Token is set when creating a new session.
	// This is left empty when looking up an existing session since
	// only the hash is stored
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
	//BytesPerToken - used to determine # of bytes to use when generating
	//session token. Ignored if not set or < MinBytesPerToken; at which
	//point the value will be MinBytesPerToken
	BytesPerToken int
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	session := Session{
		UserID:    userID,
		Token:     token,
		TokenHash: ss.hash(token),
	}

	row := ss.DB.QueryRow(`
    UPDATE sessions
    SET token_hash = $2
    WHERE user_id = $1
    RETURNING id;`, session.UserID, session.TokenHash)
	err = row.Scan(&session.ID)
	if err == sql.ErrNoRows {
		row = ss.DB.QueryRow(`
      INSERT into sessions (user_id, token_hash)
      VALUES ($1, $2)
      RETURNING id`, session.UserID, session.TokenHash)
		err = row.Scan(&session.ID)
	}
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	tokenHash := ss.hash(token)
	// Query DB for session that matches the hash
	var user User
	row := ss.DB.QueryRow(`
    SELECT user_id
    FROM sessions
    WHERE token_hash = $1`, tokenHash)
	err := row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}
	// Using above matched userid, query for that user
	row = ss.DB.QueryRow(`
    SELECT email, password_hash
    FROM users
    WHERE id = $1;`, user.ID)
	err = row.Scan(&user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}
	return &user, nil
}

func (ss *SessionService) Delete(token string) error {
	tokenHash := ss.hash(token)
	_, err := ss.DB.Exec(`
    DELETE FROM sessions
    WHERE token_hash = $1;`, tokenHash)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return nil
}

func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}