package generator

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/BagusAK95/zaun/domain/route"
	"github.com/BagusAK95/zaun/domain/target"
)

//GeneratorService : set route service
type GeneratorService struct {
	route  route.RouteService
	target target.TargetService
}

//NewService : instantiate service
func NewService(route route.RouteService, target target.TargetService) GeneratorService {
	return GeneratorService{route, target}
}

//MatchingRoute : mathcing route
func (c *GeneratorService) MatchingRoute(path string) (matchedRoute route.Route, mappedParams map[string]string, err error) {
	routes := c.route.FindAll()

	for _, route := range routes {
		pathMatch, paramsMap := mappingParams(route.Endpoint, path)
		if len(pathMatch) != 0 {
			return route, paramsMap, nil
		}
	}

	return matchedRoute, nil, errors.New("Route not found")
}

//SendToTarget : send to service target
func (c *GeneratorService) SendToTarget(routeData route.Route, requestData map[string]interface{}) (result interface{}) {
	var routeTarget map[string]interface{}

	errUnmarshalRouteTarget := json.Unmarshal([]byte(routeData.Target), &routeTarget)
	if errUnmarshalRouteTarget != nil {
		return nil
	}

	targetData, err := c.target.GetByName(routeTarget["name"].(string))
	if err != nil {
		return nil
	}

	reqBody, err := json.Marshal(map[string]string{
		"username": "BagusAK95",
		"password": "q1w2e3r4t5",
	})

	resp, err := http.Post(targetData.URL+requestData["path"].(string), "application/json", bytes.NewBuffer(reqBody))
	body, err := ioutil.ReadAll(resp.Body)

	errUnmarshalRespBody := json.Unmarshal(body, &result)
	if errUnmarshalRespBody != nil {
		return nil
	}

	return result
}

func mappingParams(registeredPath, matchingPath string) (pathMatch []string, paramsMap map[string]string) {
	var regexParams = regexp.MustCompile(`\{([a-zA-Z]*)\}`) //Find {parameter} on route
	var regexSlices = regexp.MustCompile(`\/`)              //Find character / on path

	var replaceParams = regexParams.ReplaceAllString(registeredPath, `(?P<$1>[a-zA-Z0-9-]*)`) //Replace to another regex
	var replaceSlices = regexSlices.ReplaceAllString(replaceParams, `\/`)                     //Replace character / to be \/ on route
	var strRegex = `^` + replaceSlices + `$`

	var compRegEx = regexp.MustCompile(strRegex)
	match := compRegEx.FindStringSubmatch(matchingPath)
	pathMatch = match

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return
}
