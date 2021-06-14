# hyperf| 快速搜索 hyperf.wiki

平时使用 [hyperf wiki](https://hyperf.wiki) 比较多, 一直有一个需求: 怎么快速搜索? 

目前的路径是: 打开网页 -> 搜索栏 -> 结果显示只有侧边栏

## hyperf.wiki 生成 dash docset

最近用 dash 比较多, 查文档还是蛮方便的, 就打算看能不能结合起来

- docsify -> dash docset ? 一番搜索无果, 此路不通
- dash docset 官网 -> Any HTML Documentation, 貌似可以, docsify 的 markdown 文件可以转为 html 文件
  - dashing: 基于 go 语言, 看 readme 中 example 的示例, 可以将文件夹下所有 html 文件转为 docset
  - dash-docset-builder: 基于 PHP 语言, 看 readme 和 源码, 这东东是真的牛逼, 可以直接一个 url 把文档的 html 网站爬完, 真屠龙宝刀 -- 可惜我这就文档搜索, 折腾爬虫万一掉坑里了...
- markdown -> html, 折腾了不少工具, 悲剧的发现直接 **文件夹 -> 文件夹** 并不行
  - pandoc: 社区小伙伴推荐, 长长的 readme...
  - gomd: 搜索到的工具, readme 非常简单, 基本就 **一行用法**
  - other: 喵了一眼不太「简单」就 pass 了
- 文件夹 -> 文件夹, 这个就比较简单了, **PHP 毕竟用的很 6**, 直接掏出 `模板`
  - 把模板中的 `// file` 的部分替换为业务逻辑即可

```php
// 基于 readdir() 的目录递归访问文件
function readdir_r($path) {
    $fd = opendir($path);
    while (($f = readdir($fd)) !== false) {
        if ($f == '.' || $f == '..') continue;
        $t = $path . '/' . $f;
        if (is_dir($t)) { // dir
            readdir_r($t);
            continue; // be care
        }
        // file
    }
}

// 基于 scandir() 的目录递归访问文件
function scandir_r($path) {
    $a = scandir($path);
    foreach ($a as $v) {
        if ($v == '.' || $v == '..') continue;
        $t = "$path/$v";
        if (is_dir($t)) { // dir
            scandir_r($t);
            continue; // be care
        }
        // file
    }
}
```

- 调试后的完整代码
  
```php
$from = '/Users/dayday/hub/hyperf/docs/en';
$to = '/Users/dayday/hub/tmp/docset';

function gomd($from, $to) {
    $a = scandir($from);
    foreach ($a as $v) {
        if ($v == '.' || $v == '..') continue;
        $md = "$from/$v";
        $html = "$to/" . str_replace('.md', '.html', $v);
        if (is_dir($md)) {
            if (!is_dir($html)) mkdir($html);
            gomd($md, $html);
            continue;
        }
        $cmd = "gomd $md -o $html";
        echo "$cmd \n"; // 先调试这里, OK 了再开 shell_exec()
        shell_exec($cmd);
    }
}
gomd($from, $to);
```

- 使用 dashing

```sh
dashing init # 生成配置文件, 根据 readme 修改, 这里有知识点: css selector

dashing build # 根据配置文件生成 docset
```

- dash 中验证: 遗憾的发现, 并没有达到要求
  - 中文文档会乱码, 打开 chrome 访问找到问题: charset(字符集) 并没有使用 utf-8
  - 英文文档不会有乱码问题, 但是英文文档不全
  - **搜索** 依赖 **索引构建**, 上面的 **css selector** 可以起部分作用, 制作麻烦且离理想状态甚远

## 命令行下搜索其实也还行

平时开发过程中, 其实也经常使用搜索, 主要有 2 类

- 命令行下使用 `ag`, 不用管那么多花里胡哨的, 常用就这几个

```sh
ag xxx # 当前文件夹下递归搜索文件内容
ag --html xxx # 限制文件类型
ag -g xxx # 查找文件名
```

- ag 还可以, 不过多一次 `cd` 操作, 速度就慢下来了, 当然 `cd` 也有快捷工具 `z-jump`

```sh
cd xxx # 一波 cd 操作后
z xxx # 会根据使用使用习惯, 跳转到最常用的目录, 如果不对, 加个 tab 也能很快找到
```

- 看起来 `ag+z-jump` 似乎满足要求了, 但是命令行下有缺陷: **刷屏 + 快速打开查找到的文件**

## vscode 的搜索其实很香

仔细一想, vscode 才是平时搜索体验拉满的状态:

- `F1` 打开 command 面板, 使用 `add folder to workspace`, 添加 `hyperf/docs/zh-cn` (中文版是最全的)
- `cmd+p` 搜索文件名, 对于熟悉 hyperf.wiki 的我而言, 这个是最常用的方法
- `cmd+shift+f` 全局搜索, 记忆力也不是一直 **有效**, 全局搜索+模糊搜索 就可以派上用场, 这里还有很多相关功能可以使用

## 写在最后

- 因为最近使用 dash 比较多, 典型的 心理误区+路径依赖: 手里拿着锤子, 看啥都是钉子
- 一个新的问题, 不断调用现有的知识, 能快速判断所有的分支/选项, 不断的寻找最优解
- 尽可能多利用「知识」来做决策, 而非「猜」 -- 后者需要花费的时间太不可控, 成功的概率非常低
- 准确和效率往往是一对反义词, 用生物学来解释 -- 数万年的进化让大脑形成的节能模式 + 身体的非条件反射(条件反射的自己去面壁)

link:

- [hyperf wiki](https://hyperf.wiki)
- [docsify](https://docsify.js.org): 使用 markdown 编写开发文档
- [dash docset](https://kapeli.com/docsets#dashDocset)
  - [Any HTML Documentation](https://kapeli.com/docsets#dashDocset)
    - [dashing](https://github.com/technosophos/dashing#readme)
    - [dash-docset-builder](https://github.com/godbout/dash-docset-builder)
- [ag, the_silver_searcher](https://github.com/ggreer/the_silver_searcher)
- [z jump](https://github.com/jethrokuan/z)
