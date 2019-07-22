package route

import (
	"github.com/BagusAK95/zaun/config"
	"github.com/jinzhu/gorm"
)

//Bootstrap : user domain bootstrap
type Bootstrap struct {
	Controller RouteController
	Repository RouteRepo
	Service    RouteService
}

//Init : user bootstrap instantiate
func Init(db *gorm.DB, cfg *config.Configuration) *Bootstrap {
	repository, err := NewRepo(db)
	if err != nil {
		return nil
	}
	service := NewService(repository)
	controller := NewController(service)

	return &Bootstrap{
		Controller: controller,
		Repository: repository,
		Service:    service,
	}
}
