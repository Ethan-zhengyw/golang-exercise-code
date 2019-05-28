package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", timed(helloWithDB("localhost:3306")))
	http.ListenAndServe(":3000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}

func timed(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		end := time.Now()
		fmt.Printf("/hello called, ")
		fmt.Println("Request cost", end.Sub(start))
	}
}

// add dbURL parameter to http handler by using anonymous function closure
func helloWithDB(dbURL string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Accessing db: ", dbURL, " ")
		fmt.Fprintln(w, "<h1>Hello!</h1>")
	}
}
