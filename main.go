package main

import (
	"html/template"
	"io"
	"os"
	"log"
	"net/http"
)

func myFunc2(w http.ResponseWriter, _ *http.Request){
	t, _ := template.ParseFiles("main.html")
	t.Execute(w, "fff")
}

func myFunc3(w http.ResponseWriter, _ *http.Request){
	t, _ := template.ParseFiles("map.html")
	t.Execute(w, "fff")
}

func main(){
	port := os.Getenv("PORT")
    	if port == "" {
        	log.Fatal("error")
    	}
	http.Handle("/", http.HandlerFunc(myFunc2))
	http.Handle("/3", http.HandlerFunc(myFunc3))
	http.ListenAndServe(":"+port, nil)
}
