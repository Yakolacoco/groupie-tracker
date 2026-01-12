package internal

import (
	"encoding/json"
	"net/http"
)

const baseURL = "https://groupietrackers.herokuapp.com/api"

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
