package helper

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// RespondWithError : Respond with error
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON : Respond with json
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.Encode(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(buf.Bytes())
}
