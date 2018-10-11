package main

import (
    "fmt"
    "os"
    "flag"
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

    fmt.Println(address)
}

func processFile(file string, key string) {
    fmt.Println("TODO")
}
