package models

type NodeDecentralization struct {
	Country   string
	Host      string
	Isp       string
	ASN       float64
	Validator bool //Required for PoS system. Defaults to false for PoW
	Longitude float64
	Latitude  float64
}
