package main

import (
	"log"
	"net/http"
	"github.com/go-martini/martini"
	"os"
)

// The one and only martini instance.
var m *martini.Martini

func init() {
	m = martini.New()

	// Setup middleware
	m.Use(martini.Recovery())
	m.Use(MapEncoder)

	// Setup routes
	r := martini.NewRouter()
	r.Get("/restaurants", getRestaurants)
  	r.Get("/menu/:restaurant/:id", getMenu)
	r.Get("/menu/:restaurant/:id/:day", getMenu)
	r.Get("/menu/:restaurant/:id/:day/:month", getMenu)
	r.Get("/menu/:restaurant/:id/:day/:month/:year", getMenu)

	// Add the router action
	m.Action(r.Handle)
}

// MapEncoder sets content type to xml or json depending on requested type
func MapEncoder(c martini.Context, w http.ResponseWriter, r *http.Request) {
	qsXML := r.URL.Query().Get("type")
	
	if qsXML == "xml" {
		w.Header().Set("Content-Type", "application/xml")
	} else {
		w.Header().Set("Content-Type", "application/json")
	}
}


func main() {
	port := os.Getenv("PORT");
	
	// If port is not defined
	if port == "" {
        port = "8000"
	}
	
  	// Listen on http: with the preconfigured martini instance
  	if err := http.ListenAndServe(":"+port, m); err != nil {
		log.Fatal(err)
	}
}
