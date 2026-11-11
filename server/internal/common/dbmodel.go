package common

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey;default:(-)"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}
