package mooxy

type HTTPMethod string

type HttpMethods struct {
    methods []HTTPMethod
}

func (h HttpMethods) Has(method HTTPMethod) (bool) {
    for _, val := range h.methods {
		if val == method {
			return true
		}
	}

	return false
}

func NewHttpMethods(m ...HTTPMethod) (HttpMethods) {
    return HttpMethods{methods: m} 
}

