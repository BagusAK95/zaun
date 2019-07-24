package target

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Target : structure table targets
type Target struct {
	ID        *uuid.UUID `gorm:"primary_key;type:varchar(36)"`
	Name      string     `gorm:"type:varchar(100)"`
	URL       string     `gorm:"type:varchar(255)"`
	Token     string     `gorm:"type:varchar(36)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate : set id before create
func (r *Target) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}
