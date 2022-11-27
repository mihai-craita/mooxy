package mooxy

import (
    "testing"
)

func TestHttpMethods(t *testing.T) {
    sut := NewHttpMethods()
    sut.Add("GET")
    if !sut.Has("GET") {
        t.Errorf("Method GET shoul be available")
    }

    sut.Add("POST")
    if !sut.Has("POST") {
        t.Errorf("Method POST should be available")
    }

    sut.Add("PUT")
    if !sut.Has("PUT") {
        t.Errorf("Method PUT should be available")
    }

    if sut.Has("DELETE") {
        t.Errorf("Method DELETE should be available")
    }
}

