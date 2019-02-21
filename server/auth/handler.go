package auth

import (
	"context"
	"github.com/robojones/graphql/prisma"
	"net/http"
)

const CookieKey = "session"
const UserContextKey = "user"
const ResponseContextKey = "r"

type Handler struct {
	Prisma *prisma.Client
	Next   http.Handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(CookieKey)
	ctx := context.WithValue(r.Context(), ResponseContextKey, w)

	if err == http.ErrNoCookie {
		h.Next.ServeHTTP(w, r.WithContext(ctx))
		return
	} else if err != nil {
		panic(err)
	}

	token := cookie.Value

	user, err := h.Prisma.Session(prisma.SessionWhereUniqueInput{
		Token: &token,
	}).User().Exec(r.Context())

	if err != nil {
		panic(err)
	}

	if user == nil {
		// session expired
		h.Next.ServeHTTP(w, r.WithContext(ctx))
		return
	}

	ctx = context.WithValue(ctx, UserContextKey, user)

	h.Next.ServeHTTP(w, r.WithContext(ctx))
}
