package util

import (
    "encoding/json"
    "log"
    "os"

    "github.com/thnery/go-google-geocoding/data"
)

func ParseResponseBody(body []byte) data.Result {
    var results data.Results
    json.Unmarshal(body, &results)
    result := results.Results[0]
    return result
}

func PrintResult(result data.Result) {
    log.Print("Address: ", result.FormattedAddress)
    log.Print("Latitude: ", result.Geometry.Location.Latitude)
    log.Print("Longitude: ", result.Geometry.Location.Longitude)
}

func HandleError(err error) {
    if err != nil {
        log.Print(err)
        os.Exit(1)
    }
}

