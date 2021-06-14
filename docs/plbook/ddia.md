# book| 设计数据密集型应用

- [中文版翻译](https://github.com/Vonng/ddia)
- [中文版gitbook](https://vonng.gitbooks.io/ddia-cn/content)
- [读书笔记| 设计数据密集型应用](https://www.jianshu.com/p/9c465f674971)

## 想法

好书, 已入手英文版准备开刷

- 如果你需要一份数据相关内容/技术的坐标或者地图, 这本书绝对可以排在明显靠前的位置, 如果再考虑到时效性, 我建议你立刻打开来看看
- 译者称这是 **2017年度读过的最好的一本技术书籍**. 这是明显的推荐词, 我们既无法知道译者究竟读了多少书, 也不知道译者读的书又和我们自己有多大的联系. 当我真的想到这些问题时, 我又该以多大的可信度来 **接受** 这句话呢? 倘若把这个问题再推广开, 如果一句话都要这么 **费力** 一番, 读书就真成了一件费力的事了, 所以会慢慢形成一种倾向, 选择相信 KOL(key opinion leader, 关键意见领袖). 讨论这些并不想表达某种倾向性, 只是想指出读书中存在 **思考与相信**.

## 笔记

在我们的社会中，技术是一种强大的力量。数据、软件、通信可以用于坏的方面：不公平的阶级固化，损害公民权利，保护既得利益集团。但也可以用于好的方面：让底层人民发出自己的声音，让每个人都拥有机会，避免灾难。本书献给所有将技术用于善途的人们。
​计算是一种流行文化，流行文化鄙视历史。 流行文化关乎个体身份和参与感，但与合作无关。流行文化活在当下，也与过去和未来无关。 我认为大部分（为了钱）编写代码的人就是这样的， 他们不知道自己的文化来自哪里。
互联网做得太棒了，以至于大多数人将它看作像太平洋这样的自然资源，而不是什么人工产物。上一次出现这种大规模且无差错的技术， 你还记得是什么时候吗？
——阿兰·凯接受Dobb博士[^1]的杂志采访时（2012年）

语言的边界就是思想的边界。
—— 路德维奇·维特根斯坦，《逻辑哲学》（1922）

与可能出错的东西比，'不可能'出错的东西最显著的特点就是：一旦真的出错，通常就彻底玩完了。
——道格拉斯·亚当斯（1992）

如果船长的终极目标是保护船只，他应该永远待在港口。
——圣托马斯·阿奎那《神学大全》（1265-1274）

数据密集型(data-intensive) 计算密集型(科学计算) IO密集型(文件读写/网络)
百分位点: 99.9%, 指 99.9% 的请求都能在设置的延迟值内返回. 属于数据型应用中非常好的一个衡量指标. 类似的指标在高可用中也有, 比如3个, 就是全年 99.9% 的时间内服务可用.

数据编码: 让不同的组织达成一致的难度超过了其他大多数问题
这种构建应用程序的方式传统上被称为面向服务的体系结构（service-oriented architecture，SOA），最近被改进和更名为微服务架构
本地函数调用 -> RPC -> 网络失败/超时 -> 重试 -> 幂等（idempotence）
Actor模型是单个进程中并发的编程模型
如果停止写入数据库并等待一段时间，从库最终会赶上并与主库保持一致。出于这个原因，这种效应被称为最终一致性（eventually consistency）
二阶段提交（2PC, two-phase commit） 快照隔离（snapshot isolation）
悲观与乐观的并发控制

[^1]: Dobb博士的杂志的 Jolt Awards, 素有「软件业界的奥斯卡」之美誉