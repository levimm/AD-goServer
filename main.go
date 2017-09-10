package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Subdomains map[string]http.Handler

func (subdomains Subdomains) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	domainParts := strings.Split(r.Host, ".")
	
	if mux := subdomains[domainParts[0]]; mux != nil {
		// Let the appropriate mux serve the request
		mux.ServeHTTP(w, r)
	} else {
		// Handle 404
		http.Error(w, "Not found", 404)
	}
}

type Mux struct {
	http.Handler
}

func (mux Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to blog subdomain, Hello, %q", r.URL.Path[1:])
}

// func adminHandlerTwo(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "It's adminHandlerTwo , Hello, %q", r.URL.Path[1:])
// }

func analyticsHandlerOne(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's analyticsHandlerOne , Hello, %q", r.URL.Path[1:])
}

func analyticsHandlerTwo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's analyticsHandlerTwo , Hello, %q", r.URL.Path[1:])
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello AntDiaries from my VPS!") // send data to client side
}

func main() {
	blogMux := http.NewServeMux()
	blogMux.HandleFunc("/blog", blogHandler)
	// blogMux.HandleFunc("/admin/pathtwo", adminHandlerTwo)

	analyticsMux := http.NewServeMux()
	analyticsMux.HandleFunc("/analytics/pathone", analyticsHandlerOne)
	analyticsMux.HandleFunc("/analytics/pathtwo", analyticsHandlerTwo)

	mainMux := http.NewServeMux()
	mainMux.HandleFunc("/", sayhelloName)

	subdomains := make(Subdomains)
	subdomains["blog"] = blogMux
	subdomains["paipaidai"] = analyticsMux

	http.ListenAndServe(":8080", subdomains)
}
