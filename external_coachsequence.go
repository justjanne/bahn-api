package bahn

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func CoachSequenceFromReader(source io.Reader) (CoachSequence, error) {
	var raw rawCoachSequence
	if err := json.NewDecoder(source).Decode(&raw); err != nil {
		return CoachSequence{}, err
	}
	return parseCoachSequence(raw), nil
}

func CoachSequenceFromBytes(source []byte) (CoachSequence, error) {
	var raw rawCoachSequence
	if err := json.Unmarshal(source, &raw); err != nil {
		return CoachSequence{}, err
	}
	return parseCoachSequence(raw), nil
}

type rawCoachSequence struct {
	Meta rawCoachSequenceMeta `json:"meta"`
	Data rawCoachSequenceData `json:"data"`
}

func parseCoachSequence(data rawCoachSequence) CoachSequence {
	return CoachSequence{
		Meta: parseCoachSequenceMeta(data.Meta),
		Data: parseCoachSequenceData(data.Data),
	}
}

type rawCoachSequenceMeta struct {
	Id          string    `json:"id"`
	Owner       string    `json:"owner"`
	Format      string    `json:"format"`
	Version     string    `json:"version"`
	Correlation []string  `json:"correlation"`
	Created     time.Time `json:"created"`
	Sequence    int       `json:"sequence"`
}

func parseCoachSequenceMeta(data rawCoachSequenceMeta) CoachSequenceMeta {
	return CoachSequenceMeta{
		Id:          data.Id,
		Owner:       data.Owner,
		Format:      data.Format,
		Version:     data.Version,
		Correlation: data.Correlation,
		Created:     &data.Created,
		Sequence:    data.Sequence,
	}
}

type rawCoachSequenceData struct {
	ActualFormation rawCoachSequenceFormation `json:"istformation"`
}

func parseCoachSequenceData(data rawCoachSequenceData) CoachSequenceData {
	return CoachSequenceData{
		ActualFormation: parseCoachSequenceFormation(data.ActualFormation),
	}
}

const (
	rawDirectionForwards  = "VORWAERTS"
	rawDirectionBackwards = "RUECKWAERTS"
	rawDirectionUndefined = "UNDEFINIERT"
)

func parseDirection(data string) CoachSequenceFormationDirection {
	switch data {
	case rawDirectionForwards:
		return DirectionForwards
	case rawDirectionBackwards:
		return DirectionBackwards
	case rawDirectionUndefined:
		return DirectionUndefined
	default:
		return DirectionUnknown
	}
}

type rawCoachSequenceFormation struct {
	Direction          string                       `json:"fahrtrichtung"`
	Groups             []rawCoachSequenceCoachGroup `json:"allFahrzeuggruppe"`
	Stop               rawCoachSequenceStop         `json:"halt"`
	Line               string                       `json:"liniebezeichnung"`
	Type               string                       `json:"zuggattung"`
	TrainId            string                       `json:"zugnummer"`
	ServiceId          string                       `json:"serviceid"`
	StartingDate       bahnDate                     `json:"planstarttag"`
	JourneyId          string                       `json:"fahrtid"`
	IsPlannedFormation bool                         `json:"istplaninformation"`
}

func parseCoachSequenceFormation(data rawCoachSequenceFormation) CoachSequenceFormation {
	return CoachSequenceFormation{
		Direction:          parseDirection(data.Direction),
		Groups:             parseCoachSequenceCoachGroups(data.Groups),
		Stop:               parseCoachSequenceStop(data.Stop),
		Line:               data.Line,
		Type:               data.Type,
		TrainId:            data.TrainId,
		ServiceId:          data.ServiceId,
		StartingDate:       data.StartingDate.Value(),
		JourneyId:          data.JourneyId,
		IsPlannedFormation: data.IsPlannedFormation,
	}
}

type rawCoachSequenceStop struct {
	Departure        timeMedium                        `json:"abfahrtszeit"`
	Arrival          timeMedium                        `json:"ankunftszeit"`
	Station          string                            `json:"bahnhofsname"`
	EvaId            string                            `json:"evanummer"`
	Platform         string                            `json:"gleisbezeichnung"`
	StopId           string                            `json:"haltid"`
	Rl100            string                            `json:"rl100"`
	PlatformSections []rawCoachSequencePlatformSection `json:"allSektor"`
}

func parseCoachSequenceStop(data rawCoachSequenceStop) CoachSequenceStop {
	return CoachSequenceStop{
		Arrival:          data.Arrival.Value(),
		Departure:        data.Departure.Value(),
		Station:          data.Station,
		EvaId:            data.EvaId,
		Platform:         data.Platform,
		StopId:           data.StopId,
		Rl100:            data.Rl100,
		PlatformSections: parseCoachSequencePlatformSections(data.PlatformSections),
	}
}

type rawCoachSequencePlatformSection struct {
	Position rawCoachSequencePlatformPosition `json:"positionamgleis"`
	Name     string                           `json:"sektorbezeichnung"`
}

func parseCoachSequencePlatformSections(data []rawCoachSequencePlatformSection) []CoachSequencePlatformSection {
	result := make([]CoachSequencePlatformSection, len(data))
	for i, element := range data {
		result[i] = parseCoachSequencePlatformSection(element)
	}
	return result
}

func parseCoachSequencePlatformSection(data rawCoachSequencePlatformSection) CoachSequencePlatformSection {
	return CoachSequencePlatformSection{
		Name:     data.Name,
		Position: parseCoachSequencePlatformPosition(data.Position),
	}
}

type rawCoachSequenceCoachGroup struct {
	Coachs      []rawCoachSequenceCoach `json:"allFahrzeug"`
	Description string                  `json:"fahrzeuggruppebezeichnung"`
	To          string                  `json:"zielbetriebsstellename"`
	From        string                  `json:"startbetriebsstellename"`
	TrainId     string                  `json:"verkehrlichezugnummer"`
}

func parseCoachSequenceCoachGroups(data []rawCoachSequenceCoachGroup) []CoachSequenceCoachGroup {
	result := make([]CoachSequenceCoachGroup, len(data))
	for i, element := range data {
		result[i] = parseCoachSequenceCoachGroup(element)
	}
	return result
}

func parseCoachSequenceCoachGroup(data rawCoachSequenceCoachGroup) CoachSequenceCoachGroup {
	return CoachSequenceCoachGroup{
		TrainId:     data.TrainId,
		Description: data.Description,
		Coachs:      parseCoachSequenceCoachs(data.Coachs),
		From:        data.From,
		To:          data.To,
	}
}

type rawCoachSequenceCoach struct {
	Equipment        []rawCoachSequenceCoachEquipment `json:"allFahrzeugausstattung"`
	Category         string                           `json:"kategorie"`
	CoachId          string                           `json:"fahrzeugnummer"`
	Orientation      string                           `json:"orientierung"`
	GroupPosition    string                           `json:"positioningruppe"`
	PlatformSection  string                           `json:"fahrzeugsektor"`
	CoachType        string                           `json:"fahrzeugtyp"`
	CoachOrdinal     string                           `json:"wagenordnungsnummer"`
	PlatformPosition rawCoachSequencePlatformPosition `json:"positionamhalt"`
	Status           string                           `json:"status"`
}

func parseCoachSequenceCoachs(data []rawCoachSequenceCoach) []CoachSequenceCoach {
	result := make([]CoachSequenceCoach, len(data))
	for i, element := range data {
		result[i] = parseCoachSequenceCoach(element)
	}
	return result
}

func parseCoachSequenceCoach(data rawCoachSequenceCoach) CoachSequenceCoach {
	return CoachSequenceCoach{
		Equipment:        parseCoachSequenceCoachEquipments(data.Equipment),
		Category:         data.Category,
		CoachId:          data.CoachId,
		Orientation:      data.Orientation,
		GroupPosition:    data.GroupPosition,
		CoachTypeInfo:    parseCoachTypeInfo(data.CoachType),
		CoachOrdinal:     data.CoachOrdinal,
		PlatformSection:  data.PlatformSection,
		PlatformPosition: parseCoachSequencePlatformPosition(data.PlatformPosition),
		Status:           data.Status,
	}
}

type rawCoachSequenceCoachEquipment struct {
	Count       string `json:"anzahl"`
	Type        string `json:"ausstattungsart"`
	Description string `json:"bezeichnung"`
	Status      string `json:"status"`
}

func parseCoachSequenceCoachEquipments(data []rawCoachSequenceCoachEquipment) []CoachSequenceCoachEquipment {
	result := make([]CoachSequenceCoachEquipment, len(data))
	for i, element := range data {
		result[i] = parseCoachSequenceCoachEquipment(element)
	}
	return result
}

func parseCoachSequenceCoachEquipment(data rawCoachSequenceCoachEquipment) CoachSequenceCoachEquipment {
	return CoachSequenceCoachEquipment{
		Count:       data.Count,
		Type:        data.Type,
		Description: data.Description,
		Status:      data.Status,
	}
}

type rawCoachSequencePlatformPosition struct {
	EndMeter     string `json:"endemeter"`
	EndPercent   string `json:"endeprozent"`
	StartMeter   string `json:"startmeter"`
	StartPercent string `json:"startprozent"`
}

func parseCoachSequencePlatformPosition(data rawCoachSequencePlatformPosition) CoachSequencePlatformPosition {
	var result CoachSequencePlatformPosition

	result.StartMeter, _ = strconv.ParseFloat(data.StartMeter, 64)
	result.EndMeter, _ = strconv.ParseFloat(data.EndMeter, 64)

	result.StartPercent, _ = strconv.ParseInt(data.StartPercent, 10, 64)
	result.EndPercent, _ = strconv.ParseInt(data.EndPercent, 10, 64)

	return result
}

func extractClass(text *string, prefix string) bool {
	if strings.Contains(*text, prefix) {
		*text = strings.Replace(*text, prefix, "", 1)
		return true
	} else {
		return false
	}
}

func extractFeature(text *string, prefix string) bool {
	if strings.Contains(*text, prefix) {
		*text = strings.Replace(*text, prefix, "", 1)
		return true
	} else {
		return false
	}
}

func parseCoachTypeInfo(coachType string) CoachTypeInfo {
	var coachInfo CoachTypeInfo

	coachInfo.RawType = coachType

	var class string
	var features string
	for i, char := range coachType {
		if !unicode.IsUpper(char) {
			class = coachType[0:i]
			if unicode.IsLetter(char) {
				features = coachType[i:]
			}
			break
		}
	}
	if class == "" {
		class = coachType[0:]
	}

	if extractClass(&class, "A") {
		coachInfo.FirstClass = true
	}

	if extractClass(&class, "B") {
		coachInfo.SecondClass = true
	}

	if extractClass(&class, "R") {
		if extractClass(&class, "W") {
			coachInfo.Restaurant = true
		} else {
			coachInfo.Bistro = true
		}
	}

	if extractClass(&class, "D") {
		coachInfo.DoubleDeck = true
		if extractClass(&class, "D") {
			coachInfo.CarTransport = true
		}
	}

	if extractClass(&class, "G") {
		coachInfo.Saloon = true
	}

	if extractClass(&class, "L") {
		coachInfo.Sleeping = true
	}

	if extractClass(&class, "E") {
		coachInfo.ElectricLocomotive = true
	}

	if extractClass(&class, "I") {
		coachInfo.DieselElectricLocomotive = true
	}

	if extractClass(&class, "V") {
		coachInfo.DieselLocomotive = true
	}

	if extractFeature(&features, "b") {
		coachInfo.Accessible = true
	}

	if extractFeature(&features, "c") {
		coachInfo.Couchette = true
	}

	if extractFeature(&features, "k") {
		coachInfo.Bistro = true
		coachInfo.Restaurant = false
	}

	if extractFeature(&features, "m") {
		coachInfo.Compartments = true
		if extractFeature(&features, "m") {
			coachInfo.InterCity = true
		}
		if extractFeature(&features, "i") {
			coachInfo.InterRegio = true
		}
		if extractFeature(&features, "d") {
			coachInfo.Bicycle = true
		}
		if extractFeature(&features, "v") {
			coachInfo.Compartments = true
		}
	}

	if extractFeature(&features, "p") {
		coachInfo.AirConditioning = true
		coachInfo.OpenCoach = true
	}

	if extractFeature(&features, "f") {
		coachInfo.ControlCar = true
	}

	if extractFeature(&features, "s") {
		coachInfo.ServicePoint = true
	}

	if extractFeature(&features, "w") {
		coachInfo.ReducedCompartmentCount = true
	}

	// TAV-enabled
	extractFeature(&features, "a")
	// Electrical heating, power supply through head-end power
	extractFeature(&features, "z")
	// Electrical heating, power supply through axial generators
	extractFeature(&features, "h")
	// Light
	extractFeature(&features, "l")

	// apparently only in use in the EuroCityExpress with ETR 610
	extractFeature(&features, "e")
	// Control car, apparently only in use in the EuroCityExpress with ETR 610
	extractFeature(&features, "t")

	return coachInfo
}
