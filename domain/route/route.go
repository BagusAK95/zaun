package route

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Route : structure table routes
type Route struct {
	ID          *uuid.UUID `gorm:"primary_key;type:varchar(36)"`
	Group       string     `gorm:"type:varchar(100)"`
	Title       string     `gorm:"type:varchar(255)"`
	Description string     `gorm:"type:varchar(500)"`
	Method      string     `sql:"type:enum('GET','POST','PUT','DELETE')"`
	Endpoint    string     `gorm:"type:varchar(255)"`
	Headers     string     `gorm:"type:mediumtext(0)"`
	Query       string     `gorm:"type:mediumtext(0)"`
	Body        string     `gorm:"type:mediumtext(0)"`
	Response    string     `gorm:"type:text(0)"`
	Middleware  string     `gorm:"type:mediumtext(0)"`
	Service     string     `gorm:"type:mediumtext(0)"`
	Status      string     `sql:"type:ENUM('active','deactive')"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// BeforeCreate : set id before create
func (r *Route) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}
