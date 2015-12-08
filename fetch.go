package main

import (
    "encoding/xml"
    "encoding/json"
    "net/http"
    "time"
    "strings"
    "io/ioutil"
    "strconv"
)

func getJSON(url string) ([]byte, error) {
    // Try to get data from url
    r, err := http.Get(url)
    if isErr(err) { return nil, err }
    
    // Close response after return
    defer r.Body.Close()
    
    // Read body
    body, err := ioutil.ReadAll(r.Body)
    if isErr(err) { return nil, err }
    
    return body, nil
}

// Get food from sodexo restaurant
//func getSodexo(restaurant Restaurant, useXML bool) (string, error) {
func getSodexo(useXML bool) (string, error) {
    var sodexo Sodexo
    
    // Change to use URL from database!!!!!!!
    url := "http://www.sodexo.fi/ruokalistat/output/daily_json/5859/{year}/{month}/{day}/fi"
    
    // Turn api to correct format
    now := time.Now()
	formatted := strings.Replace(url, "{day}", strconv.Itoa(now.Day()), 1)
	formatted = strings.Replace(formatted, "{month}", strconv.Itoa(int(now.Month())), 1)
	formatted = strings.Replace(formatted, "{year}", strconv.Itoa(now.Year()), 1)
    
    // Decode json to struct
    js, err := getJSON(formatted)
    if isErr(err) { return "", err }
    
    // Decode json
    erro := json.Unmarshal(js, &sodexo)
    if isErr(erro) { return "", erro } 
    
    // Turn sodexo json to json we want
    var foods []Food
    for _, course := range sodexo.Courses {
        foods = append(foods, Food {
            XMLName: xml.Name{ Local: "food" },
            Name: course.TitleFi,
            Allergies: course.Properties,
            Price: course.Price,
        })
    }
    
    // return xml if defined
    if(useXML) {
        foodXML := Foods{
            XMLName: xml.Name{ Local: "foods" },
            Foods: foods,
        } 
        decoded, err := xml.Marshal(foodXML)
        if isErr(err) { return "", err } 
        
        return string(decoded), nil
    }
    
    // Return json
    decoded, err := json.Marshal(foods)
    if isErr(err) { return "", err } 
    
    return string(decoded), nil
}