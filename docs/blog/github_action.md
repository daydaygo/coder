# github action 实战

入门可以优先参考 [阮一峰 - github action 入门](http://www.ruanyifeng.com/blog/2019/09/getting-started-with-github-actions.html), 快速运行跑起来, 会带来极大的信心, 进行后面的折腾

接下来是官方的资源:

- [官方文档](https://docs.github.com/en/free-pro-team@latest/actions)
- [action market](https://github.com/marketplace?type=actions)

建议的食用方式是: 
- 先食用官方文档中的 demo 部分, 这里基本是最佳实践, 而众所周知, 二八定律的普世特性, 基本上 demo 中的内容或者 demo 中的内容稍微改一改, 就几乎够用了
- 概念相关的部分: 阮一峰的教程非常不错, 几乎完全够用
- 官方 demo 没有提供的最佳实践? 使用 `action market` 或者 `awesome-action`, 记住: **不要重复造轮子**

踩到的坑:
- 注意环境 `env`: 默认 action 其实是在 `runs-on: ubuntu-latest` 的 Ubuntu 环境下, 引用的某个 action 使用自己构建 docker 镜像, 这样就形成 `本地 Ubuntu + docker 容器` 2套环境, 那么 `uses run` 等指令, 也要区分是在哪个环境中执行的

比如:

```yaml
name: Build and Deploy
on: [push]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2 # 这里 uses, 是在 Ubuntu 本地的

    - name: vuepress-deploy
      uses: jenkey2011/vuepress-deploy@master # 这个 action 使用了自定义的 docker 镜像
      env:
        ACCESS_TOKEN: ${{ secrets.CODING_TOKEN }}
        TARGET_LINK: https://cOZajDiTyR:${{ secrets.CODING_TOKEN }}@e.coding.net/daydaychen/mac/coder.git
        TARGET_BRANCH: master
        BUILD_SCRIPT: yarn && yarn docs:build # 这里 yarn 是容器内的
        BUILD_DIR: docs/.vuepress/dist/
```

## 写在最后

可供参考的资料:
- 入门教程: http://www.ruanyifeng.com/blog/2019/09/getting-started-with-github-actions.html
- action market: https://github.com/marketplace?type=actions
- 官方文档: https://docs.github.com/en/free-pro-team@latest/actions
- awesome-action: https://github.com/sdras/awesome-actions