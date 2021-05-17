package main

import (
	"io"
	"os"
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
	port := os.Getenv("PORT")
    	if port == "" {
        	log.Fatal("error")
    	}
	http.Handle("/", http.HandlerFunc(myFunc2))
	http.Handle("/3", http.HandlerFunc(myFunc3))
	http.ListenAndServe(":"+port, nil)
}
