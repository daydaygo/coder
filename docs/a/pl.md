# 编程语言基础

- todo
  - [已拿下总包60W+腾讯offer！我决定公开这份资源包](https://mp.weixin.qq.com/s/uDgzmNB7Yu8KBOA4X2WtnA)
  - <https://github.com/donnemartin/system-design-primer>
  - <https://github.com/kamranahmedse/developer-roadmap>
  - 每个程序员都该尝试的项目 <https://mp.weixin.qq.com/s/kZtQX21mr2APe25JtkoDZQ>
  - 提升代码力 <https://mp.weixin.qq.com/s/aCcL1P4_UxMfg_sJ8uOr7g>
  - it开源书 <https://csgordon.github.io/books.html>
  - 14 种模型思维: <https://mp.weixin.qq.com/s/gwcQU6KTJqQ_uFpEUIB5kQ>
  - 技术人快速成长: <https://mp.weixin.qq.com/s/EMZrnDRGCoIAloSFr1js5w>
  - 大学如何自学 cs: <https://www.zhihu.com/question/403657999/answer/1585740856>
  - 量化技术团队效能 <https://mp.weixin.qq.com/s/DCpre2j7AkjI4mhQq410_g>
  - 技术方案设计方法论 <https://mp.weixin.qq.com/s/Q94f0Y-lAWjuBrHdNFFIVQ>
  - <https://github.com/chaos-mesh/chaos-mesh>
  - 技术管理
    - 给ITer的技术前沿课.pdf
    - 职业发展黄金手册.pdf
    - 软件技术职业选择之道.pdf
    - 阿里工程师的自我修养.pdf
    - 技术TL: <https://mp.weixin.qq.com/s/U-hqectN-fes7Td6Osut7Q>
    - [关于架构师：角色、能力和挑战](https://mp.weixin.qq.com/s/v_HkBfmGkjHgWKG8oncm6A)

## 计算机资料

- 公众号 小林coding 注意：压缩包我都是加密的，解压密码「xiaolin」。
  - Java 链接： <https://pan.baidu.com/s/1vB2a1GYEVpfpgkk2r7Qaig> 提取码：2f1g
  - C 和 C+- 链接： <https://pan.baidu.com/s/1r_bMH-tcG4FQAK4ZpmgVng> 提取码：wc45
  - Linux 链接： <https://pan.baidu.com/s/1qJigxjzMrJrfFgKPFbS3Jw> 提取码：9a23
  - 操作系统链接： <https://pan.baidu.com/s/1StDTkn-wEGRtcoQSqwGsuQ> 提取码：kmgj
  - 计算机网络链接： <https://pan.baidu.com/s/1D7-Yj5CgdtCshPIyQ4iu4w> 提取码：7c4o
  - 算法链接： <https://pan.baidu.com/s/1DYg1O8PFwu_1c6gbuKGpOw> 提取码：qln8
  - 数据库链接： <https://pan.baidu.com/s/1xrYlstgZDlVVfYXOFxvebQ> 提取码：lkf5
  - 面试题链接： <https://pan.baidu.com/s/1bpEs58x89qZpyUH973GjVw> 提取码：v2jo
  - 设计模式链接： <https://pan.baidu.com/s/14NwqRbQ6AY5nKL9nn8sjaw> 提取码：exix

---

- how: 最大公共子集; show me the code; 学习打基础是确定性, 解决问题是不确定性; [附录/词汇/索引](appendix.md)
- why: for who/what; 前景 能做什么 适合做什么
- common: scope范围(public/protected/private/default 大小写) order顺序 int数值 lifecycle生命周期
  - 通信本质是状态同步; allIsXXX(object file)
  - search->coding->note
  - 静态类型: compile验证类型; 动态类型: runtime解析方法和字段引用
- env环境: install/config run/debug IDE
  - install
  - run: 能命令行跑起来; hello world
    - command line: option; argument; iostream; interactive(ncurses readline); shell tab complete
  - config: env version 命令行参数 配置文件/常用配置
  - debug: 在 run 的基础上, 编程语言的 debug 工具 + IDE 的 debugger 工具
    - 原理+操作: [hyperf | 使用 yasd debug](https://www.yuque.com/docs/share/627e3c89-6b8e-489d-b6b3-f1c1fa2e74ba)
    - [vscode DAP](https://code.visualstudio.com/blogs/2018/08/07/debug-adapter-protocol-website/)
  - ide: [idea](idea.md) [vscode](vscode.md) [vim](vim.md)
- syntax 语法
  - comment 注释
  - reserved key world 保留关键字
  - type 基础数据类型 literalValue 字面量
    - bool
    - int: base进制 `INI_MAX/INI_MIN` unsigned-溢出 可读性`1_000` 结尾`100L` 开头`0X0F`
      - bit: 8-256 10-1024 16-65536-port 32-42亿-溢出 64-太阳寿命100亿
      - math: bit位运算 GMP任意长度 BCMath任意精度 rand
    - float: precision精度
    - string: printf格式`%ns %ni %m.nf` 转义字符 `\`
      - regex 正则: `^ $ . * + ? {m} {m,n} \ [^0-9] | \b` lua使用`%`转义
    - 类型转换: auto force(big->small)
    - code编码: bit->int->string(Unicode/charset) morse case(upper lower)
    - 泛型
    - compound types: array/link stack/queue hash/map tree graph
  - var
    - init-default 初始化与默认值 inferred类型推断
    - predefined预定义变量
    - val&reference 值与引用: 拷贝 深/浅/写时; 引用计算 refcount; reference-`&` 引用
      - explain: C pointer(arithmetic); memory address; symbol table alias; var(reference)-name/content file(hardlink)-name/content
      - do: assign(`&` `new`) pass return unset(unix unlink)
      - type: 强(默认) soft软 weak弱 phantom虚
    - 作用域: global/function/block/const
    - const 常量
  - op operator 操作符
    - precedence 优先级
    - arithmetic 算术: `a++` `++a` `%`mod取模-环/进位/奇偶
    - assignment 赋值: `a=5` `a+=5` `b = &a` `c = new C` `a.=b` `a??=b`
    - bitwise 位运算: `& | ^ ~ << >>` 原码/补码/反码(统一加减法) 二进制->思维
    - Comparison `comparable` 比较: 地址`===` 值`==` `<>` `?:` `??` `<=>`
    - logic: `and/or/xor` `! && ||`
    - string: `+ .` -> stringBuilder
    - type: instanceof `x.(T)`
  - expr expression statement 表达式: `anything that has a value` `$a=5;` `$a=foo();`
    - Lambda
  - ctl control 流程控制
    - if-elseif-else end endif
    - switch-case-break-default
    - for while do-while: 小心死循环
    - foreach fori
    - break continue
    - return
    - goto 慎用
  - fn function 函数
    - main entrypoint
    - function/func scope/visible overload
    - parameter/argument参数: var/default reference`&` arg-list`...$nums` type-declare defaultValue named
    - return: chain 链式调用 `return this`
    - recursive 递归: 100-200level smashStack 尾递归优化
    - built-in/internal func
    - var func: `$a='foo'; $a();`
    - anonymous/callbale/Closure
    - arrow: `a=fn(x)=>x+y`
  - OO OOP class类 object对象
    - class this new/clone(深/浅) visible(public/protected/private internal模块) static(late static bindings)
      - object: objects are passed by references by default; instance实例
      - property属性: const var, attribute/field enum枚举; type init/getter/setter 幕后属性`private var _table`
      - method方法=action行为
    - inherit继承: extends self/parent(super)/static(current class in runtime) Polymorphism多态
      - override重写=不同实现 overload重载=不同输入 final
    - abstract抽象
    - interface接口 封装
    - predefined
    - serialize序列化 ClassLoader类加载
    - [DP](dp.md), design pattern 设计模式: proxy facade
      - container-服务容器 DI依赖注入 IoC控制反转
    - nullsafe op: `A?->foo()?->a;`
    - magic method: `__construct __destruct __get __isset __callstatic`
    - nest嵌套类/内部类; immutable不可变类
  - namespace/package module模块
    - like fs(relative/absolute); a way of encapsulating items; name collisions + alias
    - namespace/package import/use as
    - scope
    - component
  - error 错误 / exception 异常 / log 日志
    - Sadly, no matter how careful we are when writing our code, errors are a fact of life
    - type: debug/info...
    - c: errno + errstr; go: err
    - throw try-catch-finally
    - predefined: error hierarchy
    - error handle & log
  - generator 生成器
    - range-foreach-yield-Iterator
    - generators vs Iterator: simple, no rewound, rebuild
  - attributes / annotation 注解
    - reflect反射: 运行时动态获取对象的类型和值以及动态创建对象 -> 抽象+简化 -> 提高开发效率
    - proxy代理 动态代理
    - declare
    - AOP: spring
- fn, function reference 功能参考
  - stdlib thirdlib第三方库
  - cli, command line: interactive shell(ncurses readline)
  - textProcess: strings(printf `str_xxx` buffer build) rege(Perl PCRE; POSIX regex)
  - code编码: bin/hex charset protocol(bin二进制 text文本) dataFormat
  - dataFormat: JSON/YAML/TOML/xml/html/csv/markdown
  - datetime: [fromat](https://www.php.net/manual/zh/datetime.formats.relative.php); timezone
  - fs fileSystem: dir mine(magic.mime database) fd bigFile
  - event eventDriven事件驱动: event->dispatcher->listener/handler
  - db database: abstract-layer vendor-specific query/statement->queryBuilder->ORM/DAO/model
  - bigData: map/reduce/filter
  - cache: redis memcache http anywhere
  - mq, message queue: rabbitmq/amqp kafka
  - search engine: ES elasticsearch; solr; sphinx
  - crypto: hash(MessageDigest)/openssl/password_hash
  - compression & archive: rar zip zlib(.gz)
  - human language & character encoding: i18n internationalization / l10n localization(gettext); multibyte string
  - image
  - mail: imap pop3 exchange
  - context上下文: 对象和数据分离
  - web: oauth soap rpc route-MVC cors template rfc(rfc1867=fileUpload)
  - net: socket/stream ip(4type icmp dns/ares) tcp(protocol encode/decode pack/unpack) udp http websocket(socket.io)
  - processControl 进程=动态的程序: PCNTL POSIX EIO(libeio, async POSIX io) EV(libev, event loop) IPC(sem.semaphore信号量.PV操作 shm/msg/pipe/socket namedPipe.mkfifo) sync ENV fork processModel(main manager worker task user) 静/动态优先级
  - thread 线程=cpu+执行: sharedMemory+lock cpuExecuteUnit aio async eventDriven
  - co go goroutine 协程: context pool(object connect) channel select defer
  - service/server
    - io模型: BIO阻塞 nio非阻塞(bio+while) ioMultiplexing多路复用(select/poll/epoll/reactor) aio signalDrivenIO
    - c10k(kqueue select/poll/epoll iocp) reactor(add set del callback) sync/async/co
    - curl event(io time signal) ftp LDAP memcache(session) network(dns host ip2long header cookie) session zookeeper
  - other
    - basic: GeoIP stream(`scheme://target` socket transport)/socket tokenizer/parle(lexical level; parser token)
      - urls: base64(mime base64); parse url; urlencode; build query
    - var & type related: ctype(`ctype.h` in C) filter(valid sanitize) `func_xxx` reflection(ide helper); var handle
    - stateMachine状态机: 状态迁移与监听
  - GUI
- project 项目管理
  - code: style/fmt/lint 编码规范; analysis/staticCheck/vet 分析; tran 版本兼容; sonarQube代码质量管理
  - vcs git: monorepo
  - build构建
    - tool: buck/bazel makefile(c) cmake(c++) cocoaPods/Carthage(ios) clang(c family)
    - fn: dep 快速编译/增量编译 tag->platform/use ci
  - dep依赖管理 layout项目组织
  - test TDD: unit(自动生成)-func interface 集成 auto自动化 feature-api mock(远程调用/依赖 数据生成) chaos
  - Performance性能 prof/benchmark压测: gprof/oprofile jdk相关
  - log
  - doc: readme api flow
  - framework框架
  - onlineError: appErrorLog plErrorLog dataLayerError
- arch architect architecture 架构
  - mind
    - lifecyle+run 生命周期+代码跑起来
    - layer 分层 边界 隔离 抽象
    - landscape 全景图/整体
    - tool 工具: 擅长/适合 区分业务/技术
  - ms, microservice 微服务
    - config 配置中心: watch merge env hotload热加载 switch开关
  - cloudnative 云原生
  - devops: [git/github](git.md) container CI(image build; lint; static check; unittest; notfiy)/CD(k8s: vension rollback) coding/阿里云云效(teambition)
    - ops: monitor(sys log flow流量 interface db) alarm/alert告警 log日志
- security 安全
- core 语言核心
  - memory manage
    - Persistent
    - gc Garbage Collection 垃圾回收
      - refcount
      - performance: mem usage; runtime delay/slowdown
  - var: array hashtable object
- eco 生态
  - history product book publication changelog
  - package doc 社区 大咖
- CS basic 计算机基础: 数据结构与算法 网络 OS(file/fd) 编译原理

```sh
# bitwise in conf
E_ALL & ~E_NOTICE

# regex
[0-9]{4}-[0-9]{2}-[0-9]{2} # Y-m-d
[0-9]{1,3}\.[0-9]{1.3}\.[0-9]{1.3}\.[0-9]{1.3} # ip

# crypto
# type: -t rsa; key_format: -m; import: -i; export: -e; change passphrase: -p
ssh-keygen # create
ssh-keygen -e [-f input_keyfile] [-m key_format] # change format
```

## topic

- aop
  - 场景: 日志记录 性能统计 安全控制 事务处理 异常处理 缓存
  - aspect切面
    - `$class` TargetObject joinpoint连接点=pointcut切入点 introduction引介=weaving织入
    - advice通知/增强: `around环绕` before after afterThrowing finally
  - ProxyClass代理类 最终生成代理类来执行切面方法的目标
  - `$annotation`; asm->CGlib(codeGeneration) javasist Instrumentation(premain静态方式 agentmain动态方式)
- gc
  - 算法篇
    - 引用计数法: 互相引用
    - 可达性分析: GCRoot+引用链
    - markSweep标记清除算法: 标记=效率问题 空间问题
    - copying复制算法=newGen: 活动区间/空闲区间 浪费一半内存/对象存活率必须低
    - markCompact标记整理算法=oldGen: 标记=GCRoot遍历 整理=内存地址次序排序
    - 保守式GC 分代GC(old+new) 增量式GC RC-immix算法
  - scavenge收集器: serial串行(单线程=stw=stopTheWorld) parallel并行 并发=CM->GI
  - 实现篇: GC 在 python/dalviKVM/rubinius/v8 等几种语言处理程序中的实现
- log日志
  - logger(producer) -> dispatcher(flush) -> target(consumer)
    - format: time=ms group level category msg context(traceId spanId env pid 分隔符)
    - level级别: debug info感兴趣=用户登录+sql日志 notice重大意义 warning异常=过时api critical关键错误=组件不可用 alert报警=网站挂了+db不可用=sms通知
    - target: console=彩色 file=轮转+自动删除 mongo
- hook 机制
- dialect 方言: 兼容差异的部分
- option/argument
- codec encode/decode 编码
  - ASCII
  - [unicode](unicode.org): utf16(16bit=\uhhhh) int32=utf32=ucs4(32bit=\Uhhhhhhhh) utf8=1-4byte(1byte `0xxxxxxx` nbtye `1110xxxx 10xxxxxx 10xxxxxx` \xhh\xhh)
  - base64: binData -> byte
- timeout 超时

公司项目 和 个人项目分开
项目管理/研发 分离
重构 - 不改老代码/新老校验; ide 的重构功能; 重构前的大纲->重构后的大纲 帮助理清思路, 后续反复查找
产品快速试错与重复博弈，商业模式与纳什均衡

- 工具思维
  - 黑箱思维: IO dep conf->class->component->power example/readme/quickStart
  - mind思维导图; 记忆->知识内化; 笔记; 画图
- 失败的思维方式
  - 高考思维->成事而非分数
  - 被动思维 做什么都需要一个理由->被动vs主动
- 约定大于配置
- 技术分享 - 锻炼演讲

使用框架 -> 懂框架 -> 造轮子
设计哲学 源码分析 工程角度 底层原理
程序: 正确->可维护->高效

## cs知识体系

- UC Berkeley EECS <https://mp.weixin.qq.com/s/P2tNxWQW8nIewvw_jHdhFQ>
- MIT EECS <https://cloud.tencent.com/developer/article/1646448>
- cs自救指北 <https://survivesjtu.gitbook.io/survivesjtumanual/fu-lu/ben-ke-sheng-zhuan-ye-jie-shao-todo/cs-zi-jiu-zhi-bei>
- [技能图谱](https://github.com/TeamStuQ/skill-map)
- net
  - tcp/ip
    - 协议: 哑协议vs智能协议
- 同步/异步 进程/线程/协程 并发concurrent/并行parallel
- lock锁
  - Pessimistic悲观锁
    - spinlock CPU CAS PAUSE慢等待 锁住代码的执行时间很短
    - mutex 线程上下文切换 几十纳秒-几微秒
    - rwlock 读多写少
  - optimistic乐观锁 假设冲突概率低-回滚代价很大 MVVC版本号/CAS算法 write_condition机制 `java.util.concurrent.atomic` ABA问题
  - lockFree: CAS atomic
  - deadlock死锁: 顺序化 baike-数据库死锁
    - [coffmanCodition](http://www.ccs.neu.edu/home/pjd/cs7600-s10/Tuesday_January_26_01/p67-coffman.pdf): MutualExclusion互斥 WaitForCondition持有资源 NoPreemption不可剥夺 CircularWait循环等待
  - 公平=FIFO=队列 vs 抢占=优先

---

- 理解计算机系统: <https://www.nand2tetris.org> <http://csapp.cs.cmu.edu>
- 多动手: 数据结构(B+Tree) 计算机系统结构(编程综合实践 <https://acm.sjtu.edu.cn/wiki/PPCA_2019>)
  - os: <https://pdos.csail.mit.edu/6.828/2019/schedule.html> <http://pages.cs.wisc.edu/~remzi/OSTEP/>
- 编程语言与工具 <https://missing.csail.mit.edu/>
- [freeCodecamp 中文社区](https://learn.freecodecamp.one/)
