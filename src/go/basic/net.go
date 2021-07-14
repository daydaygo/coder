package basic

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

func netDemo() {
	// net.FlagUp // %b
}

func tcpServer() {
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Print(err) // eg: conn abort
			continue
		}
		go handleCoon(c)
	}
}

// echo
func handleCoon(c net.Conn) {
	defer c.Close()
	// io.Copy(c, c) // ignore err
	input := bufio.NewScanner(c) // 回声echo
	for input.Scan() {
		go echo(c, input.Text(), time.Second)
	}
}

func echo(c net.Conn, s string, t time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(s))
	time.Sleep(t)
	fmt.Fprintln(c, "\t", s)
	time.Sleep(t)
	fmt.Fprintln(c, "\t", strings.ToLower(s))
}

// nc
func tcpClient() {
	c, err := net.Dial("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	go mustCopy(os.Stdout, c) // 回声echo
	mustCopy(c, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err) // must
	}
}

// http server
var mu sync.Mutex
var cnt int

func httpServer() {
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

// http client
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
