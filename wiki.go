package main

import (
    "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MyJsonName struct {
	Markers []struct {
		Fee        string  `json:"fee"`
		Latitude   float64 `json:"latitude"`
		Longitude  float64 `json:"longitude"`
		Wheelchair string  `json:"wheelchair"`
	} `json:"markers"`
}

func main() {
	 var s MyJsonName
  resp, err := http.Get("http://amenimaps.com/amenimapi.php?amenity=toilet&mylat=51.50784&mylon=-0.127324&mode=json&name=rara_pirates&key=773afa6b5638d5ce24e12e9acbe30bb2")
	//Fill the record with the data from the JSON
	if err != nil {
		// An error occurred while converting our JSON to an object
		fmt.Printf("%s", err)
	}
	defer resp.Body.Close()
	contents, err:= ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	
    err = json.Unmarshal(contents, &s)
    if err != nil {
        fmt.Printf("%s", err)
    }
    fmt.Println(s)
}
		