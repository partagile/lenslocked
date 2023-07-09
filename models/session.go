package models

import "database/sql"

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
	// TODO: create session token
	// TODO: implement SessionService.Create
	return nil, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement SessionService.User
	return nil, nil
}
