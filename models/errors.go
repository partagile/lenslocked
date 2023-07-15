package models

import "errors"

var (
	ErrEmailTaken = errors.New("models: email address already in use")
	ErrNotFound   = errors.New("models: gallery by id not found")
)
