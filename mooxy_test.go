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
}

var testCases = []TestCase {
    {"/foo", "/foo", "simple route match"},
    {"/bar", "/bar", "second simple route match"},
    {"/post/list", "/post/list", "third simple route match"},
    {"/foo/bar", "/foo/bar", "fourth simple route match with parents already defined"},
    {"/long/path/to/url/long", "/long/path/to/url/long", "long path to url"},
    // {"/bar/{id}", "/bar/1", "route with param"},
}

func TestRouter(t *testing.T) {
    var router = NewRouter()

    for _, test := range testCases{
        var handler = Controller{test.handlerOutput}
        router.Handle(NewRoute(test.handlerPath), handler)
    }

    // t.Log(router.NextRoutes)

    for _, test := range testCases{
        req := httptest.NewRequest(http.MethodGet, test.requestPath, nil)
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
