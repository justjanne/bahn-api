package bahn

import "time"

type Station struct {
	Platforms   []string   `json:"platforms,omitempty"yaml:"platforms,omitempty"`
	Meta        []string   `json:"meta,omitempty"yaml:"meta,omitempty"`
	StationName string     `json:"station_name,omitempty"yaml:"station_name,omitempty"`
	EvaId       string     `json:"eva_id,omitempty"yaml:"eva_id,omitempty"`
	StationCode string     `json:"station_code,omitempty"yaml:"station_code,omitempty"`
	Db          bool       `json:"db,omitempty"yaml:"db,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"yaml:"created_at,omitempty"`
}
