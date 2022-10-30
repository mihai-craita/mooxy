package mooxy

import (
	"net/http"
	"strings"
)

type Router struct {
    NextRoutes map[string]*Router
    Handler http.Handler
}

func (r *Router) Handle (route *Route, handler http.Handler) {
    var pathParts = getPathParts(route.Path)
    var lastElementIndex = len(pathParts) - 1

    currentMatrix := r.NextRoutes
    for index, p := range pathParts{
        var c = currentMatrix[p]
        if c == nil {
            currentMatrix[p] = NewRouter()
        }
        if (index == lastElementIndex) {
            currentMatrix[p].Handler = handler
        }

        currentMatrix = currentMatrix[p].NextRoutes
    }
}

func (router *Router) GetServer(w http.ResponseWriter, r *http.Request) {
    // here we will match to the correct handler and ServeHTTP
    var path = r.URL.Path
    var pathParts = getPathParts(path)
    var currentMatrix = router.NextRoutes
    var handler = router.Handler
    for _, p := range pathParts {
        handler = currentMatrix[p].Handler
        if len(currentMatrix[p].NextRoutes) > 0 {
            currentMatrix = currentMatrix[p].NextRoutes
        }
    }
    handler.ServeHTTP(w, r)
}

func getPathParts(path string) []string {
    path = strings.Trim(path, "/")
    var pathParts = strings.Split(path, "/")
    return pathParts
}

func NewRouter() *Router {
    return &Router{ NextRoutes: make(map[string]*Router)}
}
