package route

import (
	"github.com/jinzhu/gorm"
)

//RouteRepo : set route repository
type RouteRepo struct {
	db *gorm.DB
}

//NewRepo : instantiate repository
func NewRepo(db *gorm.DB) (RouteRepo, error) {
	return RouteRepo{db}, nil
}

//FindAll : get list route
func (r *RouteRepo) FindAll() (routes []Route) {
	r.db.Find(&routes)

	return routes
}
