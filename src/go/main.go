package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://baidu.com")
	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(b))
}
