package main

import (
    "fmt"
    "os"
    "log"
    "flag"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "encoding/csv"
    "path"
    "bufio"
    "strings"

    "github.com/thnery/go-google-geocoding/reader"
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

        // readFile(*key, *file)
        reader.ReadFile(*key, *file)
    } else {
        body := getAddressFromGoogle(*address, *key)
        result := parseResponseBody(body)
        printResult(result)
    }
}

func parseResponseBody(body []byte) Result {
    var results Results
    json.Unmarshal(body, &results)
    result := results.Results[0]
    return result
}

func printResult(result Result) {
    log.Print("Address: ", result.FormattedAddress)
    log.Print("Latitude: ", result.Geometry.Location.Latitude)
    log.Print("Longitude: ", result.Geometry.Location.Longitude)
}

func handleError(err error) {
    if err != nil {
        log.Print(err)
        os.Exit(1)
    }
}

