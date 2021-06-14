# coder| 基础必知必会之换行符

换行符是很 **基础** 的概念, 所以也是很 **容易** 遇到问题. 这里简单扫除一下盲点.

内容摘要:

- 换行符的概念
- 换行符的引发的问题: git 文件全文冲突
- 换行符最佳实践

## 换行符的基本概念

[阮一峰blog - 回车和换行](http://www.ruanyifeng.com/blog/2006/04/post_213)

同时推荐:

- [阮一峰大大的blog](http://www.ruanyifeng.com/blog/)
- [书 - 如何变得有思想](http://www.ituring.com.cn/book/1533)
- [书 - 黑客与画家](http://www.ituring.com.cn/book/1171)

换行符有 3 种:

- `\n` linux 风格换行符, 也写作 `LF`
- `\r` Unix(早期mac) 风格换行符, 也写作 `CR`
- `\r\n` win 风格换行符, 也写作 `CRLF`

这里补充下 **编码** 的知识:

> 计算机只能识别二进制 -> 二进制可以转十进制(对人友好) -> 字符集将十进制和字符进行对应

- ASCII: 最简单的字符集就是 [百度百科 - ASCII](https://baike.baidu.com/item/ASCII), 只定义了 128 个 10 进制和字符的对应关系
- 乱码: 有很多不同的字符集(不同语言使用不同的字符), 存储的最终其实是二进制, 如果字符集错了, 十进制由字符集匹配到的字符就不对了, 于是乱码
- 不可见字符: 学过 c 语言的小伙伴知道, 有 **转义字符** `\` 这个概念, 就是用来表示不可见字符.

在你眼里的回车(换行), 在计算机里其实 **换行符**, 最终都是 **0和1**

## 换行符引发的问题: git 全文件冲突

如果不注意换行符, 很容易发生这样的情况: `git merge` 一下, 提示有文件冲突了, 打开一看, 吓一跳: 全文件都冲突了.

```
<<<<<<< HEAD
    // 全文件
=======
    // 全文件
>>>>>>> xxx-branch
```

冲突原因: **2个分支上的文件版本使用了不同的换行符**

## 换行符最佳实践

先来解决上面的问题:

- 首先, coding 时就需要注意, 一般 IDE 或者编辑器的 **右下角**都可以查看当前文件的换行符
![编辑器 - sublime3](http://upload-images.jianshu.io/upload_images/567399-63f2dfd6211e1eb3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![IDE - phpstorm](http://upload-images.jianshu.io/upload_images/567399-3c68c36f9850754a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

同时在配置中配置上默认使用 Linux 风格换行符:

```
"default_line_ending": "unix", // sublime3 中添加此配置

phpstorm -> setting -> editor -> code style -> line separator // phpstorm 中的设置方式
```
不过推荐通过 `Ctrl + shift + a` 使用命令面板打开:

![phpstorm - 打开命令(action)面板](http://upload-images.jianshu.io/upload_images/567399-03524a5f3e5733ac.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
![phpstorm - 通过命令(action)直接打开配置](http://upload-images.jianshu.io/upload_images/567399-85941761c8428e4d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 其次, git 有配置可以处理这种情况

```
git config -l # 查看版本库的 git 配置
git config -h # 查看 git config 的帮助, 使用 --help 可以查看 html 帮助页面

git config core.autocrlf input # 配置见下图
git config core.safecrlf true # 换行符有冲突时禁止 commit
```

![autocrlf 配置项](http://upload-images.jianshu.io/upload_images/567399-635a68bf2e57ceb1.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 写在最后

这里只用 sublime3 和 phpstorm 来举例, 毕竟对换行符的支持, 已经是 **现代** 编辑器必备的功能了 -- 无论文本文件使用的何种换行符, 可以正确的显示 **换行**.

对 **不可见字符** 想要印象更深刻一点, 可以加上配置让编辑器显示出来:

```
"draw_white_space": "all",  // sublime3 显示 tab 和 空格
```
![sublime3 显示 tab 和 空格](http://upload-images.jianshu.io/upload_images/567399-55c4af311862a81b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![phpstorm 显示 空格](http://upload-images.jianshu.io/upload_images/567399-2d8917bf6bc3a2a2.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
