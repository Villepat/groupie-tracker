package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
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
	url := "https://groupietrackers.herokuapp.com/api/artists"
	res, err := http.Get(url) //get contents of API
	if err != nil {
		fmt.Println("ERRRRR")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body) //read body of contents
	if err != nil {
		// handle error
		return
	}
	var AllArtists []Artist
	json.Unmarshal(body, &AllArtists)
	fmt.Println(AllArtists[0].Members[0]) //first member or first artist -> Freddieeee
	fmt.Println(AllArtists[0].Name)
	fmt.Print(AllArtists[0])

}
