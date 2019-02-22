package session_cookie

import (
	"context"
	"github.com/robojones/graphql/server/session_context"
	"net/http"
)

// Unset the session cookie
func Unset(ctx context.Context) {
	w := session_context.Writer(ctx)

	cookie := &http.Cookie{
		Name:   CookieKey,
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
}
