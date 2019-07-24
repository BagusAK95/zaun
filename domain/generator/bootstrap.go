package generator

import (
	"github.com/BagusAK95/zaun/common"
	"github.com/BagusAK95/zaun/config"
	"github.com/BagusAK95/zaun/domain/route"
	"github.com/BagusAK95/zaun/domain/target"
	"github.com/jinzhu/gorm"
)

//Bootstrap : user domain bootstrap
type Bootstrap struct {
	Controller GeneratorController
	Service    GeneratorService
}

//Init : user bootstrap instantiate
func Init(db *gorm.DB, cfg *config.Configuration, cache common.Cache) *Bootstrap {
	routeRepo, err := route.NewRepo(db)
	if err != nil {
		return nil
	}
	targetRepo, err := target.NewRepo(db)
	if err != nil {
		return nil
	}

	route := route.NewService(routeRepo, cache)
	target := target.NewService(targetRepo, cache)

	service := NewService(route, target)
	controller := NewController(service)

	return &Bootstrap{
		Controller: controller,
		Service:    service,
	}
}
