package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>test</h1>")
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health check</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/health_check", check)
	http.HandleFunc("/test", test)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)
}
