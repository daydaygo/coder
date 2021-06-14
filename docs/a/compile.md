# compile 编译原理

- [Go程序编译流程](https://golang.design/under-the-hood/zh-cn/part1basic/ch02life/compile/)
  - 词法分析 src->token; 语法分析 src->ast(元素节点=表达式+申明+陈述+错误/调试信息)
    - main -> gc.Main -> amd64.Init -> amd64.LinkArch.Init -> typecheck -> typecheck -> saveerrors -> typecheckslice -> checkreturn -> checkMapKeys -> capturevars -> typecheckinl -> inlcalls -> escapes -> newNowritebarrierrecChecker -> transformclosure
  - 语义分析: 逃逸分析 变量捕获 函数内联 闭包处理
    - ast类型检查: 名称解析 类型推断
    - ast变换: 类型信息细化 死代码消除 函数调用内联 转义分析
  - SSA生成
    - ast->ssa ssa传递与规则
    - initssaconfig -> peekitabs -> funccompile -> finit -> compileFunctions -> compileSSA -> buildssa -> genssa -> typecheck -> checkMapKeys -> dumpdata -> dumpobj
  - 机器码生成
    - 底层ssa和架构特定的传递 go函数->`obj.Prog`
    - 生成机器码

- ast abstractSyntaxCode 抽象语法树
- ssa staticSingleAssignment 静态单一分配
- rt runtime 运行时
- argc argCnt; argv argVector
- tls threadLocalStorage 线程本地存储
