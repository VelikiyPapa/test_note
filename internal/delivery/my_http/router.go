package my_http

import "net/http"

// DONE переписать роутер
func NewRouter(ns *NoteHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("POST /notes", LoggingMiddleware(http.HandlerFunc(ns.CreateNote)))
	mux.Handle("GET /notes/{id}", LoggingMiddleware(http.HandlerFunc(ns.GetNote)))
	mux.Handle("DELETE /notes/{id}", LoggingMiddleware(http.HandlerFunc(ns.DeleteNote)))
	mux.Handle("PUT /notes/{id}", LoggingMiddleware(http.HandlerFunc(ns.PutNote)))

	return mux
}
