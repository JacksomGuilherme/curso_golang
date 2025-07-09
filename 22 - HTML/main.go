package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	NomeUsuario string
	email       string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		u := usuario{
			"Jacksom",
			"jacksom@golang.com",
		}

		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Listening at: 127.0.0.1:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
