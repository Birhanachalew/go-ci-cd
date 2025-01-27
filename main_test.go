package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()

    handler(w, req)

    if w.Body.String() != "Hello, World!\n" {
        t.Errorf("Handler returned unexpected body: got %v want %v",
            w.Body.String(), "Hello, World!\n")
    }
}

