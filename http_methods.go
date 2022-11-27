package mooxy

type HTTPMethod string

type HTTPMethods struct {
    methods map[HTTPMethod]bool
}

func (h HTTPMethods) Has(method HTTPMethod) (bool) {
	v, ok := h.methods[method]
	if ok {
		return v
	}

	return false
}

func (h HTTPMethods) Add(method HTTPMethod) {
	h.methods[method] = true;
}

func NewHttpMethods() (HTTPMethods) {
	return HTTPMethods{methods: make(map[HTTPMethod]bool)} 
}
