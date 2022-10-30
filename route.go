package mooxy

import "net/http"

type Route struct {
    Path string
    methods []string
}

// set allowed methods for route
func (r *Route) Methods(methods []string) *Route {
    r.methods = methods
    return r
}

func NewRoute(path string) *Route {
    return &Route { Path: path, methods: []string{http.MethodGet, http.MethodPost} }
}
