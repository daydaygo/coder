package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var cnt int

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/count", counter)
	http.ListenAndServe(":8000", nil)
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", hello)
	// http.ListenAndServe(":8000", mux)
}

func hello(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	cnt++
	mu.Unlock()
	// w.Write([]byte("hello"))
	fmt.Fprintf(w, "%q\n", r.URL.Path) // %q safely
	fmt.Fprintln(w, r.Method, r.URL, r.Proto, r.Host, r.RemoteAddr)
	fmt.Println(r.URL.Query().Get("item"))
	for k, v := range r.Header {
		fmt.Fprintln(w, k, v)
	}
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "cnt: %d\n", cnt) // %q safely
	mu.Unlock()
}
