package display

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Text  string `json:"text"`
	Error string `json:"error"`
}

func DisplayName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	w.Header().Set("Content-Type", "application/json")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(Response{Error: "empty name"})

		return
	}

	hmessage := fmt.Sprintf("Привет от %s!", name)
	json.NewEncoder(w).Encode(Response{Text: hmessage})

	w.WriteHeader(http.StatusOK)
}
