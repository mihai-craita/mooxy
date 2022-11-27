package mooxy

import "net/http"

type Route struct {
    Path string
    methods []HTTPMethod
}

// set allowed methods for route
func (r *Route) Methods(methods ...HTTPMethod) *Route {
    r.methods = methods
    return r
}

func NewRoute(path string) *Route {
    return &Route { Path: path, methods: []HTTPMethod{http.MethodGet} }
}
