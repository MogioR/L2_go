package restutills

import (
	"encoding/json"
	"net/http"
)

// Cообщение о результате
type Result struct {
	Msg interface{} `json:"result"`
}

// Успешный ответ
func SendResponce(w http.ResponseWriter, code int, msg interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)

	resResp := Result{msg}
	if err := json.NewEncoder(w).Encode(resResp); err != nil {
		http.Error(w, "responseJson error", http.StatusInternalServerError)
	}
}
