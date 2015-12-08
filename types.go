// Defines types for returning data and fetching data

package main

import (
	"encoding/xml"
)

// database
type (
	Restaurant struct {
		XMLName xml.Name `json:"-" xml:"restaurant"`
		Name     string `json:"name" xml:"name"`
		Location string `json:"location" xml:"location"`
		Website  string `json:"website" xml:"website"`
		Image    string `json:"image" xml:"image"`
		API		 string `json:"api" xml:"api"`
		Custom   string `json:"custom" xml:"custom"`
	}
	Foods struct {
		XMLName xml.Name `json:"-" xml:"foods"`
		Foods     []Food `json:"foods" xml:"foods"`
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

// External restaurant api's
type (
	Sodexo struct {
		XMLName xml.Name `json:"-" xml:"sodexo"`
		Meta     SodexoMeta `json:"meta" xml:"meta"`
		Courses  []SodexoCourse `json:"courses" xml:"courses"`
	}
	SodexoMeta struct {
		XMLName xml.Name `json:"-" xml:"meta"`
		GeneratedTimestamp int `json:"generated_timestamp" xml:"generated_timestamp"`
		RequestedTimestamp int `json:"requested_timestamp" xml:"requested_timestamp"`
		RefURL string `json:"ref_url" xml:"ref_url"`
		RefTitle string `json:"ref_title" xml:"ref_title"`
	}
	SodexoCourse struct {
		XMLName xml.Name `json:"-" xml:"food"`
		TitleFi    string `json:"title_fi" xml:"title_fi"`
		TitleEn    string `json:"title_en" xml:"title_en"`
		Category   string `json:"category" xml:"category"`
		Price      string `json:"price" xml:"price"`
		Properties string `json:"properties" xml:"properties"`
		DescFi     string `json:"desc_fi" xml:"desc_fi"`
		DescEn     string `json:"desc_en" xml:"desc_en"`
		DescSe     string `json:"desc_se" xml:"desc_se"`
	}
)