package internal

// Artist représente la structure d'un artiste telle qu'elle est renvoyée par l'API Groupie Tracker.
type Artist struct {
	ID           int      `json:"id"`           // Identifiant unique de l'artiste
	Image        string   `json:"image"`        // URL de l'image de l'artiste
	Name         string   `json:"name"`         // Nom du groupe ou de l'artiste
	Members      []string `json:"members"`      // Liste des membres du groupe
	CreationDate int      `json:"creationDate"` // Année de création du groupe
	FirstAlbum   string   `json:"firstAlbum"`   // Date de sortie du premier album
	Locations    string   `json:"locations"`    // URL vers l'API des lieux
	ConcertDates string   `json:"concertDates"` // URL vers l'API des dates de concerts
	Relations    string   `json:"relations"`    // URL vers l'API des relations (lieux + dates)
}

// Relations contient les lieux et les dates de concerts d'un artiste
type Relations struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}
