package bahn

import "time"

type CoachSequence struct {
	Meta CoachSequenceMeta `json:"meta"yaml:"meta"`
	Data CoachSequenceData `json:"data"yaml:"data"`
}

type CoachSequenceMeta struct {
	Id          string     `json:"id"yaml:"id"`
	Owner       string     `json:"owner"yaml:"owner"`
	Format      string     `json:"format"yaml:"format"`
	Version     string     `json:"version"yaml:"version"`
	Correlation []string   `json:"correlation"yaml:"correlation"`
	Created     *time.Time `json:"created"yaml:"created"`
	Sequence    int        `json:"sequence"yaml:"sequence"`
}

type CoachSequenceData struct {
	ActualFormation CoachSequenceFormation `json:"actual_formation"yaml:"actual_formation"`
}

type CoachSequenceFormationDirection string

const (
	DirectionForwards  CoachSequenceFormationDirection = "FORWARDS"
	DirectionBackwards CoachSequenceFormationDirection = "BACKWARDS"
	DirectionUndefined CoachSequenceFormationDirection = "UNDEFINED"
	DirectionUnknown   CoachSequenceFormationDirection = "UNKNOWN"
)

type CoachSequenceFormation struct {
	Direction          CoachSequenceFormationDirection `json:"direction"yaml:"direction"`
	Groups             []CoachSequenceCoachGroup       `json:"groups"yaml:"groups"`
	Stop               CoachSequenceStop               `json:"stop"yaml:"stop"`
	Line               string                          `json:"line"yaml:"line"`
	Type               string                          `json:"type"yaml:"type"`
	TrainId            string                          `json:"train_id"yaml:"train_id"`
	ServiceId          string                          `json:"service_id"yaml:"service_id"`
	StartingDate       *time.Time                      `json:"starting_date"yaml:"starting_date"`
	JourneyId          string                          `json:"journey_id"yaml:"journey_id"`
	IsPlannedFormation bool                            `json:"is_planned_formation"yaml:"is_planned_formation"`
}

type CoachSequenceStop struct {
	Arrival          *time.Time                     `json:"arrival"yaml:"arrival"`
	Departure        *time.Time                     `json:"departure"yaml:"departure"`
	Station          string                         `json:"station"yaml:"station"`
	EvaId            string                         `json:"eva_id"yaml:"eva_id"`
	Platform         string                         `json:"platform"yaml:"platform"`
	StopId           string                         `json:"stop_id"yaml:"stop_id"`
	Rl100            string                         `json:"rl100"yaml:"rl100"`
	PlatformSections []CoachSequencePlatformSection `json:"platform_sections"yaml:"platform_sections"`
}

type CoachSequencePlatformSection struct {
	Name     string                        `json:"name"yaml:"name"`
	Position CoachSequencePlatformPosition `json:"position"yaml:"position"`
}

type CoachSequenceCoachGroup struct {
	TrainId     string               `json:"train_id"yaml:"train_id"`
	Description string               `json:"description"yaml:"description"`
	Coachs      []CoachSequenceCoach `json:"coachs"yaml:"coachs"`
	From        string               `json:"from"yaml:"from"`
	To          string               `json:"to"yaml:"to"`
}

type CoachSequenceCoach struct {
	Equipment        []CoachSequenceCoachEquipment `json:"equipment"yaml:"equipment"`
	Category         string                        `json:"category"yaml:"category"`
	CoachId          string                        `json:"coach_id"yaml:"coach_id"`
	Orientation      string                        `json:"orientation"yaml:"orientation"`
	GroupPosition    string                        `json:"group_position"yaml:"group_position"`
	CoachTypeInfo    CoachTypeInfo                 `json:"coach_type"yaml:"coach_type"`
	CoachOrdinal     string                        `json:"coach_ordinal"yaml:"coach_ordinal"`
	PlatformSection  string                        `json:"platform_section"yaml:"platform_section"`
	PlatformPosition CoachSequencePlatformPosition `json:"platform_position"yaml:"platform_position"`
	Status           string                        `json:"status"yaml:"status"`
}

type CoachSequenceCoachEquipment struct {
	Count       string `json:"count"yaml:"count"`
	Type        string `json:"type"yaml:"type"`
	Description string `json:"description"yaml:"description"`
	Status      string `json:"status"yaml:"status"`
}

type CoachSequencePlatformPosition struct {
	StartMeter   float64 `json:"start_meter"yaml:"start_meter"`
	EndMeter     float64 `json:"end_meter"yaml:"end_meter"`
	StartPercent int64   `json:"start_percent"yaml:"start_percent"`
	EndPercent   int64   `json:"end_percent"yaml:"end_percent"`
}
