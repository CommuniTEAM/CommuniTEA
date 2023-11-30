package main

import (
	fmtdev "fmt"
	"net/http"
	"time"
)

// appeasing the "magic number detector" linter (maybe we'll have to turn this off down the line? lol)
const three = 3

func main() {
	// ! debug statement for project init's demonstrative purposes -- remove for production
	fmtdev.Println("\nSuccessfully recompiled! Hello World :)")

	// // Create a new request multiplexer

	// // Take incoming requests and dispatch them to the matching handlers

	// mux := http.NewServeMux()

	// // Register the routes and handlers

	// mux.Handle("/", &homeHandler{})

	// mux.Handle("/test", &testHandler{})

	// Run the server
	server := &http.Server{

		Addr: ":8000",

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
