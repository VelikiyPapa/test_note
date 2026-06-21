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

func (nr *NoteRepo) Create(ctx context.Context, note models.Note) error {
	noteDB := ToNoteDB(note)

	if err := nr.DB.WithContext(ctx).Create(&noteDB).Error; err != nil {
		return err
	}

	return nil
}

func (nr *NoteRepo) Get(ctx context.Context, userID, noteID int) (models.Note, error) {
	var noteDB NoteDB

	if err := nr.DB.WithContext(ctx).Where("user_id = ? AND id = ?", userID, noteID).First(&noteDB).Error; err != nil {
		return models.Note{}, err
	}

	note := ToNoteModel(noteDB)

	return note, nil
}

func (nr *NoteRepo) Delete(ctx context.Context, userID, noteID int) error {
	var noteDB NoteDB

	d := nr.DB.WithContext(ctx).Where("user_id = ? AND id = ?", userID, noteID).Delete(&noteDB)

	if err := d.Error; err != nil {
		return err
	}

	if d.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (nr *NoteRepo) Put(ctx context.Context, userID, noteID int, note models.Note) error {

	p := nr.DB.WithContext(ctx).Model(&NoteDB{}).Where("user_id = ? AND id = ?", userID, noteID).Update("text", note.Text)

	if p.Error != nil {
		return p.Error
	}

	if p.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
