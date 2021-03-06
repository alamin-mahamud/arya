package model

import (
	"time"

	"github.com/go-pg/pg/orm"
)

// Base contains common fields for all tables
type Base struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty" pg:",soft_delete"`
}

// BeforeInsert hooks into insert operations, setting createdAt and updatedAt to current time
func (b *Base) BeforeInsert(_ orm.DB) error {
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now
	return nil
}

// BeforeUpdate hooks into update operations, setting updatedAt to current time
func (b *Base) BeforeUpdate(_ orm.DB) error {
	b.UpdatedAt = time.Now()
	return nil
}
