package models

import (
	"database/sql"
	"fmt"

	"github.com/partagile/lenslocked/rand"
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
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	token, err := rand.SessionToken()
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	// TODO: hash the session token
	session := Session{
		UserID: userID,
		Token:  token,
		// TODO: set TokenHash
	}
	// Store session in the DB
	// TODO: implement SessionService.Create
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement SessionService.User
	return nil, nil
}
