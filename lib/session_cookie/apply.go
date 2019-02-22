package session_cookie

import (
	"context"
	"github.com/robojones/graphql/lib/session_context"
	"github.com/robojones/graphql/prisma"
	"net/http"
	"time"
)

const CookieKey = "session"

// Set the cookie for the session
func Set(ctx context.Context, session *prisma.Session) {
	w := session_context.Writer(ctx)

	cookie := &http.Cookie{
		Value:    session.Token,
		Name:     CookieKey,
		HttpOnly: true,
		// TODO: Secure: env.Env == env.Production,
		Expires:  time.Now().AddDate(1, 0, 0),
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)
}
