package models

const (
	MessageStatusCreated = iota
	MessageStatusProcessed
)

// Message ...
type Message struct {
	ID       int    `json:"id"`
	Content  string `json:"string"`
	Status   int    `json:"status"`
	Priority int    `json:"priority"`
	BundleID int    `json:"bundle_id"`
}

func (m *Message) Creating() {
	if m.Status > MessageStatusProcessed || m.Status < MessageStatusCreated {
		m.Status = 0
	}
	if m.Priority > PriorityHigh || m.Priority < PriorityLow {
		m.Priority = 0
	}
}
