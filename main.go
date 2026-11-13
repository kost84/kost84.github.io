package main

import (

	//"encoding/json"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Projet struct {
	Titre      string `json:"Titre"`
	Contexte   string `json:"Contexte"`
	LienGithub string `json:"LienGithub"`
	Image      string `json:"Image"`
}

func main() {

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.Handle("/imageetdoc/", http.StripPrefix("/imageetdoc/", http.FileServer(http.Dir("imageetdoc/"))))

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
func renderPage(w http.ResponseWriter, tmpl string, data any) {

	t, err := 	template.ParseFiles("page/" + tmpl)
	if err != nil{
		http.Error(w, err.Error(), http.StatusAlreadyReported)
		return
	}
	t.Execute(w, data)
}

func index(w http.ResponseWriter, r *http.Request){
	renderPage(w, "index.html", nil)
}

func acceuil(w http.ResponseWriter, r *http.Request){
	renderPage(w, "acceuil.html",nil)
}

func project(w http.ResponseWriter, r *http.Request){
	renderPage(w, "project.html",nil)
}

func contact(w http.ResponseWriter, r *http.Request){
	renderPage(w, "contact.html",nil)
}


func project(w http.ResponseWriter, r *http.Request){

	fichier, err := os.ReadFile("prokect.json")
	if err != nil {
		http.Error(w, "Erreur de lecture du JSON", http.StatusInternalServerError)
		return
	}

	var mesProjects []Projet
	err = json.Unmarshal(fichier, &mesProjects)
	if err != nil {
		http.Error(w, "Erreur de décodage JSON", http.StatusInternalServerError)
		return
	}

	renderPage(w, "project.html", mesProjects)
}