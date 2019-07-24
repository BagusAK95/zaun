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
func (c *GeneratorService) SendToTarget(routeData route.Route, requestData map[string]interface{}) (result interface{}, err error) {
	var routeTarget route.RouteTarget

	errUnmarshal := json.Unmarshal([]byte(routeData.Target), &routeTarget)
	if errUnmarshal != nil {
		return result, errUnmarshal
	}

	requestBody, _ := json.Marshal(routeTarget.Body)
	requestQuery, _ := json.Marshal(routeTarget.Query)
	requestHeader, _ := json.Marshal(routeTarget.Headers)

	stringBody := replaceVariable(string(requestBody), requestData)
	stringQuery := replaceVariable(string(requestQuery), requestData)
	stringHeader := replaceVariable(string(requestHeader), requestData)
	stringPath := replaceVariable(requestData["path"].(string), requestData)

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
	switch requestData["method"] {
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

func mappingParams(registeredPath, matchingPath string) (pathMatch []string, paramsMap map[string]string) {
	var regexParams = regexp.MustCompile(`\{([a-zA-Z]*)\}`) //Find {parameter} on route
	var regexSlices = regexp.MustCompile(`\/`)              //Find character / on path

	var replaceParams = regexParams.ReplaceAllString(registeredPath, `(?P<$1>[a-zA-Z0-9-:]*)`) //Replace to another regex
	var replaceSlices = regexSlices.ReplaceAllString(replaceParams, `\/`)                      //Replace character / to be \/ on route
	var strRegex = `^` + replaceSlices + `$`                                                   //Set start and end

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

func replaceVariable(object string, requestData map[string]interface{}) string {
	findRegEx := regexp.MustCompile(`\$\{([a-zA-Z.]*)\}`) //Find variable {key.value} on route
	match := findRegEx.FindAllStringSubmatch(object, -1)
	for _, obj := range match {
		keyVal := strings.Split(obj[1], ".")

		replaceRegEx := regexp.MustCompile(`\$\{` + keyVal[0] + `\.` + keyVal[1] + `\}`)

		switch keyVal[0] {
		case "body":
			object = replaceRegEx.ReplaceAllString(object, requestData[keyVal[0]].(map[string]interface{})[keyVal[1]].(string)) //Replace variable with real value
		default:
			object = replaceRegEx.ReplaceAllString(object, requestData[keyVal[0]].(map[string]string)[keyVal[1]]) //Replace variable with real value
		}
	}

	return object
}
