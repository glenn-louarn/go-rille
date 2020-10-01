package json

import (
	"encoding/json"
	"net/http"
)

// WriteJSON : write a json to the HTTP response from
// a go structure
func WriteJSON(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// ReadJSON : convert a json from an HTTP request to
// a go structure
func ReadJSON(data interface{}, r *http.Request) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(data)
}
