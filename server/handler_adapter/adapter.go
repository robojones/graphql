package handler_adapter

import "net/http"

type HandlerFuncAdapter struct {
	NextFunc http.HandlerFunc
}

func (h *HandlerFuncAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.NextFunc(w, r)
}
