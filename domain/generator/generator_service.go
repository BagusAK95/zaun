package generator

//GeneratorService : set route service
type GeneratorService struct{}

//NewService : instantiate service
func NewService() GeneratorService {
	return GeneratorService{}
}

//RequestToService : request to service
func (c *GeneratorService) RequestToService() interface{} {

	return "Success"
}
