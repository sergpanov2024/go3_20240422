package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	templateNames := []string{"aboutme", "skills"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("./tmpl/_layout.html", "./tmpl/"+name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

type taboutme struct {
	Name  string
	Email string
	Phone string
	Addr  string
}

func aboutme(w http.ResponseWriter, r *http.Request) {
	templates["aboutme"].Execute(w, &taboutme{
		Name:  "Панов Сергей Михайлович",
		Email: "mail@ya.ru",
		Phone: "+7(123)456-78-90",
		Addr:  "Архангельск",
	})
}

// listHandler handles "/list" URL
func skills(w http.ResponseWriter, r *http.Request) {
	skills := []string{}
	skills = append(skills, "SQL")
	skills = append(skills, "go")
	templates["skills"].Execute(w, skills)
}

func main() {
	loadTemplates()

	http.HandleFunc("/", aboutme)
	http.HandleFunc("/skills", skills)

	err := http.ListenAndServe(":1313", nil)
	if err != nil {
		log.Fatal(err)
	}
}
