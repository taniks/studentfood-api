// Functions for importing data from database (WIP)

package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func getRestaurants() ([]*Restaurant, error) {
	con, err := sql.Open("mymysql", database+"/"+user+"/"+password)
	if isErr(err) { return nil, err }

	// Close connection when value is returned
	defer con.Close()

	rows, err := con.Query("SELECT name, location, website, image FROM restaurant WHERE p1=? and p2=?", p1, p2)
	if isErr(err) { return nil, err }

	// Define variables here
	restaurants := make([]*Restaurant)
	var name, location, website, image string

	for rows.Next() {
		// Sace row values into variables
	    err = rows.Scan(&name, &location, &website, &image)
	    if isErr(err) { return nil, err }

	    // Save restaurant to restaurants list
	    restaurants = append(restaurants, &Restaurant{name, location, website, image})
	}

	return restaurants, nil
}

func foodFromDB() ([]*Restaurant, error) {
	con, err := sql.Open("mymysql", database+"/"+user+"/"+password)
	if isErr(err) { return nil, err }

	// Close connection when value is returned
	defer con.Close()

	rows, err := con.Query("SELECT name, location, website, image FROM restaurant WHERE p1=? and p2=?", p1, p2)
	if isErr(err) { return nil, err }

	// Define variables here
	restaurants := make([]*Restaurant)
	var name, location, website, image string

	for rows.Next() {
		// Sace row values into variables
	    err = rows.Scan(&name, &location, &website, &image)
	    if isErr(err) { return nil, err }

	    // Save restaurant to restaurants list
	    restaurants = append(restaurants, &Restaurant{name, location, website, image})
	}

	return restaurants, nil
}
