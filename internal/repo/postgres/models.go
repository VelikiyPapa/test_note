package postgres

import (
	"test/internal/models"
	"time"
)

type NoteDB struct {
	ID        int       `gorm:"column:id;primaryKey"`
	Text      string    `gorm:"column:text"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatetAt time.Time `gorm:"column:updated_at;not null"`
}

func(NoteDB) TableName() string {
	return "notes"
}

func ToNoteDB(mn models.Note) NoteDB {
	return NoteDB{
		Text: mn.Text,
	}
}

func ToNoteModel(nb NoteDB) models.Note {
	return models.Note{
		ID: nb.ID,
		Text: nb.Text,
	}
}
