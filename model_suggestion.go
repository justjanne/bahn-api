package bahn

import "time"

type Suggestion struct {
	Value            string     `json:"value,omitempty"yaml:"value,omitempty"`
	Cycle            string     `json:"cycle,omitempty"yaml:"cycle,omitempty"`
	Pool             string     `json:"pool,omitempty"yaml:"pool,omitempty"`
	Id               string     `json:"id,omitempty"yaml:"id,omitempty"`
	TrainLink        string     `json:"train_link,omitempty"yaml:"train_link,omitempty"`
	PublishedTime    *time.Time `json:"published_time,omitempty"yaml:"published_time,omitempty"`
	DepartureStation string     `json:"departure_station,omitempty"yaml:"departure_station,omitempty"`
	DepartureTime    *time.Time `json:"departure_time,omitempty"yaml:"departure_time,omitempty"`
	ArrivalStation   string     `json:"arrival_station,omitempty"yaml:"arrival_station,omitempty"`
	ArrivalTime      *time.Time `json:"arrival_time,omitempty"yaml:"arrival_time,omitempty"`
}
