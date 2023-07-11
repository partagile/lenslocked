package context

import (
	"context"

	"github.com/partagile/lenslocked/models"
)

type key string

const (
	userKey key = "user"
)

func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func User(ctx context.Context) *models.User {
	val := ctx.Value(userKey)
	user, ok := val.(*models.User)
	if !ok {
		// This shouldn't happen. If it does, it's because:
		// 1. Nothing was stored in the context and thus no *models.User type
		// 2. (Unlikely but..) other code wrote an invalid value using user key.
		return nil
	}
	return user
}
