package bahn

import (
	"encoding/json"
	"io"
	"time"
)

func SuggestionsFromReader(source io.Reader) ([]Suggestion, error) {
	var raw rawSuggestions
	if err := json.NewDecoder(source).Decode(&raw); err != nil {
		return make([]Suggestion, 0), err
	}
	return parseSuggestions(raw), nil
}

func SuggestionsFromBytes(source []byte) ([]Suggestion, error) {
	var raw rawSuggestions
	if err := json.Unmarshal(source, &raw); err != nil {
		return make([]Suggestion, 0), err
	}
	return parseSuggestions(raw), nil
}

type rawSuggestions struct {
	Suggestions []rawSuggestion
}

func parseSuggestions(data rawSuggestions) []Suggestion {
	result := make([]Suggestion, len(data.Suggestions))
	for i, element := range data.Suggestions {
		result[i] = parseSuggestion(element)
	}
	return result
}

type rawSuggestion struct {
	Value            string `json:"value"`
	Cycle            string `json:"cycle"`
	Pool             string `json:"pool"`
	Id               string `json:"id"`
	TrainLink        string `json:"trainLink"`
	JourneyParams    string `json:"journParam"`
	PublishedTime    string `json:"pubTime"`
	PublishedDate    string `json:"pubDate"`
	DepartureStation string `json:"dep"`
	DepartureDate    string `json:"depDate"`
	DepartureTime    string `json:"depTime"`
	ArrivalStation   string `json:"arr"`
	ArrivalTime      string `json:"arrTime"`
	ArrivalDate      string `json:"arrDate"`
}

func parseTime(dateStr string, timeStr string) *time.Time {
	DateFormat := "02.01.2006"
	DateTimeFormat := "02.01.2006 15:04"
	if dateStr == "" {
		dateStr = time.Now().Format(DateFormat)
	}

	dateTime, err := time.Parse(DateTimeFormat, dateStr+" "+timeStr)
	if err != nil {
		return nil
	}
	return &dateTime
}

func parseSuggestion(data rawSuggestion) Suggestion {
	return Suggestion{
		Value:            data.Value,
		Cycle:            data.Cycle,
		Pool:             data.Pool,
		Id:               data.Id,
		TrainLink:        data.TrainLink,
		PublishedTime:    parseTime(data.PublishedDate, data.PublishedTime),
		DepartureStation: data.DepartureStation,
		DepartureTime:    parseTime(data.DepartureDate, data.DepartureTime),
		ArrivalStation:   data.ArrivalStation,
		ArrivalTime:      parseTime(data.ArrivalDate, data.ArrivalTime),
	}
}
