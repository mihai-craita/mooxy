package mooxy

type Route struct {
    Path string
    methods HTTPMethods
}

// set allowed methods for route
func (r *Route) Methods(methods ...HTTPMethod) *Route {
	for _, method := range methods {
		r.methods.Add(method)
	}
    return r
}

func NewRoute(path string) *Route {
    return &Route { Path: path, methods: NewHttpMethods() }
}
