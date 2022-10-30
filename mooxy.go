package mooxy

import (
	"net/http"
	"strings"
)

type Router struct {
    NextRoutes map[string]Router
    Handler http.Handler
}

func (r *Router) Handle (route *Route, handler http.Handler) {
    var path = route.getPath()
    var pathParts = getPathParts(path)
    var lastElementIndex = len(pathParts) - 1

    currentMatrix := r.NextRoutes
    for index, p := range pathParts{
        nodeToBeAdded := Router{NextRoutes: make(map[string]Router)}
        if (index == lastElementIndex) {
            nodeToBeAdded.Handler = handler
        }

        currentMatrix[p] = nodeToBeAdded
        currentMatrix = nodeToBeAdded.NextRoutes
    }
}

func (router *Router) GetServer(w http.ResponseWriter, r *http.Request) {
    // here we will match to the correct handler and ServeHTTP
    var path = r.URL.Path
    var pathParts = getPathParts(path)
    var currentMatrix = router.NextRoutes
    for _, p := range pathParts {
        if len(currentMatrix[p].NextRoutes) > 0 {
            currentMatrix = currentMatrix[p].NextRoutes
        } else {
            currentMatrix[p].Handler.ServeHTTP(w, r)
        }
    }
}

func getPathParts(path string) []string {
    path = strings.Trim(path, "/")
    var pathParts = strings.Split(path, "/")
    return pathParts
}

func NewRouter() *Router {
    return &Router{ NextRoutes: make(map[string]Router)}
}

type Route struct {
    path string
    methods []string
}

func (r *Route) getPath() string {
    return r.path
}

// set path for route
func (r *Route) Path(path string) *Route {
    r.path = path
    return r
}

// set allowed methods for route
func (r *Route) Methods(methods []string) *Route {
    r.methods = methods
    return r
}

func NewRoute(path string) *Route {
    return &Route { path: path, methods: []string{http.MethodGet, http.MethodPost} }
}
