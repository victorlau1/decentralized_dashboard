package models

import "time"

type NodeDecentralization struct {
	ASN        float64 //Grouping of IP addresses managed by 1 service
	Blockchain string  `validate:"required"`
	City       string
	Country    string
	IPAddress  string
	Longitude  float64
	Latitude   float64
	Timestamp  time.Time
	Validator  bool //Required for PoS system. Defaults to false for PoW
}
