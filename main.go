package main

import(
  "fmt"
  "net/http"
)

func main() {
  
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    fmt.Println("fasafiso")
  })
  
  s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		fmt.Print("There is a Error from http.ListenAndServe it is ", err)
	}
}
