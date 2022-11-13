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
    route *Route
    requestPath string
    requestMethod string
    handlerOutput string
}

var testCases = []TestCase {
    {NewRoute("/foo").Methods([]string {http.MethodGet}),                   "/foo",                   http.MethodGet, "simple route match",                         },
    {NewRoute("/bar").Methods([]string {http.MethodGet}),                   "/bar",                   http.MethodGet, "second simple route match",                  },
    {NewRoute("/post/list").Methods([]string {http.MethodGet}),             "/post/list",             http.MethodGet, "third simple route match",                   },
    {NewRoute("/foo/bar").Methods([]string {http.MethodGet}),               "/foo/bar",               http.MethodGet, "simple route with root already defined",     },
    {NewRoute("/long/path/to/url/long").Methods([]string {http.MethodGet}), "/long/path/to/url/long", http.MethodGet, "long path to url",                           },
    {NewRoute("/trailing/slash").Methods([]string {http.MethodGet}),        "/trailing/slash/",       http.MethodGet, "should match a request with trailing slash", },
    // {"/multi-method-request",  "/multi-method-request",        "post", http.MethodPost},
    // {"/multi-method-request",  "/multi-method-request",        "get", http.MethodGet},
    // {"/bar/{id}", "/bar/1", "route with param"},
}

func TestRouter(t *testing.T) {
    var router = NewRouter()

    for _, test := range testCases{
        var handler = Controller{test.handlerOutput}
        router.Handle(test.route, handler)
    }

    for _, test := range testCases{
        req := httptest.NewRequest(test.requestMethod, test.requestPath, nil)
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

    req := httptest.NewRequest(http.MethodGet, "/missing-route", nil)
    w := httptest.NewRecorder()

    // make the call
    router.GetServer(w, req)

    resp := w.Result()
    body, _ := io.ReadAll(resp.Body)

    // t.Log(resp.StatusCode)
    // t.Log(resp.Header.Get("Content-Type"))
    // t.Log(string(body))

    if (resp.StatusCode != 404) {
        t.Errorf("Status code should be 404 got %d", resp.StatusCode)
    }
    output := "Page not found.\n"
    if (string(body) != output) {
        t.Errorf("Response should be |" + output + "| got |%s|", string(body))
    }

    // var methods = []string{http.MethodGet}

    // r.Handle({ path: '/login', method: ['GET']}, handler) 
    // r.Handle({ path: '/img/*', method: ['GET']}, handler)
}
