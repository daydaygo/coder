package basic

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"testing"
	"time"
	"unicode/utf8"
	"unsafe"

	"github.com/pkg/errors"
)

func Test_type(t *testing.T) {
	fmt.Println(-10%3, -10%-3) // -1 -1, same sign with dividend

	// s string
	s := "hello"
	for k, v := range s { // k int; v int32
		fmt.Println(k, v)
	}
	// c := 'c' // int32
	// var c2 byte = 'c' // uint8
	fmt.Println('0') // '0'-'9'
	// s := `[1, null, 2, 3]` // json -> [1, 0, 2, 3]

	// arr array: fixed-length; Using pointer
	arr1 := [...]int{1, 2, 3}
	var arr2 [3]int = [3]int{1, 2, 3}
	fmt.Printf("%x %x %t %T\n", arr1, arr2, arr1 == arr2, arr1) // %x array/slice %t bool %T type

	// a slice: pointer+len()+cap(); not comparable; nil
	// make([]int, 0, 3) // make([]int, 3) 会导致前 3 个已赋值, append() 从第 4 个开始
	// https://ueokande.github.io/go-slice-tricks/
	a := arr1[:]
	a[2] = 4
	// var a []int // a=nil; a = []int{} // a!=nil
	// make([]T, len, cap) = make([]T, cap)[:len]
	// append() cap扩容
	a = append(a[:10], a[10+1:]...) // 删除第i个值
	var x int
	a = append(a, x)                 // push
	x, a = a[len(a)-1], a[:len(a)-1] // stack
	x, a = a[0], a[1:]               // queue

	// m map[K]V: unordered/random; K == comparable; cannot &V
	m := map[string]int{"a": 1, "b": 2} // m := make(map[string]int)
	delete(m, "a")
	if _, ok := m["b"]; ok {
	}
	// graph := make(map[string]map[string]bool) // graph[from][to] edge
	seen := make(map[string]bool)
	if !seen["i"] {
		seen["i"] = true
		// do
	}

	// p struct
	// p := &Point{1, 2} // %#v struct; p=new(Point); %p
	// w = Wheel{Circle{Point{8, 8}, 5}, 20} // embed %#v
	// w.X = 8 // equivalent to w.circle.point.X = 8
	// Year int `json:"released"` // field tag
}

func Test_func(t *testing.T) {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

// squares: function value with state
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

// math big bits cmplx rand
func Test_math(t *testing.T) {
	const pi = math.Pi
	// math.Hypot()

	// rand
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ru := rune(r.Intn(0x1000)) // rand rune up to '\u0999'
	fmt.Println(ru)
}

// container: heap(intHeap priorityQueue) list(doublyLinkedList) ring(circularList)
func Test_container(t *testing.T) {
}

type qsort []int // sort.Interface

func (q qsort) Len() int {
	return len(q)
}

func (q qsort) Less(i, j int) bool {
	return q[i] < q[j] // 升序
}

func (q qsort) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q qsort) String() string { // fmt.Stringer
	return ""
}

func Test_qsort(t *testing.T) {
	a := qsort{1, 3, 2, 4}
	sort.Sort(a)
	// sort.IntsAreSorted(a)
	fmt.Println(a)
}

// https://github.com/adonovan/gopl.io/blob/master/ch5/toposort/main.go
func Test_toposort(t *testing.T) {
}

func Test_bfs(t *testing.T) {
	q := make([]string, 0)                         // work queue
	f := func(item string) []string { return nil } // handle item
	seen := make(map[string]bool)                  // 去重
	for len(q) > 0 {
		a := q
		q = nil
		for _, i := range a { // 处理当前层
			if !seen[i] {
				seen[i] = true
				q = append(q, f(i)...) // 添加下一层
			}
		}
	}
}

func Test_strings(t *testing.T) {
	s := "hello"
	fmt.Println(s[0]) // byte 104; range -> rune; 不可以直接修改 s[0]='a'

	strings.Join([]string{"hello", "world"}, ";") // os.Args[1:] flag.Args()
	strings.Split(s, ",")
	// strings.HasPrefix()
	strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")
	strings.LastIndexAny(s, "/")

	var b strings.Builder
	b.Grow(len(s)) // 多次拼接, 可以预先分配内存
	b.WriteString(s)
	b.String()
}

func Test_bytes(t *testing.T) {
	a := []int{1, 2, 3}
	var buf bytes.Buffer
	buf.WriteByte('[') // buf.WriteRune('世')
	for i, v := range a {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	fmt.Println(buf.String())

	// bytes.Compare() bytes.Equal()
}

func Test_bufio(t *testing.T) {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		input.Text()
	}
	// f, err := os.Open(arg) // os.create()
	// n, err := io.Copy(f, data)
	// defer f.Close()
	// data, err := ioutil.ReadFile(filename)
}

func Test_strconv(t *testing.T) {
	strconv.Itoa(123) // int -> ascii
	fmt.Sprintf("%d", 123)

	strconv.Atoi("123")
	strconv.ParseInt("123", 10, 64)

	strconv.FormatInt(123, 2) // fmt.Printf verbs %b, %d, %o, and %x
}

// unicode: utf16 utf8 https://github.com/adonovan/gopl.io/blob/master/ch4/charcount/main.go
func Test_unicode(t *testing.T) {
	s := "hello, 世界"
	fmt.Println(utf8.RuneCountInString(s))
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}

// encoding(byte->txt): ascii85 asn1 base32 base64 binary(int<->byte) csv gob hex pem(PrivacyEnhancedMail) xml
// https://mholt.github.io/json-to-go
// https://github.com/mailru/easyjson
func Test_encoding(t *testing.T) {
	var p []int
	// TotalCount bool `json:"total_count,omitempty"`
	json.Marshal(p) // reflection: struct field <-> json obj
	json.MarshalIndent(p, "", "    ")
	json.Unmarshal([]byte(""), &p) // []byte
	// json.NewEncoder().Decode() // streaming
	json.Marshal(map[int]string{2: "a", 3: "b"}) // {"2":"a","3":"b"} key.int 会被转为 string
}

// https://github.com/adonovan/gopl.io/blob/master/ch4/issuesreport/main.go
func Test_template(t *testing.T) {
	const tpl = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	template.Must(template.New("escape").Parse(tpl)) // auto escape
}

func Test_time(t *testing.T) {
	fmt.Println(time.Now().Format(time.RFC3339Nano))

	const timeout = 1 * time.Minute
	// time.Since(time.Now()).Hours()
	fmt.Printf("%T %[1]v\n", timeout) // "time.Duration 5m0s"

	deadline := time.Now().Add(timeout)
	for try := 0; time.Now().Before(deadline); try++ {
		log.Printf("try %d", try)
		time.Sleep(time.Second << uint(try)) // exponential back-off 指数退避 指数补偿
	}

	r := new(Rocket)                      // r := &Rocket{}
	time.AfterFunc(time.Second, r.Launch) // method value syntax: func() { r.Launch() }
}

type Rocket struct{}

func (r *Rocket) Launch() {}

func Test_errors(t *testing.T) {
	err := errors.New("test") // fmt.Errorf log.Fatalf=err+os.Exit
	err = errors.Wrap(err, "wrap")
	fmt.Printf("%+v", err)
}

func Test_defer(t *testing.T) {
	defer trace("bigSlowOperation")() // f := trace("bigSlowOperation"); defer f()
	time.Sleep(1 * time.Second)
}

func trace(s string) func() {
	defer printStack()
	start := time.Now()
	log.Printf("enter %s", s) // log.syslog
	return func() { log.Printf("exit: %s (%s)", s, time.Since(start)) }
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func Test_panic(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("recover panic: %v", p)
		}
	}()
	panic("test panic") // special type: panic(bailout{}) + switch
}

func Test_path(t *testing.T) {
	fmt.Println(path.Base("a/b/c.com"))
}

func Test_regexp(t *testing.T) {
	regexp.Compile(`^https?:`)
	regexp.MustCompile(`^https?:`) // panic
}

type IntSet struct{ words []uint64 }

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) { // non-negative x
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] |= word
		} else {
			s.words = append(s.words, word)
		}
	}
}

func Test_test(t *testing.T) {
	var out io.Writer = os.Stdout // modify during testing
	// var out = new(bytes.Buffer) // captured output
	got := out.(*bytes.Buffer).String()
	fmt.Println(out, got)

	// reflect.DeepEqual(x, y)
}

func Test_crypto(t *testing.T) {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func Test_reflect(t *testing.T) {
	// reflect.DeepEqual() // 深度比较, 单测中会使用
	// reflect.Indirect(reflect.ValueOf(dest)).Type() // ORM 中大量使用

	v := reflect.ValueOf(3)
	v.Kind()
	v.Index(0) // slice/arr/string
	v.Field(0) // struct
	// v.MapKeys() v.MapIndex(key)
	// v.IsNil() v.Elem() // ptr/interface
	// v.NumMethod() v.Method(0)
}

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}

func cycleCheck(x, y reflect.Value, seen map[comparison]bool) bool {
	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true // identical references
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}

	return false
}

func Test_os(t *testing.T) {
	// signal
	e := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		e <- fmt.Errorf("%s", <-c)
	}()
}
