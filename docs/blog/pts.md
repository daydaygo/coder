# devops| PTS 性能测试一览

- [tech| PTS 性能测试一览](https://www.jianshu.com/p/d31155c404ff)

综合尝试了以下性能测试工具, 对性能测试至少不是 **盲人摸象** 的状态了.

- httprunner
- jmeter
- 阿里云pts

## PTS 基础: http 协议

http协议基础知识: tcp/ip 4层网络结构; http协议响应码; http method; http post body format

> 关于 http 协议, 推荐看这篇 [alipay ILLEGAL_SIGN 错误解决](https://www.jianshu.com/p/28585a6454b2), 无论 http 相关的什么难题, 最终都能在 http 协议细节找到答案.

http 的使用可能是编程中常见的技能点了, 会接触到各种 http 协议上的封装的, 比如 wget/curl, 怎么判断这样一层封装到底好与不好呢? 我给出自己的答案 -- 使用频次 + 贴近协议:

- curl: 几乎都会有基于 curl 的封装与实现, 内容也极其丰富, 可以根据使用频次, 总结出最常用的几个放着
- [PSR-7: HTTP message interfaces](https://www.php-fig.org/psr/psr-7/): 根据 RFC 中 http 协议相关内容设计
- [guzzlephp](http://docs.guzzlephp.org/en/stable/): 写法些许陌生, 但这就是 http 本来的样子, 简单举2个例子, 区分 http method + post body format

## httprunner 快速上手

- HttpRuner: https://cn.httprunner.org/
- locust: https://docs.locust.io

- 环境安装, 关键点只有一行:

```
# 安装
pip install httprunner locustio

# 验证
har2case --version # HAR 转换成测试用例
hrun --version # 运行测试用例
locust -V # 性能测试时使用
```

- charles: 录制请求为 HAR(HTTP Archive)

正常 charles 抓包即可, 选中抓到的请求 -> 右键 -> export -> 选中 har 格式

如果遇到 charles 无法抓取 localhost, 请参考这篇: [Localhost traffic doesn't appear in Charles](https://www.charlesproxy.com/documentation/faqs/localhost-traffic-doesnt-appear-in-charles/)

- 使用 har2case 自动生成测试用例

```
har2case test.har -2y # 生成 yml 格式的测试用例
```

- 测试用例优化

自动生成的测试用例如下, 以官网 demo 为例, 可以进行的优化点: 参数提取 + base_url + 使用变量(包括全局变量) + 参数提取 + 函数热加载 + 参数化数据驱动

```yml
- config:
    name: testcase description
    base_url: http://127.0.0.1:5000 # 使用 base_url
    variables:
        app_version: 2.8.6 # 定义变量; 变量可以是当前测试用例, 也可以使公共变量
        device_sn: ${gen_random_string(15)} # 热加载, 从 debugtalk.py 中获取函数

- test:
    name: /api/get-token
    request:
        headers:
            Content-Type: application/json
            User-Agent: python-requests/2.18.4
            app_version: $app_version # 使用变量
            device_sn: FwgRiO7CNA50DSU
            os_platform: ios
        json:
            sign: ${get_sign($user_agent, $device_sn, $os_platform, $app_version)}
        method: POST
        url: /api/get-token
    validate:
        - eq: [status_code, 200]
        - eq: [headers.Content-Type, application/json]
        - eq: [content.success, true]
        - eq: [content.token, baNLX1zhFYP11Seb]
    extract:
        token: content.token # 参数提取

- test:
    name: /api/users/$user_id
    parameters:
        user_id: [1000,1001,1002,1003] # 参数化数据驱动, 方便测试多个 user_id 的情况
    request:
        headers:
            Content-Type: application/json
            User-Agent: python-requests/2.18.4
            device_sn: FwgRiO7CNA50DSU
            token: $token # 参数关联
        json:
            name: user1
            password: '123456'
        method: POST
        url: /api/users/1000
    validate:
        - eq: [status_code, 201]
        - eq: [headers.Content-Type, application/json]
        - eq: [content.success, true]
        - eq: [content.msg, user created successfully.]
```

- 运行 hrun

```
hrun test.yml # 进行测试

ll reports # 查看生成的报告, html 文件
```

- httprunner 配合 locust 进行性能测试

## jmeter
> apache jmeter: http://jmeter.apache.org/
- [慕课网-JMeter性能测试入门篇](https://www.imooc.com/view/735): JMeter 工具使用的展示
- [JMeter之HTTP协议接口性能测试](https://www.imooc.com/learn/791): 主要讲解 http 协议的基础知识
- [高性能产品的必由之路—性能测试工具](https://www.imooc.com/learn/278): 业务测试(JMeter) 占大部分篇幅,  还有系统性能/相关服务性能的知识

取样器: 业务流程
线程组: 场景设置->访问量
监视器: 监控->性能指标

业务流程 业务永远是基础 -> 录制工具 badboy/代理录制 -> 脚本制作 模拟流量 -> 性能测试 加压策略

用户自定义变量: `$[var]`
参数化: 需要重复多次测试的场景 -> 读取csv(csvReader / csvDataSetConfig)
关联: 上文获取的值需要在下文使用

类似工具对比: loadrunner

## aliyun PTS

阿里云性能测试 PTS: https://help.aliyun.com/product/29260.html

- 才并发100, 居然有不少超时
![才并发100, 居然有不少超时](http://qiniu.dayday.tech/pts-baogao.png)
- 原来是入门主机的 1M 带宽压满了
![原来是入门主机的 1M 带宽压满了](http://qiniu.dayday.tech/pts-ecs.png)

我经常向身边的人推荐云平台, 就像编程为人工作一样, 云也是一样的定位, 在接触云的过程中, 一群做技术的人怎么把技术产品化, 本身就是一件很有意思的事.

> 敢收费, 就是要证明技术的价值, 不服, 先看看 [阿里云帮助文档-PTS-名词解释](https://help.aliyun.com/document_detail/74223.html)

## 写在最后

性能测试, 新领域? 是的, 确实很新, 以前也就 ab 压完看 qps, 停留在这样的阶段, 真要跑 **大规模/自动化/性能测试更多指标**, 其实并没有多少沉淀. 这样折腾一波下来, **不虚了**.

性能测试三部曲:
- 流量录制: 代理工具记录 http 访问, 导出为 har/jmx
- 脚本编制: 自定义变量 + 参数化 + 上下文关联 + 热加载函数
- 测试报告: QPS 成功/失败 其他指标