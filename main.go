package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Creation of a simple web server listening on port 1337.
	srv := &http.Server{
		Addr:    ":1337",
		Handler: handleFunc(),
	}

	fmt.Printf("go to http://localhost%s\n", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}

// handleFunc creates the router.
// Not very useful here as we only have one path.
func handleFunc() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)

	return r
}

// home page handler. It parses and executes home.html.
func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("home.html")
	if err != nil {
		log.Fatalf("Can not parse home page : %v", err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("Can not execute templates for home page : %v", err)
	}
}
