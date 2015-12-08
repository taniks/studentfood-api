package main

import (
	"log"
	"net/http"
	"github.com/go-martini/martini"
)

// The one and only martini instance.
var m *martini.Martini

func init() {
	m = martini.New()

	// Setup middleware
	m.Use(martini.Recovery())
	m.Use(martini.Logger())

	// Setup routes
	r := martini.NewRouter()
	r.Get("/restaurants", getRestaurants)
  	r.Get("/food", getFood)
	r.Get("/food/xml", getFoodXML)
  	r.Post("/review", postReview)

	// Add the router action
	m.Action(r.Handle)
}

func main() {
  	// Listen on http: with the preconfigured martini instance
  	if err := http.ListenAndServe(":8000", m); err != nil {
		log.Fatal(err)
	}
}
