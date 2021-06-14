# coder| 支付系统设计

## 巨人的肩膀
> In me the tiger sniffs the rose. —— 英国诗人西格夫里·萨松代表作《于我，过去，现在以及未来 》

- [凤凰牌老熊 - 支付老兵的 blog](http://blog.lixf.cn/)
- [龙果开源支付系统](http://git.oschina.net/roncoocom/roncoo-pay)
- [Ping++ - 聚合支付](https://www.pingxx.com)

## 从入门到精通：支付系统设计
> [ping++ - 从入门到精通：关于支付系统设计的 17 篇文章](http://www.jianshu.com/p/954872cba5d2)

### 趋势

支付的本质，是价值的转移

聚合支付：

- 原因：支付账户和银行账户垂直交叉
- 服务(>=)：支付通道服务、集合对账服务、技术对接服务、差错处理服务、运行维护服务以及其他增值服务等服务内容

![](http://qiniu.dayday.tech/pay-system-design/pay-account)

支付服务其实需要一套完整的业务系统，**包括且不限于账务系统、账户系统、路由系统、风控／反欺诈系统、运营系统等，要为商户提供完整的“一站式”支付服务**

企业级支付突围：

- 加强支付渠道能力，提供更完善的基础支付服务
- 拓展支付场景，提供场景下的支付解决方案：红包、打赏、平台分账、增值服务
- 提供安全稳定的支付服务

支付市场未来格局：

- 电子支付账户的多维化呈现不可逆转之势
- 利用自身的技术与服务集成能力，越来越多的技术供应商开始着眼于聚合支付领域
- 聚合支付的系统安全技术已日益成熟，让企业无论大小，都能获得安全、稳定、便捷的支付系统

### 支付系统的专业与商业

支付的商业化探索：

- 从功能到场景的转换：广告、粉丝经济、社区电商、平台类商城
- 支付场景的构建：账户体系、营销功能、层级管理、账务管理
- 收费服务：好的服务一定是需要利润支撑的、低价值的数据增长会对发展带来致命的影响、付费是对可从事商业模式的认可

**从功能到场景的转变：**

- 支付方式：网关、快捷、支付账户、信用分期；支付账户 和 银行账户
- 支付服务：身份认证、银行卡信息认证、海关报关、物流信息查询
- 场景：C2B2B、C2B2C、B2B2C、B2B2B；功能 -》组织成场景 -》用户黏性和付费意愿 -》创造交易
- 互联网的消费中，有很大比例的冲动型消费，通过优惠券可以有效的刺激用户的消费意愿
- 账户系统（应该具备）：现金账户系统，卡券账户系统，支付系统和帐号系统

支付营销：

- 用户对于货币电子红包的接受度已经达到了 **88.3%** 的高度
- 没人再敢拿互联网的流量红利来说事了

![](http://qiniu.dayday.tech/pay-system-design/pay-operation1)

![](http://qiniu.dayday.tech/pay-system-design/pay-operation2.jpg)

### 不同行业支付系统设计

映客APP：

- 区别化对待的优惠机制：从优惠策略层面引导用户关注微信服务号，鼓励用户进行大额充值（还有首充）
- 微信红包提现（当时没有手续费）
- 设置不同账户出口：消费、提现
- 用兑换比例控制运营成本
- 紧扣用户消费线的经验值系统

![](http://qiniu.dayday.tech/pay-system-design/inke)

教育行业：

- 互联已经形成，支付还会远吗？
- 体验机制：小额定金 / 免费试课；教育是一个双向选择、互相信任的过程；支付行为本身就能让平台的存在感骤增
- 饥饿营销：限量预售；让预售行为本身就带有自传播性
- 优惠券机制：打造口口相传的熟人效应；最重要的获客渠道之一
- 重线上，规范线下流程：线上专家、拜师，线下资质认证、退款保障等

社区 + 消费（母婴）：

- 想法很美好，现实很骨感：市场大，但是做好难
- 先做社区，在做电商：打赏 -》 余额 -》 不能提现，只能消费
- 抱团：非标产品的促销与拼团机制结合起来
- 用会员制建立核心用户群：‘专家组’

旅游 app：

- 利用出行场景的周期性，增强消费转化：APP 的平均打开间隔和日启动次数可以看出其使用场景和用户特性
- 支持第三方支付账号登录
- 绑定渠道，争取渠道用户流量：‘银联专属’
- 利用互联网优势，打造消费体验生态链：优惠卡券前置，引导用户消费行为；多个账户体系，满足多样人群；加强支付环节体验设计

![](http://qiniu.dayday.tech/pay-system-design/consume-flow)

## 「支付力量」系列

嘛，带些软广，就不列举了

## 凤凰牌老熊 - 支付系统设计系列
> [支付系统架构](http://blog.lixf.cn/essay/2016/08/08/payment-arch/)

### 支付的典型架构

- 某团

![](http://qiniu.dayday.tech/pay-system-design/arch_meituan.png)

- 某Q旅游公司

![](http://qiniu.dayday.tech/pay-system-design/arch_qunar.png)

- 某东金融

![](http://qiniu.dayday.tech/pay-system-design/arch_jd.png)

- 业界最强的某金服金融

![](http://qiniu.dayday.tech/pay-system-design/arch_alipay.png)


## 龙果开源支付系统
> http://git.oschina.net/roncoocom/roncoo-pay

- 架构

![](http://qiniu.dayday.tech/pay-system-design/roncoo-architecture.jpg)

- 支付时序

![](http://qiniu.dayday.tech/pay-system-design/roncoo-pay-flow.png)

## sfy 支付系统升级

### 支付系统现状

sfy 支付相关整理：使用**百度脑图**进行整理

相关问题：

- 没有接查询接口进行支付结果确认
- 支付后执行的逻辑没有提取出来，在同步和异步函数中都写了一遍
- 支付后执行的逻辑全到写到支付 sdk 中，需要 sdk 自己判断支付场景
- 项目依赖 `WEB_LIB` 常量来加载第三方 sdk（短信、支付、二维码）

**数据库事务** 是必须要加的，除了防止支付渠道的异步、同步同时到达，也要防止业务系统内部的并行操作

支付场景：

- 分期订单首付（pay-www-v2、pay-mobile-v2）：连连wap、连连web、alipay（即时到账）、wxpay（扫码支付）、alipay-mobile（wap 支付）
- 还款（bill-mobile）：jdpay-mobile、alipay-mobile
- 优惠券（mobile）：alipay-mobile

支付方式交互字典：

- `shoufuyou_fas.FasPaymtCmpyAcntInf` 以此表的信息作为 `业务 - 支付系统 - 账务系统` 间支付方式交互字典

## 设计方案：独立出支付系统

支付系统当做支付渠道，包含一套 sdk 可以供业务系统使用（类似shoufuyou-sdk），sdk 包括 4 部分内容：下单、同步、异步、查询（这个可以延后再做）

需要考虑 2 点：

- 支付网关前置、后置：采取后置，这样业务系统只用给支付系统发起请求，用户在支付系统提供的页面选择支付方式
- 支付后的逻辑由业务系统还是支付系统处理：支付成功后支付系统通知业务系统，由业务系统完成支付之后的逻辑

最终细化设计如下：

- 业务系统只用展示 ‘去支付’ 按钮，并传递相应支付参数和用户可以使用的支付方式给支付系统，支付系统验证后生成 PayOrder，支付系统按照业务系统传递的 `order_no` 进行去重
- 支付系统展示收银台（**设计 + 前端**）
- 根据用户选择的支付方式，调起支付渠道
- 接收支付渠道的同步和异步通知，解析出支付结果，更新 PayOrder 后转发给业务系统
- 对接支付渠道查询接口，异步通知业务系统结果

- 安全：和业务交互的接口都进行验签

- 支付流水记录：`PayOrder` 跟踪一笔支付单全部生命周期

![](http://qiniu.dayday.tech/pay-system-design/pay-system-flow.png)


![](http://qiniu.dayday.tech/pay-system-design/roncoo-pay-flow.png)

**盗用一下别人的图，抽空再练练 Axure**

## 支付系统前期工作：代码拆分，进行分层

### 分层设计

- 业务系统：各业务系统 -》biz_common -》php-lib
- 支付sdk：pay-sdk -》pay-lib（主要 PayClien.php，业务系统调用这个来接入支付系统）-》php-lib
- 支付系统网关：pay-gateway -》pay-lib -》php-lib
- 支付系统后台任务：pay-task -》pay-lib -》php-lib

确定 **代码分层依赖关系** 后，实现起来的难点主要在于一个接一个业务项目的 **拆分** 代码，并清理支付方式中的业务代码

## 支付系统网关 - shoufuyou-pay-gateway

业务系统和支付系统API交互约定规则：

- 返回：json 格式，6 位 code + message，成功返回 `000000`，失败返回 `28xxxx`（支付系统错误码段）
- 验参：必填参数、参数长度（大部分参数会落地到存储层），返回第一条匹配失败的信息
- 验签：添加 sign 字段进行签名、验签

### 收单接口

- 业务系统向支付系统收单接口发起请求
- 支付系统收单接口 验参、验签、去重（逻辑待定），验证通过后生成 **支付流水记录**（PayOrder）
- 展示支付方式给用户，用户选择后，通过支付渠道 sdk 调起支付
- **注意**：只有一种支付方式时，直接通过支付渠道 sdk 调起支付，不显示支付方式选择页面

### 同步接口

- 支付系统接收到支付渠道的同步通知，通过支付渠道 sdk 获取支付结果，更新支付流水记录
- 支付系统跳转到业务系统下单时提供的 `return_url`，并带上支付信息

### 支付系统 - 异步接口

- 支付系统接收到支付渠道的异步通知，通过支付 sdk 获取支付结果，更新支付流水记录
- 支付系统向业务系统发送异步通知

### 查询接口（**可以延后**）

业务通过下单使用的单号，调用支付系统支付查询接口获取支付结果
业务系统和支付系统属于内网交互，可靠性更高，可以考虑不做

## 支付系统后台任务 - shoufuyou-pay-task

### 查询支付渠道支付结果

- 通过支付 sdk 查询接口获取支付结果，更新支付流水记录
- 获取结果后向支付系统发送异步通知（需要和之前的异步通知进行去重处理）

## 数据层设计

增加 PayOrder 用来跟踪支付单

### PayOrder

支付单：

```
DROP TABLE IF EXISTS PayOrder;
CREATE TABLE `PayOrder` (
`id`  BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT ,
`pay_channel_list` VARCHAR(64) not null DEFAULT '' COMMENT '业务传递的支付渠道',
`choose_pay_channel` VARCHAR(25) not null DEFAULT '' COMMENT '用户选择的支付方式',
`biz_order_amount` INT(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单金额，单位分',
`biz_order_no` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '业务系统订单号',
`biz_order_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '订单名称',
`pay_no` VARCHAR(64) not null DEFAULT '' COMMENT '支付系统单号',
`channel_pay_no` VARCHAR(64) not null DEFAULT '' COMMENT '支付渠道单号',
`return_url` VARCHAR(255) not null DEFAULT '' COMMENT '同步url',
`notify_url` VARCHAR(255) not null DEFAULT '' COMMENT '异步url',
`pay_status` enum('initial','wait_pay','success', 'fail') NOT NULL DEFAULT 'initial' COMMENT '支付结果',
`order_status` enum('initial','finished','notify_finished') NOT NULL DEFAULT 'initial' COMMENT '支付单状态',
`query_cnt` SMALLINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '查询支付渠道次数',
`created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
`expiry_time` datetime NULL DEFAULT NULL COMMENT '过期时间',
`paid_time` datetime NULL DEFAULT NULL COMMENT '支付时间',
`biz_content` text COMMENT '创单数据',
PRIMARY KEY (`id`),
KEY `biz_order_no` (`biz_order_no`) USING BTREE,
KEY `pay_no` (`pay_no`) USING BTREE,
KEY `channel_pay_no` (`channel_pay_no`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

biz_content - 包含不同支付渠道需要的额外字段:

- 业务系统中的用户id user_id，以下连连支付（wap+web）需要
- 用户姓名 user_name
- 身份证 id_no
- 手机号 phone
- 出发机场三字码 airport_from
- 到达机场三字码 airport_to

## 日志

日志实现 `PayLog`，打在关键节点

- 支付系统：收到业务系统下单请求、给支付渠道发送下单请求、支付渠道同步请求、支付渠道异步请求、调用支付渠道查询接口
- 业务系统：发起支付、同步通知、异步通知

## 补充说明

- 更新了 wxpay、alipay、alipay-mobile 的 sdk，新版 alipay-mobile 支持手机上调起支付宝钱包，不需要传递 `show_url` 字段
- llpay、llpayweb、wxpay、alipay-mobile、jdpay 增加了查询接口
- 微信支付目前是二维码支付，更新使用新的二维码生成工具
- 同步通知是同步执行的，用户需要等待业务系统处理完业务逻辑，才能看到支付结果

## php 编程相关

### 编码规范
> [PHP 标准规范](https://psr.phphub.org/)

部分通过配置 phpstorm IDE `code style` 实现（快捷键 `ctrl-alt-l`）

遵循的编码规范：

- 类的命名 必须 遵循 StudlyCaps 大写开头的驼峰命名规范；
- 类中的常量所有字母都 必须 大写，单词间用下划线分隔；
- 方法名称 必须 符合 camelCase 式的小写开头驼峰命名规范。

- 代码 必须 使用4个空格符而不是「Tab 键」进行缩进。
- 每个 namespace 命名空间声明语句和 use 声明语句块后面，必须 插入一个空白行。
- 类的开始花括号（{） 必须 写在函数声明后自成一行，结束花括号（}）也 必须 写在函数主体后自成一行。
- 方法的开始花括号（{） 必须 写在函数声明后自成一行，结束花括号（}）也 必须 写在函数主体后自成一行。
- 类的属性和方法 必须 添加访问修饰符（private、protected 以及 public），abstract 以及 final 必须 声明在访问修饰符之前，而 static 必须 声明在访问修饰符之后。
- 控制结构的关键字后 必须 要有一个空格符，而调用方法或函数时则 一定不可 有。
- 控制结构的开始花括号（{） 必须 写在声明的同一行，而结束花括号（}） 必须 写在主体后自成一行。

违反的编码规范：

- PHP代码中 应该 只定义类、函数、常量等声明，或其他会产生 副作用 的操作（如：生成文件输出以及修改 .ini 配置文件等），二者只能选其一；（sdk 里面 12、13 年的代码）

### 自动加载
> http://docs.phpcomposer.com/

- 项目中 composer 版本为 1.0-dev，我本地环境为 1.2，导致生成的文件不同；
- 修改 `composer.json` 文件中 `autoload` 后，执行 `php composer.phar.php aumpautoload` 来生成类的映射文件

### 第三方依赖
> https://packagist.org

- [QrCode](https://github.com/endroid/QrCode): 生成微信二维码时使用

### 开发环境 docker + docker-compose

### 使用 RAML 写 api
> http://raml.org/
