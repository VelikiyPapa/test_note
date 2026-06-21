package my_http

import "net/http"

func NewRouter(ns *NoteHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users/{user_id}/notes", ns.CreateNote)
	mux.HandleFunc("GET /users/{user_id}/notes/{id}", ns.GetNote)
	mux.HandleFunc("DELETE /users/{user_id}/notes/{id}", ns.DeleteNote)
	mux.HandleFunc("PUT /users/{user_id}/notes/{id}", ns.PutNote)

	return Chain(mux, LoggingMiddleware, RecoverMiddleware)
}
