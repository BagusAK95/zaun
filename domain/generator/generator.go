package generator

//HttpRequest :
type HttpRequest struct {
	Path    string
	Headers map[string]string
	Params  map[string]string
	Method  string
	Query   map[string]string
	Body    interface{}
}
