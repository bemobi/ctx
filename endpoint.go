package ctx

var endpoints = []*Endpoint{}

type Endpoint struct {
	Public   bool
	Path     string
	Handlers map[string]ContextHandler
}

func Register(endpoint *Endpoint) {
	endpoints = append(endpoints, endpoint)
}
