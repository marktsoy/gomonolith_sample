package models

const (
	PriorityLow = iota
	PriorityMedium
	PriorityHigh
)

const (
	BundleStatusCreated = iota
	BundleStatusIsProcessing
	BundleStatusProcessed
)

// Bundle ...
type Bundle struct {
	ID       int `json:"id"`
	Priority int `json:"priority"`
	Size     int `json:"size"`
	Status   int `json:"status"`
}
