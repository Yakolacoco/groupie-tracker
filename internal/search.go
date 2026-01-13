package internal

import (
	"strings"

	"fyne.io/fyne/v2/widget"
)

// Filtre les artistes
func FilterArtists(all []Artist, query string) []Artist {
	t := strings.ToLower(strings.TrimSpace(query))
	filtered := []Artist{}

	if t == "" {
		return append(filtered, all...)
	}

	for _, ar := range all {
		nameMatch := strings.Contains(strings.ToLower(ar.Name), t)
		memberMatch := false

		for _, m := range ar.Members {
			if strings.Contains(strings.ToLower(m), t) {
				memberMatch = true
				break
			}
		}

		if nameMatch || memberMatch {
			filtered = append(filtered, ar)
		}
	}

	return filtered
}

// Crée la barre de recherche + gère le filtrage
func NewSearchBar(artists []Artist, filtered *[]Artist, list *widget.List) *widget.Entry {
	search := widget.NewEntry()
	search.SetPlaceHolder("Rechercher un artiste ou un membre...")

	search.OnChanged = func(text string) {
		*filtered = FilterArtists(artists, text)
		list.Refresh()
	}

	return search
}
