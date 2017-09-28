package main

import "fmt"
import "net/http"
import "log"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Header().Add("Content-Type", "text/plain")
	//w.Write([]byte("Hello, World!"))
	fmt.Fprintf(w, "Hello %s!", name)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello/", helloHandler)
	fmt.Printf("server is listening at http://localhost:4000\n")
	log.Fatal(http.ListenAndServe("localhost:4000", mux))
}
