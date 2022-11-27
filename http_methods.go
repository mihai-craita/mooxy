package mooxy

import "net/http"

type HTTPMethod string

type HTTPMethods struct {
    methods map[HTTPMethod]http.Handler
}

func (h HTTPMethods) Has(method HTTPMethod) (bool) {
	_, ok := h.methods[method]
	if ok {
		return true
	}

	return false
}

func (h HTTPMethods) Add(method HTTPMethod, handler http.Handler) {
	h.methods[method] = handler;
}

func NewHttpMethods() (HTTPMethods) {
	return HTTPMethods{methods: make(map[HTTPMethod]http.Handler)} 
}
