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


func myFunc5(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("/deneme/deneme.html")
	t.Execute(w, "fff")
}

func template_page(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("hello template <3"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("erol")
	}

	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.Handle("/deneme/", http.StripPrefix("/deneme", http.FileServer(http.Dir("./deneme"))))
	
	http.Handle("/", http.HandlerFunc(myFunc4))
	http.Handle("/2", http.HandlerFunc(myFunc5))
	http.Handle("/3", http.HandlerFunc(myFunc3))
	http.Handle("/template", http.HandlerFunc(template_page))
	http.ListenAndServe(":"+port, nil)
	//http.ListenAndServe(":8080", nil)
}
