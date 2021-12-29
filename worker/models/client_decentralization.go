package models

type ClientDecentralization struct {
	ASN           float64 //Autonomous System Number
	Blockchain    string  `validate:"required"`
	City          string
	Country       string `validate:"required"`
	Client        string `validate:"required"`
	ClientVersion string `validate:"required"`
	IPAddress     string
	Longitude     float64
	Latitude      float64
	Organization  string // Organization who manages this
	Region        string `validate:"required"`
	Validator     bool   //Required for PoS system. Defaults to false for PoW Systems
	Timestamp     string `validate:"required"` // Format must be time.RFC3339
	Timezone      string
}
