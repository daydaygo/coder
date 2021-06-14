# rust

内存管理 并发模型 实用设施

- why: 学习曲线突兀 底层控制能力 顶尖工具链(包管理/依赖管理/构建工具) 系统层面(memManage dataRepresent co)
  - 让程序员更快乐: 安全 并发 高性能; 程序3大定律=正确+可维护+高效; 2大难题=内存安全+线程安全; 3zen=内存安全+零成本抽象+实用性; LLVM 一次编译到处运行
- env
  - tool chain
    - cargo dep/build
    - rustfmt
    - rust language server complete/errInfo
- syntax
  - type: `:`
  - var: let mut
  - fn: `fn` `->`
  - macro `println!()`

```sh
rustup # install/update rust
rustc # rust compile
~/.cargo/bin # rust toolchain
cargo # rust package manager; cargo mirror: https://mirrors.ustc.edu.cn/help/crates.io-index.html
cargo build/update/run
cargo doc --open
```

- rust 开发实时通信 https://mp.weixin.qq.com/s/p38QdkIksy6fI5jHbpnFsw
- The Rust Programming Language Rust核心团队成员
  - https://doc.rust-lang.org/stable/book/ https://kaisery.github.io/trpl-zh-cn/
  - rust权威指南 2020.6
- the tao of rust, rust编程之道; 张汉东 2019.1 设计哲学 源码分析 工程角度 底层原理
- 深入浅出rust; 范长春 2018.8 基本语法 内存管理 抽象表达 并发模型 实用设施
