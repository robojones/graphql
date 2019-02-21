package auth

import (
	"github.com/robojones/graphql/prisma"
	"net/http"
	"time"
)

func CreateCookie(session *prisma.Session) *http.Cookie {
	return &http.Cookie{
		Value:    session.Token,
		Name:     CookieKey,
		HttpOnly: true,
		// TODO: Secure: env.Env == env.Production,
		Expires:  time.Now().AddDate(1, 0, 0),
		SameSite: http.SameSiteStrictMode,
	}
}
