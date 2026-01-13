package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"

	"groupie-tracker/internal"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Petit buffer pour encoder l'image en mémoire
type buffer struct{ b *[]byte }

func (w *buffer) Write(p []byte) (int, error) {
	*w.b = append(*w.b, p...)
	return len(p), nil
}

// Charge une image depuis une URL et la transforme en fyne.Resource
func loadImageFromURL(url string) fyne.Resource {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	img, format, err := image.Decode(resp.Body)
	if err != nil {
		return nil
	}

	buf := []byte{}
	writer := &buffer{&buf}

	if format == "png" {
		_ = png.Encode(writer, img)
	} else {
		_ = jpeg.Encode(writer, img, nil)
	}

	return fyne.NewStaticResource("img", buf)
}

// créer app du projet
func main() {
	a := app.New()
	w := a.NewWindow("Groupie Tracker")

	// Charger les artistes depuis l'API Groupie Tracker
	artists, err := internal.LoadArtists()
	if err != nil {
		w.SetContent(widget.NewLabel("Erreur de chargement des artistes"))
		w.ShowAndRun()
		return
	}

	// Liste filtrée (modifiable)
	filtered := make([]internal.Artist, len(artists))
	copy(filtered, artists)

	// Liste des artistes (gauche)
	list := widget.NewList(
		func() int { return len(filtered) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i int, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(filtered[i].Name)
		},
	)

	// Barre de recherche
	search := internal.NewSearchBar(artists, &filtered, list)

	// bouton filtre
	years := []string{"Toutes", "1960", "1970", "1980", "1990", "2000", "2010"}

	yearSelect := widget.NewSelect(years, func(value string) {
		filtered = internal.FilterByYear(artists, value)
		list.Refresh()
	})

	filters := container.NewHBox(
		yearSelect,
	)

	// Panneau de droite : image + texte
	img := canvas.NewImageFromResource(nil)
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(300, 300))

	details := widget.NewLabel("Sélectionne un artiste.")
	details.Wrapping = fyne.TextWrapWord

	rightPanel := container.NewVBox(img, details)

	// Clic sur un artiste → image + détails + concerts
	list.OnSelected = func(i int) {
		if i < 0 || i >= len(filtered) {
			return
		}

		ar := filtered[i]

		// Charger l'image
		res := loadImageFromURL(ar.Image)
		if res != nil {
			img.Resource = res
			img.Refresh()
		}

		// Infos de base
		info := fmt.Sprintf(
			"Nom : %s\nCréation : %d\nPremier album : %s\nMembres : %v",
			ar.Name, ar.CreationDate, ar.FirstAlbum, ar.Members,
		)

		// Charger les relations (lieux + dates)
		rel, err := internal.LoadRelations(ar.Relations)
		if err == nil {
			info += "\n\nConcerts :\n\n"
			for city, dates := range rel.DatesLocations {
				info += city + " :\n"
				for _, d := range dates {
					info += "  - " + d + "\n"
				}
				info += "\n"
			}
		} else {
			info += "\n\nImpossible de charger les concerts."
		}

		// Afficher dans le panneau de droite
		details.SetText(info)
	}

	// Colonne gauche : recherche + filtres + liste
	leftTop := container.NewVBox(
		widget.NewLabel("Recherche :"),
		search,
		widget.NewLabel("Filtres :"),
		filters,
		widget.NewLabel("Artistes :"),
	)

	listScroll := container.NewVScroll(list)
	listScroll.SetMinSize(fyne.NewSize(600, 500))

	left := container.NewBorder(
		leftTop,
		nil,
		nil,
		nil,
		listScroll,
	)

	// Split gauche/droite
	content := container.NewHSplit(left, rightPanel)
	content.SetOffset(0.35)

	w.SetContent(content)
	w.Resize(fyne.NewSize(1000, 600))
	w.ShowAndRun()
}
