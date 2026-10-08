package models

import "time"

type Policy struct {
	ID        string
	Number    string
	StartDate time.Time
	EndDate   time.Time
}
