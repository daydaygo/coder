# yii| 框架简析

- [yii 框架解析](https://www.jianshu.com/p/c7ad75ce5cef)

因为工作原因需要重拾 yii 框架, 而之前一直使用的 [hyperframework](http://hyperframework.com/cn/manual) -- 公司技术团队内部开发的框架, 需要什么服务, 直接往框架上添加即可. hyperframework 底层是服务容器, 需要添加新的服务很简单, 这个在我之前的 blog [hyperframework WebClient 源码解读](http://www.jianshu.com/p/cf39804b7c04) / [用 yii 框架 10 分钟开发 blog 系统?](http://www.jianshu.com/p/9740878ea07b) 都有提到, 不熟悉的同学可以移步一览. 所以思路上需要做一点改变: **yii 已经封装好了很多常用服务, 开箱即用**.

> PS: 这篇博客的英文版标题是 `get_yii.md`, 纪念一下我在 github 上加入的第一个开源组织以及开源项目 [iiYii/getyii](https://github.com/iiYii/getyii). 一晃三年过去了, 在此对开源作者表达诚挚敬意!

之前也提到过, 这样的重型框架之所以入门比较困难, 很大一部分原因是**功能太多**, 导致难以分清主次和记忆. 这里**记忆**不是死记硬背, 而是知识的内化, 不过知识的内化说起来更难以理解, 倒不如说是想要达到知识随用随取信手拈来的境界, **你起码得记得吧**. 当时也提到一些方法, 这篇 blog 会进行完善并实践.

简析 yiii 框架的方法(类似的重型框架都可以采用这个思路):

- 生命周期
- 核心架构 & 模块划分
- 脑图/笔记 等工具
- 实践, 比如开发 blog 系统

## 生命周期

通过生命周期来 **解读源码/定位问题** 是非常非常重要的手段. [鸟哥](http://www.laruence.com) 在他的博客中, 无论是源码分析, 还是问题解决, 多次实践, 这里摘录 [思考能力何其重要..](http://www.laruence.com/2009/06/17/951.html) 中的一段话:

> 没有, 好吧, 如果说一定要有, 那就是:vim + grep + “大胆推论,小心验证”, 我知道一个c写的可执行文件, 是从main开始的, 我知道对于mod_php来说, 开始点必然在apache将控制权交给它的那一刻开始, 有了这些, 就可以使用vim徜徉在海一样的代码中, 而不会迷路. 有了这些, 不就足够了么?

yii 框架的生命周期, 虽然在应用上会做如下分类:

- web 应用: 一次 http 请求的生命周期, 抽象一点就是 request + response
- console 应用: 一次脚本的执行过程.

但是其实是统一的, 都是对应程序的输入与输出(input/output), 只是在 http 请求这里, 使用 request/response 来表示, 而在后台脚本这里, 使用 argument/option 表示输入, 脚本中直接 echo/print 表示输出

- php cli lifecycle: 仔细看这张图, 在执行的过程中, 也有 `request` 的概念

![php cli lifecycle](http://upload-images.jianshu.io/upload_images/567399-52f8e2d7e756e578.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 官方供图 - yii request lifecycle: 很明显的 request / response

![运行进制](http://upload-images.jianshu.io/upload_images/567399-7ede62dde4cd1c8b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 核心架构 & 模块划分

- 官方供图: application structure

![image](http://upload-images.jianshu.io/upload_images/567399-5b6e7d33257f8d35.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

这张图里, 最核心的其实是: **应用主体**, 其他部分, 都是通过应用主体来协调调度, 甚至你可以认为, `yii = application(container) + component(service)`

我制作的 [百度脑图: get_yii](http://www.laruence.com/2008/09/19/520.html), 带上了标识表示我理解的重要性.

- entry: 入口脚本, 这是一切的开端, 从这里开始可以对框架有一个「全景式」的张开
- application(container): 框架的核心, 容器机制已经是现代化框架的代名词了
- component(service): 框架提供的功能(服务), 由 application 来协调调度, 所有可见的功能, 大家经常提的 MVC, 其实也属于这里
- extension(vendor): 扩展(依赖), 虽然需要和框架配合, 可能也需要改造成 component, 但是它的核心理念其实是 **依赖管理**, 从而 `站在巨人的肩膀上`
- concept: 框架需要涵盖的设计理念, 这部分的内容希望大家可以好好阅读, 因为 **可能你一直徜徉在业务的海洋, 却没有花时间思考技术本身**

## 工具 & 实践

熟悉我的同学可能都知道, 我比较喜欢用 **思维导图**, 平时也喜欢 **记笔记**, 部分观点在之前也提到过, 这里再赘述一下:

- 关于记忆: 知识的内化是需要过程的, 这个过程可以简单理解为 **记忆**, 很多时候, 其实 **记着了**, 也就渐渐掌握了. 所以有时候大牛会跟你说, 下去多看看就行了. 为什么, 因为多看看就是一个重复的过程, 一回生二回熟. 当然, 也可以用点文雅的词 -- **潜移默化**
- 笔记: **好记性不如烂笔头**, github 的 gist 来记代码片段, 各种笔记应用, 这个就不用多言了吧. 提醒一点, **注意生态, 不然最后都是杂草了**
- 思维导图: 开始使用思维导图的契机其实蛮简单, 这玩意可以帮助记忆, 当时还特地买了一套书来看这个. 其实并没有必要特地去研究这个, 我一般这样的场景下使用 -- **一个事物太复杂, 我需要花时间来整理清楚它**

当然还有一些其他的工具, 主要遵循 **一图胜千言** 的理念, 比如流程图, 时序图, 这就看场景了, 不过使用频率没那么高.

关于实践, 我们前任 CTO 的建议我会一直铭记:

> 多读源码
