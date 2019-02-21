package auth

import (
	"context"
	"github.com/pkg/errors"
	"github.com/robojones/graphql/prisma"
	"net/http"
)

var WriterMissingInContext = errors.Errorf("response writer missing in context (the auth handler should add it)")

func ApplySession(ctx context.Context, session *prisma.Session) {
	w, ok := ctx.Value(ResponseContextKey).(http.ResponseWriter)

	if !ok {
		panic(WriterMissingInContext)
	}

	http.SetCookie(w, CreateCookie(session))
}
