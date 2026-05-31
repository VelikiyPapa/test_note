package postgres

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open(dsn string, maxOpenConns, maxIdleConns int, connMaxLifeTime time.Duration) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sql, err := db.DB()
	if err != nil {
		return nil, err
	}

	// TODO - почитать другие set
	sql.SetMaxOpenConns(maxOpenConns)
	sql.SetMaxIdleConns(maxIdleConns)
	sql.SetConnMaxLifetime(connMaxLifeTime)

	return db, nil
}

// DONE - добавить миграцию
func Migrate(db *gorm.DB) error {
	log.Println("автомиграция началась")

	err := db.AutoMigrate(&NoteDB{})
	if err != nil {
		return err
	}

	log.Println("автомиграция успешно завершилась")

	return nil
}
