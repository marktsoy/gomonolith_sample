package sqlstore

import (
	"github.com/marktsoy/gomonolith_sample/internal/app/models"
	"github.com/marktsoy/gomonolith_sample/internal/app/store"
)

// MessageRepository ..
type MessageRepository struct {
	store *Store
}

// Create ...
func (rep *MessageRepository) Create(msg *models.Message) error {
	db := rep.store.db
	msg.Creating()
	err := db.QueryRow(
		"INSERT INTO messages(content,priority,status,bundle_id) VALUES ($1,$2,$3,$4) RETURNING id,priority,status",
		msg.Content, msg.Priority, msg.Status, msg.BundleID,
	).Scan(&msg.ID, &msg.Priority, &msg.Status)

	if err != nil {
		return err
	}

	return nil
}

// Update ...
func (rep *MessageRepository) Update(msg *models.Message) error {
	db := rep.store.db

	sqlStatement := "UPDATE messages  SET content=$2 , priority=$3,status=$4 where messages.id = $1 "
	res, err := db.Exec(sqlStatement, msg.ID, msg.Content, msg.Priority, msg.Status)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return store.ErrRecordNotFound
	}
	return nil
}
