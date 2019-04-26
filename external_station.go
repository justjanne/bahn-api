package bahn

import (
	"encoding/xml"
	"io"
)

func StationsFromReader(source io.Reader) ([]Station, error) {
	var raw rawStations
	if err := xml.NewDecoder(source).Decode(&raw); err != nil {
		return make([]Station, 0), err
	}
	return parseStations(raw), nil
}

func StationsFromBytes(source []byte) ([]Station, error) {
	var raw rawStations
	if err := xml.Unmarshal(source, &raw); err != nil {
		return make([]Station, 0), err
	}
	return parseStations(raw), nil
}

type rawStations struct {
	XMLName  xml.Name     `xml:"stations"`
	Stations []rawStation `xml:"station"`
}

func parseStations(data rawStations) []Station {
	result := make([]Station, len(data.Stations))
	for i, element := range data.Stations {
		result[i] = parseStation(element)
	}
	return result
}

type rawStation struct {
	Platforms   *bahnStringList `xml:"p,attr,omitempty"`
	Meta        *bahnStringList `xml:"meta,attr,omitempty"`
	StationName string          `xml:"name,attr,omitempty"`
	EvaId       string          `xml:"eva,attr,omitempty"`
	StationCode string          `xml:"ds100,attr,omitempty"`
	Db          bool            `xml:"db,attr,omitempty"`
	CreatedAt   *bahnTime       `xml:"creationts,attr,omitempty"`
}

func parseStation(data rawStation) Station {
	return Station{
		Platforms:   data.Platforms.Value(),
		Meta:        data.Meta.Value(),
		StationName: data.StationName,
		EvaId:       data.EvaId,
		StationCode: data.StationCode,
		Db:          data.Db,
		CreatedAt:   data.CreatedAt.Value(),
	}
}
