package store

import "github.com/marktsoy/gomonolith_sample/internal/app/models"

// Store interface
type Store interface {
	Bundle() BundleRepository
	Message() MessageRepository
}

// BundleRepository ..
type BundleRepository interface {
	Create(*models.Bundle) error
	Update(*models.Bundle) error
}

// MessageRepository ..
type MessageRepository interface {
	Create(*models.Message) error
	Update(*models.Message) error
}
