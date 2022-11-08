package server

import (
	"html/template"
	"net/http"
)

type ErrorInfo struct {
	ErrorNum     int
	ErrorMessage string
}

var Err ErrorInfo

func errorPage(w http.ResponseWriter) {
	tmpl, _ := template.ParseFiles("Front-end/error.html")
	tmpl.ExecuteTemplate(w, "error.html", Err)
}

func error400(w http.ResponseWriter) {
	w.WriteHeader(400)
	Err.ErrorMessage = "Bad request"
	Err.ErrorNum = 400
	//w.Write([]byte("400 Bad request")
	errorPage(w)
}
func error404(w http.ResponseWriter) {
	w.WriteHeader(404)
	Err.ErrorMessage = "Page not found"
	Err.ErrorNum = 404
	//w.Write([]byte("404 Page not found"))
	errorPage(w)
}
func error500(w http.ResponseWriter) {
	w.WriteHeader(500)
	Err.ErrorMessage = "Internal Server Error"
	Err.ErrorNum = 500
	errorPage(w)
	//w.Write([]byte("500 Internal Server Error"))
}
