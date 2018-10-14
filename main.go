package main

import (
    "fmt"
    "os"
    "log"
    "flag"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "path"
)

type Results struct {
    Results []Result `json:"results"`
}

type Result struct {
    FormattedAddress string `json:"formatted_address"`
    Geometry Geometry `json:"geometry"`
}

type Geometry struct {
    Location Location `json:"location"`
}

type Location struct {
    Latitude float32 `json:"lat"`
    Longitude float32 `json:"lng"`
}

func main() {
    key := flag.String("key", "", "Google Geocoding API Key")
    file := flag.String("file", "", "Addresses file path")
    address := flag.String("address", "", "Address string to be found")

    flag.Usage = func() {
        fmt.Println("Welcome to Go Geocoder")
        os.Exit(1)
    }

    flag.Parse()

    if file != nil && *file != "" {
        if address != nil && *address != "" {
            log.Print("[WARNING] arg -address will be ignore due to arg -file has been set")
        }

        readFile(*file)
    } else {
        body := getAddressFromGoogle(*address, *key)
        result := parseResponseBody(body)
        fmt.Println("Address: " + result.FormattedAddress)
        fmt.Printf("Latitude: %f\n", result.Geometry.Location.Latitude)
        fmt.Printf("Longitude: %f", result.Geometry.Location.Longitude)
    }
}

func getAddressFromGoogle(address, key string) []byte {
    if address == "" {
        os.Exit(1)
    }

    client := &http.Client{}
    req, err := http.NewRequest("GET", "https://maps.googleapis.com/maps/api/geocode/json", nil)

    handleError(err)

    query := req.URL.Query()
    query.Add("key", key)
    query.Add("address", address)

    req.URL.RawQuery = query.Encode()

    log.Print(req.URL.String())
    log.Print(client)

    resp, err := client.Do(req)

    handleError(err)

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    handleError(err)

    return body
}

func parseResponseBody(body []byte) Result {
    var results Results
    json.Unmarshal(body, &results)
    result := results.Results[0]
    return result
}

// TODO: read and address file (txt or csv) and fetch the coordinates
// from Google Geocoding API
func readFile(file string) {
    switch path.Ext(file) {
    case ".txt":
        processTxt(file)
    case ".csv":
        processCSV(file)
    default:
        log.Print("File not supported. Please use .txt or .csv files")
    }
}

func processTxt(file string) {
    log.Print(file)
}

func processCSV(file string) {
    log.Print(file)
}

func handleError(err error) {
    if err != nil {
        log.Print(err)
        os.Exit(1)
    }
}

