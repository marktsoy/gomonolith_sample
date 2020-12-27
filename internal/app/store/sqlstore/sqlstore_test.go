package sqlstore_test

import (
	"database/sql"
	"strconv"
	"testing"

	"github.com/marktsoy/gomonolith_sample/internal/app/models"
	"github.com/marktsoy/gomonolith_sample/internal/app/store"
	"github.com/marktsoy/gomonolith_sample/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestStore_CreateBundle(t *testing.T) {
	s, clear := sqlstore.TestStore(t, "postgres://localhost/sandbox?sslmode=disable")

	defer func() {
		clear("bundles")
	}()

	b := &models.Bundle{
		Size:     3,
		Priority: models.PriorityLow,
	}
	rep := s.Bundle()
	assert.Implements(t, (*store.BundleRepository)(nil), rep)
	assert.IsType(t, &sql.DB{}, s.DB())
	err := s.Bundle().Create(b)
	assert.NoError(t, err)
	assert.Equal(t, 0, b.Status)

}

func generateMessage(t *testing.T, b *models.Bundle) []*models.Message {
	t.Helper()
	q := make([]*models.Message, 0)
	for i := 0; i < 10; i++ {
		m := &models.Message{
			Content:  "Sample Content " + strconv.Itoa(i),
			Priority: i,
			BundleID: b.ID,
		}
		q = append(q, m)
	}
	return q
}

func TestStore_CreateMessage(t *testing.T) {
	s, clear := sqlstore.TestStore(t, "postgres://localhost/sandbox?sslmode=disable")

	defer clear("messages", "bundles")
	b := &models.Bundle{
		Size:     3,
		Status:   models.BundleStatusCreated,
		Priority: models.PriorityLow,
	}
	err := s.Bundle().Create(b)
	assert.NoError(t, err)

	msgs := generateMessage(t, b)

	rep := s.Message()
	assert.Implements(t, (*store.MessageRepository)(nil), rep)
	for _, m := range msgs {
		assert.IsType(t, &models.Message{}, m)
		err := rep.Create(m)
		assert.NoError(t, err)
		assert.NotNil(t, m.ID)
		assert.Equal(t, m.Status, 0)
		assert.LessOrEqual(t, m.Priority, models.PriorityHigh)
		assert.GreaterOrEqual(t, m.Priority, models.PriorityLow)
	}
}
