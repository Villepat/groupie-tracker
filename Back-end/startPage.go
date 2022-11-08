package server

import (
	"html/template"
	"net/http"
)

func StartPage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" || r.Method != "GET" {
		error404(w)
		return
	} else {
		tmpl, err := template.ParseFiles("Front-end/index.html")
		if err != nil {
			error500(w)
			return
			//log.Fatalln(err) //deal with tomorrow lol
		}
		//fmt.Println(AllArtists)
		GetRelations()
		tmpl.ExecuteTemplate(w, "index.html", PageData(w))
		//w.WriteHeader(200)
		//tmpl.Execute(w, nil)
	}
}
