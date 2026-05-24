package shortcut

import (
	"encoding/json"
	"net/http"
)

func SendJSON(w http.ResponseWriter, code int, data any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

// TODO подумать насчет дженерика

// TODO подумать насчет интерфейса в аргументах

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		http.Error(w, "Не смог распарсить json", http.StatusBadRequest)
		return
	}
}
