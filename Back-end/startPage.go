package server

import (
	"html/template"
	"log"
	"net/http"
)

func StartPage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" || r.Method != "GET" {
		error404(w)
		return
	} else {
		tmpl, err := template.ParseFiles("Front-end/index.html")
		if err != nil {
			log.Fatalln(err) //deal with tomorrow lol
		}
		//fmt.Println(AllArtists)
		GetRelations()
		tmpl.ExecuteTemplate(w, "index.html", GetArtists())
		//w.WriteHeader(200)
		//tmpl.Execute(w, nil)
	}
}
