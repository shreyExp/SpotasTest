package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "sort"
    "math"
    "strings"
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
func makeQueryStringForCircle (long float64, lat float64, dist float64) string {
    distance := fmt.Sprintf("ST_Distance(coordinates, ST_GeomFromText('POINT(%f %f)', 4326))", long, lat)
    azimuth := fmt.Sprintf("ST_Azimuth(ST_GeomFromText('POINT(%f %f)',4326), coordinates)", long, lat)
    part1 := fmt.Sprintf("select *, %s , %s from ", distance, azimuth)
    part2 := `"MY_TABLE" `
    part3 := "where ST_DWithin(coordinates, "
    part4 := fmt.Sprintf("ST_GeomFromText('Point(%f %f)', 4326), %f);", long, lat, dist)
    query := part1 + part2 + part3 + part4
    return query
}

func makeQueryStringForSquare (long float64, lat float64, dist float64) string {
    distance := fmt.Sprintf("ST_Distance(coordinates, ST_GeomFromText('POINT(%f %f)', 4326))", long, lat)
    azimuth := fmt.Sprintf("ST_Azimuth(ST_GeomFromText('POINT(%f %f)',4326), coordinates)", long, lat)
    part1 := fmt.Sprintf("select *, %s , %s from ", distance, azimuth)
    part2 := `"MY_TABLE" `
    var polygon_details string
    var square_corners []string
    bearings := []float64{45, 135, 225, 315, 45}

    for _, bearing := range bearings {
        corner_detail := fmt.Sprintf("ST_Project(ST_GeomFromText('POINT(%f %f)',4326) , %f, radians(%f)) :: geometry", long, lat, math.Sqrt(2)*dist, bearing)
        square_corners = append(square_corners, corner_detail)
    }
    polygon_details = strings.Join(square_corners, ",")
    part3 := fmt.Sprintf("where ST_Within( coordinates :: geometry, ST_MakePolygon(ST_MakeLine(Array[%s", polygon_details)
    part4 := "])));"
    query := part1 + part2 + part3 + part4
    //fmt.Println(query)
    return query
}

type query_row struct {
    id string
    name string
    website string
    coordinates string
    description string
    rating float64
    distance float64
    azimuth float64
    fifty_proximity []string
}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    var resp string
    var lat float64
    var long float64
    var distance float64
    var shape string
    lats, ok := r.URL.Query()["Latitude"]
    if !ok || len(lats[0]) < 1 {
        log.Printf("Latitude is not provided correctly")
    }else{
        lat, _ = strconv.ParseFloat(lats[0], 64)
    }
    longs, ok := r.URL.Query()["Longitude"]
    if !ok || len(longs[0]) < 1 {
        log.Printf("Longitude is not provided correctly")
    }else{
        long, _ = strconv.ParseFloat(longs[0], 64)
    }

    distances, ok := r.URL.Query()["Radius"]
    if !ok || len(distances[0]) < 1 {
        log.Printf("Radius is not provided correctly")
    }else{
        distance, _ = strconv.ParseFloat(distances[0], 64)
    }

    shapes, ok := r.URL.Query()["Type"]
    if !ok || len(shapes[0]) < 1 {
        log.Printf("Shape is not provided correctly")
    }else{
        shape = shapes[0]
    }


    db := setupDB()
    //long := -1.922959
    //lat := 52.468337
    var query string
    if shape == "circle" {
        query = makeQueryStringForCircle(long, lat, distance)
    } else if shape == "square" {
        //fmt.Println(shape)
        query = makeQueryStringForSquare(long, lat, distance)
    }
    //fmt.Println(query)
    rows, err := db.Query(query)
    defer rows.Close()
    if err != nil {
        resp = "Oops"
    }else {
        var row_list []query_row
        var row_list_of_map []map[string]string
        for rows.Next() {
            var ro query_row
            rows.Scan(&ro.id, &ro.name, &ro.website, &ro.coordinates, &ro.description, &ro.rating, &ro.distance, &ro.azimuth)
            row_map := map[string]string{"id": ro.id, "name": ro.name, "website": ro.website, "coordinates": ro.coordinates, "rating": fmt.Sprintf("%f", ro.rating)}
            row_list_of_map = append(row_list_of_map, row_map)
            row_list = append(row_list, ro)
            //fmt.Println(ro.name, ro.distance)
        }
        for i, l_row := range row_list {
            for _, ll_row := range row_list {
                if (is_proximity_fifty_meters(l_row.azimuth, l_row.distance, ll_row.azimuth, ll_row.distance)) {
                    row_list[i].fifty_proximity = append(row_list[i].fifty_proximity, ll_row.id)
                }
            }
        }
        sort.Slice(row_list, func(i, j int) bool {
            var return_value bool
            //if math.Abs(row_list[i].distance - row_list[j].distance) < 50 {
            if contains(row_list[i].fifty_proximity, row_list[j].id) {
                return_value = row_list[i].rating < row_list[j].rating
            } else {
                return_value = row_list[i].distance < row_list[j].distance
            }
            return return_value
        })
        respByte, _ := json.Marshal(row_list_of_map)
        resp = string(respByte)
    }
    w.Write([]byte(resp))
}
func contains (list []string, value string) bool {
  for _, x := range list {
      if value == x {
          return true
      }
  }
  return false
}
func is_proximity_fifty_meters(azimuth2 float64, chain2 float64, azimuth1 float64, chain1 float64) bool {
    diff_az := azimuth2 - azimuth1
    distance_spots := math.Sqrt(math.Pow(chain1, 2) + math.Pow(chain2, 2) - 2 * math.Abs(chain1) * math.Abs(chain2) * math.Cos(diff_az))
    return_value := false
    if distance_spots < 50 {
        return_value = true
    }
    return return_value
}

func main () {
    mux := http.NewServeMux()
    mux.HandleFunc("/", HelloHandler)
    addr := ":5000"
    log.Printf("server is listening at %s...", addr)
    log.Fatal(http.ListenAndServe(addr, mux))
}
