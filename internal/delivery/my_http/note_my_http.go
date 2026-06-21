package my_http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"test/internal/models"
	"test/internal/shortcut"
)

type NoteService interface {
	CreateNote(ctx context.Context, note models.Note) error
	GetNote(ctx context.Context, userID, noteID int) (models.Note, error)
	DeleteNote(ctx context.Context, userID, noteID int) error
	PutNote(ctx context.Context, userID, noteID int, note models.Note) error
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

	shortcut.ReadJSON(w, r, &note)

	err := nh.Service.CreateNote(r.Context(), note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := map[string]string{
		"msg": "Заметка успешно создана",
	}

	shortcut.SendJSON(w, http.StatusCreated, result)
}

func (nh *NoteHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	rawNoteID := r.PathValue("id")
	noteID, err := strconv.Atoi(rawNoteID)
	if err != nil {
		http.Error(w, "некорректный note id", http.StatusBadRequest)
		return
	}

	rawUserID := r.PathValue("user_id")
	userID, err := strconv.Atoi(rawUserID)
	if err != nil {
		http.Error(w, "некорректный user id", http.StatusBadRequest)
		return
	}

	note, err := nh.Service.GetNote(r.Context(), userID, noteID)
	if err != nil {
		http.Error(w, "не удалось получить карту", http.StatusInternalServerError)
		return
	}

	shortcut.SendJSON(w, http.StatusOK, note)
}

func (nh *NoteHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	rawNoteID := r.PathValue("id")
	noteID, err := strconv.Atoi(rawNoteID)
	if err != nil {
		http.Error(w, "некорректный note id", http.StatusBadRequest)
		return
	}

	rawUserID := r.PathValue("user_id")
	userID, err := strconv.Atoi(rawUserID)
	if err != nil {
		http.Error(w, "некорректный user id", http.StatusBadRequest)
		return
	}

	err = nh.Service.DeleteNote(r.Context(), userID, noteID)
	if err != nil {
		http.Error(w, "не получилось удалить заметку", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (nh *NoteHandler) PutNote(w http.ResponseWriter, r *http.Request) {
	rawNoteID := r.PathValue("id")
	noteID, err := strconv.Atoi(rawNoteID)
	if err != nil {
		http.Error(w, "некорректный note id", http.StatusBadRequest)
		return
	}

	rawUserID := r.PathValue("user_id")
	userID, err := strconv.Atoi(rawUserID)
	if err != nil {
		http.Error(w, "некорректный user id", http.StatusBadRequest)
		return
	}

	var note models.Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "не удалось распарсить json", http.StatusBadRequest)
		return
	}

	err = nh.Service.PutNote(r.Context(), userID, noteID, note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := map[string]string{
		"msg": "Заметка успешно обновлена",
	}

	shortcut.SendJSON(w, http.StatusOK, result)
}
