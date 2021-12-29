package models

type ClientDecentralization struct {
	ASN           float64 `fake:"{number:1,99999}"` //Autonomous System Number
	Blockchain    string  `validate:"required" fake:"{randomstring:[Solana, Ethereum, Bitcoin]}"`
	City          string  `fake:"{city}"`
	Country       string  `validate:"required" fake:"{countryabr}"`
	Client        string  `validate:"required" fake:"{randomstring:[Rust, Geth, Satoshi]}"`
	ClientVersion string  `validate:"required" fake:"{streetnumber}"`
	IPAddress     string  `fake:"{ipv6address}"`
	Longitude     float64 `fake:"{longitude}"`
	Latitude      float64 `fake:"{latitude}"`
	Organization  string  `fake:"{gamertag}"` // Organization who manages this
	Region        string  `validate:"required" fake:"{timezoneregion}"`
	Validator     bool    `fake:"{bool}"`                   //Required for PoS system. Defaults to false for PoW Systems
	Timestamp     string  `validate:"required" fake:"skip"` // Format must be time.RFC3339
	Timezone      string  `fake:"{timezone}"`
}
