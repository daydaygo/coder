# devops| 开源之路: github PR 走起

> 想快速提高编程能力, 还不快来 「全球最大同性交友社区」~

## PR 只需几步

github 开源之路, 从 PR 开始, 只需要如下简单的几步:

- 找到心仪的项目, **fork** 它

![fork](https://upload-images.jianshu.io/upload_images/567399-7e1c7e601f13d8e8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 自己的仓库里就有了, **clone** 它

![clone](https://upload-images.jianshu.io/upload_images/567399-c1979393d0d839a1.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 最新的 master 拉一个分支, 修改代码

```
# 更新 master
git merge upstream/master

# 基于最新的 master 拉出新的分支进行开发
git checkout -b feat-xxx

# coding

# 提交
git add
git commit
git push

# PR 神器
```

- PR 神器参上 [github desktop](https://desktop.github.com/)

![github desktop](https://upload-images.jianshu.io/upload_images/567399-c374c98e927084da.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

完成 PR 只需要一步: `cmd+r` 快捷键

## 参与 PR

- 参与 PR

![参与 PR](https://upload-images.jianshu.io/upload_images/567399-28e5e817a2d9fe92.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

使用 `cmd+b` 快捷键, 切换不同分支, 包括 PR

在 [hyperf-cloud/hyperf](https://github.com/hyperf-cloud/hyperf) 主项目上, 使用 github desktop 就可以切换到 PR, 参与到 PR 中

## github desktop 小技巧

- cmd+1 : change 界面, 查看代码修改
- cmd+2 : history 界面, 查看历史提交
- cmd+t : 切换不同仓库
- cmd+b : 切换不同分支, 包括 PR
- cmd+r : 提交 PR

其他功能自行探索, 不过我感觉上面几个就够用了, 其他我都是使用命令行, 键盘啪啪啪就行了

## tips

- upstream 是啥?

git 是 **分布式** 的, 本地就是一个完整的仓库, 那么如何和其他人一起工作呢? -- 需要一个远程仓库, 它也是一个完整的仓库, 用它来和我们的仓库之间进行同步

```
# 查看 git 本地配置
git config --local -l

# 可以看到这个
remote.origin.url=git@github.com:daydaygo/hyperf.git
remote.origin.fetch=+refs/heads/*:refs/remotes/origin/*
```

这个 `origin` 分支就是我们用来同步的远程的分支, 正常情况下, 我们只需要一个 `origin` 就可以了, 至于为啥叫 origin, 约定俗成.

正常情况下, 我们有 origin 就够用了, 但是在 PR 的场景下, 我们的 origin 是自己账号里的仓库, 比如 `daydaygo/hyperf` 是从 `hyperf-cloud/hyperf` fork 出来的, 我们推送到 origin, 只是推送到自己的仓库, 并没有到开源项目里, 这个时候就需要 `upstream` 这个分支了

怎么玩了呢? 如果使用 github desktop, 这步自动搞定了, 如果遇到没有的情况

```
# 添加 upstream 
git remote add upstream https://github.com/hyperf-cloud/hyperf

# 如果需要修改
git remote set-url upstream https://github.com/hyperf-cloud/hyperf

# 添加好后, 会看到这个
git clone --local -l

remote.upstream.url=https://github.com/hyperf-cloud/hyperf
remote.upstream.fetch=+refs/heads/*:refs/remotes/upstream/*
```

- 提 pr 时分支是不是需要同名

稍微熟悉一点 git, 就会知道 branch/tag 等都只是指向对应的 commit, 提 PR 时, 其实是把你分支上的 commit 和你需要提交 PR 的分支上的 commit 进行合并, 所以从 master 上拉新分支时, 想取啥名字无所谓, 更为重要的是, 你要更新到最新的 master, 和开源项目的 master 保持一致:

```
# 和开源项目的最新代码保持一致
git merge upstream/master

# 再拉分支进行开发
git checkout feat-xxx # 想叫啥都行, 关键是你要合并到开源项目的哪个分支里
```

至于为啥叫 upstream, `约定俗成` 而已, 这样的套路, 在 coding 的过程中比比皆是, 看多了自然就有感觉了.

- 愉快的参与开源, 成为 contribute

设置 `用户名 + email` 和 github 账号里保持一致, 这样在 PR 被合并进去的时候, 就可以在开源项目的 contribute 里看到自己啦

![github contribute](http://qiniu.dayday.tech/20190910232920.png)

- PR 收到了修改意见怎么办

在自己原来的分支上继续修改, push 过后, PR 里就会自动同步啦

## 写在最后

简单吧, 交友哪有那么难, 只要掌握一点 **小技巧**

国内开源的开发组都很热烈欢迎 **小鲜肉** 的加入, 几乎很容易拿到开发组的联系方式, 然后进行 **深入交流**, PR 从来都不是开源之路上的障碍.