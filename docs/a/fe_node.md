# node

- express koa egg

```js
$('.thunderhref').each(function() {this.value})
```

## npm

```sh
# npm 加速 http://npm.taobao.org
npm i -g nrm --registry=https://registry.npm.taobao.org
nrm ls
nrm use taobao

# update
npm i npm@latest -g # -g global

# permission
npm config get prefix # npm's dir
npm config set prefix '~/.npm-global'

# locally package
npm i lodash # or change package.json
var lodash = require('lodash'); used in js

# package.json
npm init
npm set init.xxx
npm i xxx --save # dependencies
npm i xxx --save-dev # devdependencies
npm i
npm rm xxx
npm update

# publish
npm publish
npm version patch # README displayed by the version

http-server -p 8080
```

## yarn
- yarn 社区: we don't want to work for you, we want to work with you

```sh
yarn global add xxx
yarn add -D vuepress # add to dev
yarn docs:dev # package.json > scripts
```

## node

- js *关联数组*
- Buffer：js 原生只有字符串数据类型，没有二进制数据类型，但是处理 TCP流、文件流 时必须使用二进制数据
- Stream：抽象接口，比如 http.request、stdout；4 种类型 RWDT；都是 EventEmitters 实例，常用事件有 data、end、error、finish
- 函数：函数作为参数传递、回调 事件驱动 将函数作为参数传递
- 函数式编程 -》 行为驱动执行 -》 router -> requestHandlers
- 全局变量：`__filename, __dirname`；全局函数：`setTimeout(), clearTimeout(), setInterval()`；全局对象：`console, process, Process`
- 模块系统：require + exports + npm
- 路由：url + queryString 模块
- http: method + url.parse + querystring
- 常用工具： util 模块、OS、Path、Net、DNS、Domain
- 文件系统：fs 模块；所有方法都有 同步 + 异步 版本
- 获取 get、post 请求：难道不能用其他 method 方法
- 处理 POST 数据：requestHandler 上添加 listener，监听 data 和 end 事件
- web 模块：http 模块；web 应用架构 client -> server -> business -> data
- 依赖注入来实现 server 模块组合 router 模块
- 应用的不同模块分析，比如：http 服务器、路由、请求处理程序、请求数据处理能力、视图逻辑、上传功能
- Express 框架：nodejs web 应用框架，中间件 + 路由表 + 模板渲染html
- restful api：Representational State Transfer；GET、PUT、DELETE、POST；软件架构风格，设计风格而非标准
- nodejs 多进程：exec、spawn、fork
- JXcore：基本不需要对你现有的代码做任何改动就可以直接线程安全地以多线程运行

## socket.io - ws
> https://socket.io/docs/
> https://hyperf.wiki/2.0/#/zh-cn/socketio-server

- socket: connect/disconnect sid(私聊就是给 roomid=sid 的房间发消息)
- namespace
- room join/to
- emit-event onEvent

## deno
> https://deno.land