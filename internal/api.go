package internal

import (
	"encoding/json"
	"net/http"
)

// StartAPI démarre une petite API interne
func StartAPI(artists []Artist) {
	http.HandleFunc("/artists", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(artists)
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// Lance le serveur interne sur le port 8080
	go http.ListenAndServe(":8080", nil)
}
