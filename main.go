package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"strings"

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

	// Barre de recherche
	search := widget.NewEntry()
	search.SetPlaceHolder("Rechercher un artiste ou un membre...")

	// Panneau de droite : image + texte
	img := canvas.NewImageFromResource(nil)
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(300, 300))

	details := widget.NewLabel("Sélectionne un artiste.")
	details.Wrapping = fyne.TextWrapWord

	rightPanel := container.NewVBox(img, details)

	// Liste des artistes (gauche)
	list := widget.NewList(
		func() int { return len(filtered) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i int, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(filtered[i].Name)
		},
	)

	// Clic sur un artiste → image + détails
	list.OnSelected = func(i int) {
		if i < 0 || i >= len(filtered) {
			return
		}

		ar := filtered[i]

		res := loadImageFromURL(ar.Image)
		if res != nil {
			img.Resource = res
			img.Refresh()
		}

		details.SetText(fmt.Sprintf(
			"Nom : %s\nCréation : %d\nPremier album : %s\nMembres : %v",
			ar.Name, ar.CreationDate, ar.FirstAlbum, ar.Members,
		))
	}

	// 🔍 Filtrage par recherche (nom + membres)
	search.OnChanged = func(text string) {
		t := strings.ToLower(strings.TrimSpace(text))
		filtered = filtered[:0]

		if t == "" {
			filtered = append(filtered, artists...)
		} else {
			for _, ar := range artists {
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
		}

		list.Refresh()
	}

	// Colonne gauche : recherche + liste
	leftTop := container.NewVBox(
		widget.NewLabel("Recherche :"),
		search,
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
