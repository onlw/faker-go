package model

type AddressInfo struct {
	Country   string  `json:"country" xml:"country"`
	Address   string  `json:"address" xml:"address"`
	Street    string  `json:"street" xml:"street"`
	City      string  `json:"city" xml:"city"`
	State     string  `json:"state" xml:"state"`
	Zip       string  `json:"zip" xml:"zip"`
	Latitude  float64 `json:"latitude" xml:"latitude"`
	Longitude float64 `json:"longitude" xml:"longitude"`
}
