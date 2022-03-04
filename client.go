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
    values := u.Query()
    values.Add("name", "Shreyansh_Singh")
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
