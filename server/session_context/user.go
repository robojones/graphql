package session_context

import (
	"context"
	"github.com/pkg/errors"
	"github.com/robojones/graphql/prisma"
)

const UserContextKey = "user"

var UserNotLoggedInError = errors.Errorf("user not logged in")

// SetUser returns a context that includes the user value.
func SetUser(ctx context.Context, user *prisma.User) context.Context {
	return context.WithValue(ctx, UserContextKey, user)
}

// User returns the user value from the context.
// If the user value is missing UserNotLoggedInError is returned.
func User(ctx context.Context) (*prisma.User, error) {
	user, ok := ctx.Value(UserContextKey).(*prisma.User)

	if !ok {
		return nil, UserNotLoggedInError
	}

	return user, nil
}
