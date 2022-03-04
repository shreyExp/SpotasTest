package main

import (
    "database/sql"
    //"encoding/json"
    "fmt"
    "log"
    "net/http"
    //"net/url"
    //"github.com/gorilla/mux"
    _ "github.com/lib/pq"
)

const (
    DB_USER = "postgres"
    //DB_PASSWORD = ""
    DB_NAME = "test_db"
)
func setupDB() *sql.DB {
    dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable", DB_USER, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    if err != nil {
        panic(err)
    }
    return db
}
//POINT(-1.922959 52.468337)
func makeQueryString (long float64, lat float64) string {

    part1 := "select * from "
    part2 := `"MY_TABLE" `
    part3 := "where ST_DWithin(coordinates, "
    part4 := fmt.Sprintf("ST_GeomFromText('Point(%f %f)', 4326), 1000.0);", long, lat)
    query := part1 + part2 + part3 + part4
    return query
}

type query_row struct {
    id string
    name string
    website string
    coordinates string
    description string
    rating float32
}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    var resp string = ""
    //names, ok := r.URL.Query()["name"]
    //if !ok || len(names[0]) < 1 {
    //    resp = "Get request not properly made"
    //}else{
    //    resp = names[0]
    //}
    //mapv := map[string]int{"apple": 2, "peer": 3}
    //mapJ, _ := json.Marshal(mapv)
    //resp = string(mapJ)

    db := setupDB()
    //query := `select upper(website) from "MY_TABLE" where name = 'Simpsons Restaurant';`
    long := -1.922959
    lat := 52.468337
    query := makeQueryString(long, lat)
    fmt.Println(query)
    //var avRat string
    //row := db.QueryRow(query)
    rows, err := db.Query(query)
    defer rows.Close()
    if err != nil {
    //panic(err)
        resp = "Oops"
    }else {
        for rows.Next() {
            var ro query_row
            rows.Scan(&ro.id, &ro.name, &ro.website, &ro.coordinates, &ro.description, &ro.rating)
            fmt.Println(ro.name)
        }
        resp = "Whatever"
        //resp = query
    }
    w.Write([]byte(resp))
}
func main () {
    mux := http.NewServeMux()
    mux.HandleFunc("/", HelloHandler)
    addr := ":5000"
    log.Printf("server is listening at %s...", addr)
    log.Fatal(http.ListenAndServe(addr, mux))
}
