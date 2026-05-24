package handlers

import (
	"log"
	"net/http"
	"strings"
	"test/internal/models"
	"test/internal/shortcut"
)

func (ns *NoteStorage) CreateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method != http.MethodPost {
		http.Error(w, "Неверный метод", http.StatusMethodNotAllowed)
		return
	}

	// Задача хэндлера - обработать json файл и отдать дальше

	var note models.Note
	shortcut.ReadJSON(w, r, &note)
	log.Println("Декодировали json", note)

	// Бизнес логика - проверка бизнес требований

	note.Text = strings.TrimSpace(note.Text)
	if note.Text == "" || len(note.Text) < 2 {
		http.Error(w, "Введи текст", http.StatusBadRequest)
		return
	}

	// Репозиторий - создание записи

	ns.mu.Lock()

	ns.nextId++
	note.ID = ns.nextId
	ns.notes = append(ns.notes, note)

	log.Println(ns.nextId)
	log.Println(ns.notes)
	ns.mu.Unlock()

	shortcut.SendJSON(w, http.StatusCreated, note)
}
