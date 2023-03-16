package restutills

import (
	"encoding/json"
	"net/http"
)

// Cообщение об ошибке
type Error struct {
	Msg string `json:"error"`
}

// Сообщение об ошибке
func SendError(w http.ResponseWriter, code int, msg interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)

	errResp := Error{msg.(string)}
	if err := json.NewEncoder(w).Encode(errResp); err != nil {
		http.Error(w, "responseJson error", http.StatusInternalServerError)
	}
}
