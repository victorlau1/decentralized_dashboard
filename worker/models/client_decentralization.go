package models

import (
	"time"
)

type ClientDecentralization struct {
	Country    string `validate:"required"`
	Region     string `validate:"required"`
	Timezone   string
	City       string
	Blockchain string    `validate:"required"`
	Timestamp  time.Time `validate:"required"`
}
