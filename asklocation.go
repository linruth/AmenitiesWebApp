 package main

 import (
         "fmt"
         "html/template"
         "net"
         "net/http"
         "encoding/json"
    "io/ioutil"
    "log"
    "html"
 )

var s MyJsonName

type MyJsonName struct {
    Markers []struct {
        Fee        string  `json:"fee"`
        Latitude   float64 `json:"latitude"`
        Longitude  float64 `json:"longitude"`
        Wheelchair string  `json:"wheelchair"`
    } `json:"markers"`
}

type Page struct {
    Title string
    Body  []byte
}

func main() {
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
    //http.HandleFunc("/test/", viewHandler)
    //http.ListenAndServe(":8080", nil)
/*
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(":8081", nil))
    */
}

/*
 func home(w http.ResponseWriter, r *http.Request) {
         var templates = template.Must(template.New("locateip").ParseFiles("asklocation.html"))
         templates.ExecuteTemplate(w, "asklocation.html", s)
       
 }
*/
 /*

 func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/test/"):]
    p, _ := loadPage("test")
    fmt.Fprintf(w, "<h1>hello</h1>")
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
*/
