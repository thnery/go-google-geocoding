package data

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
