package my_http

import "net/http"

// DONE переписать руками роутер
func NewRouter(ns *NoteHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /note", ns.CreateNote)
	mux.HandleFunc("GET /note/{id}", ns.GetNote)
	mux.HandleFunc("DELETE /note/{id}", ns.DeleteNote)

	return mux
}
