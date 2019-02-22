package session_context

import (
	"context"
)

const TokenContextKey = "session_token"

// SetToken returns a context that contains the token value.
func SetToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, TokenContextKey, token)
}

// Token returns the token from the context.
// If the token is missing UserNotLoggedInError is returned.
func Token(ctx context.Context) (string, error) {
	user, ok := ctx.Value(TokenContextKey).(string)

	if !ok {
		return "", UserNotLoggedInError
	}

	return user, nil
}
