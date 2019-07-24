package generator

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

//GeneratorController : set route controller
type GeneratorController struct {
	service GeneratorService
}

//NewController : instantiate controller
func NewController(service GeneratorService) GeneratorController {
	return GeneratorController{service}
}

//Process : process
func (controller *GeneratorController) Process(c echo.Context) error {
	path := c.Request().URL.Path
	method := c.Request().Method
	query := c.QueryParams()
	body := new(interface{})
	c.Bind(body)

	route, params, err := controller.service.MatchingRoute(path)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	request := map[string]interface{}{
		"path":   path,
		"params": params,
		"method": method,
		"query":  query,
		"body":   body,
	}

	result := controller.service.SendToTarget(route, request)

	return c.JSON(http.StatusOK, result)
}
