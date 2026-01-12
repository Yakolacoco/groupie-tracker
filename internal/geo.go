package internal

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type GeoResult struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func GeoCity(city string) (string, string, error) {
	endpoint := "https://nominatim.openstreetmap.org/search?q=" +
		url.QueryEscape(city) + "&format=json&limit=1"

	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("User-Agent", "GroupieTracker-App")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	var results []GeoResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return "", "", err
	}

	if len(results) == 0 {
		return "", "", nil
	}

	return results[0].Lat, results[0].Lon, nil
}
