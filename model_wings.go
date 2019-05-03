package bahn

type WingDefinition struct {
	Start WingDefinitionElement `json:"start"yaml:"start"`
	End   WingDefinitionElement `json:"end"yaml:"end"`
}

type WingDefinitionElement struct {
	EvaId       string `json:"eva_id,omitempty"yaml:"eva_id,omitempty"`
	StationName string `json:"station_name,omitempty"yaml:"station_name,omitempty"`
	PlannedTime string `json:"planned_time,omitempty"yaml:"planned_time,omitempty"`
	Fl          bool   `json:"fl,omitempty"yaml:"fl,omitempty"`
}
