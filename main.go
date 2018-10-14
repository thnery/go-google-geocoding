package main

import (
    "fmt"
    "os"
    "log"
    "flag"
    "net/http"
    "io/ioutil"
)

type simple_address struct {
    FormattadedAddress string
    Latitude float32
    Longitude float32
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

    fmt.Println(*key)
    fmt.Println(*file)

    body := fetchAddress(*address, *key)

    fmt.Println(body)
}

func fetchAddress(address, key string) string {
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

    return string(body)
}

func parseResponseBody(body string) string {
    // TODO: parse the response body
}

// TODO: read and address file (txt or csv) and fetch the coordinates
// from Google Geocoding API
func processFile(file string, key string) {
    fmt.Println("TODO")
}

func handleError(err error) {
    if err != nil {
        log.Print(err)
        os.Exit(1)
    }
}

