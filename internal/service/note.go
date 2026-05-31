package service

import (
	"context"
	"errors"
	"strings"
	"test/internal/models"
)

type NoteRepo interface {
	Create(ctx context.Context, note models.Note) error
	Get(ctx context.Context, id int) (models.Note, error)
	Delete(ctx context.Context, id int) error
}

type NoteService struct {
	Repo NoteRepo
}

func NewNoteService(ns NoteRepo) *NoteService {
	return &NoteService{
		Repo: ns,
	}
}

func (ns *NoteService) CreateNote(ctx context.Context, note models.Note) error {
	if len(note.Text) < 3 {
		return errors.New("слишком короткий текст")
	}

	note.Text = strings.TrimSpace(note.Text)

	return ns.Repo.Create(ctx, note)
}

// DONE прописать get и delete
func (ns *NoteService) GetNote(ctx context.Context, id int) (models.Note, error) {
	if id <= 0 {
		return models.Note{}, errors.New("некорректный id")
	}

	return ns.Repo.Get(ctx, id)
}

func (ns *NoteService) DeleteNote(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("некорректный id")
	}

	return ns.Repo.Delete(ctx, id)
}
