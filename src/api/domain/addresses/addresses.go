package addresses

type Address struct {
    Id           int64  `json:"id"`
    StreetName   string `json:"street_name"`
    StreetNumber int32  `json:"street_number"`
    City         string `json:"city"`
}
