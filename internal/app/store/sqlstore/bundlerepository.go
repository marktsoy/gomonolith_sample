package sqlstore

import (
	"github.com/marktsoy/gomonolith_sample/internal/app/models"
	"github.com/marktsoy/gomonolith_sample/internal/app/store"
)

// BundleRepository ..
type BundleRepository struct {
	store *Store
}

// Create ...
func (rep *BundleRepository) Create(bundle *models.Bundle) error {
	db := rep.store.db

	err := db.QueryRow(
		"INSERT INTO bundles(priority,size,status) VALUES ($1,$2,$3) RETURNING id,status",
		bundle.Priority, bundle.Size, bundle.Status,
	).Scan(&bundle.ID, &bundle.Status)

	if err != nil {
		return err
	}

	return nil
}

// Update ...
func (rep *BundleRepository) Update(bundle *models.Bundle) error {
	db := rep.store.db

	sqlStatement := "UPDATE bundles SET priority=$2 , size=$3, status=$4 where decks.id = $1 "
	res, err := db.Exec(sqlStatement, bundle.ID, bundle.Priority, bundle.Size, bundle.Status)
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
