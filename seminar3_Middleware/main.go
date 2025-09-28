package main

import (
	"fmt"
	"net/http"
)

// "<host:ip>/buy": --> (request) --> Middleware (before) --> Handler() --> Middleware (after) --> (response) -->

func firstMiddleware(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("1 - middleware (before handler) ---> ")
		f(writer, request) // Handler
		fmt.Println("<--- 1 - middleware (after handler)")
	}
}

func secondMiddleware(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("2 - middleware (before handler) ---> ")
		f(writer, request) // Handler
		fmt.Println("<--- 2 - middleware (after handler)")
	}
}

type welcome string

func (wc welcome) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to our Golang application!")
}

func main() {
	router := http.NewServeMux()

	var wc welcome

	router.Handle("/", wc)
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the login page")
	})

	router.HandleFunc("/buy", firstMiddleware(secondMiddleware(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Success!")
	})))

	server := http.Server{
		Addr:    ":8070",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server starting error: ", err)
	}
}
