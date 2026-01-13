package internal

import (
	"strconv"
	"strings"

	"fyne.io/fyne/v2/widget"
)

// ---------------------------------------------------------
// Recherche principale (nom + membres)
// ---------------------------------------------------------

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

// ---------------------------------------------------------
// Barre de recherche
// ---------------------------------------------------------

func NewSearchBar(artists []Artist, filtered *[]Artist, list *widget.List) *widget.Entry {
	search := widget.NewEntry()
	search.SetPlaceHolder("Rechercher un artiste ou un membre...")

	search.OnChanged = func(text string) {
		*filtered = FilterArtists(artists, text)
		list.Refresh()
	}

	return search
}

// ---------------------------------------------------------
// Filtre par décennie (1960 → 1960 à 1969)
// ---------------------------------------------------------

func FilterByYear(all []Artist, year string) []Artist {
	if year == "Toutes" {
		return all
	}

	var filtered []Artist

	// Convertit "1960" → 1960
	start, _ := strconv.Atoi(year)
	end := start + 9 // 1960 → 1969

	for _, a := range all {
		if a.CreationDate >= start && a.CreationDate <= end {
			filtered = append(filtered, a)
		}
	}

	return filtered
}
