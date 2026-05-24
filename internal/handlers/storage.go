package handlers

import (
	"sync"
	"test/internal/models"
)

type NoteStorage struct {
	notes  []models.Note
	nextId int
	mu     sync.Mutex
}

func NewNoteStorage(n []models.Note) *NoteStorage {
	return &NoteStorage{
		notes: n,
	}
}
