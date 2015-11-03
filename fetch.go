// Functions for importing data from external sources (WIP)

package main

import (
    "encoding/json"
    "net/http"
)

func getJson(url string, target interface{}) error {
    r, err := http.Get(url)
    if isError(err) { return err }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}
