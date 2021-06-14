# go| beego 速览

快速学习 beego 官方文档与 demo,  积累 web 应用及 go 程序开发知识

本来想做一期 beego源码解读 的, 不过作为 go 初学者, 代码量还没提上来前, 正应当多看多写, 做源码解读实在火候不够. 这里综合自己学习 beego 的一些感受:

- [beego 官方文档](https://beego.me/docs/intro/)
- [beego demo](https://beego.me/docs/examples/)

融入到自己对 **Web 应用框架** 开发知识的积累中, 以期可以应用到自己参与的 [开源项目 - Swoft](https://github.com/swoft-cloud/swoft) 中.

> web framework 知识图谱: http://naotu.baidu.com/file/a045802b858c57c56c73dd2e7bda50b5?token=fccd16c1458c78bd
> Modern High performance AOP and Coroutine PHP Framework - Swoft: https://github.com/swoft-cloud/swoft

## 架构

说架构的时, 往往会 **面向业务** 进行抽象, 所以经常看到的架构, 大抵都是 **分层/分子系统**, 一个又一个 **豆腐块**. web 应用其实相当 **成熟** 了, 需要哪些功能已经十分清楚. 这个时候, 也许更应该来关注代码.

![The architecture of Beego](https://beego.me/docs/images/architecture.png)

从图中可以看到, beego 架构很简单, 一言以蔽之: **组件化**(或者说 模块化). 一个 web 应用, 其实就是 **一组组件协同工作** 的结果.

[Swoft1.0](https://github.com/swoft-cloud/swoft) 的一大重要重构, 就是实现组件化. 目前在 PHP 框架中, 组件化实现最好的要算 laravel 和 symfony 了.

PS: Swoft 的现状是一个组件一个 git 仓库. 更正规的方式是一个中央仓库(例如叫 framework), 所有的修改都是 fork/push 到这个仓库, 每次发版时, 中央仓库执行一个 build 脚本(使用 [https://github.com/dflydev/
git-subsplit](https://github.com/dflydev/git-subsplit)) , 将 framework 下的代码推到不同的子项目仓库, 并且给每个组件仓库打一个统一的 tag. 最终所有的修改都在 framework 仓库, 组件仓库只读状态, 不接受 pr/push, 只接受 build 脚本的 split. 可以参考的 laravel 的 build 脚本: [https://github.com/laravel/framework/blob/5.1/build/illuminate-split-full.sh](https://github.com/laravel/framework/blob/5.1/build/illuminate-split-full.sh)

组件化后, 可以明显减少重复造轮子, 更好的代码复用. 组件化其实很好的回答下面 2 个问题:

- 依赖如何管理
- 如何分享代码

## 生命周期

快速上手一个框架, 优先关注生命周期, 了解代码的执行逻辑.

![The execution logic of Beego](https://beego.me/docs/images/flow.png)

beego 框架生命周期: 启动应用(app) -> 一次请求进来 -> 路由(route) -> 过滤器(filter, 有的框架使用 Middleware, 作用相同) -> 控制器(Controller) -> 和各模块交互 -> 返回(response)

![Introduction to Beego’s MVC](https://beego.me/docs/images/detail.png)

一次请求的生命周期: context(请求上下文) -> 路由(Route, 静态文件/动态) -> 路由匹配(fixed/regex/auto) -> filter -> Controller(和各模块交互) -> filter -> 返回(response)

PS: beego 中 module 意为 **组件**, 而在 yii 框架中, module 意为 **模块**, 包含一个完整的 MVC 应用, 而组件用的单词为 **component**. yii 的项目中是可以存在多个 module 的, 即 **一个项目多个应用**. 可想而知, yii 这样的功能, 最终会导致一个 **超大规模的项目**.

## app & context

app 通常作为整个应用的抽象. PHP框架中的常见划分是 **web应用** 和 **console应用**. 不过和传统 PHP 框架最大的不同是, 传统 fpm 多进程模式下, 每次请求都需要初始化一次应用. 而在 beego 和 Swoft 下, 应用是常驻内存的, 只需要启动时初始化一次 app 即可.

这也就造成了很多使用 fpm 的 PHPer 不知道 context 这个概念. context, 请求上下文, 包含一次请求的 request/response. 因为 fpm 是多进程的, 每次请求都是独立在每个进程中并且重新初始化一次应用. beego 或 Swoft 只初始化一次 app, 所以需要 context 来作为每次请求的抽象. 这也是为什么传统框架中, PHPer 依然可以使用 `$_GET/$_POST` 等超全局变量, 而在 Swoft 这样基于 Swoole 的框架中, 必须使用框架提供的方法. 其实质就是对 context 封装.

app 和 context 的分离, 其实可说是必然, 2 者有着不同的生命周期.

## 围绕 web 请求

Model-View-Controller 这样的架构可以说是 **深入人心** 了, 以至于不这样的 web应用, 还会不习惯. 现实中的 web应用, 当然不止 MVC 这三个模块.

### url 的解析

url 的解析, 包含 2 个方向:

- 从 url 到业务代码 -> 路由模块(route)
- 反向生成 url -> URL build 模块

路由模块相对复杂, 功能会多一些:

- 在形式上, 支持 fixed/regex/auto 方式进行匹配, 匹配顺序 fixed -> regex -> auto
- 支持 http method, 比如某个请求只允许 GET 请求
- 支持 restful
- namespace, 命名空间, 比如版本控制, 一部分 url 都使用 `v1/` 前缀, 另一部分使用 `v2/` 前缀
- 注解(annotation) 和 自动参数匹配(auto parameter)

### 路由注解和自动参数匹配

先看代码:

```go
// @router /tasks/:id [get]
func (c *TaskController) MyMethod(id int, field string) (map[string]interface{}, error) {
    if u, err := getObjectField(id, field); err == nil {
        return u, nil
    } else {
        return nil, context.NotFound
    }
}

// 省去的路由代码
beego.Router("/task/:id", &TaskController{}, "get:MyMethod")

// 省去的获取参数(路由参数/请求参数)的代码
id := c.GetInt("id")
field := c.GetString("field")
```

使用注解后, 就省去了在路由中的代码, 如果全局都使用注解, 就可以省去路由文件, Swoft 就是如此做的. 至此, 我们可以对比一下路由解析形式的 4 种方式:

- fixed: 所见即所得, 缺点是 url 太多了就需要写很多
- regex: 正则匹配, 不常用, 大部分也只用在路由参数的地方, 比如 `/user/[\d]+/task/[\d]+`
- auto: 自动匹配, yii 中就是使用这种方式, 约定 url 优先匹配 `Controller + Action`, laravel 中的也支持 Controller 级别的自动路由
- annotation: 注解, 既能保持代码的简洁, 又可以保留路由模块的灵活性

> 所以我理想中的路由形式: 大部分可以直接自动匹配, 少部分的特殊需求使用注解, 独立的路由文件可有可无.

路由中使用注解, 要遵守 **约定大于配置** 的思想. 比如上面的示例改成这样:

```go
// @router /mission/:id [get]
func (c *TaskController) MyMethod(id int, field string) (map[string]interface{}, error) {
    if u, err := getObjectField(id, field); err == nil {
        return u, nil
    } else {
        return nil, context.NotFound
    }
}
```

原来我只需要使用 IDE 的文件查找功能, 查找类似 `TaskController` 的文件, 现在就得先全局搜索 `/mission`, 再根据搜索结果定位到代码.

自动参数匹配比较适合在参数比较少的时候使用, 这个就看个人编码习惯了, **加一个参数可以少写一行代码**, 这个功能还是有必要支持的.

### restful

restful 是现在比较流行的一种软件架构风格/设计风格. **以资源为核心**, 配合 http method(GET/POST/PUT/DELETE) 进行资源的管理.

不过现实是大部分情况下, 大家只用 GET/POST. 面对业务进行设计的时, 也很少会 **以资源为核心**, 怎么简单怎么来或者说怎么快怎么来. 而且要发起 PUT 请求, 你需要这样的表单:

```html
<form method="post" ...>
  <input type="hidden" name="_method" value="put" />
  ...
</form>
```

虽然现在 restful 很流行, 很多框架也都标榜 **轻松** 就可以实现. 是否需要用, 还是自己判断.

### URL Building

beego 中使用 URL Building 来反向生成 url, 比如:

```go
// route
beego.Router("/api/list", &TestController{}, "*:List")
beego.Router("/person/:last/:first", &TestController{})
beego.AutoRouter(&TestController{})

// URL Building
URLFor("TestController.List") // Output /api/list
URLFor("TestController.Get", ":last", "xie", ":first", "asta") // Output /person/xie/asta
URLFor("TestController.Myext") // Output /Test/Myext
URLFor("TestController.GetUrl") // Output /Test/GetUrl
```

其他框架可能会使用 `url() route()`, 一个表示 url, 一个表示 route(路由), 不过我倾向于 **简单点**, 只用 url. 虽然有些时候用 route 来表达(比如 restful) 会简单一些.

### 其他 web 相关服务一览

- Form validation: 参数校验, 这个模块主要解决一些常用的参数校验, 比如 e-mail, 金额等
- response format: 返回值格式, 常见的有 html(view) json xml jsonp. [Swoft](https://doc.swoft.org/master/zh-CN/http-server/controller.html) 中推荐 **返回的格式类型, 不应该由服务端指定, 而是根据客户端请求时的 Header 里面的 Accept 决定**, 控制器中只用返回数据, 框架自动进行格式转换
- view: 视图, beego 中包含 2 个功能, 模板引擎(temple) 和 分页(pagination). 模板引擎方便组织前端代码, 这样 Controller 中只要返回数据即可; 分页也是常见的功能, yii/laravel 在这方面做得很极致, 不仅有分页, 还有其他页面相关的组件
- Filters: 过滤器, 其他框架中可能使用 Middleware(中间件), 功能类似, 其实就是将一部分功能从业务中隔离, 比如网站开启维护状态, user authentication, log visiting, compatibility switching.
- Flash Messages: 一次性通知消息, 其实是一次性的 **消息 - 订阅** 功能, 比如页面上通知操作成功
- Error Handling: 错误处理, 在 web 请求中, 通常配合 `http code + redirect + error page`
- Session: session/cookie 其实很简单, 无论什么框架, 变来变去就那么几个方法, 更重要的是知道 -- 有什么用; 有哪些方法; 使用不同的驱动

## Model

简单说: **如何和数据打交道**. 其实上面的 `session / cookie / flash message` 等, 都在不同场合下和数据打交道, 这里讲讲数据更原始的地方:

- 和关系型数据库打交道, 如 mysql
- 和文档型数据库打交道, 如 MongoDB
- 和键值型数据库打交道, 如 redis, 通常用作缓存, 也可用作 session

主要说一下和关系型数据库打交道, 通常涉及一下几个方面:

- 对数据库连接进行抽象, 这样可以方便 **切换数据库**, **切换主从** 等
- 对数据的抽象, 比如一个 Model 类对一个张数据表, 从而方便的实现 **CRUD** 功能
- 对 sql 语句的抽象: raw, 支持原生语句的执行; QuerySeter, 对 sql 原语的抽象, 比如 `> = < like` 等操作符; queryBuilder, 拼接 sql 语句

- 支持数据库事务(transaction), 要注意事务的写法, 如果发生嵌套如何写

```go
// beego 中的 transaction
o := NewOrm()
err := o.Begin()
// transaction process
...
...
// All the query which are using o Ormer in this process are in the transaction
if SomeError {
    err = o.Rollback()
} else {
    err = o.Commit()
}
```

- 支持表的关联(relation), 比如在 yii 中, 只要定义一个方法, 就可以在查询当前表时, 将关联表的数据, 自动赋值给当前表对象的属性中

```go
// beego 中的 relation
type Post struct {
    Id    int    `orm:"auto"`
    Title string `orm:"size(100)"`
    User  *User  `orm:"rel(fk)"`
}

var posts []*Post
qs := o.QueryTable("post")
num, err := qs.Filter("User__Name", "slene").All(&posts
```

## 更多 more

想要做好一个框架, 还需要在很多地方下功夫.

### 配置管理

beego 的配置模块非常简单易用:

- 单层配置, 没有复杂的嵌套, 如果希望嵌套, 也推荐 `db.user / db.password` 这样用统一分隔符的写法
- 对不同环境支持
- 通过 `include` 关键字, 加载其他配置文件
- 不同格式(format)支持: ini / xml / yaml / json

配置的划分: basic/app/web/http/session/log. 配置的划分, 在一定程度上提现了架构的设计和模块的划分. 一个的好的框架, 只需要简单的修改配置, 就可以实现不同的功能. **推荐好好看一下配置**.

### 安全

安全也是框架必须提供的功能之一, 一个原因是很多人可能没有积累这一块的知识. 不过安全相关的功能, 通常是散落在框架之中:

- csrf(xsrf): 非 GET 请求要注意是否开启 XSRF 防护. 开启则需要给表单或者 ajax 请求中添加 csrf token
- sql 注入: 使用参数绑定即可解决, 要注意在写 **原生 sql** 时是否可能导致 sql 注入

还有很业务关联性比较强的:

- 图片验证码
- 手机验证码防刷

其他更多安全的内容, 可以参考 laravel/yii 框架的官方文档, 平时也要注意积累这方面的知识.

### 日志

日志也是必备功能, 做好日志, 既可以方便了解系统运行状态, 也方便在出问题的时候还原问题发生的轨迹.

- 用法: 日志通常是全局可用, 一般简单的 `Log::info($message)` 即可
- 日志级别: log level, 标准给日志划分了很多级别, 但实际中, 通常开始都使用简单的划分, info 记录信息, error 需要马上处理
- 日志输入: output, 可以将日志输出到不同地方

yii 框架中日志功能更加精细, 包含 `logger - dispatch - target` 三个角色, 同时日志在 level 的基础上, 还可以细分 category / tag 等. 另外日志可能还需要有的功能:

- flush, 刷新, 比如 1000 条后再输出(落地). **缓冲(buffer)** 的思想可以说在系统设计中比比皆是.
- 切片, 比如说日志按照时间日切, 或者按照大小 10m 一切

### 其他

- Live Monitor: 实时监控
- 热更新: 一方面是开发过程中, 如果修改后就需要去喝杯咖啡, 效率就可想而知了, 在传统 fpm 下, 不用担心这个问题, 但是在应用常驻内存情况下, 或者像 go 这样需要编译后运行, 就很有必要做好热更新了; 另一方面是上线, 传统 fpm 是进程逐步重启, [Swoft 基于 Swoole, 也支持这样的机制](https://wiki.swoole.com/wiki/page/20.html), beego 这部分还在开发中
- 文档: 可以编辑(支持 github 更好); 支持上一页/下一页; 支持文档总目录, 支持当前文档目录(TOC); 支持搜索功能. 当然, 功能之外还需要美观, beego 目前支持的功能不多, Swoole 和 Swoft 文档功能支持还不错
- api doc: 自动生成 api 文档, beego 使用的 [swagger](http://swagger.io/swagger-ui/), 简单且功能满足大部分场景
- 辅助工具(命令行工具): beego 提供了 `bee` 工具, 可以用来快速生成空项目, 生成代码. 类似工具 laravel/yii 框架都有支持, Swoft 则正在完善这部分
- i18n: 国际化, 使用和实现上都简单
- CI, 目前 github 上的开源项目, 基本都在使用 travis CI, 可以参考项目下的 `.travis.yml` 文件, 也可以使用 gitlab Jenkins 等开源工具进行集成
- deploy: 发布, go 程序复制编译的可执行文件即可. 不过推荐还是使用 [docker](https://gitee.com/daydaygo/docker)

其他可以做的, 还有很多, 比如:

- 添加 demo, 学习完文档后可以练手
- toolbox / util 等更多辅助功能
- 第三方(third-part)功能集成

