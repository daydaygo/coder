package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "80"
    }
    username := os.Getenv("USERNAME")
    if username == "" {
        username = "world"
    }
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        fmt.Fprintf(writer, "hello %s\n\n", username)
    })
    http.ListenAndServe(":" + port, nil)
}