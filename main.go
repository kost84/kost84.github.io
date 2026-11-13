package main

import (

	"html/template"
	"log"
	"net/http"
	"fmt"
)

func main() {

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))

	http.HandleFunc("/", index)

	http.HandleFunc("/acceuil", acceuil)

	http.HandleFunc("/project", project)

	http.HandleFunc("/contact", contact)

	fmt.Println("Serveur lancé à l'adresse suivante : http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request){
	renderPage(w, "index.html")
}

func acceuil(w http.ResponseWriter, r *http.Request){
	renderPage(w, "acceuil.html")
}

func project(w http.ResponseWriter, r *http.Request){
	renderPage(w, "project.html")
}

func contact(w http.ResponseWriter, r *http.Request){
	renderPage(w, "contact.html")
}

func renderPage(w http.ResponseWriter, tmpl string) {

	t, err := 	template.ParseFiles("page/" + tmpl)
	if err != nil{
		http.Error(w, err.Error(), http.StatusAlreadyReported)
		return
	}
	t.Execute(w, nil)
}