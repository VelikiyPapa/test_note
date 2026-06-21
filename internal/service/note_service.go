package service

import (
	"context"
	"errors"
	"strings"
	"test/internal/models"
)

type NoteRepo interface {
	Create(ctx context.Context, note models.Note) error
	Get(ctx context.Context, userID, noteID int) (models.Note, error)
	Delete(ctx context.Context, userID, noteID int) error
	Put(ctx context.Context, userID, noteID int, note models.Note) error
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
	note.Text = strings.TrimSpace(note.Text)

	if len(note.Text) < 3 {
		return errors.New("слишком короткий текст")
	}

	return ns.Repo.Create(ctx, note)
}

func (ns *NoteService) GetNote(ctx context.Context, userID, noteID int) (models.Note, error) {
	if noteID <= 0 || userID <= 0 {
		return models.Note{}, errors.New("некорректный id")
	}

	return ns.Repo.Get(ctx, userID, noteID)
}

func (ns *NoteService) DeleteNote(ctx context.Context, userID, noteID int) error {
	if noteID <= 0 || userID <= 0 {
		return errors.New("некорректный id")
	}

	return ns.Repo.Delete(ctx, userID, noteID)
}

func (ns *NoteService) PutNote(ctx context.Context, userID, noteID int, note models.Note) error {
	if noteID <= 0 || userID <= 0 {
		return errors.New("некорректный id")
	}

	note.Text = strings.TrimSpace(note.Text)

	if len(note.Text) < 3 {
		return errors.New("слишком короткий текст")
	}

	return ns.Repo.Put(ctx, userID, noteID, note)
}
