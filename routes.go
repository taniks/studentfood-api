// Route functions are defined here

package main

import (
	"encoding/xml"
    "encoding/json"
	"net/http"
	"github.com/go-martini/martini"
	"time"
	"strconv"
	"strings"
)

func getRestaurants(r *http.Request) string {

	var restaurants []Restaurant
	restaurants = append(restaurants, Restaurant {
		XMLName: xml.Name{Local: "Restaurant"},
		Name: "Jamk Dynamo",
		API: "menu/Sodexo/5865",
	})
	restaurants = append(restaurants, Restaurant {
		XMLName: xml.Name{Local: "Restaurant"},
		Name: "Jamk Pääkampus",
		API: "menu/Sodexo/5859",
	})
	restaurants = append(restaurants, Restaurant {
		XMLName: xml.Name{Local: "Restaurant"},
		Name: "Jamk Rajacafé",
		API: "menu/Sodexo/5861",
	})
	restaurants = append(restaurants, Restaurant {
		XMLName: xml.Name{Local: "Restaurant"},
		Name: "Jamk Musiikkikampus",
		API: "menu/Sodexo/5868",
	})
	
	qs := r.URL.Query().Get("type")
	if qs == "xml" {
		restaurantXML := Restaurants{
            XMLName: xml.Name{ Local: "restaurants" },
            Restaurants: restaurants,
        } 
		
		// Decode it restaurants to xml
		decoded, err := xml.Marshal(restaurantXML)
		if isErr(err) { return handleError(err.Error(), true) }
		
		// Return json
		return string(decoded)
	}
	
	// Decode it restaurants to json
	decoded, err := json.Marshal(restaurants)
	if isErr(err) { return handleError(err.Error(), false) }
	
	// Return json
	return string(decoded)
}


func getMenu(r *http.Request, params martini.Params) string {
	
	qs := r.URL.Query().Get("type")
	isXML := qs == "xml"
	
	restaurant := strings.ToLower(params["restaurant"])
	if restaurant == "" { return handleError("Restaurant not defined", isXML) }
	
	id := params["id"]
	if id == "" { return handleError("ID not defined", isXML) }
	
	// Get time now
	now := time.Now()
	
	// Get date
	day := params["day"]
	if day == "" { day = strconv.Itoa(now.Day()) }
	
	// Get month
	month := params["month"]
	if month == "" { month = strconv.Itoa(int(now.Month())) }
	
	// Get month
	year := params["year"]
	if year == "" { year = strconv.Itoa(now.Year()) }
	
	switch restaurant {
		case "sodexo": {
			if isXML {
				sodexo, err := getSodexo(id, day, month, year, true)
				if isErr(err) { return handleError(err.Error(), isXML) }
				return sodexo
			}
			sodexo, err := getSodexo(id, day, month, year, false)
			if isErr(err) { return handleError(err.Error(), isXML) }
			return sodexo
		}
		default: {
			return "Invalid restaurant" + restaurant
		}
	}
}
