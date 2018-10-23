package reader

import (
    "encoding/csv"
    "path"
    "bufio"
    "strings"
    "log"

    "github.com/thnery/go-google-geocoding/geocoding"
)

func ReadFile(key, filePath string) {
    switch path.Ext(filePath) {
    case ".txt":
        processTxt(key, filePath)
    case ".csv":
        processCSV(key, filePath)
    default:
        log.Print("File not supported. Please use .txt or .csv files")
    }
}

func processTxt(key, filePath string) []string {
    file, err := os.Open(filePath)
    handleError(err)
    log.Print(file)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    results := []string
    for scanner.Scan() {
        body := geocoding.GetAddressFromGoogle(scanner.Text(), key)
        result := parseResponseBody(body)
        printResult(result)
    }
}

func processCSV(key, filePath string) []string {
    file, err := os.Open(filePath)
    handleError(err)
    defer file.Close()

    reader := csv.NewReader(bufio.NewReader(file))
    reader.Comma = ';'

    records, err := reader.ReadAll()
    handleError(err)

    results := []string

    for i := range(records) {
        if i == 0 {
            // skip header
            continue
        }

        address := strings.Join(records[i], ",")
        if address != "" {
            body := geocoding.GetAddressFromGoogle(address, key)
            result := parseResponseBody(body)
            printResult(result)
        }
    }

    return results
}

