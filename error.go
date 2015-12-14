// Error handling system

// Custom error type to return correct JSON
package main

import (
	"encoding/xml"
    "encoding/json"
)

type (
	// Error struct to return message in xml or json
	Error struct {
		XMLName xml.Name `json:"-" xml:"error"`
		Message     string `json:"message" xml:"message"`
	}
)

func handleError(message string, isXML bool) string {

	if isXML {	
		// Decode it error to xml
		decoded, err := xml.Marshal(Error{XMLName: xml.Name{ Local: "error" }, Message: message})
		if isErr(err) { return handleError(err.Error(), isXML) }
		
		// Return json
		return string(decoded)
	}
	
	// Decode it restaurants to json
	decoded, err := json.Marshal(Error{XMLName: xml.Name{ Local: "error" }, Message: message})
	if isErr(err) { return handleError(err.Error(), isXML) }
	
	// Return json
	return string(decoded)
}

// Simple is error check
func isErr(err error) bool {
    if err != nil {
        return true
    }

    return false
}
