package main

import (
	    "net/http"
        "fmt"
)

func main() {
fmt.Println("HELLO WORLDS!!!!!!!!!!!!!")
 // Create a new request multiplexer
 // Take incoming requests and dispatch them to the matching handlers
 mux := http.NewServeMux()

 // Register the routes and handlers
 mux.Handle("/", &homeHandler{})
 mux.Handle("/test", &testHandler{})

 // Run the server
 http.ListenAndServe(":8000", mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
 w.Write([]byte("This is my home page"))
}

type testHandler struct{}

func (t *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
 w.Write([]byte("This is a newer test message"))
}
