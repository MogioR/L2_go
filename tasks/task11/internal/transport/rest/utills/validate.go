package restutills

import (
	"fmt"
	"net/http"
)

// ValidateQuery для валидации метода и URL-query
func ValidateQuery(w http.ResponseWriter, r *http.Request, validateQuery ...string) bool {
	if r.Method != validateQuery[0] {
		SendError(w, http.StatusMethodNotAllowed, fmt.Sprintf("bad %v method", r.Method))
		return false
	}
	for _, v := range validateQuery[1:] {
		if !r.URL.Query().Has(v) {
			SendError(w, http.StatusServiceUnavailable, "not parameters "+v)
			return false
		}
	}
	return true
}
