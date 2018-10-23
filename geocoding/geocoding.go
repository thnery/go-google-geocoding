package geocoding

import(
    "net/http"
    "log"
    "io/ioutil"
)

func GetAddressFromGoogle(address, key string) []byte {
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

