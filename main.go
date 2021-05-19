package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
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

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("error")
	}
	http.Handle("/", http.HandlerFunc(myFunc2))
	http.Handle("/3", http.HandlerFunc(myFunc3))
	http.Handle("/option", http.HandlerFunc(myFunc4))
	http.Handle("/heroku/assets/", http.StripPrefix("/heroku/assets/", http.FileServer(http.Dir("./heroku/assets"))))
	http.ListenAndServe(":"+port, nil)
}
