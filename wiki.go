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
    
    http.HandleFunc("/view", handler)
    http.ListenAndServe(":8080", nil)
    //http.HandleFunc("/showimage", showimage)
}


func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, rootForm)
}

const rootForm = `
  <!DOCTYPE html>
    <!DOCTYPE html>
<html>
  <head>
    <style>
      #map {
        width: 500px;
        height: 400px;
      }
    </style>
  </head>
  <body>
    <div id="map"></div>
    <script>

    var positionLat;
var positionLong;
var map;

window.onload = function() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(initMap);

    } else {
        console.log("not using geolocation");
    };
};
/*
//gets your lat and long
function showPosition(position) {
	positionLat = position.coords.latitude;
	positionLong = position.coords.longitude;
   // x.innerHTML = "Latitude: " + position.coords.latitude + 
    //"<br>Longitude: " + position.coords.longitude; 
	console.log(positionLat);
	//console.log(positionLong);
};
*/

     function initMap(position) {
	positionLat = position.coords.latitude;
	positionLong = position.coords.longitude;
  map = new google.maps.Map(document.getElementById('map'), {center: {lat: parseFloat(positionLat), lng: parseFloat(positionLong)},zoom: 15});
}
    </script>
    <script src="https://maps.googleapis.com/maps/api/js?callback=initMap"
        async defer></script>
  </body>
</html>
`
		