package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"fmt"
)

type HtmlData struct {
	ArtistData   []Artist
	RelationData Relation
}

type relation struct {
	Index []struct {
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type AllData struct {
	Id             int
	Image          string
	Name           string
	Members        []string
	CreationDate   int
	FirstAlbum     string
	DatesLocations map[string][]string `json:"datesLocations"`
}

var IndexPageData HtmlData

//var Popup []AllData
var ArtistAPI = GetArtists()
var RelationAPI = GetRelations()

func PageData(w http.ResponseWriter) []AllData {
	//IndexPageData.ArtistData = GetArtists()
	//IndexPageData.RelationData = GetRelations()
	//relData := GetRelations()
	popup := []AllData{}
	Testo := relation{}
	data, err1 := Harvest("https://groupietrackers.herokuapp.com/api/artists")
	data2, err2 := Harvest("https://groupietrackers.herokuapp.com/api/relation")
	if err1 != nil || err2 != nil {
		error500(w)
	}
	json.Unmarshal(data, &popup)
	json.Unmarshal(data2, &Testo)
	//fmt.Println(Testo)
	test, _ := json.Marshal(Testo.Index)
	//fmt.Println("test marshalled", string(test))
	json.Unmarshal(test, &popup)
	//fmt.Println("test",string(test))
	//fmt.Println("relData:",len(relData))
	//fmt.Println("Alldata:", len(Popup))
	//fmt.Println("shouldWork", Popup[0])
	//fmt.Println("???", string(data2))

	//fmt.Println(Popup)
	//artData := GetArtists()
	//relData := GetRelations()
	//Popup = GetArtists()
	// for i, strct := range artData {
	// 	Popup.id = strct.Id
	// 	Popup.Image = strct.Image
	// 	Popup.Name = strct.Name
	// 	Popup.Members = strct.Members
	// 	Popup.CreationDate = strct.CreationDate
	// 	Popup.FirstAlbum = strct.FirstAlbum
	// }
	//Popup.DatesLocations = artData.DatesLocations

	//fmt.Println()
	fmt.Println("POPOPOPOPOPOPOPOPOPOPOPUP", popup)
	//fmt.Println()
	return popup
}

// func MasterData(IDnm int) AllData {
// 	Popup.Id = IDnm
// 	Popup.Name = ArtistAPI[IDnm].Name
// 	Popup.Image = ArtistAPI[IDnm].Image
// 	Popup.Members = ArtistAPI[IDnm].Members
// 	Popup.CreationDate = ArtistAPI[IDnm].CreationDate
// 	Popup.FirstAlbum = ArtistAPI[IDnm].FirstAlbum
// 	Popup.DatesLocations = RelationAPI.Index[IDnm].DatesLocations
// }
