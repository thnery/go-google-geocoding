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

        readFile(*key, *file)
    } else {
        body := getAddressFromGoogle(*address, *key)
        result := parseResponseBody(body)
        printResult(result)
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

func readFile(key, filePath string) {
    switch path.Ext(filePath) {
    case ".txt":
        processTxt(key, filePath)
    case ".csv":
        processCSV(key, filePath)
    default:
        log.Print("File not supported. Please use .txt or .csv files")
    }
}

func processTxt(key, filePath string) {
    file, err := os.Open(filePath)
    handleError(err)
    log.Print(file)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        body := getAddressFromGoogle(scanner.Text(), key)
        result := parseResponseBody(body)
        printResult(result)
    }
}

func processCSV(key, filePath string) {
    file, err := os.Open(filePath)
    handleError(err)
    defer file.Close()

    reader := csv.NewReader(bufio.NewReader(file))
    reader.Comma = ';'

    records, err := reader.ReadAll()
    handleError(err)

    for i := range(records) {
        if i == 0 {
            // skip header
            continue
        }

        log.Print(records[i])
        address := strings.Join(records[i], ",")
        log.Print(address)
        if address != "" {
            body := getAddressFromGoogle(address, key)
            result := parseResponseBody(body)
            printResult(result)
        }
    }
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

