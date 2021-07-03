# swoole| swoole wiki(老版) 笔记

- [swoole| swoole wiki 笔记](https://www.jianshu.com/p/12d645ac02b2)

## 初心

PHP想做微服务, 离不开 swoole. 而 swoole 进入协程时代后, 是时候抛开历史的包袱, 轻装上阵了. 希望这样过一份 swoole 知识的梳理, 能对 PHPer 有所帮助.

> 梳理 swoole wiki 的方向建议: 尽量用协程, 实在不行就同步, 所有的知识都围绕这展开.

## swoole 功能一览

重点

- 协程 server(继承关系): server(tcp/udp) http(http/http2) websocket
- 协程 client: client http http2 socket [zookeeper](https://github.com/swoole/ext-zookeeper)
- 协程 runtime
- process: `Server->addProcess()` process process-pool

辅助

- memory: table atomic mma
- timer: tick after; server 中的 Timer 也会自动创建协程
- event
- 同步 client: client

历史版本, 只做了解, 不要造成心智负担(可以不看)

- ext: 使用扩展方式, 保持主项目 focus 协程, 提高稳定性
  - ext-async([非协程特性扩展](https://wiki.swoole.com/wiki/page/p-async_ext.html))
  - [ext-pgsql](https://github.com/swoole/ext-postgresql)
  - lock: 协程中请使用 chan / 协程调度器
- 异步 client: client mysql redis http

## swoole 协程

- [swoole| swoole 协程初体验](https://www.jianshu.com/p/745b0b3ffae7): 辅助理解协程的基本概念
- [swoole| swoole 协程知识点小结](https://www.jianshu.com/p/b620836c461a): 辅助理解 swoole 中的协程知识点
- [swoole| swoole 协程用法笔记](https://www.jianshu.com/p/28e882352da5): swoole 协程 API 梳理/速查

## swoole 协程 server

![运行流程图](https://wiki.swoole.com/static/uploads/swoole.jpg)
![进程/线程结构图](https://wiki.swoole.com/static/image/process.jpg)
![进程/线程结构图](https://wiki.swoole.com/static/uploads/wiki/201808/03/635680420659.png)

- setting
  - [常用配置](https://wiki.swoole.com/wiki/page/13.html)
  - [配置选项](https://wiki.swoole.com/wiki/page/274.html)
  - dispatch_mode: 1-轮询(平均)分配 2-按fd取模(默认) 3-争抢分配 [7-stream](https://wiki.swoole.com/wiki/page/838.html)
  - task
    - task_worker_num
    - [task_ipc_mode](https://wiki.swoole.com/wiki/page/212.html)
    - [task_enable_coroutine](https://wiki.swoole.com/wiki/page/p-task_enable_coroutine.html): task支持同步/协程, API有差异
  - ssl
    - 'ssl_cert_file' => $key_dir.'/ssl.crt',
    - 'ssl_key_file' => $key_dir.'/ssl.key',
  - 固定包头协议自动分包
    - open_length_check => true
    - package_length_type => 'N' // pack()
    - package_length_offset => 10
    - package_body_offset => 120
    - package_max_length => 800000
- [多端口监听](https://wiki.swoole.com/wiki/page/525.html): 一个swoole, 多个server
- [server中对象的4层生命周期](https://wiki.swoole.com/wiki/page/354.html): 注意和传统 fpm 的区别

## swoole 协程 client

- [是否可以共用同一个redis/mysql连接](https://wiki.swoole.com/wiki/page/325.html): 连接池是标配

## swoole 进程 & 进程池

- [Swoole| Swoole 中 Process](https://www.jianshu.com/p/4b6326cdaaa7): 主要内容是进程间通信部分
- `->addProcess()`: [广播消息](https://wiki.swoole.com/wiki/page/390.html)
- process: [使用](https://wiki.swoole.com/wiki/page/221.html)
- process-pool: [需要长期运行的脚本, 比如 mq 消费者](https://wiki.swoole.com/wiki/page/901.html)
- [process+协程](https://wiki.swoole.com/wiki/page/p-process_coro.html)

## swoole more

- [swoole相关函数](https://wiki.swoole.com/wiki/page/548.html)
- 附录: linux信号 linux错误码 swoole错误码 tcp连接状态 tcpdump/strace/gdb/lsof/perf
- [swoole内核开发](https://wiki.swoole.com/wiki/index/prid-11)
  - 开源项目工程实践: 工程哲学 性能优化
  - 版本规划: 严格的版本更新记录; 迭代周期(大小版本); 版本类型/单双数
  - 编码风格
  - 工程结构
- [php-x](https://wiki.swoole.com/wiki/index/prid-15): 使用 c++ 开发 php 扩展

## swoole dashboard

- wiki: https://www.yuque.com/swoole-wiki/dam5n7
- 申请试用: https://www.swoole-cloud.com/dashboard/catdemo/

## 编程指引

- swoole 快速起步 [daydaygo/php-note](https://github.com/daydaygo/php-note/tree/master/swoole): swoole API 简单 demo, 方便速查
- [学习 swoole 需要掌握的基础知识](https://wiki.swoole.com/wiki/page/487.html): linux高性能服务器编程; UNP; tcp/ip详解
- [swoole 编程注意事项](https://wiki.swoole.com/wiki/page/p-instruction.html): 语言相关/协程编程/并发编程/内存管理/进程隔离
- [版本更新记录](https://wiki.swoole.com/wiki/page/p-project/change_log.html): 仔细看看, 就能体会到开发组的努力
- 版本路线路线图: [trello](https://trello.com/b/SEdDCrCu/swoole-kernel-developer) [rfc](https://github.com/swoole/rfc)
- 编程建议(个人看法)
  - 基础知识的补充和扎实, 远比对 API 的熟悉要强, 建议花时间理解: 网络编程基础; 同步/异步/协程基本概念; swoole进程/线程模型; swoole各组件+组件功能+组件生命周期
  - 编程方式: **尽量使用协程, 实在不行就同步**; 好处很明显, 关注swoole功能的子集, 理解和熟悉API 都容易不少
  - 建议不使用别名, 直接使用原写法, 如 `Server->tick()` 是 `\Swoole\Timer::tick()` 的别名
  - 建议除了 `go() chan()` 等十分常用的函数, 不要使用简写, 直接使用原写法, 如 `Co::sleep()` 是 `\Swoole\Coroutine::sleep()` 的简写

## php

- [callback](https://wiki.swoole.com/wiki/page/458.html)
- 信号处理 ext-pcntl
- 原生socket `ext-socket` `stream_socket_client()`
- `memory_get_usage() + memory_limit` 来做内存检测 -> 防止 worker 内存溢出, 配合 `exit()` 退出worker进程, manager 进程会自动拉起新的 worker 进程

## NP(network programming)基础

![tcp 三次握手四次挥手](https://www.swoole.com/static/image/tcp_syn.png)

- [网络通信协议设计](https://wiki.swoole.com/wiki/page/484.html): 为什么需要协议->tcp是流式(stream), 需要协议设定边界; 设计->换行符/固定包头
- [多进程共享数据](https://wiki.swoole.com/wiki/page/836.html): 为什么->进程隔离; ext-apcu/yac(仅作缓存)/swoole memory系列组件
- 进程信号处理: `pcntl_signal() / pcntl_signal_dispatch()`
- 并发编程: 并发执行->为每个客户端/请求, 创造不同资源和上下文(context)
- [内核参数调整](https://wiki.swoole.com/wiki/page/p-server/sysctl.html): `sysctl -a` 查看所有内核参数; `/proc/sys/` 文件夹

## devops

- 实践推荐
  - 项目git -> 绑定阿里云容器镜像服务 -> 关联容器服务(小型项目可以先用 swarm 验证, 大型复杂项目转k8s)
  - git 更新 -> 触发镜像自动构建 -> 镜像构建完后, 触发容器自动部署
  - 开发: 同一份镜像, 开发环境统一
  - 使用容器来管理服务生命周期, 替代传统的服务管理方式([systemd](https://wiki.swoole.com/wiki/page/14.html); [开机启动](https://wiki.swoole.com/wiki/page/19.html);  [`->reload()` 等服务管理函数](https://wiki.swoole.com/wiki/page/p-server/reload.html))
- swoole 安装
  - PHP版本([PHP版本支持](https://www.php.net/supported-versions.php)): 推荐 >=7.2; 只有同步 client 可以在 fpm 下使用
  - 安装实例: [docker - swoft/alphp](https://github.com/swoft-cloud/alphp/blob/master/alphp-cli.Dockerfile)
  - 编译参数: 上面 Dockerfile 中的示例就够了, 还可添加 `--enable-coroutine-postgresql`
- 监控/日志/性能/工具
  - `Server->stats`
  - swoole日志 task日志
  - [ab 压力测试](https://wiki.swoole.com/wiki/page/62.html): 还是那句话, **你用 swoole 达不到的性能, 换个语言, 呵呵呵**
  - [timer压测 10w个 0.08s](https://wiki.swoole.com/wiki/page/p-timer.html)
  - [tcp/udp压测工具](https://wiki.swoole.com/wiki/page/197.html)
- 案例
  - [组播实现](https://mp.weixin.qq.com/s/N8cgZtR-7COZET6zXbJohQ): httpserver 登录服务 + proxy层转发信息 + WS聊天室
  - [tech| 再探grpc](https://www.jianshu.com/p/f3221df39e6f)
