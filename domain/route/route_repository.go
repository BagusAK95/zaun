package route

import "github.com/jinzhu/gorm"

//RouteRepo : set route repository
type RouteRepo struct {
	db *gorm.DB
}

//NewRepo : instantiate repository
func NewRepo(db *gorm.DB) (RouteRepo, error) {
	return RouteRepo{db}, nil
}

//ListAll : get list route
func (r *RouteRepo) ListAll() (routes []Route) {
	r.db.Find(&routes)

	return routes
}
