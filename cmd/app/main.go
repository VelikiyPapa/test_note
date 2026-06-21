package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"test/internal/config"
	"test/internal/delivery/my_http"
	"test/internal/repo/postgres"
	"test/internal/service"
	"time"
)

// DONE Добавить сущность user
// TODO Учитывать в заметках userа (заметка должна принадлежать ему)
// DONE Использовать пакет sc

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

	log.Println("реализуем транспортный слой (роутер)")
	noteHandler := my_http.NewNoteHandler(noteService)
	mux := my_http.NewRouter(noteHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		log.Println("Сервер запустился на 8080")

		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

	log.Println("получен сигнал остановки сервера")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatal(err)
	}

	log.Println("сервер корректно остановлен")
}
