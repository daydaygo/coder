# fishshell

* <http://fishshell.com/docs/current/tutorial.html>
* 使用 `help xxx`, 直接查看 fish 帮助, 如 `help set`

## use

``` sh
# fix mac tab slow https://github.com/fish-shell/fish-shell/issues/6270
function __fish_describe_command; end

# fisher: https://github.com/jorgebucaran/fisher/releases/latest
fisher install jethrokuan/z
fisher install evanlucas/fish-kubectl-completions

# env 环境变量
set # -x export; -u unexport; -U universal; -g global
set -U fish_user_paths /usr/local/bin $fish_user_paths # fish_user_paths 会添加到 PATH, 一次添加, 永久有效
```

* config
  * ~/.config/fish
    * config.fish 主配置
    * conf.d 拆分配置
    * fish_variables
    * functions
    * completions

## mark

```sh
# set fish_greeting 'Talk is cheap. Show me the code. -- Linus'
function fish_greeting
  set a 'Stay hungry, Stay foolish. -- Whole Earth Catalog' 'You build it, You run it. -- Werner Vogels, Amazon CTO' 'Talk is cheap. Show me the code. -- Linus' 'go big or go home; done is better than perfect -- facebook' 'Eating our own dog food -- Microsoft'
  echo $a[(random 0 4)]
end

abbr -a rm 'rm -rf'
abbr -a cp 'cp -r'
# 防止 PATH 重复设置
set PATH ~/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin
# https://github.com/fish-shell/fish-shell/issues/6270
function __fish_describe_command; end

# cli tool: https://mp.weixin.qq.com/s/i3qEnIF9XeKstKFebKlsDw
# https://github.com/jorgebucaran/awesome.fish
set -gx __done_min_cmd_duration 5000  # default: 5000 ms
alias fda 'fd -IH'
alias rga 'rg -uuu'

# git
abbr -a 'gc' 'git checkout'
abbr -a 'gb' 'git branch'
abbr -a 'gp' 'git pull'
abbr -a 'gm' 'git merge'
abbr -a 'gs' 'git status'
abbr -a 'gd' 'git diff'

# docker
abbr -a d docker
abbr -a doc docker-compose
abbr -a k kubectl

# php
set -g PATH /usr/local/opt/php@7.2/bin $PATH
abbr -a 'c' 'composer'
# hyperf
abbr -a 'h' 'php bin/hyperf.php'
abbr -a 'ht' 'php bin/hyperf.php t'
abbr -a 'hs' 'php bin/hyperf.php start'

# go
set -gx GOPATH ~/go
set -gx GO111MODULE on
# set -gx GOPROXY https://mirrors.aliyun.com/goproxy/
set -gx GOPROXY https://goproxy.cn,direct
set -g PATH ~/go/bin $PATH

# py
set -gx PATH /usr/local/anaconda3/bin $PATH

# gitlab 快速发起 pr
function pr
    set from (git symbolic-ref --short -q HEAD)
    set to $argv[1]
    open 'https://${project_url}/-/merge_requests/new?merge_request%5Bsource_branch%5D='$from'&merge_request%5B&merge_request%5Btarget_branch%5D='$to
end
```
