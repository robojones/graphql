package auth

import (
	"context"
	"github.com/robojones/graphql/lib/session_context"
	"github.com/robojones/graphql/prisma"
	"net/http"
	"time"
)

const CookieKey = "session"

// SetCookie the cookie for the session
func SetCookie(ctx context.Context, session *prisma.Session) {
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

// UnsetCookie the session cookie
func UnsetCookie(ctx context.Context) {
	w := session_context.Writer(ctx)

	cookie := &http.Cookie{
		Name:   CookieKey,
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
}
