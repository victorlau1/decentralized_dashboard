package models

import (
	"time"
)

type ClientDecentralization struct {
	Country       string `validate:"required"`
	Client        string `validate:"required"`
	Region        string `validate:"required"`
	Timezone      string
	City          string
	Blockchain    string    `validate:"required"`
	Timestamp     time.Time `validate:"required"`
	ClientVersion string    `validate:"required"`
	Organization  string    // Organization who owns this
	IPAddress     string
	Longitude     float64
	Latitude      float64
	ASN           float64 //Autonomous System Number
}
