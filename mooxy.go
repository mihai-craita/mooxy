package mooxy

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

type Router struct {
    Children map[string]*Router
    Handler *http.Handler
    AvailableMethods HttpMethods
}

func (r *Router) Handle (route *Route, handler http.Handler) {
    var pathParts = getPathParts(route.Path)
    var lastElementIndex = len(pathParts) - 1

    currentMatrix := r.Children
    for index, p := range pathParts{
        var c = currentMatrix[p]
        if c == nil {
            currentMatrix[p] = NewRouter()
        }
        if (index == lastElementIndex) {
            currentMatrix[p].Handler = &handler
            currentMatrix[p].AvailableMethods = NewHttpMethods(route.methods...)
        }

        currentMatrix = currentMatrix[p].Children
    }
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // here we will match to the correct handler and ServeHTTP
    var foundRouter, err = router.getRouterForUrl(*r.URL)

    if (err != nil) {
        http.Error(w, err.Error(), 404)
        return
    }

    rt := *foundRouter
    if (rt.Handler == nil) {
        http.Error(w, "Page not found.", 404)
        return
    }
    handler := *rt.Handler

    if (!rt.AvailableMethods.Has(HTTPMethod(r.Method))) {
        http.Error(w, "Method not available", 405)
        return
    }
    handler.ServeHTTP(w, r)
}

func (router *Router) getRouterForUrl(u url.URL) (h *Router, er error) {
    var path = u.Path
    var pathParts = getPathParts(path)

    var currentMatrix = router.Children
    var routerForPath = router
    for _, p := range pathParts {
        routerForPath = currentMatrix[p]
        if routerForPath == nil {
            routerForPath = getVariableRouter(currentMatrix)
            if routerForPath == nil {
                return nil, errors.New("Page not found.")
            }
        }
        if len(routerForPath.Children) != 0 {
            currentMatrix = routerForPath.Children
        }
    }
    return routerForPath, nil
}

func getPathParts(path string) []string {
    path = strings.Trim(path, "/")
    var pathParts = strings.Split(path, "/")
    return pathParts
}

func getVariableRouter(routerList map[string]*Router) (*Router) {
    for k, r := range routerList {
        if strings.HasPrefix(k, "{") && strings.HasSuffix(k, "}") {
            return r
        }
    }
    return nil
}

func NewRouter() *Router {
    return &Router{ Children: make(map[string]*Router)}
}
