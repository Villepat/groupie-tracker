package server

import (
"html/template"
"log"
"net/http"
)

func main () {
	func startPage(w http.ResponseWriter, r *http.Request) {
		fs := http.FileServer(http.Dir("Front-end/"))
		http.Handle("/src/", http.StripPrefix("/src/", fs))
		http.HandleFunc("/", StartPage)
		http.HandleFunc("/ascii-art", SubmitData)
		fmt.Println("Server is listening to port #8080...")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
		fmt.Printf("Server error: %v", err)
}