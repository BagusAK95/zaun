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

	headers := make(map[string]string)
	for key, val := range c.Request().Header {
		headers[key] = val[0]
	}

	query := make(map[string]string)
	for key, val := range c.QueryParams() {
		query[key] = val[0]
	}

	matchedRoute, mappedParams, errMatching := controller.service.MatchingRoute(method, path)
	if errMatching != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": errMatching.Error(),
		})
	}

	httpRequest := HttpRequest{
		Path:    path,
		Headers: headers,
		Params:  mappedParams,
		Method:  method,
		Query:   query,
		Body:    *body,
	}

	httpResponse, errSendToTarget := controller.service.SendToTarget(matchedRoute, httpRequest)
	if errSendToTarget != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": errSendToTarget.Error(),
		})
	}

	return c.JSON(http.StatusOK, httpResponse)
}
