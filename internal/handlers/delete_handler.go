package handlers

import (
	"log"
	"net/http"
	"strconv"
)

func (ns *NoteStorage) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// ns.mu.Lock()
	// defer ns.mu.Unlock()
	
	log.Println(r.Method)
	if r.Method != http.MethodDelete {
		http.Error(w, "Неверный метод", http.StatusMethodNotAllowed)
		return
	}

	idFromURL := r.PathValue("id")
	id, err := strconv.Atoi(idFromURL)
	if err != nil {
		http.Error(w, "Неправильный ID", http.StatusBadRequest)
		return
	}

	for index, note := range ns.notes {
		if note.ID == id {
			ns.notes = append(ns.notes[:index], ns.notes[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			log.Println(ns.notes)
			return
		}
	}

	http.Error(w, "Заметка на найдена", http.StatusNotFound)
}
