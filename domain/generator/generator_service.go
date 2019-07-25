package generator

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"

	"github.com/BagusAK95/zaun/domain/route"
	"github.com/BagusAK95/zaun/domain/target"
	"github.com/go-resty/resty/v2"
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
func (c *GeneratorService) MatchingRoute(method string, path string) (matchedRoute route.Route, mappedParams map[string]string, err error) {
	routes := c.route.FindAll()

	for _, route := range routes {
		if method == route.Method {
			matchedPath, mappedParams := mappingParams(route.Endpoint, path)
			if len(matchedPath) != 0 {
				return route, mappedParams, nil
			}
		}
	}

	return matchedRoute, nil, errors.New("Cannot " + method + " " + path)
}

//SendToTarget : send to service target
func (c *GeneratorService) SendToTarget(routeData route.Route, httpRequest HttpRequest) (result interface{}, err error) {
	var routeTarget route.RouteTarget

	errUnmarshal := json.Unmarshal([]byte(routeData.Target), &routeTarget)
	if errUnmarshal != nil {
		return result, errUnmarshal
	}

	requestBody, _ := json.Marshal(routeTarget.Body)
	requestQuery, _ := json.Marshal(routeTarget.Query)
	requestHeader, _ := json.Marshal(routeTarget.Headers)

	stringBody := httpRequest.replaceVariable(string(requestBody))
	stringQuery := httpRequest.replaceVariable(string(requestQuery))
	stringHeader := httpRequest.replaceVariable(string(requestHeader))
	stringPath := httpRequest.replaceVariable(httpRequest.Path)

	setQuery := make(map[string]string)
	json.Unmarshal([]byte(stringQuery), &setQuery)

	setHeaders := make(map[string]string)
	json.Unmarshal([]byte(stringHeader), &setHeaders)

	targetData, errGetTarget := c.target.GetByName(routeTarget.Name)
	if errGetTarget != nil {
		return result, errGetTarget
	}

	setPath := targetData.URL + stringPath
	if len(setQuery) > 0 {
		setPath += "?"
		for key, val := range setQuery {
			setPath += key + "=" + val + "&"
		}
	}

	client := resty.New()

	switch httpRequest.Method {
	case "GET":
		client.R().
			SetHeaders(setHeaders).
			SetBody(stringBody).
			SetResult(&result).
			Get(setPath)
	case "POST":
		client.R().
			SetHeaders(setHeaders).
			SetBody(stringBody).
			SetResult(&result).
			Post(setPath)
	case "PUT":
		client.R().
			SetHeaders(setHeaders).
			SetBody(stringBody).
			SetResult(&result).
			Put(setPath)
	case "DELETE":
		client.R().
			SetHeaders(setHeaders).
			SetBody(stringBody).
			SetResult(&result).
			Delete(setPath)
	}

	return result, nil
}

func mappingParams(registeredPath, path string) (matchedPath []string, mappedParams map[string]string) {
	var regexParams = regexp.MustCompile(`\{([a-zA-Z]*)\}`) //Find {parameter} on path
	var regexSlices = regexp.MustCompile(`\/`)              //Find character / on path

	var replaceParams = regexParams.ReplaceAllString(registeredPath, `(?P<$1>[a-zA-Z0-9-:]*)`) //Replace to another regex for group params
	var replaceSlices = regexSlices.ReplaceAllString(replaceParams, `\/`)                      //Replace character / to be \/ on route

	var regexMapping = regexp.MustCompile(`^` + replaceSlices + `$`)
	matching := regexMapping.FindStringSubmatch(path)
	matchedPath = matching

	mappedParams = make(map[string]string)
	for i, param := range regexMapping.SubexpNames() {
		if i > 0 && i <= len(matching) {
			mappedParams[param] = matching[i]
		}
	}
	return
}

func (h *HttpRequest) replaceVariable(object string) string {
	regexFind := regexp.MustCompile(`\$\{([a-zA-Z.]*)\}`) //Find variable {key.value} on route
	matching := regexFind.FindAllStringSubmatch(object, -1)

	for _, obj := range matching {
		keyVal := strings.Split(obj[1], ".")

		regexReplace := regexp.MustCompile(`\$\{` + keyVal[0] + `\.` + keyVal[1] + `\}`)

		switch keyVal[0] {
		case "headers":
			object = regexReplace.ReplaceAllString(object, h.Headers[keyVal[1]]) //Replace variable with real value
		case "params":
			object = regexReplace.ReplaceAllString(object, h.Params[keyVal[1]]) //Replace variable with real value
		case "query":
			object = regexReplace.ReplaceAllString(object, h.Query[keyVal[1]]) //Replace variable with real value
		case "body":
			object = regexReplace.ReplaceAllString(object, h.Body.(map[string]interface{})[keyVal[1]].(string)) //Replace variable with real value
		}
	}

	return object
}
