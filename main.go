package main

import (
	"io"
	"net/http"
)

func myFunc(w http.ResponseWriter, _ *http.Request){
	io.WriteString(w, "deneme")
}
func myFunc2(w http.ResponseWriter, _ *http.Request){
	io.WriteString(w, "deneme2")
}
func myFunc3(w http.ResponseWriter, _ *http.Request){
	io.WriteString(w, "deneme3")
}

func main(){
	http.Handle("/", http.HandlerFunc(myFunc2))
	http.Handle("/", http.HandlerFunc(myFunc3))
	http.ListenAndServe(":8080", http.HandlerFunc(myFunc))
}
