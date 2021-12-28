package models

import (
	"time"
)

type DevDecentralization struct {
	CommitID	string
	Committer	string
	Email	string
	TimeStamp	time.Time
	Client	string
	Blockchain	string
}