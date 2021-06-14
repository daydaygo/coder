package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        r := "productpage"
        r1, _ := getSrv("review")
        r = r + "|with review: " + r1
        r2, _ := getSrv("detail")
        r = r + "|with detail: " + r2
        fmt.Fprintf(writer, r)
    })
    http.ListenAndServe(":80", nil)
}

func getSrv(srv string) (string, error) {
    r, err := http.Get("http://" + srv)
    if err != nil {
        return "", err
    }
    defer r.Body.Close()
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return "", nil
    }
    return string(b), nil
}