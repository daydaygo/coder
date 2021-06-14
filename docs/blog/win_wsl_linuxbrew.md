# coder| 折腾回 win 开发环境之 linuxbrew

> 是否推荐 linuxbrew: **不推荐**

内容简介:

- why: 为什么尝试 linuxbrew
- how: 如何使用
- WTF: 踩到了哪些坑
- mark: 总结

## why: 为什么尝试 linuxbrew

- homebrew 官网推荐
- 刚从 mac 切换过来, 还是想使用熟悉的工具
- 由于国内感人的网速, 我手里积累不少常用的镜像源, 平时在使用 homebrew 的过程中, 也稍微接触过 linuxbrew, 有点印象

## how: 如何使用

目前可用的 linuxbrew 源:

> 清华大学开源软件镜像站 - linuxbrew: https://mirror.tuna.tsinghua.edu.cn/help/linuxbrew-bottles/

参考 homebrew 官网和镜像站提供的帮助文档, 就可以安装上 linuxbrew

- 安装

```bash
# 1. 安装 ruby
sudo apt install -y ruby

# 2. 获取 homebrew 安装脚本
# 官网提供的脚本
# /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
mkdir tmp
cd tmp
wget https://raw.githubusercontent.com/Homebrew/install/master/install.sh

# 3. 使用镜像源加速安装 linuxbrew
sed -i 's/https://github.com/Homebrew/brew/https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/brew.git/g' install.sh
bash install.sh
# 执行到 “==> Tapping homebrew/core” 时 Ctrl-C
export PATH=/home/linuxbrew/.linuxbrew/Homebrew/bin:$PATH # 将 brew 添加到 PATH
git clone https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/linuxbrew-core.git "$(brew --repo homebrew/core)"
# 再次执行 install
bash install.sh
```

- brew 简单使用

```bash
brew search xxx # 查找
brew info xxx # 查看软件包的信息
brew install xxx # 安装
brew install php@7.2 # 安装不同版本
```

`brew info` 查看到的信息很有用:

- 比如软件的官网
- 比如需要设置 `PATH` 等信息

## WTF: 踩到了哪些坑

- 第一个坑是 homebrew 和 linuxbrew 傻傻分不清楚, 是的, 因为犯了这个 **低级错误**, 安装使用 linuxbrew 环节浪费了大量的时间
- 第二个坑是镜像源并不是完备的, 经常会遇到软件 `failed to download` 的情况, 由于还耦合了上面的问题, 这个问题也导致花了大量时间
- 第三个坑是 linuxbrew 和原生包管理, 在文档等方面有明显的劣势, 遇到问题的 debug 难度以及解决手段自然就少了

最后贴一下 linuxbrew **调教一番** 之后的简单配置:

```fish
# 下面使用的 fishshell, 比 bash/zsh 等都好用, 推荐~
# ~/.config/fish/config.fish, 相当于 .bashrc / .zshrc

# 安装 linuxbrew 后, 设置使用 linuxbrew 生效
set -gx HOMEBREW_PREFIX "/home/linuxbrew/.linuxbrew";
set -gx HOMEBREW_CELLAR "/home/linuxbrew/.linuxbrew/Cellar";
set -gx HOMEBREW_REPOSITORY "/home/linuxbrew/.linuxbrew/Homebrew";
set -q PATH; or set PATH ''; set -gx PATH "/home/linuxbrew/.linuxbrew/bin" "/home/linuxbrew/.linuxbrew/sbin" $PATH;
set -q MANPATH; or set MANPATH ''; set -gx MANPATH "/home/linuxbrew/.linuxbrew/share/man" $MANPATH;
set -q INFOPATH; or set INFOPATH ''; set -gx INFOPATH "/home/linuxbrew/.linuxbrew/share/info" $INFOPATH;

# 下载失败时, 切换到原生源
function brew_ori
    set -e HOMEBREW_BOTTLE_DOMAIN
    git -C (brew --repo) remote set-url origin https://github.com/Homebrew/brew.git
    git -C (brew --repo homebrew/core) remote set-url origin https://github.com/Homebrew/linuxbrew-core.git
    brew update
end

# 切换到镜像源
function brew_thu
    set -gx HOMEBREW_BOTTLE_DOMAIN https://mirrors.tuna.tsinghua.edu.cn/linuxbrew-bottles
    git -C (brew --repo) remote set-url origin https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/brew.git
    git -C (brew --repo homebrew/core) remote set-url origin https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/linuxbrew-core.git
    brew update
end

# 使用 linuxbrew 安装的 gcc
set -g LDFLAGS "-L/home/linuxbrew/.linuxbrew/opt/isl@0.18/lib"
set -g CPPFLAGS "-I/home/linuxbrew/.linuxbrew/opt/isl@0.18/include"

# 安装 PHP7.2 后, 设置 PATH 生效
# set -g fish_user_paths "/home/linuxbrew/.linuxbrew/opt/php@7.2/bin" $fish_user_paths
# set -g fish_user_paths "/home/linuxbrew/.linuxbrew/opt/php@7.2/sbin" $fish_user_paths

# librdkafka
set -gx LDFLAGS "-L/home/linuxbrew/.linuxbrew/Cellar/librdkafka/1.4.4/lib"
set -gx CPPFLAGS "-I/home/linuxbrew/.linuxbrew/Cellar/librdkafka/1.4.4/include"
```

## mark: 总结

不推荐使用 linuxbrew 替代原生包管理工具