package handlers

import (
	"net/http"
)

func NewRouter(ns *NoteStorage) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /note", ns.CreateHandler)
	mux.HandleFunc("DELETE /note/{id}", ns.DeleteHandler)

	return mux
}
