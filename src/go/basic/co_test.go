package basic

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"sync"
	"testing"
	"time"
)

func Test_go(t *testing.T) {
	// ch chan
	ch := make(chan struct{}) // 0-mem-use
	ch <- struct{}{}
	close(ch) // close write
	_, ok := <-ch
	if !ok { // chan close
		return
	}

	var tokens = make(chan struct{}, 20) // token sem 类似 sync.Mutex
	tokens <- struct{}{}

	// select: timeout abort
	ch1, ch2 := make(<-chan struct{}), make(chan<- struct{})
	select {
	case val := <-ch1: // 从 ch1 读取数据
		fmt.Printf("get value %d from ch1\n", val)
	case ch2 <- struct{}{}: // 使用 ch2 发送消息
		fmt.Println("send value by ch2")
	case <-time.After(2 - time.Second): // 超时设置
		fmt.Println("Time out")
		return
	}

	// buffered & unidirectional & close() & bench & pipeline: https://github.com/adonovan/gopl.io/tree/master/ch8/cake
	// broadcast: https://github.com/adonovan/gopl.io/tree/master/ch8/chat/chat.go
	// monitor go: https://github.com/adonovan/gopl.io/tree/master/ch9/bank1/bank.go

	fmt.Println(runtime.GOOS, runtime.GOARCH, runtime.NumGoroutine())
}

// animation
func Test_spinner(t *testing.T) {
	go func() {
		for {
			for _, r := range `-\|/` {
				fmt.Printf("\r%c", r)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("\rmain")
}

// wait go end
func Test_co(t *testing.T) {
	// co1
	go task("co1")
	time.Sleep(time.Second) // only for test

	// co2
	done := make(chan struct{})
	go func() {
		task("co2")
		done <- struct{}{}
	}()
	<-done

	// co3
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		task("co3")
	}()
	wg.Wait()

	fmt.Println("main")
}

func task(s string) {
	fmt.Println(s + " start")
	time.Sleep(time.Second)
	fmt.Println(s + " end")
}

func Test_close(t *testing.T) {
	s := []string{"a", "b", "c"}
	ch := make(chan string)
	for _, i := range s {
		go task2(i, ch)
	}
	// go func() { // 使用 waitGroup 关闭 ch
	// 	wg.wait()
	// 	close(ch)
	// }()
	// for i := range ch{ // ch 未 close, 会导致主协程一直 wait
	// 	fmt.Println(i)
	// }
	for range s {
		fmt.Println(<-ch)
	}
}

func task2(s string, ch chan<- string) {
	time.Sleep(time.Second)
	ch <- s
}

func Test_tick(t *testing.T) {
	tick := time.Tick(time.Second)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		<-tick
	}
}

func Test_abort(t *testing.T) {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	fmt.Println("countdown. press return to abort")
	select {
	case <-time.After(3 * time.Second): // do nothing
	case <-abort:
		return
	}
	fmt.Println("launch")
}

func Test_context(t *testing.T) {
	ctx := context.Background()          // new empty context
	ctx = context.WithValue(ctx, "a", 1) // add context info
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	go testContext(ctx, time.Second*4) // task1
	go testContext(ctx, time.Second*3) // task2

	time.Sleep(time.Second * 5) // let all go run
}

func testContext(ctx context.Context, t time.Duration) {
	select {
	case <-time.After(t): // 模拟耗时任务
	case <-ctx.Done(): // ctx cancel
		fmt.Println(ctx.Err()) // ctx err
		// clean
	}
}

func Test_animation(t *testing.T) {
	go func() {
		for {
			for _, r := range `-\|/` {
				fmt.Printf("\r%c", r)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	time.Sleep(time.Second * 3)
	fmt.Println("\rmain")
}

func Test_fetchAll(t *testing.T) {
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
	for i := range ch { // 程序会 hang
		fmt.Println(i)
	}
	// for range urls {
	// 	fmt.Println(<-ch)
	// }
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
