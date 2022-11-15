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

type TestOutput struct {
    Body string
    StatusCode int
    Handler Controller
}

type TestCase struct {
    route *Route
    requestPath string
    requestMethod string
    output TestOutput
}

func NewTestOutput(body string, statusCode int) (TestOutput) {
    var handler = Controller{body}
    return TestOutput{body, statusCode, handler}
}

var testCases = []TestCase {
    {NewRoute("/foo").Methods(http.MethodGet),                   "/foo",                   http.MethodGet, NewTestOutput("simple route match", 200)},
    {NewRoute("/bar").Methods(http.MethodGet),                   "/bar",                   http.MethodGet, NewTestOutput("second simple route match", 200)},
    {NewRoute("/post/list").Methods(http.MethodGet),             "/post/list",             http.MethodGet, NewTestOutput("third simple route match", 200)},
    {NewRoute("/foo/bar").Methods(http.MethodGet),               "/foo/bar",               http.MethodGet, NewTestOutput("simple route with root already defined", 200)},
    {NewRoute("/long/path/to/url/long").Methods(http.MethodGet), "/long/path/to/url/long", http.MethodGet, NewTestOutput("long path to url", 200)},
    {NewRoute("/trailing/slash").Methods(http.MethodGet),        "/trailing/slash/",       http.MethodGet, NewTestOutput("should match a request with trailing slash", 200)},
    {nil,                                                        "/missing-route",         http.MethodGet, NewTestOutput("Page not found.\n", 404)},
    {NewRoute("/foo/method/post").Methods(http.MethodPost, http.MethodConnect),      "/foo/method/post", http.MethodGet, NewTestOutput("Method not available\n", 405)},
    {NewRoute("/foo/method/put").Methods(http.MethodPut),        "/foo/method/put",        http.MethodPost, NewTestOutput("Method not available\n", 405)},
    // {NewRoute("/foo/bar/{id}").Methods(http.MethodGet),        "/foo/bar/1",        http.MethodGet, NewTestOutput("Route with param", 200)},
}

func TestRouter(t *testing.T) {
    var router = NewRouter()

    for _, test := range testCases{
        if test.route != nil {
            router.Handle(test.route, test.output.Handler)
        }
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

        if (resp.StatusCode != test.output.StatusCode) {
            t.Errorf("Status code should be %d got %d", test.output.StatusCode, resp.StatusCode)
        }
        if (string(body) != test.output.Body) {
            t.Errorf("Response should be |" + test.output.Body + "| got |%s|", string(body))
        }
    }
}
