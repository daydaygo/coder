# coder| 支付系统0x00: 支付系统预研

最近在写支付系统, 之前公司写了一版, 量级比较小, 纯同步, 应用层就简单的 api + task, 项目结构也简单:

```
├── lib
├── pay-lib
├── pay-gateway
└── pay-task
```

- lib: 用来存放项目核心依赖, 包括框架和 Util(工具助手类)
- pay-lib: 包括支付公共逻辑, 包括支付渠道 sdk, 数据库模型(model)
- pay-gateway: 支付网关, 用来处理所有 http 请求相关的内容, 已实现 普通支付(支付宝等) / 支付结果查询 / 代扣代发(包括批量) / 部分对账获取 等功能
- pay-task: 支付脚本, 已实现 异步通知 / 查询 / 结算 / 对账 等功能

虽然当时在开发的时候, 也参考了不少资料, 并出了一份设计文档(之前的 blog - [支付系统设计](http://www.jianshu.com/p/05300f67cccc)), 但是当时并没有好好阅读源码, 所以有点浅尝辄止. 这次重拾支付系统, 自然是想做得 **更好一点**, 不然也就同 [rango](http://rango.swoole.com/) 说的那样了:

> 无法进入一个拥有大规模并发请求的项目中得到历练，不坚持提升自己，那也只能在小公司混日子了。

这次的技术预研花了不少时间, 一方面是由于现在开源出来的系统, 都是基于 java 的, 自己确实 java 水平有限. java 的语法看了不止 4 遍(学校里的 java 课, [慕课网](http://www.imooc.com/) java 视频教程, 以及 [<算法>](http://www.ituring.com.cn/book/1807) 这本书是用 java 实现的), 不过离开学校之后就再没有用 java 写过项目. 有一次看 java 同事改请求地址在改动 xml 文件, 当时我还特别纳闷, 毕竟 php 的逻辑是 **什么都可以 php 文件来处理, 最典型的是直接在 php 文件中写 html**. 所以为了能看这些开源系统的源码轻松一点, 花了不少时间 **继续** 熟悉 java 的语法, 并快速看了几个主流框架的简单例子, 差不多能看懂大部分 **黑话**(各种词汇缩写) 的程度为止. 另一方面就比较客观了, **量变到质变**, 没有足够的积累就是做不到想要的程度.

blog 的脉络:

- 通过教程, 尝试阅读 java 源码
- 龙果支付系统源码解读
- XxPay 支付系统源码解读
- 凤凰牌老熊的支付系统设计
- php 自研支付系统设计

## 教程: java SSH 项目使用 dubbo 进行服务化改造
> 龙果学院 - Dubbo视频教程（Dubbo项目实战）: http://www.roncoo.com/course/view/f614343765bc4aac8597c6d8b38f06fd

拿项目中比较常见的用户注册登录功能来说, 如果量级比较小, 注册登录只是整个项目中的一个模块(module), php 各类大型框架(laravel, yii, thinkphp 等) 也都是这样设计, 而且开箱即用, 功能全面. 这方面可以参考 laravel 框架, 功能点很全面, 包括 认证(authentication)/授权(authorization)/密码安全(password)/流量控制(rate limit) 等.

当然, 如果量级比较大, 就需要一个单独的用户(账户)系统了. 独立出来的好处很明显, 比如统一登录, 这样可以踢人下线, 比如安全和风控. 独立出来作为单独的系统, 其实就很接近现在流行的 **微服务设计**.

回到正题, 罗列一下简单思路(恕我 java 水平有限, 欢迎指正):

- 原项目: user-demo, 包含用户注册登录等功能
- 公共基础项目: user-common(user-base), 包含公共基础功能
- Facade 项目: user-facade, 用户服务接口,  这个需要了解一下 facade 设计模式, 熟悉 laravel 对这个应该不陌生, 几乎框架提供的所有服务都使用 facede 进行了一层封装
- service 项目: user-service, 用户服务实现, 服务提供者(provider)
- boss 项目: user-web-boss, 服务消费者(consumer)

更近一步, 基础部分(user-common)可以做的更彻底:

- common-parent: maven 父配置
- common: 公共工程
- common-config: 公共配置工程 -> 以后可以演化为配置中心
- common-core: 公共 core 工程
- common-web: 公共 web 工程 -> 接入层更细粒度的控制

当然少不了服务化的基础设施:

- dubbo: 上面 facade-service-boss 的拆分, 就来自于 dubbo 架构中 provider(container)-registry-consumer 这样的设计

![dubbo architecture](http://dubbo.io/images//dubbo-architecture.png)

- zookeeper: 服务注册与发现

![zookeeper service](http://zookeeper.apache.org/doc/trunk/images/zkservice.jpg)

> 一点个人观点: 要深入一个工具可能需要与时间长跑, 但是, 先用起来吧, 体验一下技术的畅快.

这里举一个之前开发 tcp server 的例子, 当时服务注册发现这块使用的 `etcd`, 我这里只需要在 server 启动的时候, 发一条 http 给 etcd 注册一下即可:

```php
use Swoole\Server;

$s = new Server("0.0.0.0", 9501); // 这就是下面的 gameServer
$s->set([
    'work_num' => 1,
    'task_work_num' => 1,
]);
$s->on('connect', function (Server $s, $fd) {
    echo "connect| $fd \n";
});
$s->on('start', function (Server $s) {
    // server 启动时注册到注册中心
    // curl -vvv http://10.0.1.48:2379/v2/keys/gameServerList/$zoneId/$nodeId -XPUT -d value="x.x.x.x:9501"

    // 获取 gameServer
    // curl http://10.0.1.48:2379/v2/keys/gameServerList/1/1
});
...
$s->start();
```

是不是 **hin简单**?

> 此亦无他, 唯手生尔.

## roncoo-pay: 龙果开源支付系统源码解读
> 龙果支付系统视频教程: http://www.roncoo.com/course/view/a09d8badbce04bd380f56034f8e68be0

首先参考的源码是 [龙果支付系统](http://www.roncoo.com/course/view/a09d8badbce04bd380f56034f8e68be0):

- 核心支付流程

![核心支付流程](http://static.roncoo.com/images/PmwcQyNNrRDxEDTkHswXCSHKKYHTzQzk.png)

下面是代码结构:

```
├── LICENSE
├── README.md
├── UPDATELOG.md
├── database.sql
├── pom.xml
├── roncoo-pay-app-notify
├── roncoo-pay-app-order-polling
├── roncoo-pay-app-reconciliation
├── roncoo-pay-app-settlement
├── roncoo-pay-common-core
├── roncoo-pay-service
├── roncoo-pay-web-boss
├── roncoo-pay-web-gateway
├── roncoo-pay-web-merchant
└── roncoo-pay-web-sample-shop
```

- common-core: 公共组件

没啥可说的, 不过列举一下, 以后设计的时候可以参考: config(配置); BaseDao(数据模型基类); 接口公共返回值/分页等请求相关; 数据库枚举类型(enum); 业务异常(BizException); 工具助手类(Util)

这里有个细节: 数据库里的字段用的 string 数据类型, 使用类来存储枚举类型

- service: 核心业务

核心业务实现, 包括 account(账户, 支付系统中标识交易实体, 参与交易和结算) / notify(异步通知商户) / permission(权限管理) / reconciliation(对账) / trade(交易) / user(用户信息/支付方式, 产品信息) 模块

每个模块按照如下方式组织代码: dao(数据层) + entity(实体, 和数据模型对比) + enums(类型枚举) + exception(异常) + service(具体业务实现) + util(工具助手类) + vo()

- app-notify: 异步通知商户

从 db (`startInitFromDB()`) 中读取数据(status 状态 / notifyTime 通知次数), 使用 线程池(threadPool) 和 队列(notifyQueue) 完成异步通知, 最后会写到数据库(notifyPersist)
需要注意的是, 最后**数据库落库并不是直接更新**, 先将数据写入到 ActiveMQ, 再由 db 来消费

- app-order-polling: 查询支付结果

和上面类似, 线程池 + 队列 来查询支付结果
这里的区别是: 订单支付后添加消息到 ActiveMQ, `app-order-polling` 设置了 `PollingMessageListener` 来订阅这个消息, 从而触发上面的查询

- app-reconciliation: 对账

这个项目比较简单, 因为对账流程是固定的: 判断是否对账 -> 获取对账数据 -> 解析 -> 对账流程
值得借鉴的是: 对账的每个步骤都抽象成相应的 `xxxBiz` 类来处理; 解析这一步定义了 `ParserInterface`, 确保对账流程使用的格式一致

![对账流程](http://static.roncoo.com/images/CJZzhFsfiWDhdp4rAfnEhPfzsjHFdyFT.png)

- app-settlement: 结算

这个项目也比较简单, 维护了 `SettThreadPool` 来处理结算任务, 按照 每日/自动 2个维度进行处理

![结算流程](http://static.roncoo.com/images/RA3jrJZxy26sCWkT6RYRazRPcrrF7zxF.png)

- web-gateway: 支付网关

基础的MVC, 项目并没有实现复杂的路由功能, 只包含 单个支付渠道/多个支付渠道

- web-boss: 运营管理后台

也是基础的MVC, 包含 权限管理 + 支付/对账/清算等支付功能的可视化

- web-sample-shop: 模拟商城

下单并对接支付系统

- web-merchant: 商户后台

订单 / 支付 / 结算

- 数据层简单划分

![龙果支付系统数据库](http://qiniu.dayday.tech/roncoo-pay-db.png)

## XxPay开源支付系统源码解读
> XxPay官网：http://www.xxpay.org

另一个参考的开源支付系统时 [XxPay支付系统](http://www.xxpay.org)

> PS: 此系统还在保持更新, 感兴趣的同学可以加群提 feature, 很可能就加到之后的发版计划了哦

XxPay 支付系统在业务上只包含核心的支付流程, 准确说业务功能实现简单, 但是在系统架构实现上提供了更多选择:

- spring-boot-dubbo架构实现
- spring-cloud架构实现
- spring-boot架构实现

而且开发上也引入了 docker 来解决环境问题, 用一句话来总结的话: **更加现代化**. 推荐阅读源码.

```
xxpay-master
├── xxpay4dubbo -- spring-boot-dubbo架构实现
|    ├── xxpay4dubbo-api -- 接口定义
|    ├── xxpay4dubbo-service -- 服务生产者
|    ├── xxpay4dubbo-web -- 服务消费者
├── xxpay4spring-cloud -- spring-cloud架构实现
|    ├── xxpay-config -- 配置中心
|    ├── xxpay-gateway -- API网关
|    ├── xxpay-server -- 服务注册中心
|    ├── xxpay-service -- 服务生产者
|    └── xxpay-web -- 服务消费者
├── xxpay4spring-boot -- spring-boot架构实现
├── xxpay-common -- 公共模块
├── xxpay-dal -- 数据持久层
├── xxpay-mgr -- 运营管理平台
├── xxpay-shop -- 演示商城
```

对应的数据层也非常简单:

![](http://qiniu.dayday.tech/xxpay-db.png)


## 基于老熊的支付系统设计
> 凤凰牌老熊 - 现代支付系统设计 - 基于微服务的实践: http://blog.lixf.cn/

老熊的博客里面干货很多, 还建立了金融产品技术群长期分享干货, 强烈推荐.

## php 自研支付系统设计

综合来看, 实现微服务化, 是 php 自研支付系统设计的正确方向.

> 常规的 web/api/task/mysql/redis 就不赘述, 你不应该只停留在这个水平

### 基础设施

- docker: 开发部署环境的统一, 也是部署微服务的前置条件
- git: 可以采用 [git flow 工作流](https://www.cnblogs.com/cnblogsfans/p/5075073.html)
- gitlab: codereview
- CI/CD: jenkins
- 系统监控: 比如 zabbix
- 日志分析: 比如 ELK, 通过 sequenceID 进行全链路分析
- doc: api doc, 比如 swagger; 开发设计文档, 流程图等
- test: phpunit

更多准备工作可以参考: [重构中的内部准备工作](http://blog.lixf.cn/essay/2016/08/06/microservice-3/)

更多技术要求:

- 多进程, 保证服务可以横向扩展
- 任务队列: 查询支付渠道; 异步通知商户; 请求接收后立即返回结果, 异步处理请求内容
- 消息队列: 比如支付成功后, 推送支付结果到消息队列中, 多个子系统(BI/风控/账务)消费这个消息

### 系统设计

支付产品:

- 支付流程: 参数校验 -> 支付路由 -> 风控 -> 支付 -> 更新订单 -> 发送消息(多系统订阅此消息)
- 系统边界:  支付网关(路由/接口安全) -> 支持产品(渠道无关) -> 支付渠道
- 支付网关参数: 公共参数 + 业务参数 + 业务扩展参数
- 支付路由: 计算因子(渠道是否可以, 银行卡是否支持, 营销, 限额, 费率等); 手工 vs 权重
- 风控决策: 增强验证 + 拒绝 + 落地支付
- 支付产品模块功能: 签约/解约(绑卡); 快捷支付(支付宝等); 代扣代付(单笔批量); 撤销; 退款; 查单; 结算; 对账

账户&账务:

- 三户模型: 客户 - 用户(商户) - 账户(支付账户)
- 账户建模: 基本属性; 账户控制(提现/冻结等); 资金相关; 银行卡/第三发支付信息
- 账户体系(可以根据会计科目制定): 资产类 / 负债类 / 所有者权益类 / 损益类 / 成本类 / 共同类
- 记账: 单边记账 / 复式记账

风控:

- 场景分析
- 数据仓库
- 风控模型

核心业务流程:

- 交易: 商户 - 支付系统 - 第三方支付 -> 同步/异步/查询(确保获取支付状态)
- 退款: 商户 - 支付系统 - 第三方支付 -> 将退款和交易拆开, 可以冗余退款金额等信息得到交易记录中
- 结算: 商户 - 支付系统, 支付系统 - 第三方支付 -> 采取日结的方式, 结算日期或者周期不同, 都可以在日结的基础上汇总
- 对账(差错): 商户 - 支付系统, 支付系统 - 第三方支付 -> 对账类目, 各类目流水, 差错表, 差错池(自动处理隔天问题)

备注: 隔天问题, 由于双方记录的时间可能隔天, 导致对账日切文件差异, 可以通过对比多天(视支付渠道而定, 大部分隔天, 通常设置2天暂存池即可)数据进行消除

## 写在最后

写这篇 blog 的初衷, 多少有点来自于现在开源支付系统 java 一家独大, 但就提供服务而言(服务器领域), php 都应该有一席之地才对. 有个说法对一个语言了解的越多, 就越清楚这个语言擅长(适合)干什么, 所以我认为 php 也完全可以写出很好的支付系统.

题外话, 关于语言的选择, 可以参看 鸟哥的这篇 blog: [关于语言的选择-选易用的](http://www.laruence.com/2012/08/06/2681.html)
