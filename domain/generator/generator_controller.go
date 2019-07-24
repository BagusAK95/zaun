package generator

import (
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
	body := new(interface{})
	c.Bind(body)

	header := make(map[string]string)
	for key, val := range c.Request().Header {
		header[key] = val[0]
	}

	query := make(map[string]string)
	for key, val := range c.QueryParams() {
		query[key] = val[0]
	}

	route, params, errMatching := controller.service.MatchingRoute(path)
	if errMatching != nil {
		return errMatching
	}

	request := map[string]interface{}{
		"path":   path,
		"header": header,
		"params": params,
		"method": method,
		"query":  query,
		"body":   *body,
	}

	result, errSendData := controller.service.SendToTarget(route, request)
	if errSendData != nil {
		return errSendData
	}

	return c.JSON(http.StatusOK, result)
}
