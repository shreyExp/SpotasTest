package main

import (
    //"database/sql"
    //"encoding/json"
    "fmt"
    //"log"
    "io/ioutil"
    "net/http"
    "net/url"
    //"github.com/gorilla/mux"
    //_ "github.com/lib/pq"
)

func main() {
    s := "http://localhost:5000"
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }
    long := -1.922959
    lat := 52.468337
    radius := 1000.0
    //shape := "circle"
    shape := "square"
    values := u.Query()
    values.Add("Longitude", fmt.Sprintf("%f",long))
    values.Add("Latitude", fmt.Sprintf("%f",lat))
    values.Add("Radius", fmt.Sprintf("%f",radius))
    values.Add("Type", shape)
    u.RawQuery = values.Encode()
    //fmt.Println(u.String())
    resp, err := http.Get(u.String())
    if err != nil {
        panic(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    response_from_server := string(body)
    fmt.Println(response_from_server)
}
