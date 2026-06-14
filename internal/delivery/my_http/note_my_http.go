package my_http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"test/internal/models"
)

type NoteService interface {
	CreateNote(ctx context.Context, note models.Note) error
	GetNote(ctx context.Context, id int) (models.Note, error)
	DeleteNote(ctx context.Context, id int) error
	PutNote(ctx context.Context, id int, note models.Note) error
}

type NoteHandler struct {
	Service NoteService
}

func NewNoteHandler(ns NoteService) *NoteHandler {
	return &NoteHandler{
		Service: ns,
	}
}

func (nh *NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "не удалось распарсить json", http.StatusBadRequest)
		return
	}

	err := nh.Service.CreateNote(r.Context(), note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := map[string]string{
		"msg": "Заметка успешно создана",
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (nh *NoteHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	rawID := r.PathValue("id")
	if rawID == "" {
		http.Error(w, "некорректный id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(rawID)
	if err != nil {
		http.Error(w, "некорректный id", http.StatusBadRequest)
		return
	}

	note, err := nh.Service.GetNote(r.Context(), id)
	if err != nil {
		http.Error(w, "не удалось получить карту", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func (nh *NoteHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	rawId := r.PathValue("id")
	if rawId == "" {
		http.Error(w, "некорректный id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		http.Error(w, "некорректный id", http.StatusBadRequest)
		return
	}

	err = nh.Service.DeleteNote(r.Context(), id)
	if err != nil {
		http.Error(w, "не получилось удалить заметку", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (nh *NoteHandler) PutNote(w http.ResponseWriter, r *http.Request) {
	rawID := r.PathValue("id")
	if rawID == "" {
		http.Error(w, "некорректный id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(rawID)
	if err != nil {
		http.Error(w, "некорректный id", http.StatusBadRequest)
		return
	}

	var note models.Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "не удалось распарсить json", http.StatusBadRequest)
		return
	}

	err = nh.Service.PutNote(r.Context(), id, note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := map[string]string{
		"msg": "Заметка успешно обновлена",
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}