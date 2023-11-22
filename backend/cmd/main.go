package main

import (
	"net/http"
	"time"
)

func main() {
	// // Create a new request multiplexer
	// // Take incoming requests and dispatch them to the matching handlers
	// mux := http.NewServeMux()

	// // Register the routes and handlers
	// mux.Handle("/", &homeHandler{})
	// mux.Handle("/test", &testHandler{})

	// Appeasing the linter??
	const three = 3

	// Run the server
	server := &http.Server{
		Addr:              ":8000",
		ReadHeaderTimeout: three * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// type homeHandler struct{}

// func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//  w.Write([]byte("This is my home page"))
// }

// type testHandler struct{}

// func (t *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//  w.Write([]byte("This is another test message"))
// }
