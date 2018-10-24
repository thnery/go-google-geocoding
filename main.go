package main

import (
    "fmt"
    "os"
    "log"
    "flag"

    "github.com/thnery/go-google-geocoding/reader"
    "github.com/thnery/go-google-geocoding/util"
    "github.com/thnery/go-google-geocoding/geocoding"
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

    if file != nil && *file != "" {
        if address != nil && *address != "" {
            log.Print("[WARNING] arg -address will be ignore due to arg -file has been set")
        }

        reader.ReadFile(*key, *file)
    } else {
        body := geocoding.GetAddressFromGoogle(*address, *key)
        result := util.ParseResponseBody(body)
        util.PrintResult(result)
    }
}

