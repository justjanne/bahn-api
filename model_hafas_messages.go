package bahn

type HafasMessage struct {
	Id       string               `json:"id,omitempty"yaml:"id,omitempty"`
	Priority HafasMessagePriority `json:"priority,omitempty"yaml:"priority,omitempty"`
	From     string               `json:"from,omitempty"yaml:"from,omitempty"`
	To       string               `json:"to,omitempty"yaml:"to,omitempty"`
	Subject  string               `json:"subject,omitempty"yaml:"subject,omitempty"`
	Content  string               `json:"content,omitempty"yaml:"content,omitempty"`
}

type HafasMessagePriority string

const (
	HafasMessagePriorityHigh   HafasMessagePriority = "HIGH"
	HafasMessagePriorityMiddle HafasMessagePriority = "MEDIUM"
	HafasMessagePriorityLow    HafasMessagePriority = "LOW"
)
