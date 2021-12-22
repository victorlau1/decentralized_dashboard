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

type BlockDecentralization struct {
	BlockNumber	int	`validate:"required"`
	TimeStamp	int	`validate:"required"`
	BlockMiner	string	`validate:"required"`
	Blockchain	string	`validate:"required"`
}

