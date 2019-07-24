package target

import (
	"github.com/jinzhu/gorm"
)

//TargetRepo : set target repository
type TargetRepo struct {
	db *gorm.DB
}

//NewRepo : instantiate repository
func NewRepo(db *gorm.DB) (TargetRepo, error) {
	return TargetRepo{db}, nil
}

//FindAll : get list target
func (r *TargetRepo) FindAll() (targets []Target) {
	r.db.Find(&targets)

	return targets
}
