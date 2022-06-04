package main

import (
	"html/template"
	"net/http"
)

func myFunc2(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("main.html")
	t.Execute(w, "fff")
}

func myFunc3(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("map.html")
	t.Execute(w, "fff")
}

func myFunc4(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("option.html")
	t.Execute(w, "fff")
}

func template_page(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("option.html")
	t.Execute(w, "fff")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("erol")
	}


	http.Handle("/", http.HandlerFunc(myFunc4))
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.Handle("/3", http.HandlerFunc(myFunc3))
	http.Handle("/template", http.HandlerFunc(template_page))
	http.ListenAndServe(":8080", nil)
}
