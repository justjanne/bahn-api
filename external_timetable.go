package bahn

import (
	"encoding/xml"
	"fmt"
	"io"
)

func TimetableFromReader(source io.Reader) (Timetable, error) {
	var raw rawTimetable
	if err := xml.NewDecoder(source).Decode(&raw); err != nil {
		return Timetable{}, err
	}
	return parseTimetable(raw), nil
}

func TimetableFromBytes(source []byte) (Timetable, error) {
	var raw rawTimetable
	if err := xml.Unmarshal(source, &raw); err != nil {
		return Timetable{}, err
	}
	return parseTimetable(raw), nil
}

type rawTimetable struct {
	XMLName  xml.Name           `xml:"timetable"`
	Station  string             `xml:"station,attr,omitempty"`
	EvaId    int64              `xml:"eva,attr,omitempty"`
	Stops    []rawTimetableStop `xml:"s,omitempty"`
	Messages []rawMessage       `xml:"m,omitempty"`
}

func parseTimetable(data rawTimetable) Timetable {
	return Timetable{
		Station:  data.Station,
		EvaId:    data.EvaId,
		Stops:    parseTimetableStops(data.Stops),
		Messages: parseMessages(data.Messages),
	}
}

type rawMessage struct {
	MessageId           string                  `xml:"id,attr,omitempty"`
	Type                rawMessageType          `xml:"t,attr,omitempty"`
	From                *timeShort              `xml:"from,attr,omitempty"`
	To                  *timeShort              `xml:"to,attr,omitempty"`
	Code                *int                    `xml:"c,attr,omitempty"`
	InternalText        string                  `xml:"int,attr,omitempty"`
	ExternalText        string                  `xml:"ext,attr,omitempty"`
	Category            string                  `xml:"cat,attr,omitempty"`
	ExternalCategory    string                  `xml:"ec,attr,omitempty"`
	Timestamp           *timeShort              `xml:"ts,attr,omitempty"`
	Priority            rawPriority             `xml:"pr,attr,omitempty"`
	Owner               string                  `xml:"o,attr,omitempty"`
	ExternalLink        string                  `xml:"elnk,attr,omitempty"`
	Deleted             int                     `xml:"del,attr,omitempty"`
	DistributorMessages []rawDistributorMessage `xml:"dm,omitempty"`
	TripLabel           []rawTripLabel          `xml:"tl,omitempty"`
}

func parseMessages(data []rawMessage) []Message {
	result := make([]Message, len(data))
	for i, element := range data {
		result[i] = parseMessage(element)
	}
	return result
}

func parseMessage(data rawMessage) Message {
	var code int
	if data.Code != nil {
		code = *data.Code
	}
	return Message{
		MessageId:           data.MessageId,
		Type:                parseMessageType(data.Type),
		From:                data.From.Value(),
		To:                  data.To.Value(),
		Code:                code,
		InternalText:        data.InternalText,
		ExternalText:        data.ExternalText,
		Category:            data.Category,
		ExternalCategory:    data.ExternalCategory,
		Timestamp:           data.Timestamp.Value(),
		Priority:            parsePriority(data.Priority),
		Deleted:             data.Deleted != 0,
		DistributorMessages: parseDistributorMessages(data.DistributorMessages),
	}
}

type rawMessageType string

const (
	rawMessageTypeHafasInformationManager rawMessageType = "h"
	rawMessageTypeQualityChange           rawMessageType = "q"
	rawMessageTypeFreeText                rawMessageType = "f"
	rawMessageTypeCauseOfDelay            rawMessageType = "d"
	rawMessageTypeIbis                    rawMessageType = "i"
	rawMessageTypeIbisUnassigned          rawMessageType = "u"
	rawMessageTypeDisruption              rawMessageType = "r"
	rawMessageTypeConnection              rawMessageType = "c"
	rawMessageTypeUndefined               rawMessageType = ""
)

func parseMessageType(data rawMessageType) MessageType {
	switch data {
	case rawMessageTypeHafasInformationManager:
		return MessageTypeHafasInformationManager
	case rawMessageTypeQualityChange:
		return MessageTypeQualityChange
	case rawMessageTypeFreeText:
		return MessageTypeFreeText
	case rawMessageTypeCauseOfDelay:
		return MessageTypeCauseOfDelay
	case rawMessageTypeIbis:
		return MessageTypeIbis
	case rawMessageTypeIbisUnassigned:
		return MessageTypeIbisUnassigned
	case rawMessageTypeDisruption:
		return MessageTypeDisruption
	case rawMessageTypeConnection:
		return MessageTypeConnection
	case rawMessageTypeUndefined:
		return MessageTypeUndefined
	default:
		fmt.Printf("Could not parse MessageType '%s'\n", data)
		return MessageTypeUnknown
	}
}

type rawPriority string

const (
	rawPriorityHigh      rawPriority = "1"
	rawPriorityMedium    rawPriority = "2"
	rawPriorityLow       rawPriority = "3"
	rawPriorityDone      rawPriority = "4"
	rawPriorityUndefined rawPriority = ""
)

func parsePriority(data rawPriority) Priority {
	switch data {
	case rawPriorityHigh:
		return PriorityHigh
	case rawPriorityMedium:
		return PriorityMedium
	case rawPriorityLow:
		return PriorityLow
	case rawPriorityDone:
		return PriorityDone
	case rawPriorityUndefined:
		return PriorityUndefined
	default:
		fmt.Printf("Could not parse Priority '%s'\n", data)
		return PriorityUnknown
	}
}

type rawDistributorType string

const (
	rawDistributorTypeCity         rawDistributorType = "s"
	rawDistributorTypeRegion       rawDistributorType = "r"
	rawDistributorTypeLongDistance rawDistributorType = "f"
	rawDistributorTypeOther        rawDistributorType = "x"
	rawDistributorTypeUndefined    rawDistributorType = ""
)

func parseDistributorType(data rawDistributorType) DistributorType {
	switch data {
	case rawDistributorTypeCity:
		return DistributorTypeCity
	case rawDistributorTypeRegion:
		return DistributorTypeRegion
	case rawDistributorTypeLongDistance:
		return DistributorTypeLongDistance
	case rawDistributorTypeOther:
		return DistributorTypeOther
	case rawDistributorTypeUndefined:
		return DistributorTypeUndefined
	default:
		fmt.Printf("Could not parse DistributorType '%s'\n", data)
		return DistributorTypeUnknown
	}
}

type rawDistributorMessage struct {
	DistributorType rawDistributorType `xml:"t,attr,omitempty"`
	DistributorName string             `xml:"n,attr,omitempty"`
	InternalText    string             `xml:"int,attr,omitempty"`
	Timestamp       *timeShort         `xml:"ts,attr,omitempty"`
}

func parseDistributorMessages(data []rawDistributorMessage) []DistributorMessage {
	result := make([]DistributorMessage, len(data))
	for i, element := range data {
		result[i] = parseDistributorMessage(element)
	}
	return result
}

func parseDistributorMessage(data rawDistributorMessage) DistributorMessage {
	return DistributorMessage{
		DistributorType: parseDistributorType(data.DistributorType),
		DistributorName: data.DistributorName,
		InternalText:    data.InternalText,
		Timestamp:       data.Timestamp.Value(),
	}
}

type rawTimetableStop struct {
	StopId                  string                      `xml:"id,attr,omitempty"`
	EvaId                   int64                       `xml:"eva,attr,omitempty"`
	TripLabel               rawTripLabel                `xml:"tl,omitempty"`
	Ref                     *rawTimetableStop           `xml:"ref,omitempty"`
	Arrival                 *rawEvent                   `xml:"ar,omitempty"`
	Departure               *rawEvent                   `xml:"dp,omitempty"`
	Messages                []rawMessage                `xml:"m,omitempty"`
	HistoricDelays          []rawHistoricDelay          `xml:"hd,omitempty"`
	HistoricPlatformChanges []rawHistoricPlatformChange `xml:"hpc,omitempty"`
	Connections             []rawConnection             `xml:"conn,omitempty"`
}

func parseTimetableStops(data []rawTimetableStop) []TimetableStop {
	result := make([]TimetableStop, len(data))
	for i, element := range data {
		result[i] = parseTimetableStop(element)
	}
	return result
}

func parseTimetableStop(data rawTimetableStop) TimetableStop {
	var ref *TimetableStop
	if data.Ref != nil {
		it := parseTimetableStop(*data.Ref)
		ref = &it
	}
	var arrival *Event
	if data.Arrival != nil {
		it := parseEvent(*data.Arrival)
		arrival = &it
	}
	var departure *Event
	if data.Departure != nil {
		it := parseEvent(*data.Departure)
		departure = &it
	}
	return TimetableStop{
		StopId:                  data.StopId,
		EvaId:                   data.EvaId,
		TripLabel:               parseTripLabel(data.TripLabel),
		Ref:                     ref,
		Arrival:                 arrival,
		Departure:               departure,
		Messages:                parseMessages(data.Messages),
		HistoricDelays:          parseHistoricDelays(data.HistoricDelays),
		HistoricPlatformChanges: parseHistoricPlatformChanges(data.HistoricPlatformChanges),
		Connections:             parseConnections(data.Connections),
	}
}

type rawTripLabel struct {
	Messages     []rawMessage  `xml:"m,omitempty"`
	CreatedAt    *timeShort    `xml:"ct,attr"`
	FilterFlag   rawFilterFlag `xml:"f,attr,omitempty"`
	TripType     rawTripType   `xml:"t,attr,omitempty"`
	Owner        string        `xml:"o,attr,omitempty"`
	TripCategory string        `xml:"c,attr,omitempty"`
	TripNumber   string        `xml:"n,attr,omitempty"`
}

func parseTripLabel(data rawTripLabel) TripLabel {
	return TripLabel{
		Messages:     parseMessages(data.Messages),
		CreatedAt:    data.CreatedAt.Value(),
		FilterFlag:   parseFilterFlag(data.FilterFlag),
		TripType:     parseTripType(data.TripType),
		Owner:        data.Owner,
		TripCategory: data.TripCategory,
		TripNumber:   data.TripNumber,
	}
}

type rawFilterFlag string

const (
	rawFilterFlagExternal     rawFilterFlag = "D"
	rawFilterFlagLongDistance rawFilterFlag = "F"
	rawFilterFlagRegional     rawFilterFlag = "N"
	rawFilterFlagSbahn        rawFilterFlag = "S"
	rawFilterFlagUndefined    rawFilterFlag = ""
)

func parseFilterFlag(data rawFilterFlag) FilterFlag {
	switch data {
	case rawFilterFlagExternal:
		return FilterFlagExternal
	case rawFilterFlagLongDistance:
		return FilterFlagLongDistance
	case rawFilterFlagRegional:
		return FilterFlagRegional
	case rawFilterFlagSbahn:
		return FilterFlagSbahn
	case rawFilterFlagUndefined:
		return FilterFlagUndefined
	default:
		fmt.Printf("Could not parse FilterFlag '%s'\n", data)
		return FilterFlagUnknown
	}
}

type rawTripType string

const (
	rawTripTypeP         rawTripType = "p"
	rawTripTypeE         rawTripType = "e"
	rawTripTypeZ         rawTripType = "z"
	rawTripTypeS         rawTripType = "s"
	rawTripTypeH         rawTripType = "h"
	rawTripTypeN         rawTripType = "n"
	rawTripTypeUndefined rawTripType = ""
)

func parseTripType(data rawTripType) TripType {
	switch data {
	case rawTripTypeP:
		return TripTypeP
	case rawTripTypeE:
		return TripTypeE
	case rawTripTypeZ:
		return TripTypeZ
	case rawTripTypeS:
		return TripTypeS
	case rawTripTypeH:
		return TripTypeH
	case rawTripTypeN:
		return TripTypeN
	case rawTripTypeUndefined:
		return TripTypeUndefined
	default:
		fmt.Printf("Could not parse TripType '%s'\n", data)
		return TripTypeUnknown
	}
}

type rawHistoricDelay struct {
	Timestamp *timeShort     `xml:"ts,attr"`
	Arrival   *timeShort     `xml:"ar,attr"`
	Departure *timeShort     `xml:"dp,attr"`
	Source    rawDelaySource `xml:"src,attr"`
	Code      string         `xml:"cod,attr"`
}

func parseHistoricDelays(data []rawHistoricDelay) []HistoricDelay {
	result := make([]HistoricDelay, len(data))
	for i, element := range data {
		result[i] = parseHistoricDelay(element)
	}
	return result
}

func parseHistoricDelay(data rawHistoricDelay) HistoricDelay {
	return HistoricDelay{
		Timestamp: data.Timestamp.Value(),
		Arrival:   data.Arrival.Value(),
		Departure: data.Departure.Value(),
		Source:    parseDelaySource(data.Source),
		Code:      data.Code,
	}
}

type rawDelaySource string

const (
	rawDelaySourceLeibit        rawDelaySource = "L"
	rawDelaySourceIrisAutomatic rawDelaySource = "NA"
	rawDelaySourceIrisManual    rawDelaySource = "NM"
	rawDelaySourceThirdParty    rawDelaySource = "V"
	rawDelaySourceIstpAutomatic rawDelaySource = "IA"
	rawDelaySourceIstpManual    rawDelaySource = "IM"
	rawDelaySourcePrognosis     rawDelaySource = "A"
	rawDelaySourceUndefined     rawDelaySource = ""
)

func parseDelaySource(data rawDelaySource) DelaySource {
	switch data {
	case rawDelaySourceLeibit:
		return DelaySourceLeibit
	case rawDelaySourceIrisAutomatic:
		return DelaySourceIrisAutomatic
	case rawDelaySourceIrisManual:
		return DelaySourceIrisManual
	case rawDelaySourceThirdParty:
		return DelaySourceThirdParty
	case rawDelaySourceIstpAutomatic:
		return DelaySourceIstpAutomatic
	case rawDelaySourceIstpManual:
		return DelaySourceIstpManual
	case rawDelaySourcePrognosis:
		return DelaySourcePrognosis
	case rawDelaySourceUndefined:
		return DelaySourceUndefined
	default:
		fmt.Printf("Could not parse DelaySource '%s'\n", data)
		return DelaySourceUnknown
	}
}

type rawHistoricPlatformChange struct {
	Timestamp         *timeShort `xml:"ts,attr,omitempty"`
	ArrivalPlatform   string     `xml:"ar,attr,omitempty"`
	DeparturePlatform string     `xml:"dp,attr,omitempty"`
	Cause             string     `xml:"cot,attr,omitempty"`
}

func parseHistoricPlatformChanges(data []rawHistoricPlatformChange) []HistoricPlatformChange {
	result := make([]HistoricPlatformChange, len(data))
	for i, element := range data {
		result[i] = parseHistoricPlatformChange(element)
	}
	return result
}

func parseHistoricPlatformChange(data rawHistoricPlatformChange) HistoricPlatformChange {
	return HistoricPlatformChange{
		Timestamp:         data.Timestamp.Value(),
		ArrivalPlatform:   data.ArrivalPlatform,
		DeparturePlatform: data.DeparturePlatform,
		Cause:             data.Cause,
	}
}

type rawConnection struct {
	ConnectionId     string              `xml:"id,attr,omitempty"`
	Timestamp        *timeShort          `xml:"ts,attr,omitempty"`
	EvaId            int64               `xml:"eva,attr,omitempty"`
	ConnectionStatus rawConnectionStatus `xml:"cs,attr,omitempty"`
	Ref              *rawTimetableStop   `xml:"ref,omitempty"`
	Stop             *rawTimetableStop   `xml:"s,omitempty"`
}

func parseConnections(data []rawConnection) []Connection {
	result := make([]Connection, len(data))
	for i, element := range data {
		result[i] = parseConnection(element)
	}
	return result
}

func parseConnection(data rawConnection) Connection {
	var ref TimetableStop
	if data.Ref != nil {
		ref = parseTimetableStop(*data.Ref)
	}
	var stop TimetableStop
	if data.Stop != nil {
		stop = parseTimetableStop(*data.Stop)
	}
	return Connection{
		ConnectionId:     data.ConnectionId,
		Timestamp:        data.Timestamp.Value(),
		EvaId:            data.EvaId,
		ConnectionStatus: parseConnectionStatus(data.ConnectionStatus),
		Ref:              &ref,
		Stop:             &stop,
	}
}

type rawConnectionStatus string

const (
	rawConnectionStatusWaiting     rawConnectionStatus = "w"
	rawConnectionStatusTransition  rawConnectionStatus = "n"
	rawConnectionStatusAlternative rawConnectionStatus = "a"
	rawConnectionStatusUndefined   rawConnectionStatus = ""
)

func parseConnectionStatus(data rawConnectionStatus) ConnectionStatus {
	switch data {
	case rawConnectionStatusWaiting:
		return ConnectionStatusWaiting
	case rawConnectionStatusTransition:
		return ConnectionStatusTransition
	case rawConnectionStatusAlternative:
		return ConnectionStatusAlternative
	case rawConnectionStatusUndefined:
		return ConnectionStatusUndefined
	default:
		fmt.Printf("Could not parse ConnectionStatus '%s'\n", data)
		return ConnectionStatusUnknown
	}
}

type rawEventStatus string

const (
	rawEventStatusAdded     rawEventStatus = "a"
	rawEventStatusCancelled rawEventStatus = "c"
	rawEventStatusPlanned   rawEventStatus = "p"
	rawEventStatusUndefined rawEventStatus = ""
)

func parseEventStatus(data rawEventStatus) EventStatus {
	switch data {
	case rawEventStatusAdded:
		return EventStatusAdded
	case rawEventStatusCancelled:
		return EventStatusCancelled
	case rawEventStatusPlanned:
		return EventStatusPlanned
	case rawEventStatusUndefined:
		return EventStatusUndefined
	default:
		fmt.Printf("Could not parse EventStatus '%s'\n", data)
		return EventStatusUnknown
	}
}

type rawEvent struct {
	Messages []rawMessage `xml:"m,omitempty"`

	PlannedPlatform    string          `xml:"pp,attr,omitempty"`
	PlannedTime        *timeShort      `xml:"pt,attr,omitempty"`
	PlannedPath        *bahnStringList `xml:"ppth,attr,omitempty"`
	PlannedDestination string          `xml:"pde,attr,omitempty"`
	ChangedPlatform    string          `xml:"cp,attr,omitempty"`
	ChangedTime        *timeShort      `xml:"ct,attr"`
	ChangedPath        *bahnStringList `xml:"cpth,attr,omitempty"`
	ChangedDestination string          `xml:"cde,attr,omitempty"`
	PlannedStatus      rawEventStatus  `xml:"ps,attr,omitempty"`
	ChangedStatus      rawEventStatus  `xml:"cs,attr,omitempty"`
	Hidden             int             `xml:"hi,attr,omitempty"`
	CancellationTime   string          `xml:"clt,attr,omitempty"`
	Wings              string          `xml:"wings,attr,omitempty"`
	Line               string          `xml:"l,attr,omitempty"`
	Transition         string          `xml:"tra,attr,omitempty"`
}

func parseEvent(data rawEvent) Event {
	return Event{
		Messages:           parseMessages(data.Messages),
		PlannedPlatform:    data.PlannedPlatform,
		PlannedTime:        data.PlannedTime.Value(),
		PlannedPath:        data.PlannedPath.Value(),
		PlannedDestination: data.PlannedDestination,
		ChangedPlatform:    data.ChangedPlatform,
		ChangedTime:        data.ChangedTime.Value(),
		ChangedPath:        data.ChangedPath.Value(),
		ChangedDestination: data.ChangedDestination,
		PlannedStatus:      parseEventStatus(data.PlannedStatus),
		ChangedStatus:      parseEventStatus(data.ChangedStatus),
		Hidden:             data.Hidden != 0,
		CancellationTime:   data.CancellationTime,
		Wings:              data.Wings,
		Line:               data.Line,
		Transition:         data.Transition,
	}
}
