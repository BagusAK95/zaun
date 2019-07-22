package route

import (
	"net/http"

	"github.com/labstack/echo"
)

//RouteController : set route controller
type RouteController struct {
	service RouteService
}

//NewController : instantiate controller
func NewController(service RouteService) RouteController {
	return RouteController{service}
}

//ListAll : get list all route
func (controller *RouteController) ListAll(c echo.Context) error {
	result := controller.service.ListAll()
	return c.JSON(http.StatusOK, result)
}
