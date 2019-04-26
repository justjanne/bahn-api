package bahn

import "time"

type Suggestion struct {
	Value            string
	Cycle            string
	Pool             string
	Id               string
	TrainLink        string
	PublishedTime    *time.Time
	DepartureStation string
	DepartureTime    *time.Time
	ArrivalStation   string
	ArrivalTime      *time.Time
}
