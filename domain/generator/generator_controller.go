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

//Process : process request
func (controller *GeneratorController) Process(c echo.Context) error {
	path := c.Request().URL.Path
	method := c.Request().Method
	query := c.QueryParams()
	body := new(interface{})
	c.Bind(body)

	data := map[string]interface{}{
		"path":   path,
		"method": method,
		"query":  query,
		"body":   body,
	}

	return c.JSON(http.StatusOK, data)
}
