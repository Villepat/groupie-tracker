package server

import "encoding/json"

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
	Dates        string   `json:"dates"`
}

func GetArtists() []Artist {

	var AllArtists []Artist
	data, _ := Harvest("https://groupietrackers.herokuapp.com/api/artists")
	json.Unmarshal(data, &AllArtists)
	return AllArtists
}
