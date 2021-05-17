package main

import (
	"io"
	"net/http"
)

func myFunc(w http.ResponseWriter, _ *http.Request){
	io.WriteString(w, "deneme")
}

func main(){
	http.ListenAndServe(":8080",http.HandlerFunc(myFunc))
}
