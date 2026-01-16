package internal

import (
	"encoding/json"
	"net/http"
)

// URL de base de l'API Groupie Tracker
const baseURL = "https://groupietrackers.herokuapp.com/api"

// LoadArtists récupère la liste complète des artistes depuis l'API Groupie Tracker.
func LoadArtists() ([]Artist, error) {
	resp, err := http.Get(baseURL + "/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []Artist
	err = json.NewDecoder(resp.Body).Decode(&data)
	return data, err
}

// LoadRelations récupère les lieux + dates de concerts depuis l'URL fournie
func LoadRelations(url string) (Relations, error) {
	var rel Relations

	resp, err := http.Get(url)
	if err != nil {
		return rel, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&rel); err != nil {
		return rel, err
	}

	return rel, nil
}
