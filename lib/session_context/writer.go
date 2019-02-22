package session_context

import (
	"context"
	"github.com/pkg/errors"
	"net/http"
)

const WriterContextKey = "writer"

var WriterMissingInContextError = errors.Errorf("response writer missing in context (the auth handler should add it)")

func SetWriter(ctx context.Context, w http.ResponseWriter) context.Context {
	return context.WithValue(ctx, WriterContextKey, w)
}

// Writer returns the response writer from the context.
// If the writer is missing, it panics with WriterMissingInContextError.
func Writer(ctx context.Context) http.ResponseWriter {
	w, ok := ctx.Value(WriterContextKey).(http.ResponseWriter)
	if !ok {
		panic(WriterMissingInContextError)
	}
	return w
}
