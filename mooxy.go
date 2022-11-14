package mooxy

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

type Router struct {
    NextRoutes map[string]*Router
    Handler *http.Handler
    //AvailableMethods HttpMethodsArray
    AvailableMethods HttpMethods
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
            currentMatrix[p].AvailableMethods = NewHttpMethods(route.methods...)
        }

        currentMatrix = currentMatrix[p].NextRoutes
    }
}

func (router *Router) GetServer(w http.ResponseWriter, r *http.Request) {
    // here we will match to the correct handler and ServeHTTP
    var foundRouter, err = router.getRouterForUrl(*r.URL)

    if (err != nil) {
        http.Error(w, err.Error(), 404)
        return
    }

    rt := *foundRouter
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

    var currentMatrix = router.NextRoutes
    var routerForPath = router
    for _, p := range pathParts {
        routerForPath = currentMatrix[p]
        if (routerForPath == nil) {
            return nil, errors.New("Page not found.")
        }
        if len(routerForPath.NextRoutes) != 0 {
            currentMatrix = routerForPath.NextRoutes
        }
    }
    return routerForPath, nil
}

func getPathParts(path string) []string {
    path = strings.Trim(path, "/")
    var pathParts = strings.Split(path, "/")
    return pathParts
}

func NewRouter() *Router {
    return &Router{ NextRoutes: make(map[string]*Router)}
}
