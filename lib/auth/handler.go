package auth

import (
	"github.com/robojones/graphql/lib/session_context"
	"github.com/robojones/graphql/prisma"
	"net/http"
)

const CookieKey = "session"

type Handler struct {
	Prisma *prisma.Client
	Next   http.Handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(CookieKey)
	ctx := session_context.SetWriter(r.Context(), w)

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
		// session removed or invalid
		h.Next.ServeHTTP(w, r.WithContext(ctx))
		return
	}

	ctx = session_context.SetToken(ctx, token)
	ctx = session_context.SetUser(ctx, user)

	h.Next.ServeHTTP(w, r.WithContext(ctx))
}
