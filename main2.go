package main

import (
	// "fmt"
	"net/http"
	"html/template"
)

var tmpl *template.Template

func init(){
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func homePage2(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello World %s", r.URL.Path)
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func aboutPage2(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello World %s", r.URL.Path)
	tmpl.ExecuteTemplate(w, "about.html", nil)
}
func contactPage2(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello World %s", r.URL.Path)
	tmpl.ExecuteTemplate(w, "contact.html", nil)
}

func main2() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets",fs))
	http.HandleFunc("/", homePage2)
	http.HandleFunc("/contact", contactPage2)
	http.HandleFunc("/about", aboutPage2)
	http.ListenAndServe(":8888", nil)
}

