package main

import (
	"log"
	"net/http"
	"test/internal/handlers"
	"test/internal/models"
)

/* 
// TODO
скачать плагин

создать пакет с конфигами
создать методы getByID, Delete в слое репозитория
попробовать реализовать автомиграцию
создать в сервисе файл note_use_case и реализовать интерфейс
подключить git

*/



func main() {
	n := make([]models.Note, 0)

	newStorage := handlers.NewNoteStorage(n)
	mux := handlers.NewRouter(newStorage)

	log.Println("Сервер запустился на 8080")
	http.ListenAndServe(":8080", mux)
}
