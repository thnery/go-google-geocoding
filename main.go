package main

import (
    "fmt"
    "os"
    "log"
    "flag"
    "net/http"
)

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

    fetchAddress(*address, *key)
}

func fetchAddress(address string, key string) {
    if address == "" {
        os.Exit(1)
    }

    client := &http.Client{}
    req, err := http.NewRequest("GET", "https://maps.googleapis.com/maps/api/geocode/json", nil)

    if err != nil {
        log.Print(err)
        os.Exit(1)
    }

    query := req.URL.Query()
    query.Add("key", key)
    query.Add("address", address)

    req.URL.RawQuery = query.Encode()

    log.Print(req.URL.String())
    log.Print(client)
}

func processFile(file string, key string) {
    fmt.Println("TODO")
}
