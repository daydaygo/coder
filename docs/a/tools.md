# tools

> pc 效率/生产 工具, 但是手机控制了大部分时间, 要把手机的效率提上去

- [开发环境 mac](mac.md)
- vscode: 文本编辑器, [github/daydaygo - vscode](https://github.com/daydaygo/vscode)

## zen 禅

- 把事情当做产品来做
- 多使用英文
- 搜索: 好记性不如会搜索
  - 百度/Google
  - 万能淘宝: 能用钱解决的问题, 都不是问题, 不如 pmp pdu 学时
  - 知乎: 奇奇怪怪的知识又增加了
- 模糊匹配: 几乎所有搜索的地方都支持模糊匹配, 只要顺序正确就可以搜索到
- 工具化: 在重复执行不出错上, 机器绝对可以轻蔑地说「愚蠢的人类」, 我通常会回应「所以你是机器」
  - 工匠应该专注于作品的创意，不应该浪费精力，没限制地在折腾自己的工具
  - 仰之弥高，钻之弥坚 工欲善其事, 必先利其器
- 手段->目的; 教程->使用; 学习->实践; 知->行
- 自动保存: 都 9102 了, 还不用自动保存
- vim: learn once, use everywhere

---

- 人生苦短, 我有工具
- 高效, 是一种瘾

## utools

> 生产力工具集: <https://u.tools/>

- 应用切换: 可以慢慢调教到一个字母就打开常用应用
- 插件管理
- 网页快开: 快速打开 百度/谷歌/知乎/bilibili 等网页
  - google stackoverflow
  - 高德设置默认城市: `https://www.amap.com/search?query={query}&city=310000`
  - gopkg     <https://pkg.go.dev/search?q={query>}
  - runoob    <https://www.runoob.com/{query>} <https://www.runoob.com/?s={query>}
- chrome小助手: chrome书签(只能搜索标题); chrome隐身模式; chrome打开的标签页
- chrome 历史记录搜索
- github小助手
- 聚合翻译 计算稿纸 clipboard
- mverything: 文件查找
- tophub 今日热榜

## edge

- extension
  - Tampermonkey
    - 插件: baidu csdn 拒绝二维码登陆 htmlVideo
    - <https://greasyfork.org/en>
  - [谷歌访问助手](http://www.ggfwzs.com)
  - gitzip for github: double click
  - [vimium C](https://microsoftedge.microsoft.com/addons/detail/vimium-c-all-by-keyboar/aibcglbfblnogfjhbcmmpobjhnomhcdo)
- console panel `cmd-shift-p`
  - 全屏截图: `shot`
- flags `edge://flags`: download dark
- collection: learn algo
- 夸克: 下滑搜索 快搜 扫一扫 学习
- 360: 支持flash

```sh
# vimium C
# Insert your preferred key mappings here.
unmapAll
map o Vomnibar.activateTabSelection
map p togglePinTab
map m toggleMuteTab
map u scrollPageUp
map yy  copyCurrentUrl
map j scrollToTop
map k scrollToBottom
map [[ goPrevious
map ]] goNext
map f LinkHints.activateMode
map ? showHelp
map i focusInput
map h previousTab
map l nextTab
map = visitPreviousTab
map r closeTabsOnRight
```

## markdown

- syntax 语法
  - 语法基础basic: 标题`#` 列表`- 1. [ ]` code 加粗/斜体/引用/高亮`*- - > ==` 链接/图片`[]() ![]( =150x150)` 分割线/删除线`--- ~~` 表格/对齐`:--:`
  - 图片和附件格式img: <https://www.yuque.com/yuque/help/rms0so#20def794>
  - 代码code: 编程语言支持 diff [运行代码 Run Code Chunk](https://github.com/shd101wyy/markdown-preview-enhanced/blob/master/docs/code-chunk.md)
  - 文本绘图graph: [数学图表 vega](http://vega.github.io/) [ditaa](https://github.com/stathissideris/ditaa) <https://draw.io> plantUML(部署图 脑图) flowchart mermaid graphviz gantt 常见图形案例
  - 特殊sp: 脚注`[]( "name") [^1]` 目录`[toc]` 数学公式latex 注音符号`<ruby>味<rt>mǐ</rt></ruby> {喜大普奔|hē hē hē hē}` ppt/滑动图片`<![](url),![](url)> <!-- slide -->` 前置frontmatter`---yaml` [批注criticmarkup](http://criticmarkup.com/users-guide.php)
- 微信排版: <https://www.mdnice.com/>
- md+onedrive+azureWeb <https://www.madoko.net/>
- Typora: 直接复制 网页-html/Excel(不支持合并行列)
- 语雀 在线文档编辑器: <https://www.yuque.com/yuque/help/thzr79>
  - 附件attachment: pdf office(word/Excel/PPT) keynote/numbers/pages xmind/mindnode axure(.rp) 其他不能预览 是否支持下载
  - 思维导图mindmap: 多用快捷键 样式/彩虹节点
  - 流程图 日历 投票 翻译 嵌入语雀文档 加密文本 第三方服务(墨刀 processon figma canva codepen amap 媒体->卡片) 提示块 `:::tips`
  - 加密文本 语雀文档 本地文档 本地视频 在线视频(B站 youku) 第三方服务
- vscode
  - yzhang.markdown-all-in-one
  - shd101wyy.markdown-preview-enhanced
  - davidanson.vscode-markdownlint
  - josephcz.vscode-markdown-mindmap-preview
- vscode ext - markdown prview enhanced: 合并单元格 纯文本绘图 ppt
  - emoji 😁  :smile:
  - 导入文件 `@import` 导出文件 HTML/chrome/pdf/ebook/pandoc
- mubu幕布: 导入opml/xmind/freemind; [幕布精选社区](https://mubu.com/explore)
- <https://word2md.com/>

## SE 搜索引擎技巧 信息检索法

- 技巧
  - 关键词: 名词 `|+-""*《》`
  - 指令: `site:*.aliyun.com` intitle inurl filetype
  - 限定时间
  - 程序员金手指 tutorial example tricks cheatsheet cookbook awesome
- 站长之家
  - [子域名查询](https://tool.chinaz.com/subdomain/github.com)
- 开发
  - <http://kaifa.baidu.com>
  - google sf github bing
- 商业数据库
  - 万得 <http://www.wind.com.cn>
  - 彭博 <http://bloomberg.com>
  - 路透 <http://cn.reuters.com>
  - 锐思 <http://www.resset.cn>
- 学术数据库
  - [知网](http://www.cnki.net)
  - [万方](http://www.wanfangdata.com.cn)
  - [交大校园图书网](https://www.lib.sjtu.edu.cn)
  - 中国国家图书馆 <http://www.nlc.cn/>
  - [国家科技图书文献中心数据库](https://www.nstl.gov.cn)
  - [中国科学院文献服务](http://sciencechina.cn)
  - [科技方面的文献](https://www.sciencedirect.com)
  - 维普 <http://www.cqvip.com>
  - EBSCO <https://www.ebsco.com>
  - 数据圈 <http://www.shujuquan.com/>
  - 台湾学术数据库 <http://fedetd.mis.nsysu.edu.tw/>
  - 台湾大学电子书 <http://ebooks.lib.ntu.edu.tw/Home/ListBooks>
  - 外国文献搜索SCI-hub <https://sci-hub.se/> http://sci.hub.tw http://apps.webofknowledge.com https://scholar.glgoo.org https://gfsoso.99lb.net
- 网上共享资源
  - 百度文库 <http://wenku.baidu.com/>
  - 豆丁文库 <http://www.docin.com/>
  - 爱问共享 <http://ishare.iask.sina.com.cn/>
  - 道客巴巴 <http://www.doc88.com/>
  - 360个人图书馆 <http://www.360doc.com/index.html>
- 专业小众
  - 大数据导航：<http://hao.199it.com/>
  - 花朵或者绿植的品种：花伴侣APP，或者是微信小程序“形色识花”
  - 歌曲的乐谱（钢琴、吉他、提琴等），可以使用搜谱APP
  - 小猿搜题APP
  - https://zhihu.com https://www.wiki-wiki.top
- 一、资源导航网站
  - 1、书享家（电子书资源网站导航）：<http://shuxiangjia.cn/>
  - 2、学吧导航（自学资源网站导航）：<https://www.xue8nav.com/>
  - 3、科塔学术（学术资源网站导航）：<https://site.sciping.com/>
  - 4、HiPPTer（PPT资源网站导航）：<http://www.hippter.com/>
  - 5、Seeseed（设计素材资源导航）：<https://www.seeseed.com/>
- 二、工具导航网站
  - 1、阿猫阿狗导航（互联网工具导航）：<https://dh.woshipm.com/>
  - 2、创造狮（互联网工具导航）：<http://chuangzaoshi.com/index>
  - 3、addog（广告营销工具导航）：<https://www.addog.vip/>
  - 4、199it（数据导航）：<http://hao.199it.com/>
  - 5、雪球导航（财经工具导航）：<https://xueqiu.com/dh>
  - 6、打假导航（国家部门导航）：<http://www.dajiadaohang.com/>
  - 7、搜狗网址导航（地方部门导航）：<http://123.sogou.com/diqu/>
- 三、聚合搜索平台
  - 1、一个开始：<https://aur.one>
  - 2、虫部落：<https://search.chongbuluo.com/>
- [Alexa](https://www.alexa.com/topsites) [webCommic](https://xkcd.com) [openMovieDB](https://omdbapi.com)

## ticktick 滴答清单

> <https://help.dida365.com/>

- GTD 清单革命: list/smartlist task(data date/due/reminder priority tag assignee) calendar summary
- feature: search sync habit(习惯养成) pomo(番茄钟-专注-单核工作法)
  - 打卡-habit(习惯养成)
- tips: parse(自然语言处理) set(特殊字符 `~ # - ! @`) shortcut
  - [智能日期识别](https://guide.dida365.com/smartdateparsingrules.html)
- Android
  - 滑动操作: 左/右 长/短
  - 语音输入: 长按右下角「+」键进入语音输入状态
  - markdown 支持
  - project-kanban
  - 清单 list: 普通(生活/工作/会议) 智能(今天/明天/最近7天) 自定义(优先级/tag)
  - 日历: 视图 快速安排任务 时间轴缩放
  - 专注计时 pomo
  - 习惯打卡 habit
- mac
  - 摘要功能
- 玩转公众号
  - 给「滴答清单」公众号发消息, 自动创建任务到 inbox

我的使用指南:

- 不确定的 task 放到 inbox -> 定期清理到 list(todo/high/萌萌dayday)
- task: 必须包含时间; 长->小; 分类
- list-todo: 不确定时间的任务, 需要后续确定 做/不做
- list-high: 高优先级事项, 重要紧急4象限
- list-萌萌dayday: 共享清单

## 生活.书影音

- 文档写入
  - dayday个人: https://github.com/daydaygo/coder https://gitee.com/daydaygo/drive
  - 萌萌dayday: 腾讯文档 [萌萌dayday日记-无法发票圈202105已不使用](https://www.yuque.com/daydaygo/blog)
- 书影: 豆瓣
- 音: 网易云音乐 qq音乐

## wechat 微信

- 收藏: 长期/短期; 搜索; tag: me/zm/运动 whu/sjtu/wh/上海/世界 coder/思维/fun
- 笔记: 文件/截图/地点/待办; 标签; 常用信息
- 微视: 30s- 朋友圈短视频
- 朋友圈->电子书: <https://weixinshu.com/u/X6aR4Z>
- 公众号: 谣言.全民较真.谣言过滤器 疾病.菠萝因子.丁香医生 养生保健.范志红.顾中一

## mail

- foxmail
  - account
    - qq邮箱 - qq企业邮箱
    - [sjtu](https://net.sjtu.edu.cn/_mediafile/wapsjtunet/2017/05/25/25iv1rkz0k.png)
  - rule 分类规则

## doc 文档/手册

- dash <https://kapeli.com/dash_guide>
  - `preference > general > show dock icon`: show menu
  - cmd-l cmd-f `xxx:` `⌥↑`
  - search with `type`: mainpage/guide type/var/const func interface/class/property/method package playground
  - doc: html/css/javascript/typescript php/go/rust redis emoji/gitmoji
- [zeal](https://zealdocs.org/download.html): doc tool, for win & linux
- 云产品文档: 技术要点+最佳实践
- 腾讯云开发者手册 <https://cloud.tencent.com/developer/devdocs>
  - http go
- 百度文库下载: `baidu.com -> baiduvvv.com`

## epub

- 改 .zip 为 .epub 进行编辑
- epub目录导出: 阅读工具打开直接copy; 编辑 `toc.ncx` 文件(php处理xml); 通常第一个文件就是目录`text00000.html`; web阅读->展开所有目录->inspect->copy as outerHtml

## android miui app

- 小爱同学
- 小技巧: 省电.锁屏.关app 广告.通知 慢.清理 下载软件 密码 护眼模式 `设置+应用商店`
- 绑卡 挂号.微信.医疗健康 购物 外卖 地图.查.导航 打车 火车票 酒店 相册.抖音
- 截图：三指下滑
- 长截图：截图后点击缩略图，选择长截图
- 小米笔记本 Ubuntu wifi sudo vi /etc/modprobe.d/blacklist.conf blacklist acer-wmi
- 小米换机：公交卡迁移后再重置老机; 卡住->微信清理
- 传送门: 长按文本分词/搜图

---

- clipboard plus: 复制后分词选择
- ppiicc: 合成电影字母图
- 一个木函: 工具集App
- 上海公交卡: 给实体卡充值 卡面
- 讯飞有声: 英文背诵(文本->语音)
- 行者: 骑行app.轨迹记录有问题
- 微视: 短视频
- 不挂科-大学题解 招行企业银行 现代汉语词典-会员91

## blog 简史

- 现在: vuepress - coding page; 简书; 知乎专栏(知乎发汇总贴/月)
- [farbox](http://farbox.com)
- csdn blog: <http://blog.csdn.net/czl1252409767> -> 这上面保存着大部分的 oj 代码
- 新浪 blog: <http://blog.sina.com.cn/u/1805628174> -> 不适合写代码
- 百度俱乐部副主席, 短暂使用过百度空间(已下线) -> 主打的代码显示功能, 复制出来格式就完全乱了
- whu 网安 GW: 建议技术人有一个自己的 blog

## mark

- <https://post.m.smzdm.com/p/andr4rgp/>
- <https://tableconvert.com/>
- 出门清单：背包（伸手要钱/雨伞/口罩/充电宝）牙刷/浴巾/一套内衣袜子
- 线上问题排查: <https://mp.weixin.qq.com/s/tG1kOGtAJzHc37XGlsfAPQ>
- slide(reveal.js)
