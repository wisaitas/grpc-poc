package postgres

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuidv7()"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:now()"`
}
