package main

import "time"

type Timetable struct {
	Station  string          `json:"station,omitempty"yaml:"station,omitempty"`
	EvaId    int64           `json:"eva_id,omitempty"yaml:"eva_id,omitempty"`
	Stops    []TimetableStop `json:"stops,omitempty"yaml:"stops,omitempty"`
	Messages []Message       `json:"messages,omitempty"yaml:"messages,omitempty"`
}

type Message struct {
	MessageId           string               `json:"message_id,omitempty"yaml:"message_id,omitempty"`
	Type                MessageType          `json:"type,omitempty"yaml:"type,omitempty"`
	From                *time.Time           `json:"from,omitempty"yaml:"from,omitempty"`
	To                  *time.Time           `json:"to,omitempty"yaml:"to,omitempty"`
	Code                int                  `json:"code,omitempty"yaml:"code,omitempty"`
	InternalText        string               `json:"internal_text,omitempty"yaml:"internal_text,omitempty"`
	ExternalText        string               `json:"external_text,omitempty"yaml:"external_text,omitempty"`
	Category            string               `json:"category,omitempty"yaml:"category,omitempty"`
	ExternalCategory    string               `json:"external_category,omitempty"yaml:"external_category,omitempty"`
	Timestamp           *time.Time           `json:"timestamp,omitempty"yaml:"timestamp,omitempty"`
	Priority            Priority             `json:"priority,omitempty"yaml:"priority,omitempty"`
	Owner               string               `json:"owner,omitempty"yaml:"owner,omitempty"`
	ExternalLink        string               `json:"external_link,omitempty"yaml:"external_link,omitempty"`
	Deleted             bool                 `json:"deleted,omitempty"yaml:"deleted,omitempty"`
	DistributorMessages []DistributorMessage `json:"distributor_messages,omitempty"yaml:"distributor_messages,omitempty"`
}

type MessageType string

const (
	MessageTypeHafasInformationManager MessageType = "HAFAS_INFORMATION_MANAGER"
	MessageTypeQualityChange           MessageType = "QUALITY_CHANGE"
	MessageTypeFreeText                MessageType = "FREE_TEXT"
	MessageTypeCauseOfDelay            MessageType = "CAUSE_OF_DELAY"
	MessageTypeIbis                    MessageType = "IBIS"
	MessageTypeIbisUnassigned          MessageType = "IBIS_UNASSIGNED"
	MessageTypeDisruption              MessageType = "DISRUPTION"
	MessageTypeConnection              MessageType = "CONNECTION"
	MessageTypeUnknown                 MessageType = "UNKNOWN"
	MessageTypeUndefined               MessageType = ""
)

type Priority string

const (
	PriorityHigh      Priority = "HIGH"
	PriorityMedium    Priority = "MEDIUM"
	PriorityLow       Priority = "LOW"
	PriorityDone      Priority = "DONE"
	PriorityUnknown   Priority = "UNKNOWN"
	PriorityUndefined Priority = ""
)

type DistributorType string

const (
	DistributorTypeCity         DistributorType = "CITY"
	DistributorTypeRegion       DistributorType = "REGION"
	DistributorTypeLongDistance DistributorType = "LONG_DISTANCE"
	DistributorTypeOther        DistributorType = "OTHER"
	DistributorTypeUnknown      DistributorType = "UNKNOWN"
	DistributorTypeUndefined    DistributorType = ""
)

type DistributorMessage struct {
	DistributorType DistributorType `json:"distributor_type,omitempty"yaml:"distributor_type,omitempty"`
	DistributorName string          `json:"distributor_name,omitempty"yaml:"distributor_name,omitempty"`
	InternalText    string          `json:"internal_text,omitempty"yaml:"internal_text,omitempty"`
	Timestamp       *time.Time      `json:"timestamp,omitempty"yaml:"timestamp,omitempty"`
}

type TimetableStop struct {
	StopId                  string                   `json:"stop_id,omitempty"yaml:"stop_id,omitempty"`
	EvaId                   int64                    `json:"eva_id,omitempty"yaml:"eva_id,omitempty"`
	TripLabel               TripLabel                `json:"trip_label,omitempty"yaml:"trip_label,omitempty"`
	Ref                     *TimetableStop           `json:"ref,omitempty"yaml:"ref,omitempty"`
	Arrival                 *Event                   `json:"arrival,omitempty"yaml:"arrival,omitempty"`
	Departure               *Event                   `json:"departure,omitempty"yaml:"departure,omitempty"`
	Messages                []Message                `json:"messages,omitempty"yaml:"messages,omitempty"`
	HistoricDelays          []HistoricDelay          `json:"historic_delay,omitempty"yaml:"historic_delay,omitempty"`
	HistoricPlatformChanges []HistoricPlatformChange `json:"historic_platform_changes,omitempty"yaml:"historic_platform_changes,omitempty"`
	Connections             []Connection             `json:"connections,omitempty"yaml:"connections,omitempty"`
}

type TripLabel struct {
	Messages     []Message  `json:"messages,omitempty"yaml:"messages,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"yaml:"created_at,omitempty"`
	FilterFlag   FilterFlag `json:"filter_flag,omitempty"yaml:"filter_flag,omitempty"`
	TripType     TripType   `json:"trip_type,omitempty"yaml:"trip_type,omitempty"`
	Owner        string     `json:"owner,omitempty"yaml:"owner,omitempty"`
	TripCategory string     `json:"trip_category,omitempty"yaml:"trip_category,omitempty"`
	TripNumber   string     `json:"trip_number,omitempty"yaml:"trip_number,omitempty"`
}

type FilterFlag string

const (
	FilterFlagExternal     FilterFlag = "EXTERNAL"
	FilterFlagLongDistance FilterFlag = "LONG_DISTANCE"
	FilterFlagRegional     FilterFlag = "REGIONAl"
	FilterFlagSbahn        FilterFlag = "SBAHN"
	FilterFlagUnknown      FilterFlag = "UNKNOWN"
	FilterFlagUndefined    FilterFlag = ""
)

type TripType string

const (
	TripTypeP         TripType = "P"
	TripTypeE         TripType = "E"
	TripTypeZ         TripType = "Z"
	TripTypeS         TripType = "S"
	TripTypeH         TripType = "H"
	TripTypeN         TripType = "N"
	TripTypeUnknown   TripType = "UNKNOWN"
	TripTypeUndefined TripType = ""
)

type HistoricDelay struct {
	Timestamp *time.Time  `json:"timestamp,omitempty"yaml:"timestamp,omitempty"`
	Arrival   *time.Time  `json:"arrival,omitempty"yaml:"arrival,omitempty"`
	Departure *time.Time  `json:"departure,omitempty"yaml:"departure,omitempty"`
	Source    DelaySource `json:"source,omitempty"yaml:"source,omitempty"`
	Code      string      `json:"code,omitempty"yaml:"code,omitempty"`
}

type DelaySource string

const (
	DelaySourceLeibit        DelaySource = "LEIBIT"
	DelaySourceIrisAutomatic DelaySource = "IRIS_AUTOMATIC"
	DelaySourceIrisManual    DelaySource = "IRIS_MANUAL"
	DelaySourceThirdParty    DelaySource = "THIRD_PARTY"
	DelaySourceIstpAutomatic DelaySource = "ISTP_AUTOMATIC"
	DelaySourceIstpManual    DelaySource = "ISTP_MANUAL"
	DelaySourcePrognosis     DelaySource = "PROGNOSIS"
	DelaySourceUnknown       DelaySource = "UNKNOWN"
	DelaySourceUndefined     DelaySource = ""
)

type HistoricPlatformChange struct {
	Timestamp         *time.Time `json:"timestamp,omitempty"yaml:"timestamp,omitempty"`
	ArrivalPlatform   string     `json:"arrival_platform,omitempty"yaml:"arrival_platform,omitempty"`
	DeparturePlatform string     `json:"departure_platform,omitempty"yaml:"departure_platform,omitempty"`
	Cause             string     `json:"cause,omitempty"yaml:"cause,omitempty"`
}
type Connection struct {
	ConnectionId     string           `json:"connection_id,omitempty"yaml:"connection_id,omitempty"`
	Timestamp        *time.Time       `json:"timestamp,omitempty"yaml:"timestamp,omitempty"`
	EvaId            int64            `json:"eva_id,omitempty"yaml:"eva_id,omitempty"`
	ConnectionStatus ConnectionStatus `json:"connection_status,omitempty"yaml:"connection_status,omitempty"`
	Ref              *TimetableStop   `json:"ref,omitempty"yaml:"ref,omitempty"`
	Stop             *TimetableStop   `json:"stop,omitempty"yaml:"stop,omitempty"`
}

type ConnectionStatus string

const (
	ConnectionStatusWaiting     ConnectionStatus = "WAITING"
	ConnectionStatusTransition  ConnectionStatus = "TRANSITION"
	ConnectionStatusAlternative ConnectionStatus = "ALTERNATIVE"
	ConnectionStatusUnknown     ConnectionStatus = "UNKNOWN"
	ConnectionStatusUndefined   ConnectionStatus = ""
)

type EventStatus string

const (
	EventStatusAdded     EventStatus = "ADDED"
	EventStatusCancelled EventStatus = "CANCELLED"
	EventStatusPlanned   EventStatus = "PLANNED"
	EventStatusUnknown   EventStatus = "UNKNOWN"
	EventStatusUndefined EventStatus = ""
)

type Event struct {
	Messages []Message `json:"messages,omitempty"yaml:"messages,omitempty"`

	PlannedPlatform    string      `json:"planned_platform,omitempty"yaml:"planned_platform,omitempty"`
	PlannedTime        *time.Time  `json:"planned_time,omitempty"yaml:"planned_time,omitempty"`
	PlannedPath        []string    `json:"planned_path,omitempty"yaml:"planned_path,omitempty"`
	PlannedDestination string      `json:"planned_destination,omitempty"yaml:"planned_destination,omitempty"`
	ChangedPlatform    string      `json:"changed_platform,omitempty"yaml:"changed_platform,omitempty"`
	ChangedTime        *time.Time  `json:"changed_time,omitempty"yaml:"changed_time,omitempty"`
	ChangedPath        []string    `json:"changed_path,omitempty"yaml:"changed_path,omitempty"`
	ChangedDestination string      `json:"changed_destination,omitempty"yaml:"changed_destination,omitempty"`
	PlannedStatus      EventStatus `json:"planned_status,omitempty"yaml:"planned_status,omitempty"`
	ChangedStatus      EventStatus `json:"changed_status,omitempty"yaml:"changed_status,omitempty"`
	Hidden             bool        `json:"hidden,omitempty"yaml:"hidden,omitempty"`
	CancellationTime   string      `json:"cancellation_time,omitempty"yaml:"cancellation_time,omitempty"`
	Wings              string      `json:"wings,omitempty"yaml:"wings,omitempty"`
	Line               string      `json:"line,omitempty"yaml:"line,omitempty"`
	Transition         string      `json:"transition,omitempty"yaml:"transition,omitempty"`
}
