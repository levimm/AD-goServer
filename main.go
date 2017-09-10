package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func blogHandler(w http.ResponseWriter, r *http.Request) {
	subdomain := strings.Split(r.Host, ".")
	fmt.Fprintf(w, "Welcome to subdomain %s", subdomain[0])
}

func ppdHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's ppd subdomain")
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello AntDiaries from my VPS!") // send data to client side
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	blog := r.Host("{subdomain:[a-z]+}.antdiaries.com").Subrouter()
	blog.HandleFunc("/", blogHandler)

	r.HandleFunc("/", mainHandler)

	//ppd := r.Host("ppd.antdiaries.com").Subrouter()
	//ppd.HandleFunc("/", ppdHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":80", r))
}
