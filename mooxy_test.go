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

func NewTestOutput(body string, statusCode int) (TestOutput) {
    var handler = Controller{body}
    return TestOutput{body, statusCode, handler}
}


type TestRequest struct {
    path string
    method string
}

func NewTestRequest(path string, method string) (TestRequest) {
    return TestRequest{path, method}
}

type TestCase struct {
    route *Route
    request TestRequest
    output TestOutput
}

var testCases = []TestCase {
    {
        NewRoute("/foo").Methods(http.MethodGet),
        NewTestRequest("/foo", http.MethodGet),
        NewTestOutput("simple route match", 200),
    },
    {
        NewRoute("/bar").Methods(http.MethodGet),
        NewTestRequest("/bar", http.MethodGet),
        NewTestOutput("second simple route match", 200),
    },
    {
        NewRoute("/post/list").Methods(http.MethodGet),
        NewTestRequest("/post/list", http.MethodGet),
        NewTestOutput("third simple route match", 200),
    },
    {
        NewRoute("/foo/bar").Methods(http.MethodGet),
        NewTestRequest("/foo/bar", http.MethodGet),
        NewTestOutput("simple route with root already defined", 200),
    },
    {
        NewRoute("/long/path/to/url/long").Methods(http.MethodGet),
        NewTestRequest("/long/path/to/url/long", http.MethodGet),
        NewTestOutput("long path to url", 200),
    },
    {
        NewRoute("/trailing/slash").Methods(http.MethodGet),
        NewTestRequest("/trailing/slash/", http.MethodGet),
        NewTestOutput("should match a request with trailing slash", 200),
    },
    {
        nil,
        NewTestRequest("/missing-route", http.MethodGet),
        NewTestOutput("Page not found.\n", 404),
    },
    {
        NewRoute("/foo/method/post").Methods(http.MethodPost, http.MethodConnect),
        NewTestRequest("/foo/method/post", http.MethodGet),
        NewTestOutput("Method not available\n", 405),
    },
    {
        NewRoute("/foo/method/put").Methods(http.MethodPut),
        NewTestRequest("/foo/method/put", http.MethodPost),
        NewTestOutput("Method not available\n", 405),
    },
    {
        NewRoute("/foo/bar/{id}").Methods(http.MethodGet),
        NewTestRequest("/foo/bar/1", http.MethodGet),
        NewTestOutput("Route with param", 200),
    },
    {
        NewRoute("/post/{slug}").Methods(http.MethodGet),
        NewTestRequest("/post/hello-world", http.MethodGet),
        NewTestOutput("Route with string param", 200),
    },
    {
        NewRoute("/post/{slug}/comments/{userId}").Methods(http.MethodGet),
        NewTestRequest("/post/hello-world/comments/1", http.MethodGet),
        NewTestOutput("Route with interleaved params", 200),
    },
    {
        NewRoute("/img/{*}").Methods(http.MethodGet),
        NewTestRequest("/img/2022/01/01/test.jpg", http.MethodGet),
        NewTestOutput("Route that matches all subsequent paths", 200),
    },
    {
        nil,
        NewTestRequest("/img", http.MethodGet),
        NewTestOutput("Page not found.\n", 404),
    },
}

func TestRouter(t *testing.T) {
    var router = NewRouter()

    for _, test := range testCases{
        if test.route != nil {
            router.Handle(test.route, test.output.Handler)
        }
    }

    for _, test := range testCases{
        req := httptest.NewRequest(test.request.method, test.request.path, nil)
        w := httptest.NewRecorder()

        // make the call
        router.ServeHTTP(w, req)

        resp := w.Result()
        body, _ := io.ReadAll(resp.Body)

        if (resp.StatusCode != test.output.StatusCode) {
            t.Errorf("Status code should be %d got %d", test.output.StatusCode, resp.StatusCode)
        }
        if (string(body) != test.output.Body) {
            t.Errorf("Response should be |" + test.output.Body + "| got |%s|", string(body))
        }
    }
}
