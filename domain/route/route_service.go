package route

//RouteService : set route service
type RouteService struct {
	repository RouteRepo
}

//NewService : instantiate service
func NewService(repo RouteRepo) RouteService {
	return RouteService{repo}
}

//ListAll : get list route
func (c *RouteService) ListAll() []Route {
	return c.repository.ListAll()
}
