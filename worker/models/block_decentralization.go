package models

import (
	"time"
)

type BlockDecentralization struct {
	BlockNumber	int	`validate:"required"`
	TimeStamp	time.Time	`validate:"required"`
	BlockMiner	string	`validate:"required"`
	Blockchain	string	`validate:"required"`
}