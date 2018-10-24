package reader

import (
    "encoding/csv"
    "path"
    "bufio"
    "strings"
    "log"
    "os"

    "github.com/thnery/go-google-geocoding/geocoding"
    "github.com/thnery/go-google-geocoding/util"
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

func processTxt(key, filePath string) {
    file, err := os.Open(filePath)
    util.HandleError(err)
    log.Print(file)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        body := geocoding.GetAddressFromGoogle(scanner.Text(), key)
        result := util.ParseResponseBody(body)
        util.PrintResult(result)
    }
}

func processCSV(key, filePath string) {
    file, err := os.Open(filePath)
    util.HandleError(err)
    defer file.Close()

    reader := csv.NewReader(bufio.NewReader(file))
    reader.Comma = ';'

    records, err := reader.ReadAll()
    util.HandleError(err)

    for i := range(records) {
        if i == 0 {
            // skip header
            continue
        }

        address := strings.Join(records[i], ",")
        if address != "" {
            body := geocoding.GetAddressFromGoogle(address, key)
            result := util.ParseResponseBody(body)
            util.PrintResult(result)
        }
    }
}

