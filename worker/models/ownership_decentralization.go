package models

import (
	"time"
)

type OwnershipDecentralization struct {
	Address		string
	Balance		int64
	Blockchain	string
	TimeStamp	time.Time
}