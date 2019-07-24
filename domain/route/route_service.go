package route

import (
	"encoding/json"

	"github.com/BagusAK95/zaun/common"
)

//RouteService : set route service
type RouteService struct {
	repository RouteRepo
	cache      common.Cache
}

//NewService : instantiate service
func NewService(repo RouteRepo, cache common.Cache) RouteService {
	return RouteService{repo, cache}
}

//ListAll : get list route
func (c *RouteService) FindAll() []Route {
	cache, err := c.cache.Get("zaunRoutes")
	if err == nil {
		var dat []Route

		err := json.Unmarshal([]byte(cache), &dat)
		if err == nil {
			return dat
		}
	}

	routes := c.repository.FindAll()
	if routes != nil {
		c.cache.Set("zaunRoutes", routes)
	}

	return routes
}
