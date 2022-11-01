package server

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func StartPage(w http.ResponseWriter, r *http.Request) {
	var AllArtists []Artist
	data, _ := Harvest("https://groupietrackers.herokuapp.com/api/artists")
	json.Unmarshal(data, &AllArtists)
	if r.URL.Path != "/" || r.Method != "GET" {
		error404(w)
		return
	} else {
		tmpl, err := template.ParseFiles("Front-end/index.html")
		if err != nil {
			log.Fatalln(err) //deal with tomorrow lol
		}
		//fmt.Println(AllArtists)
		tmpl.ExecuteTemplate(w, "index.html", AllArtists)
		//w.WriteHeader(200)
		//tmpl.Execute(w, nil)
	}
}
