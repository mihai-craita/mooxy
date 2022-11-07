package mooxy

import (
    "testing"
    "io"
    "net/http"
    "net/http/httptest"
)

type Controller struct {
    content string
}

func (c Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, c.content)
}

type TestCase struct {
    handlerPath string
    requestPath string
    handlerOutput string
    method string
}

var testCases = []TestCase {
    {"/foo",                   "/foo",                   "simple route match",                         http.MethodGet},
    {"/bar",                   "/bar",                   "second simple route match",                  http.MethodGet},
    {"/post/list",             "/post/list",             "third simple route match",                   http.MethodGet},
    {"/foo/bar",               "/foo/bar",               "simple route with root already defined",     http.MethodGet},
    {"/long/path/to/url/long", "/long/path/to/url/long", "long path to url",                           http.MethodGet},
    {"/trailing/slash",        "/trailing/slash/",       "should match a request with trailing slash", http.MethodGet},
    // {"/multi-method-request",  "/multi-method-request",        "post", http.MethodPost},
    // {"/multi-method-request",  "/multi-method-request",        "get", http.MethodGet},
    // {"/bar/{id}", "/bar/1", "route with param"},
}

func TestRouter(t *testing.T) {
    var router = NewRouter()

    for _, test := range testCases{
        var handler = Controller{test.handlerOutput}
        var route = NewRoute(test.handlerPath);
        if (test.method != http.MethodGet) {
            route.Methods([]string {http.MethodGet, http.MethodPost})
        }
        router.Handle(route, handler)
    }

    // t.Log(router.NextRoutes)

    for _, test := range testCases{
        req := httptest.NewRequest(test.method, test.requestPath, nil)
        w := httptest.NewRecorder()

        // make the call
        router.GetServer(w, req)

        resp := w.Result()
        body, _ := io.ReadAll(resp.Body)

        // t.Log(resp.StatusCode)
        // t.Log(resp.Header.Get("Content-Type"))
        // t.Log(string(body))

        if (resp.StatusCode != 200) {
            t.Errorf("Status code should be 200 got %d", resp.StatusCode)
        }
        if (string(body) != test.handlerOutput) {
            t.Errorf("Response should be " + test.handlerOutput + " got %s", string(body))
        }
    }

    // var methods = []string{http.MethodGet}

    // r.Handle({ path: '/login', method: ['GET']}, handler) 
    // r.Handle({ path: '/img/*', method: ['GET']}, handler)
}
