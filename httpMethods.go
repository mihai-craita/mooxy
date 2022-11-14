package mooxy

import "golang.org/x/exp/slices"

type HTTPMethod string

type HttpMethods struct {
    methods []HTTPMethod
}

func (h HttpMethods) Has(method HTTPMethod) (bool) {
    if (slices.Contains(h.methods, method)) {
        return true
    }
    return false
}

func NewHttpMethods(m ...HTTPMethod) (HttpMethods) {
    return HttpMethods{methods: m} 
}

