package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	ch := make(chan string)
	start := time.Now()
	urls := []string{ // http://www.alexa.cn/siterank/
		"http://baidu.com",
		"http://tmall.com",
		"http://qq.com",
	}
	for _, url := range urls {
		go fetch(url, ch)
	}
	// for i := range ch { // 程序会 hang
	// 	fmt.Println(i)
	// }
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Println(time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	// b, err := ioutil.ReadAll(resp.Body)
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	ch <- fmt.Sprintf("%s %d %.2f", url, nbytes, time.Since(start).Seconds())
}
