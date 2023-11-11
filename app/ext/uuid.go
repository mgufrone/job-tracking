package ext

import "github.com/google/uuid"

type UUID struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid"`
}
