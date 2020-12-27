package sqlstore

import (
	"database/sql"

	"github.com/marktsoy/gomonolith_sample/internal/app/store"
)

// Store ...
type Store struct {
	db                *sql.DB
	bundleRepository  *BundleRepository
	messageRepository *MessageRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) DB() *sql.DB {
	return s.db
}

func (s *Store) Bundle() store.BundleRepository {
	if s.bundleRepository == nil {
		s.bundleRepository = &BundleRepository{
			store: s,
		}
	}
	return s.bundleRepository
}

func (s *Store) Message() store.MessageRepository {
	if s.messageRepository == nil {
		s.messageRepository = &MessageRepository{
			store: s,
		}
	}
	return s.messageRepository
}
