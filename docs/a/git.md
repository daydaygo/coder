# git

## git basic

- [git 简明指南](http://www.bootcss.com/p/git-guide/)
- repo
  - 工作区(work space, 编辑区): 当前文件都属于编辑区, 对文件编辑(增删改)后, 文件进行 `change` 状态
  - 暂存区 stash: 可以将 `change` 状态的文件, 放到暂存区
  - 提交区 commit: 可以将暂存区的文件修改, 固化为 commit
    - commit 是重点, 记录了一次修改的所有变动
    - branch/tag/head 等概念, 都只是指向 commit 的指针
    - merge & conflict: 不同的 commit 可以合并到一起, 但是不同 commit 改动了相同文件, 就可能引起冲突
  - commit: [commitlint](https://commitlint.js.org/)
  - merge: `fast forward` 删除分支后会丢掉分支信息 `--no-ff`
  - 分支管理: bug/hotfix feature rebase(commit整理成直线) [gitFlow git协作流程](http://kb.cnblogs.com/page/535581/)
  - 分布式
    - origin 默认远程仓库名
    - upstream: github pr 默认的原始项目分支
- config
  - `--global` `~/.gitconfig`; `--local` `.git/config`
  - `.git/hooks/`
  - `.gitignore`
- other
  - [submodule](http://www.jianshu.com/p/b49741cb1347): `.gitmodules`
- [monorepo](https://zhuanlan.zhihu.com/p/77577415)
  - build: buck/bazel install依赖 增量编译/编译缓存 run执行任何一个项目
  - 适合迭代: 分库后一次工作要在多库中处理完成
  - 不存在版本升级->都用最新的库
  - 第三方库升级?
  - 大文件: git lfs
- 工具链
  - [github desktop](https://desktop.github.com/): **首推** 图形化工具
    - 查看修改, 图形界面远超命令行
    - 快速的切换: 切 change/history/branch/pr/repo
    - 快捷键支持快速完成 git 操作: commit/pull/push/clone/branch/repo
    - github 支持: 快速 clone/pr
    - 设置使用 vscode 打开文件, 可以快速解决冲突
  - git 命令: 所有命令都可以通过 git 命令完成
  - idea ide 中的 git 功能
    - `action` 输入 git 即可进入 git view, 查看支持的 git 功能
    - git history: 查看当前文件的所有历史修改记录
    - git annotate: 查看代码的最后修改作者, 定位 `trouble maker`
  - gitk: 就一个功能, 命令行下查看某个文件的历史修改记录, 使用 `brew install git-gui` 安装

![Git四个区五个状态以及之间的变换](https://img-blog.csdn.net/20171212193726546?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvenJjMTk5MDIx/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![git 常用操作图](https://img-blog.csdn.net/20141119105906092?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQva2VoeXVhbnl1/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/Center)

## git cli

```sh
# repo 初始化
# 本地
git init .
# 远程
git clone <repo-address> <dir>
# 2 种 git clone 的方式
git clone git@github.com:daydaygo/coder # ssh
git clone http://git.coding.net/daydaygo/coder # http

# 编辑区
git status # 哪个文件修改了
git diff # 具体的修改是啥
git checkout . # 取消对文件的修改

# 暂存区: 文件修改 进/出 暂存区
git stash
git stash pop

# 提交区
# 之所以不把 add 放到暂存区讲, 因为 add 用法比较固定, add 和 commit 组合执行完成 commit
# 提交所有修改到暂存区
git add -A
git commit -m 'xxx' # 提交 commit 并附带 commit 信息
git log # 查看 commit(当前分支) 历史

# branch/tag/head 等都只是指向 commit 的指针
git checkout -b xxx upstream/xxx # checkout 切换分支, 本质是移动指针到不同 commit

git tag -a <name> -m <msg> # 含附注tag, 推荐

# 和远程交互
git remote add origin <url> # 添加 origin 远程
git remote set-url origin <url> # 修改
git remote -v # 查看
git pull # 获取远程 commit
git pull --rebase # fatal: Not possible to fast-forward, aborting
git fetch # 获取远程内容, 这样可以 checkout 到远程才有的分支
git push -u origin <branch> # 推送本地分支到远程
git push --tags # 推送 tag 到远程
git push origin --delete <name> # 删除 branch/tag

git commit # --amend 修改提交信息
git commit --amend --author 'daydaygo <1252409767@qq.com>' # 修改最后一次 commit 的作者信息
git cherry-pick 62ecb3 # 合并某个 commit 到当前分支
git branch -m <old> <new> # 分支重命名
git branch | xargs git branch -D # 快速删除本地分支
git remote prune origin # error: there are still refs under 'refs/remotes/origin/xxxx'
git fetch -p origin # 同步删除本地分支
git checkout --orphan test && git clean -df # 全新分支
git log --pretty=short --graph
git log --grep='xxx' # 查询 commit msg
git clean -f # 清除本地修改
git rm -r --cached .idea # 配合 .gitignore 修改
git mv -f file file2
git rebase origin/master # 使当前分支等同于 origin/master
git rebase -i # 压缩历史，比如修正拼写错误
git reset --hard origin/dev2 # 重置本地分支到远程分支

# github PR 工作流
git remote add upstream <url> # 相当于 fork
git fetch upstream master && git rebase upstream/master # 获取原库最新更新, 并 rebase 使 commit 合并到一条线里
git checkout -b <branch> # 新建工作分支
# 完成后原库提 PR

# git gc
# 查找大文件
git rev-list --objects --all | grep "$(git verify-pack -v .git/objects/pack/*.idx | sort -k 3 -n | tail -5 | awk '{print$1}')"
# 删除文件
git filter-branch --force --index-filter 'git rm -rf --cached --ignore-unmatch bin/nspatientList1.txt' --prune-empty --tag-name-filter cat -- --s
# 使用 gc 再次压缩
git gc --prune=now

# https://rtyley.github.io/bfg-repo-cleaner/
git clone --mirror git://example.com/some-big-repo.git
bfg --strip-blobs-bigger-than 100M --replace-text banned.txt repo.git
cd some-big-repo.git
git reflog expire --expire=now --all && git gc --prune=now --aggressive
git push

git cat-file # .git 目录

git submodule update --init --recursive

# clear history 清除所有历史记录: 处理完后 gitee 进行 git gc
git checkout --orphan new # orphan 孤儿
git add -A; git commit -am 'init'
git branch -D master; git branch -m master
git push -f origin master; git push -u origin master
```

## git config

```conf
user.name "daydaygo"
user.email "1252409767@qq.com"
push.default simple
color.ui "always"
credential.helper "store" # 保存 https 账号/密码
format.pretty "oneline"
core.quotepath false
core.autocrlf false # 自动转换换行符
core.safecrlf true # 混合换行符警告
core.ignorecase false # git 忽略文件名大小写，导致 window 下修改 linux 下报错
core.filemode false
alias.s status # --unset
gui.encoding utf-8 # gitk 中文乱码
help.format web # 使用 web 查看帮助
web.browser open # 设置 web
```

## github

- `?` shwo shortcut
- search: `i` `org:dotnet filename:ConfigurationBuilder.cs`
- issue: comment pr
- url：compare-commit; diff patch
- comment 评论: `R` quote `:` 表情
- [github 国内加速 9 种方式](https://mp.weixin.qq.com/s/FHCzOA72VsV4ePHA7cDjVg)
- [Git Clone加速三种方式](https://www.cnblogs.com/XT-xutao/p/12134045.html)
  - 浅拷贝: `git clone --depth=1`
  - 获取 `github.com global.ssl.fastly.net` 的 ip, 添加到 hosts 中
  - 借助 gitee
- github 数据: <https://vesoft-inc.github.io/github-statistics/#Star>
- [Trending](https://github.com/trending)
- [fossa: Open Source Management](https://fossa.com/): 包括项目使用的开源协议
- [github bot](https://github.com/hyperf/github-bot): 给 pr/issue 添加部分自动化功能
- github action
  - action market: <https://github.com/marketplace?type=actions>
  - awesome-action: <https://github.com/sdras/awesome-actions>
  - 官方文档: <https://docs.github.com/en/free-pro-team@latest/actions>
  - 入门教程: <http://www.ruanyifeng.com/blog/2019/09/getting-started-with-github-actions.html>
- 添加github小图标: <http://shields.io/>
- 在Github个人资料页面上启用自述文件: 创建一个与你的Github账户用户名相同的新仓库, 添加 `readme.md` 文件
- [dflydev/git-subsplit](https://github.com/dflydev/git-subsplit): 封装 `git subtree` 为 `git subplite` 命令, 方便使用
  - [laravel 中使用 git subsplit 的示例](https://github.com/laravel/framework/tree/5.1/build)
- [change github tab size](https://stackoverflow.com/a/23522945/15009997)

```sh
set -gx GITHUB_TOKEN xxx # gh/hub 都可以用

# https://cli.github.com/manual
brew install gh
gh auth login # GITHUB_TOKEN
# gh completion -s fish > ~/.config/fish/completions/gh.fish # brew fish auto set completion
gh gist

# https://hub.github.com
# hub --help: see github commands
brew install hub
hub delete [−y] [ORGANIZATION/]NAME
git config --global hub.protocol https # https ssh

# git.io 短网址
curl -i https://git.io -F url="https://github.com/daydaygo/coder" -F "code=CoderAtWork"
```

## gitlab

- revert 操作也可以被 revert 掉
