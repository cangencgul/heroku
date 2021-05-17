package main

import(
  "fmt"
  "net/http"
  "golang.org/x/net/html"		
)

func main() {
  
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
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
