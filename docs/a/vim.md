# vim

- normal
  - wander: `jkhl 0^$ gg/G wWbBeE %`
  - edit: `y p x d r J = u/ctrl-r gu/U`
- insert: `aA iI oO`
- visual: `v V ctrl-v`
- command
  - regex
  - find: `/? n N`
- macro 宏

```sh
100G 100gg # 跳转到第100行
y2fB # 复制光标到第二个大写B中间的内容
xp # 交换2个字符
di' # 删除引号内的内容 -> -> c/d/y/v i/a/t/f '/"
ggVG # 全选

% # 括号匹配
| # column
c # change
=<> # indent
ctrl-r # redo
J # join line

%s/$/sth/ # 在行尾追加sth
%s/\^M//g # 替换掉dos换行符，\^M使用ctrl+v  + Enter即可输入
:g/\^\s*$/d # 删除空行以及只有空格的行
%s/#.*//g # 删除#之后的字符

:verson
:!
:h xxx
:e -> c-d -> tab # completion
:sp :vsp C-w(switch) zz-光标居中 # 窗口
:set xxx
:set noxxx
:set clipboard=unnamed # use systemClipboard
:set paste # 粘贴时不自动换行
```

- /etc/vimrc ~/.vimrc

```sh
syntax on
set autoindent
set cindent
set ru
set number
set cursorline
set cursorcolumn
autocmd BufNewFile *.py,*.sh, exec ":call SetTitle()"
let $author_name="daydaygo"
let $author_email="1252409767@qq.com"

func SetTitle()
  call setline(1, "\####################################")
  call append(line("."), "\# File Name: ".expand("%"))
  call append(line(".")+1, "\# Author: ".$author_name)
  call append(line(".")+2, "\# Email: ".$author_email)
  call append(line(".")+3, "\# Created Time: ".strftime("%c"))
  call append(line(".")+4, "\#================================")
  if &filetype == 'sh'
    call append(line(".")+5, "\#!/bin/bash")
  else
    call append(line(".")+5, "\#!/usr/bin/python")
  endif
    call append(line(".")+6, "")
  autocmd BufNewFile * normal G
endfunc
```

- Treat **vim** as a programming language, and use it everywhere
- vim 键位图: http://cenalulu.github.io/linux/all-vim-cheatsheat/
- https://vim-adventures.com/
- https://github.com/mhinz/vim-galore
- vim mastery: https://www.bilibili.com/video/av9406050/
- https://github.com/amix/vimrc
