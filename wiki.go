package main

import (
    "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
   "html/template"
    "os"
)

type MyJsonName struct {
	Markers []struct {
		Fee        string  
		Latitude   float64 
		Longitude  float64
		Wheelchair string  
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
    t := template.New("rootForm")
    t, _ = t.Parse(rootForm)
    t.Execute(os.Stdout, s)
    http.HandleFunc("/view", handler)
    //http.HandleFunc("/showimage",showimage)
    http.ListenAndServe(":8080", nil)
    //http.HandleFunc("/showimage", showimage)
}


func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, rootForm)
}

const rootForm = `
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
    <p> {{with .Markers}} {{range .}} a Marker {{.Latitude}} {{.Longitude}}{{.Wheelchair}}{{end}}{{end}} </p>
    <div id="map"></div>
    <div id = "box"></div>
    <script>
    var positionLat;
var positionLong;
var map;
window.onload = function() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(initMap);
        setTimeout(showLatLong,800);
        setTimeout(createForm,800);
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
	 var myLatLng = {lat: positionLat, lng: positionLong};
  map = new google.maps.Map(document.getElementById('map'), {center: {lat: parseFloat(positionLat), lng: parseFloat(positionLong)},zoom: 15});
  
  
  var marker = new google.maps.Marker({
    position: myLatLng,
    map: map,
    title: 'Hello World!'
  });
/*
var allLocations;
  {{with .Markers}} {{range .}}
          var lat = {{.Latitude}}
          var long = {{.Longitude}}
          allLocations.push(lat long);
  {{end}}{{end}} 


    for (var i = 0; i<allLocations.length; i++ ) {
          var finalLat = i.lat;
          var finalLong = i.long;
          var myLatlng = new google.maps.LatLng(finalLat, finalLong);

          var marker = new google.maps.Marker({
              position: myLatlng,
              title:'toilets',
              map:map,
          });
      marker.setMap(map);
    }
*/
function showLatLong(){
document.getElementById("box").innerHTML = "Latitude and Longitude";
};

function createForm(){
var f = document.createElement("form");
f.setAttribute('method',"post");
f.setAttribute('action',"/showimage");

var i = document.createElement("input"); //input element, text
i.setAttribute('type',"text");
i.setAttribute('name',"positionLat");
i.setAttribute('value',positionLat);

var j = document.createElement("input"); //input element, text
j.setAttribute('type',"text");
j.setAttribute('name',"positionLong");
j.setAttribute('value',positionLong);

var s = document.createElement("input"); //input element, Submit button
s.setAttribute('type',"submit");
s.setAttribute('value',"Submit");

f.appendChild(i);
f.appendChild(j);
f.appendChild(s);

//and some more input elements here
//and dont forget to add a submit button

document.getElementsByTagName('body')[0].appendChild(f);
};
    </script>
    <script src="https://maps.googleapis.com/maps/api/js?callback=initMap"
        async defer></script>
  </body>
</html>
`


//var upperTemplate = template.Must(template.New("showimage").Parse(upperTemplateHTML))

//func showimage(w http.ResponseWriter, r *http.Request) {
        // Sample address "1600 Amphitheatre Parkway, Mountain View, CA"
        //latitude := r.FormValue("positionLat")
        //longitude := r.FormValue("positionLong");
    //}

	// QueryEscape escapes the addr string so
	// it can be safely placed inside a URL query
	// safeAddr := url.QueryEscape(addr)
        //safeAddr := url.QueryEscape(addr)
        //fullUrl := fmt.Sprintf("http://maps.googleapis.com/maps/api/geocode/json?sensor=false&address=%s", safeAddr)

        //c := appengine.NewContext(r)
        //client := urlfetch.Client(c)
    
       // resp, err := client.Get(fullUrl)
        //if err != nil {
                //http.Error(w, err.Error(), http.StatusInternalServerError)
                //return
        //}
    
	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	//defer resp.Body.Close()

	// Read the content into a byte array
	//body, dataReadErr := ioutil.ReadAll(resp.Body)
	//if dataReadErr != nil {
		//panic(dataReadErr)
	//}/

        //res := make(map[string][]map[string]map[string]map[string]interface{}, 0)

	// We will be using the Unmarshal function
	// to transform our JSON bytes into the
	// appropriate structure.
	// The Unmarshal function accepts a byte array
	// and a reference to the object which shall be
	// filled with the JSON data (this is simplifying,
	// it actually accepts an interface)
	//json.Unmarshal(body, &res)
        
	//lat, _ := res["results"][0]["geometry"]["location"]["lat"]
	//lng, _ := res["results"][0]["geometry"]["location"]["lng"]
	
	// %.13f is used to convert float64 to a string
	//queryUrl := fmt.Sprintf("http://maps.googleapis.com/maps/api/streetview?sensor=false&size=600x300&location=%.13f,%.13f", lat, lng)

        //tempErr := upperTemplate.Execute(w, queryUrl)
        //if tempErr != nil {
	       // http.Error(w, tempErr.Error(), http.StatusInternalServerError)
        //}
//}

// const upperTemplateHTML = ` 
// <!DOCTYPE html>
//   <html>
//     <head>
//       <meta charset="utf-8">
//       <title>Display Image</title>
//       <link rel="stylesheet" href="/stylesheets/goview.css">              
//     </head>
//     <body>
//       <h1><img style="margin-left: 120px;" src="images/gsv.png" alt="Street View" />GoView</h1>
//       <h2>Image at your Address</h2>
//       <img style="margin-left: 120px;" src="{{html .}}" alt="Image" />
//     </body>
//   </html>
// `
		