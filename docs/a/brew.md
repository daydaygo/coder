# homebrew

> http://docs.brew.sh/

```sh
# install
xcode-select --install # 部分软件需要 `xcode command line tool` 支持: https://developer.apple.com/download/more
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
curl -O https://raw.githubusercontent.com/Homebrew/install/master/install.sh
# 修改 BREW_REPO 为 https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/brew.git 后执行 install.sh

# fishshell 下 brew 加速
set -gx HOMEBREW_BOTTLE_DOMAIN https://mirrors.sjtug.sjtu.edu.cn/homebrew-bottles/bottles
git -C (brew --repo) remote set-url origin https://mirrors.sjtug.sjtu.edu.cn/git/brew.git
git -C (brew --repo homebrew/core) remote set-url origin https://mirrors.sjtug.sjtu.edu.cn/git/homebrew-core.git
git -C (brew --repo homebrew/cask) remote set-url origin https://mirrors.sjtug.sjtu.edu.cn/git/homebrew-cask.git
```

- ori
  - https://github.com/homebrew/brew.git
  - https://github.com/homebrew/homebrew-core.git
- sjtu
  - https://mirrors.sjtug.sjtu.edu.cn/git/brew.git
  - https://mirrors.sjtug.sjtu.edu.cn/git/homebrew-core.git
  - https://mirrors.sjtug.sjtu.edu.cn/git/homebrew-cask.git
  - https://mirrors.sjtug.sjtu.edu.cn/homebrew-bottles/bottles
- tsinghua
  - https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/brew.git
  - https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/homebrew-core.git
  - https://mirrors.tuna.tsinghua.edu.cn/homebrew-bottles
- aliyun
  - https://mirrors.aliyun.com/homebrew/brew.git
  - https://mirrors.aliyun.com/homebrew/homebrew-core.git
  - https://mirrors.aliyun.com/homebrew/homebrew-bottles

- brew 常用操作

```sh
brew search xxx
brew install/upgrade xxx # --cask
brew info xxx
brew --repo # repo
brew --repo brew/core
brew link/unlink # multi-version

brew install git fish ag # 常用
brew install graphviz # prof
```

- install: brew+homebrew-core
- bottles=binary packages
- cask: macos app, 如 go miniconda iterm2 vscode google-chrome firefox
- taps=third-party repo

- brew install 会先执行 brew update, 可使用 ctrl-c 跳过
