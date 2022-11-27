package mooxy

import (
    "testing"
)

func TestHttpMethods(t *testing.T) {
    sut := NewHttpMethods()
    handler := Controller{}
    sut.Add("GET", handler)
    if !sut.Has("GET") {
        t.Errorf("Method GET shoul be available")
    }

    sut.Add("POST", handler)
    if !sut.Has("POST") {
        t.Errorf("Method POST should be available")
    }

    sut.Add("PUT", handler)
    if !sut.Has("PUT") {
        t.Errorf("Method PUT should be available")
    }

    if sut.Has("DELETE") {
        t.Errorf("Method DELETE should be available")
    }
}
