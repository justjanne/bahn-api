package bahn

import "encoding/xml"

func WingDefinitionFromDecoder(source xml.Decoder) (WingDefinition, error) {
	var raw rawWingDefinition
	if err := source.Decode(&raw); err != nil {
		return WingDefinition{}, err
	}
	return parseWingDefinition(raw), nil
}

func WingDefinitionFromBytes(source []byte) (WingDefinition, error) {
	var raw rawWingDefinition
	if err := xml.Unmarshal(source, &raw); err != nil {
		return WingDefinition{}, err
	}
	return parseWingDefinition(raw), nil
}

type rawWingDefinition struct {
	XMLName xml.Name                 `xml:"wing-def"`
	Start   rawWingDefinitionElement `xml:"start,omitempty"`
	End     rawWingDefinitionElement `xml:"end,omitempty"`
}

func parseWingDefinition(data rawWingDefinition) WingDefinition {
	return WingDefinition{
		Start: parseWingDefinitionElement(data.Start),
		End:   parseWingDefinitionElement(data.End),
	}
}

type rawWingDefinitionElement struct {
	EvaId       string `xml:"eva,attr,omitempty"`
	StationName string `xml:"st-name,attr,omitempty"`
	PlannedTime string `xml:"pt,attr,omitempty"`
	Fl          bool   `xml:"fl,attr"`
}

func parseWingDefinitionElement(data rawWingDefinitionElement) WingDefinitionElement {
	return WingDefinitionElement{
		EvaId:       data.EvaId,
		StationName: data.StationName,
		PlannedTime: data.PlannedTime,
		Fl:          data.Fl,
	}
}
