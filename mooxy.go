package mooxy

import (
	"net/http"
	"net/url"
	"strings"
    "errors"
)

type Router struct {
    NextRoutes map[string]*Router
    Handler *http.Handler
    //AvailableMethods HttpMethodsArray
    AvailableMethods []string
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
            currentMatrix[p].Handler = &handler
        }

        currentMatrix = currentMatrix[p].NextRoutes
    }
}

func (router *Router) GetServer(w http.ResponseWriter, r *http.Request) {
    // here we will match to the correct handler and ServeHTTP
    var handler, err = router.getHandlerForUrl(*r.URL)
    if (err != nil) {
        http.Error(w, err.Error(), 404)
        return
    }
    handler.ServeHTTP(w, r)
}

func (router *Router) getHandlerForUrl(u url.URL) (h http.Handler, er error) {
    var path = u.Path
    var pathParts = getPathParts(path)

    var currentMatrix = router.NextRoutes
    var handler = router.Handler
    for _, p := range pathParts {
        var routerForPath = currentMatrix[p]
        if (routerForPath == nil) {
            return nil, errors.New("Page not found.")
        }
        handler = routerForPath.Handler
        if len(routerForPath.NextRoutes) != 0 {
            currentMatrix = routerForPath.NextRoutes
        }
    }
    return *handler, nil
}

func getPathParts(path string) []string {
    path = strings.Trim(path, "/")
    var pathParts = strings.Split(path, "/")
    return pathParts
}

func NewRouter() *Router {
    return &Router{ NextRoutes: make(map[string]*Router)}
}
