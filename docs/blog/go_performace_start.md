# go| go 性能优化入门之「Go代码重构：23倍的性能爆增」实践

最近在整理以前攒的 go 语言学习资料 -- 可能很多人都和我一样, 随手一个收藏, 不动手也不深入, 然后就过去了. 这次从故纸堆里扫出来, 当然不能错过

资料:
- blog 地址: https://www.cnblogs.com/sunsky303/p/9296188.html
- 原作者已经提供好了代码: https://github.com/Deleplace/forks-golang-good-code-bad-code

学习到的知识:
- 使用 `go test` 进行 单测/压测
- 使用 `go tool` 进行 prof/trace
- 性能问题 debug 与优化思路

## let's party

- 作者准备好了代码 https://github.com/Deleplace/forks-golang-good-code-bad-code

- 确定基准, 使用 cpuprof 中的 `ns/op` 作为比较基准

```sh
cd bad
➜  bad git:(master) ✗ go test -bench=. -cpuprofile cpu.prof
goos: darwin
goarch: amd64
pkg: test/bad
BenchmarkParseAdexpMessage-8   	   18999	     63848 ns/op
PASS
ok  	test/bad	2.007s
```

- bad & good 代码对比: good 更惯用，更易读，利用go语言的细节, 后续的修改都基于 good 代码进行

- 查看 trace, 查看 CPU 使用情况

```sh
# 使用 trace 工具
go test -bench=. -trace trace.out
go tool trace trace.out # 会在默认浏览器中打开 trace
```

![](https://images2018.cnblogs.com/blog/420532/201807/420532-20180711195735797-153557120.png)

- trace 分析: 放大 CPU 部分 -> 数千个小的彩色计算切片 + 空闲插槽 -> 一些核心处于空闲状态

- 首先进行竞争检测, **如果发生竞争, 比性能问题更严重**

```sh
# 竞争检测
go test -race
```

- 尝试 **不开协程**, 对应代码: https://github.com/Deleplace/forks-golang-good-code-bad-code/tree/nogo
```go
// 改动就一行
for _, line := range in {
	// go mapLine(line, ch)
	mapLine(line, ch)
}
```

- 使用 cpuprof, 查看热函数调用, 定位到瓶颈

```sh
# 1. 生成 cpuprof
go test -bench=. -cpuprofile cpu.prof
# 2. 生成 svg
go tool pprof -svg cpu.prof > cpu.svg
# 3. 使用 chrome 打开 svg 文件即可
```

- 根据瓶颈进行性能优化: https://github.com/Deleplace/forks-golang-good-code-bad-code/tree/performance
    - Fast custom trim func, to remove the space character only.
    - Use bytes.HasPrefix.
    - regexp.MustCompile is exactly what we need here.
    - Instead of regexp, use a loop: 10x speedupgap .
    - bytes.IndexByte is more appropriate here.
    - Small parseLine and findSubfields refactoring, same perf.
    - Remove startWith, call directly bytes.HasPrefix: slightly faster.

- 协程使用(调度)优化: 5k message + 20/100 协程
    - 20 协程: https://github.com/Deleplace/forks-golang-good-code-bad-code/commit/a2dfc2a6e8397ae1a3dd6f4be19786ebb45008be
    - 100 协程: https://github.com/Deleplace/forks-golang-good-code-bad-code/commit/26dbf25f2d01c96002a0e9ba66210a9e58ebbbbe

- 到此, 已经优化达到的效果

![](https://images2018.cnblogs.com/blog/420532/201807/420532-20180711200446618-288843359.png)

- 还能不能再过分一点: 能, 用 `Lexer + Parser` https://github.com/Deleplace/forks-golang-good-code-bad-code/tree/lexerparser

## 写在最后
> 总算把一个很久很久之前的坑给填上了, 开心😺

- prof 相关: 可以定位热点函数, 方便定位瓶颈

```sh
go test -bench=. -cpuprofile cpu.prof # 压测, 生成 prof 文件
go tool pprof -svg cpu.prof > cpu.svg # 使用 prof 工具, prof 转为 svg, svg 可以使用 chrome 打开
```

- trace 相关: 可以查看 cpu 使用状态

```sh
go test -bench=. -trace trace.out
go tool trace trace.out
```

- goroutine 相关

首先要区分 CPU密集型任务/IO密集型任务, 协程更适合处理 **IO密集型任务**, 减少 IO wait 导致的 CPU 空转, 其次协程过多会导致协程调度的开销, 同样会造成性能损失

- 推荐使用 github desktop

切换分支, 查看 commit, so easy ~