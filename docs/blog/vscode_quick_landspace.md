# tool| vscode 全景速看
> a quick view of vscode landscape

奉上 vscode 干货. 本文着重 **速看**, 详细内容, 参看 [github/daydaygo - vscode](https://github.com/daydaygo/vscode)

新问题往往都是老问题, 或者说历史总是惊人的相似, 在具体到 vscode 的问题中之前, 先把一些老问题给拎清, 扫清障碍

## 老问题
### 名与实
- 名与实的重要性: 很多问题往往只是「语义上」的, 要么是不清楚概念(名), 要么是不清楚概念具体指的是什么(实), 要么就是翻译转换过程导致的名实不符
- 很多问题其实并不难, 只是前面叠加了一层 **名与实** 的问题, 所以很难, 因为 **名实问题**: 知道就是知道, 不知道就是不知道
- 英文: 避免翻译产生的歧义, 减少思维转换; 编程世界, 对英文太友好了

### 认识论
- 认识论: 先整体后局部; 先通用后细节
- 知识成体系的重要性: 因为遗忘的存在, 学习是需要不断重复的, 不成体系的知识, 难以区分哪些是重要的, 进而陷入到细节当中

### 下层基础决定上层建筑
- 在深入细节之前, 了解一下 vscode core, 会在遇到问题时 **从容** 很多: 上层的问题往往是一层又一层的封装后导致问题的本质 **不可见**, 从 core 来理解往往 **事半功倍**

### 分层
- 计算机科学领域的任何问题都可以通过增加一个间接的中间层来解决.

## 通用技巧

在继续深入之前, 了解一些通用技巧, 可以 **事半功倍**:

- 搜索
    - 几乎所有地方都支持搜索功能
    - 搜索支持 **模糊匹配**, 技巧是 **单词前缀匹配**, 比如 `view: toggle zen mode`, 输入 `vtz` 就能找到了
    - 很多地方都支持搜索, 可以试试`cmd-f` 快捷键
    - 很多地方隐藏搜索功能, 直接输入字符进行 **前缀过滤**
- 二八法则: 只要花很少的时间掌握通用功能, 就能完成大部分工作
    - vscode 中正例: command(命令) view(视图) menu(菜单)
    - vscode 中反例: shortcut /setting 等就是反例, 大而全, 全是细节
    - 如何应用二八法则? **任自然 -- 经常使用到的, 自然会记下**
- 自动提示: 好的工具往往不需要你记忆大量的细节, 编辑过程中的自动提示会让你 **如虎添翼**

## 认识 vscode

![vscode01.jpg](https://upload-images.jianshu.io/upload_images/567399-53985322990e2c95.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

先整体后局部:

- activity bar(局部下还有很多功能)
    - explorer
    - search
    - source control
    - run
    - extension
- side bar
- editor
- status bar

![vscode02.jpg](https://upload-images.jianshu.io/upload_images/567399-402258cd7a79c096.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


先通用后细节:

- menu: 大部分功能都可以在 menu 中找到, 并且 menu 已经帮我们进行了 `分类`
- <a id="command">command palette</a>
    - vscode 中的大部分功能, 都是 `command`
    - 如图: 使用 `menu > view > command palette` 就可以打开, 后面是快捷键
- panel

![vscode03.jpg](https://upload-images.jianshu.io/upload_images/567399-49e6c18d1ac99205.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![vscode04.jpg](https://upload-images.jianshu.io/upload_images/567399-a25f8e30041b6d8e.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

`command palette` 已经非常 **通用** 了, 其实它上一级的 `view` 更通用, 几乎所有的功能, 都是一个又一个 `view`

- `?`: 查看有哪些 `view`
- `...`: 空的时候, 用来打开文件
- `>`: 上面看到的 command palette
- `view `: 打开各种界面

> 限于篇幅, vscode landscape / vscode core 等类容, 请查看 [github/daydaygo - vscode](https://github.com/daydaygo/vscode)

## 荐书

虽然你说得很有道理, 可是 vscode 的内容好多, 掌握了 **道与术**(套路), 还是要花很多时间, 有没有更快的方法?

翻译一下: 敢不敢更过分点, 给条捷径?

敢! 书籍是人类进步的阶梯. vscode 这条路, 几千万人(夸张手法)已经走过.

荐书: vscode 权威指南

- 「我知道, 我改行动了. 这是一片即将变红的蓝海」 -- 作者简述结识相伴历程, 于 2016.6 开始开发 vscode 的插件
- zen
    - 搜索: google / Stack Overflow
    - 提问
        - [提问的智慧](https://github.com/ryanhanwu/How-To-Ask-Questions-The-Smart-Way/blob/master/README-zh_CN.md)
        - 不仅需要好的问题, 也需要好的解决方案
    - 学习
        - 自己思考
        - what-how-why
        - 举一反三: 通过 **类比** 等手段, 调用自己已有的知识
- vscode 如何做开源
    - github: issue & pull request(PR)
    - 开发流程: roadmap/年 产品设计/月初 测试与发布/月末
    - 文档
    - 插件
- 想了解更多细节, 融入 vscode 大家庭, 这本书你值得拥有

## 写在最后

读到这里, 希望对你有帮助. 

我是 dayday, 读书写作敲代码, 永远在路上.