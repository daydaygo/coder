# kotlin

- why
  - Modern, concise and safe programming language; 减少模糊, 显式指定; 语法糖非常多, 非常适合 ide+coding
  - use: mobile BE=serverSide FE=webFrontend android dataScience
  - ide: idea androidStudio
  - framework: spring ktor
  - engine: netty jetty tomcat
- syntax
  - type
    - Byte/Short/Int/Long(L) Float(f/F)/Double 进制(不支持八进制) 无符号整形UByte
    - Unit=void=`omitted` 显式类型转换`toXxx()`
    - char: 不能直接当做数字 支持Unicode转义
    - Array: `arrayOf()` 原生类型数组(ByteArray IntArray)
    - string: `"${s1.replace("is", "was")}"` 模板vs拼接 多行vs嵌入'\n' trimIndent-trimMargin
      - `s[i]` `for-in` `+` 原始字符串(""") `trimMargin()` 字符串模板
    - list: `val list = listOf("a", "b", "c")` `setOf()`
    - map: `val map = mapOf("a" to 1, "b" to 2, "c" to 3)`
  - var
    - 默认为val不可变变量 var可变变量 `val a: Int = 1`
    - lazy`var p: String by lazy { /* 计算该字符串 */ }`
    - 交换变量`a = b.also { b = a }`
  - op: nullable`Int?` typeCheck`is` nullElse`?:`
    - bit位运算: shl/shr ushr and/or/xor/inv
  - expr
    - lambda: filter sortedBy map forEach `{it}`
  - ctl
    - if`fun maxOf(a: Int, b: Int) = if (a > b) a else b`
    - for-in while `when-is-else->` range`1..10 step`
    - label`loop@ break@loop`
  - fn
    - Single-expression单表达式`fun sum(a: Int, b: Int) = a + b` + when
    - filter`list.filter { x -> x > 0 } list.filter { it > 0 }`
    - ext扩展`fun String.foo() {}`
    - TODO()`fun a(): Any = TODO("todo")`
    - 作用域函数: apply/with/run/also/let
  - oo
    - Properties`class Rectangle(var height: Double, var length: Double)`
    - new`val rectangle = Rectangle(5.0, 2.0)`
    - extend`:`
    - DAO`data class` singleton`object` with调用一个对象实例多个方法 apply配置对象属性
    - constructor构造函数 init初始化代码
    - Any超类 super调用超类 open开放成员(属性/方法)
- echo
  - <https://kotlinlang.org/> <https://www.kotlincn.net/>

```kt
// try with resources
val stream = Files.newInputStream(Paths.get("/tmp/test.txt"))
stream.buffered().reader().use { reader -> println(reader.readText()) }
```

## doc

- overview: used
- what's new
- basic: synctax idiom codingConvention
- concept: type controlFlow pakage/import oo fn null = this co dsl annotation destruct reflect
- multiplatform
- platform: jvm js native
- release/roadmap
- stdlib: collection scopeFn optIn
- lib: co serial
- api: stdlib test co ktor
- pl: keyword/op grammar special
- tool: build=gradle idea(snip codeStyle) compiler(kotlinc .kt .kts)
