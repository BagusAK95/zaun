package generator

import (
	"github.com/BagusAK95/zaun/config"
	"github.com/jinzhu/gorm"
)

//Bootstrap : user domain bootstrap
type Bootstrap struct {
	Controller GeneratorController
	Service    GeneratorService
}

//Init : user bootstrap instantiate
func Init(db *gorm.DB, cfg *config.Configuration) *Bootstrap {
	service := NewService()
	controller := NewController(service)

	return &Bootstrap{
		Controller: controller,
		Service:    service,
	}
}
