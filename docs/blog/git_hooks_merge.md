# devops| git hooks 实战: 防分支 merge

基于工作中 git 工作流遇到的问题, 实战 git hooks, 防止测试分支合并到开发分支

> devops| git hooks 实战: 防分支 merge: https://www.jianshu.com/p/287eed8aa1a8

先推荐一本书, [Pro Git>](https://git-scm.com/book). 接触并使用 git 的过程中, 会不断积累一个又一个 git 命令, 一个又一个 git 知识点, 而这本书, 可以让学习到的 git 知识 **连接** 起来.

> [曾经我以为会几条git命令就算是掌握git了，然后遇到一些问题时直呼『还有这种操作』，比如SVN切换到git，比如限制分支merge，比如看段代码谁闯的祸（git blame）。是的，这一切，都是你需要这本书的理由，当然如果你对 gfs 也感兴趣的话，那可以好好折腾了。](https://book.douban.com/subject/26208470/)


## 问题

项目切换到 git 后(再次重复一遍, **人生苦短, 快用 git**), 效率大大提升, 需求和分支增长速度明显提高. 虽然切换到 git 之前进行过 **技术讨论** 协商了一套 git 工作流, 类似 [git flow](https://www.git-tower.com/learn/git/ebook/cn/command-line/advanced-topics/git-flow). 不过由于 **节奏很快, 需求很多, 同时开发/测试的分支很多**, 所以会有一个 `rtest` 分支, 作为集成测试使用.

> 问题在于 rtest 分支上可能有多个需求做集成测试, 所以这个分支不允许被合并到其他分支, 但是 git merge 合并反分支的情况时有发生.

## git hooks

熟悉 github 可能听过 webhooks, 可以在 **pull request** / **merge master** 等几个场景下, 设置异步回调通知(http 请求). 这个背后就是 git hooks 在起作用.

git hooks 采用 **事件机制**, 在相应的操作(比如 **git commit** / **git merge**)下触发, 分为 2 种:

- 服务端 hooks, github 的 webhooks 就是在此基础上建立起来的
- 客户端 hooks, 每个 git 版本库的 `.git/hooks/` 文件夹下就有可以使用的例子

**注意: 客户端 hooks 并不会同步到版本库中**

## 实战: 不允许分支合并

查看 [git hooks 官方文档](https://git-scm.com/docs/githooks) 可知, git merge 时会触发 `commit-msg` hook.

> This hook is invoked by git commit and git merge, and can be bypassed with the --no-verify option. It takes a single parameter, the name of the file that holds the proposed commit log message. Exiting with a non-zero status causes the command to abort.

git 版本库在 `.git/hooks/` 目录下都内置了几个常用 hooks 的示例, 比如 `.git/hooks/commit-msg.sample`:

```
#!/bin/sh
#
# An example hook script to check the commit log message.
# Called by "git commit" with one argument, the name of the file
# that has the commit message.  The hook should exit with non-zero
# status after issuing an appropriate message if it wants to stop the
# commit.  The hook is allowed to edit the commit message file.
#
# To enable this hook, rename this file to "commit-msg".

# Uncomment the below to add a Signed-off-by line to the message.
# Doing this in a hook is a bad idea in general, but the prepare-commit-msg
# hook is more suited to it.
#
# SOB=$(git var GIT_AUTHOR_IDENT | sed -n 's/^\(.*>\).*$/Signed-off-by: \1/p')
# grep -qs "^$SOB" "$1" || echo "$SOB" >> "$1"

# This example catches duplicate Signed-off-by lines.

test "" = "$(grep '^Signed-off-by: ' "$1" |
         sort | uniq -c | sed -e '/^[   ]*1[    ]/d')" || {
        echo >&2 Duplicate Signed-off-by lines.
        exit 1
}
```

不允许合并 **rtest** 分支, 稍微修改一下 `.git/hooks/commit-msg.sample` 即可:

```
#!/c/bin/php/php
<?php

// var_dump($argv);
$str = file_get_contents($argv[1]);
// var_dump($str);
if (strpos($str, "Merge branch 'rtest'") !== false) {
    echo "can not merge rtest \n";
    exit(1);
}

// echo 'for test'; exit(2);
```

文件重命名为 `commit-msg` 即可生效,  执行效果如下:

![git-hook-commit-msg](https://upload-images.jianshu.io/upload_images/567399-51415a46e1386000.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

返回码不为 0 就会终止 `git merge` 命令的执行, 然后执行 `git reset --hard head` 即可撤销这次操作.

## 写在最后

除了使用 git hooks, 如果使用 gitlab, 也可以通过 gitlab 提供的 `merge request`, 减少发生此错误的情况.
