# tool| 尝试黑苹果系列一

入手黑苹果主机并体验近一周, 讲讲故事与说说体验.

## why: 为什么选择黑苹果主机

- 生产力: 8k 主机(不含显示器)超 3w+ MacBook Pro 性能完全无压力, 用 MacBook Pro 也是要上各种外设来提高生产力
- 折腾: 不是鼓励折腾, 而是要习惯 **不完美**, 有问题不能心态蹦了, 遇到问题, 解决问题, **是大英雄显本色, 能折腾者享黑果**

## how: 说说经历与走的弯路

简单列列时间线:

- 最长的时间线可以追溯到大学
  - 2011 百度俱乐部主席, 在一次分享中讲了他 acer 成功安装黑苹果, 由于我当时用的 联想z465(AMD CPU), 只在virtual box 中安装尝试
  - 2012 网安(网络安全协会)社团的技术分享中, 昌力大大在他的分享中讲解了台式吃黑果以及一步解锁 宝开(popcap)游戏
- 2017 这一年参与了数个公司的多个项目, 去过好几个办公室, 都是抱着台式机去的
- 2018 在一家公司短暂兼职并使用 MacBook Pro 习惯后, 入手 MacBook Pro 作为生产力工具
- 2018-2020 MacBook Pro 不断增加外设, 几乎当做台式机在用
- 2020.4 开始了解并入坑黑苹果
- 2020.5 刷知乎了解黑苹果知识, 确定黑果所需的硬件配置
- 2020.6 刷京东确定黑果主机 -- 没必要挑战自己的装机实力了, 装机只有 0 次和一百次

遇到的问题:

- 配置环境时报错 `permission deny`, 误操作 `sudo chown dayday /`, 执行很久没执行完后 `ctrl-c` 终止掉
- 第二天来开机发现系统卡在苹果进度条进不去
- 联系客服解决: 对客服不要心存幻想...
  - 重启试试 -- 重启大法
  - 关机一会再试试
  - 那只能恢复, 用恢复U盘试试 -- 黑屏了
  - 黑屏了换显示器接口试试, 换显示器试试
  - 有 win 的电脑么? 向日葵远程 + 大白菜PE -- 这一步花了巨久的时间, 实际只是为了更换 EFI 文件
  - 邮箱多少? 发了 2 个 **EFI** 文件过来, 使用 **disk genius** 复制到 ESP 分区的 EFI 文件
- 花了一天, 种种尝试还是黑屏, 我只好先装好一个 win 先用起来, 之后尝试 wsl2 作为生产力工具
- 恶补黑果安装的知识, 开始各种尝试, 下面只讲解成功步骤

成功安装黑果:

- 我是 ssd + 机械双硬盘, 为了降低难度, ssd => macos, 机械 => win
- 使用 **分区助手**, 操作系统迁移工具, 将 win 迁移到 机械硬盘
  - 注意: 会遇到重启, 需要手动选择开机启动项进入机械盘完成迁移
- 单独一整块 ssd 给 macos 用, 对分区无需操作
  - 我看的 bilibili 教程大部分都是一块 ssd 双系统, 需要进行分区操作, 只有一个 up 主是单硬盘, 所以想简单安装, 单 ssd 首选
- **黑果小兵** 下载最新版镜像, 我选择的微信首发版, 并使用 ssd 移动硬盘刻录启动盘
  - ssd 移动硬盘 4-5 分钟刻录完成, U盘需要 15分钟+, 而且 2 次刻录都失败了...(不要对主机商送的 U盘 抱有幻想)
- 替换主机商邮件发送来的 EFI 文件, 选择硬件最接近的 EFI
  - OC EFI: z390 主板
  - clover EFI: z390主板 + rx5700xt 显卡 => 最接近我的配置
- 关键: EFI OK, 就能顺利进入到 macos 安装界面, 错了就是 黑屏/卡报错, 所以最 easy 的模式就是找对 EFI
- 关于显示器的细节: 我是双 DP 连显卡(不支持主机集显), BIOS 时只有左边的显示器亮, 安装过程 macos 时, 只有右边的显示器亮, 安装完成后, 左边显示器只有接了 HDMI 才可以显示
  - 细思极恐: 我使用恢复 U盘安装遇到苹果进度条走完黑屏, 是不是因为显示器不亮了?

目前使用正常:

- 查看 mac 设备信息: CPU / 显卡 / 主板 / 内存 / 双显示器 / 网卡(有线无线) / 蓝牙 / 声卡 都OK
- 懒得折腾 EFI 开机启动项, 重启时需要手动选择开机启动项
- 可正常重启, lock 一天后可正常点亮

## what: 讲讲黑果的一些百科

关于黑苹果的基础知识:

- mac os 系统, 推荐 **黑果小兵** 下载
- 支持 mac os 的硬件 -- 越接近 mac 产品原生越好
- EFI 引导文件, 引导 mac os 加载硬件

获取黑苹果的相关资源:

- [黑果小兵的部落阁](https://blog.daliansky.net/)
  - mac os 镜像文件
  - 黑苹果教程
  - 黑苹果长期维护机型清单
- [知乎-Mac996: 2020年黑苹果硬件配置推荐](https://zhuanlan.zhihu.com/p/139075806)
- bilibili: 黑苹果装机视频

PS: 不要为了装而装, 如果 EFI 可以解决, 选好 EFI 就对了

## 写在最后

目前使用时间还比较短, 后续继续更新, 吃黑果可以简单概括为一句话:

> 黑果的 easy 模式: EFI 对了, 苹果图标亮了, 世界就对了

黑苹果主机 -> 确定财务状况后处理, 明确主机不是瓶颈

- CPU: ix-9xxx+核显(F不带核显)
- 主板: z370/390 技嘉/华硕
- 内存: 16G+ 3200起步
- 硬盘: 250g+ M2 三星970/960
- 显卡: 蓝宝石 5700xt vega56/64 rx590 radeon-VII
- 显示器: 4k 27寸+
- 蓝牙/wifi: 博通 BCM943602CS/BCM94360CD
- 电源: 550w+
- 散热: 4管+
