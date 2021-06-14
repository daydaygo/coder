package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        fmt.Fprintf(writer, "detail")
    })
    http.ListenAndServe(":80", nil)
}