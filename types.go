// Defines types for returning data and fetching data

package main

type (
	Restaurant struct {
		XMLName xml.Name `json:"-" xml:"restaurant"`
		Name     string `json:"name" xml:"name"`
		Location string `json:"location" xml:"location"`
		Website  string `json:"website" xml:"website"`
		Image    string `json:"image" xml:"image"`
	}
	Food struct {
		XMLName xml.Name `json:"-" xml:"food"`
		Name     string `json:"name" xml:"name"`
		Allergies string `json:"allergies" xml:"allergies"`
		Price  string `json:"price" xml:"price"`
	}
	Language struct {
		XMLName xml.Name `json:"-" xml:"language"`
		English string `json:"english" xml:"english"`
		Finnish string `json:"finnish" xml:"finnish"`
	}
)
