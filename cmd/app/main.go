package main

import (
	"log"
	"net/http"
	"test/internal/config"
	"test/internal/delivery/my_http"
	"test/internal/repo/postgres"
	"test/internal/service"
	"time"
)

// TODO создать в сервисе файл note_use_case и реализовать интерфейс

// DONE добавить параметры для настрйоки бд

// DONE над каждым шагом накинуть логи

func main() {
	log.Println("получаем конфиги")
	cfg := config.Load()

	log.Println("подключаемся к бд")
	db, err := postgres.Open(cfg.DbConfig.Dsn, 10, 5, time.Hour)
	if err != nil {
		log.Fatal()
	}

	log.Println("начинаем автомиграцию")
	if err := postgres.Migrate(db); err != nil {
		log.Fatal()
	}

	log.Println("создаем слой репозитория")
	noteRepo := postgres.NewNoteRepo(db)

	log.Println("создаем слой бизнес логики")
	noteService := service.NewNoteService(noteRepo)

	// DONE реализовать роутер (транспортный слой)
	log.Println("реализуем транспортный слой (роутер)")
	noteHandler := my_http.NewNoteHandler(noteService)
	mux := my_http.NewRouter(noteHandler)

	log.Println("Сервер запустился на 8080")
	http.ListenAndServe(":8080", mux)
}
