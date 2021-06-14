# book| the go programming language

- [gopl](https://www.gopl.io) go语言圣经 go程序设计语言
- [code](https://github.com/adonovan/gopl.io/) [exercise](https://github.com/relsa/exercise-gopl.io)
- [pdf](https://github.com/dreamrover/gopl-pdf)
- [中文版](https://flydk.gitbooks.io/go)

- preface: the origin of go; the go project; book organization; more info
- program structure: name declar var assign type package&file scope
- basic type: int float complex bool string const
- composite type: arr slice map struct json text&html
- func: delcar recursion multi-return err func-var anonymous variadic(多参数) defer panic recover
- method: declare pointer-receiver struct-embed value&expr
- interface: contract&satisfy type&value typeAsset typeSwitch advice
- go&chan: go chan.buffered select.multiplex cancel.broadcast.close
- co&shareMem: raceCondition mutex rwmutex once.lazyInit raceDetector go&thread
- pkg&tool: import `_` naming GOPATH internalPackage get/build/doc/list
- test: test cover bench prof
- reflection: Type&Value Display stuctFieldTag method caution
- lowLevel: unsafe cgo Caution

- 正如Rob Pike所说，“软件的复杂性是乘法级相关的”，通过增加一个部分的复杂性来修复问题通常将慢慢地增加其他部分的复杂性。通过增加功能、选项和配置是修复问题的最快的途径，但是这很容易让人忘记简洁的内涵，即从长远来看，简洁依然是好软件的关键因素。
- 一个包由位于单个目录下的一个或多个.go源代码文件组成, 目录定义包的作用。每个源文件都以一条package声明语句开始
- 在main里的main 函数 也很特殊，它是整个程序执行时的入口
- 编译器会主动把特定符号后的换行符转换为分号
- 格式化： gofmt goimports
- 如果变量没有显式初始化，则被隐式地赋予其类型的零值（zero value）
- `The Origins of Go`
- `Organization of the Book`
- four major kinds of declarations: var, const, type, and func
- Not every value has an address, but every variable does
- In Go, the sign of the remainder is always the same as the sign of the dividend, so -5%3 and -5%-3 are both -2
- Usually when a function returns a non-nil error, its other results are undefined and should be ignored
- When the error is ultimately handled by the program’s main function, it should provide a clear causal chain from the root problem to the overall failure, reminiscent of a NASA accident investigation: `genesis: crashed: no parachute: G-switch failed: bad relay orientation`
- anonymous functions that share a variable local; an anonymous function can access its enclosing function’s variables, including named results
- the scope rules for loop variables
- as a general rule, you should not attempt to recover from another package’s panic
- two key principles of object-oriented programming, encapsulation and composition
- method receiver have both type T and *T
- We should forget about small efficiencies, say about 97% of the time: premature optimization is the root of all evil.