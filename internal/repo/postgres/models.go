package postgres

import (
	"test/internal/models"
	"time"
)

type UserDB struct {
	ID        int       `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}

func (UserDB) TableName() string {
	return "users"
}

type NoteDB struct {
	ID        int       `gorm:"column:id;primaryKey"`
	UserID    int       `gorm:"column:user_id;not null"`
	Text      string    `gorm:"column:text;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}

func (NoteDB) TableName() string {
	return "notes"
}

func ToNoteDB(mn models.Note) NoteDB {
	return NoteDB{
		Text: mn.Text,
	}
}

func ToNoteModel(nb NoteDB) models.Note {
	return models.Note{
		ID:   nb.ID,
		Text: nb.Text,
	}
}
