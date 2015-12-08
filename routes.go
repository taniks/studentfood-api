// Route functions are defined here

package main

import (
	//"encoding/json"
	"net/http"
)

func getRestaurants(r *http.Request) string {
	/*
	// Get restaurants from database
	restaurants, err := getRestaurantsFromDb()
	if isErr(err) { return json.Marshal(err) }
	
	// Turn data to json
	jsonized, err := json.Marshal(restaurants)
	if isErr(err) { return json.Marshal(err) }
	
	// Return json
	*/
	return "";
}

func getFood(r *http.Request) string {
	asd, err := getSodexo(false)
	if isErr(err) { return err.Error() }
	return asd
}

func getFoodXML(r *http.Request) string {
	asd, err := getSodexo(true)
	if isErr(err) { return err.Error() }
	return asd
}

func postReview(r *http.Request) string {
	// Should contain functionality to post review of food
	return "";
}
