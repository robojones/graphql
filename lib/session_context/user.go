package session_context

import (
	"context"
	"github.com/robojones/graphql/prisma"
	"github.com/vektah/gqlparser/gqlerror"
)

const UserContextKey = "user"

var UserNotLoggedInError = &gqlerror.Error{
	Message: "user not logged in",
	Extensions: map[string]interface{}{
		"type": "Auth",
		"name": "NotLoggedIn",
	},
}

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
