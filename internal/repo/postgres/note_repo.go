package postgres

import (
	"context"
	"test/internal/models"

	"gorm.io/gorm"
)

type NoteRepo struct {
	DB *gorm.DB
}

func NewNoteRepo(db *gorm.DB) *NoteRepo {
	return &NoteRepo{
		DB: db,
	}
}

// CreateNote
func(nr *NoteRepo) Create(ctx context.Context, note models.Note) error {
	noteDB := ToNoteDB(note)

	if err := nr.DB.WithContext(ctx).Create(noteDB).Error; err != nil {
		return err
	}
	
	return nil
}

// GetNote
func(nr *NoteRepo) Get(ctx context.Context, id int) (models.Note, error) {
	var noteDB NoteDB

	if err := nr.DB.WithContext(ctx).First(&noteDB, id).Error; err != nil {
		return models.Note{}, err
	}

	note := ToNoteModel(noteDB)

	return note, nil
}

// DeleteNote
func(nr *NoteRepo) Delete(ctx context.Context, id int) error {
	var noteDB NoteDB

	if err := nr.DB.WithContext(ctx).Delete(&noteDB, id).Error; err != nil {
		return err
	}

	return nil
}
