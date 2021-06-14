package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	u := "http://baidu.com"
	httpGet(u)
	httpReq(u)
}

func httpGet(u string) {
	resp, _ := http.Get(u) // url.QueryEscape()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK { // resp.Status
		return
	}
	b, _ := ioutil.ReadAll(resp.Body)     // txt
	json.NewDecoder(resp.Body).Decode(&b) // json
	fmt.Println(string(b))
}

func httpHead(u string) {
	_, err := http.Head(u)
	if err == nil {
		return // success
	}
}

func httpReq(u string) {
	req, _ := http.NewRequest("GET", u, nil)
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json") // github api v3
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	ct := resp.Header.Get("Content-Type")
	fmt.Println(string(b), ct)
}

func urlValue() {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Get("lang")
}
