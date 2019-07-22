package http

import (
	"github.com/BagusAK95/zaun/domain/generator"
	"github.com/labstack/echo"
)

//Handler : setup handler
type Handler struct {
	generator generator.Bootstrap
}

//New : register routes
func New(generator generator.Bootstrap) *echo.Echo {
	e := echo.New()

	e.GET("/*", generator.Controller.Process)
	e.POST("/*", generator.Controller.Process)
	e.PUT("/*", generator.Controller.Process)
	e.DELETE("/*", generator.Controller.Process)

	return e
}
