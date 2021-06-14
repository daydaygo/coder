# idea 全家桶

> 定位: IDE

- <https://www.jetbrains.com/help/>
- github 开发者申请: <https://www.jetbrains.com/opensource/>
- edu 申请: <https://www.jetbrains.com/zh/student/>
- early access program, EAP: <https://www.jetbrains.com/resources/eap/>

tips:

- 安装推荐使用 [toolbox](https://www.jetbrains.com/toolbox-app/)
- 不同语言推荐使用 idea 下不同 ide, 不同语言开发体验更友好

资源:

- <https://cn.intellij.tips/>
- B站 - JetBrains 技术布道师范圣佑: <https://space.bilibili.com/414846001>

## learn

> 官方自带的 `ide features trainer`

- edit: context/indentAction `⌥↩︎` action `⌘⇧A` selection `⌥↑↓`  comment `⌘/` line `⌘D ⌘⌫` move `⌥⇧↑ ⌘⇧↑` collapase `⌘- ⌘=` surround `⌘⌥T` unwrap `⌘⇧⌫` occur `⌃G`
- completion: basic `⌃␣` all `⌃␣␣` tab `⇥` postfix `.` smart `⌃⇧␣` F-string `${str}` smartType
- refactor: `⌃T` rename/var/method fix `⌥↩︎`
- assist: format `⌘⌥L` param `⌘P` quick `F1 ⌥␣` error`F2 ⌘F1 ⌥↩︎ ⌘6` symbol `⌘⇧F7`
- navigation: usage`F7` declar/usage`⌘B ⌥F7` tool/view`⇧⎋ pin ⌘3` struct`⌘7 ⌘F12 ⌘;` recent`⌘E ⌘⇧E` search`⇧⇧ ⌘O F1` hierarchy occur`G`
- run/debug: conf debug-workflow
- icon: 一级(c const; t type i interface; f function) 二级(f field; m method)
- other: ideFeatureTrainer liveTemplate

---

- go: exprNil errNil `.nn .rr` `.forr` `.else` `fori` `if`
- clion: rust emmylua(与其他插件不兼容)
- idea: projectSetting`cmd+;` `sout`

## mark

> 原则: 能用 `action` 解决就用 action, 常用操作 action 执行多了, 自然会记下快捷键

- **action**: 一会降十力, 四两拨千斤, 用来查找功能
  - 快捷键 `cmd-shift-a`, 支持模糊匹配
  - 下面内容无特殊说明, 可通过 action 操作
  - 搜索出来的 action, 有快捷键会显示快捷键, 可以快速配置快捷键
  - 推荐使用 `distraction free mode`
  - always select `opened file`: `project`(`cmd-1`)中自动定位到当前打开的文件, 也可以使用 project 中的 `target` 图标 :o:
  - `show diagram` 查看类图: 内容极度舒适, 还可以编辑
  - `learn` 官方交互是教程, 推荐
- coding: 生产力
  - completion: `⌃␣` `⌃⇧␣`
  - generate: `⌘N`
  - refactor: `⌃T`
  - vcs: `⌃V`
  - live template: `⇥`
  - inspection
  - intent(`action|show context action`): `⌥⏎` `⌥⇧⏎`
  - appearance:
    - `⎋`editor; `⌘+NUM` tool window; `F12` editor->tool window; `⌘E` recent file
    - `F2` error; `⌥↑` selection; `⌥F7` find usage
    - run/deubg/test/profile: `⌃R` `⌃⌥R` `⌃D` `⌃⌥D`
    - bookmark: `⌘+2` favorite; `F3` toggle bookmark; `⌘+F3` show bookmark
  - 快速学习: `plugin|feature trainer` `action|learn` `view|tool|learn`
- setting: `cmd-,`(mac 通用快捷键)
  - `Appearance & Behavior | Appearance | Always show full paths in window header`
  - keymap 中支持 `find actions from shortcut`
  - `spelling` 设置为 `application level`
  - 配置区分 `current/new project`
  - `live/file/code template`: 设置代码模板, 比如 插入文件创建信息/doc注释
  - keymap | always use fn
  - comment `leading space`
  - code style `set from`
  - `code folding`: 代码块折叠
- 搜索的地方都支持 **模糊搜索**, 技巧为 **单词前缀匹配**; 搜索还支持 **正则匹配**
  - speed search: `↑ ↓ ⏎ ⎋`
- 支持 [.editorconfig](https://editorconfig.org/) 配置文件
- 文件不同步: 文件右键 | sync
- vcs: `git history` `git annotation`

## menu

> menu 本身就暗含了 idea 的功能分类, 分类则方便了大家理解和使用

- file
- edit: 通用编辑功能, find/replace line(duplicate/join/delete)
  - 多点编辑
    - `option-鼠标左键` 多选
    - `ctrl-g` add selection for next occur
- view: 视图, 可以认为需要视图进行展示的功能, 都在这里
  - editor: 使用 `⎋` 返回
  - action indicator: `⌥⏎` `⇧⌥⏎`
  - tool window: `⌘1` `F12` `⇧⎋` `⇧⌘F12`
    - tool window bar
    - arrange tool window: view mode; resize `⇧⌘←, ⇧⌘→, ⇧⌘↑, ⇧⌘↓, ⇧⌘'`
  - appearance
    - navigation bar: `⌘↑`
    - status bar
    - tool bar: run/debug vcs
    - backgroud image
  - context menu: `right-click`
  - popup menus: `⌘N` generate; `⌃T` refactor; `⌘N` new; `⌃V` vcs -> quick list
  - 切换到其他 view 时, 通常可以使用 `esc` 快速回到 editor 继续进行编辑
- navigate: 跳转, 快速直达, **效率神器**, 快速查看 action/file/code(class/method/line/usage)
- code: 针对代码/代码块的各类编辑操作, implement/generate/completion/reformat
- refactor: 重构 -> 使用重构而不是编辑, 利用 ide 能力可以尽量减小代码改动时产生的错误
- run: run/debug
- tools: 辅助开发工具, 比如 http client(编程实现 http 请求)
- vcs
- window: 可以类比 iterm 等工具的 window 管理
- help

一些 tips:

- view 和 window 比较好区分, 明显的一大一小, 并且 window 包含的功能比较少, 可以采用 **非 window 即 view** 来帮助理解
- 不要小瞧了 run 部分的功能, `coding = writing + running`, 往往是 run 这部分花费非常多的时间, 这也是 CI/CD/devops 着力解决的问题之一
- 多点编辑是个炫酷大于实际的功能, 纯粹的多点编辑(edit 里的, code/refactor 里针对 coding 进行) 使用场景非常有限, `option-鼠标左键` 往往就足够了

## plugin

- 不要轻易关闭默认插件, 除非对插件的功能非常熟悉, 比如下面这条
- github: gitignore 忽略的文件, 在文件列表会置灰, 需要此插件支持
- acejump-lite: 只用其中一个功能, 绑定快捷键 `ctrl-;` 进行字符(char)跳转
- ide features trainer: 推荐积累一定代码量后使用, 会发现超多省时间的 ide feature
- ide setting sync: 配置同步, 一个ide开多个项目的不需要再频繁修改配置了
- key promoter x: 鼠标操作是如果有可用快捷键, 右下角弹窗提示
- save action: fmt
- ideavim

## goland

- `file watch`: 添加 go fmt current file
- `goimports`: 设置 soring type 为 gofmt

## phpstorm

- setting
  - inspections 中关闭 `phpdoc`: 不要为了些 doc 而写 doc
  - add package as lib: 允许编辑/定位 vendor 下的文件

## datagrip

> sql 从此也有了 ide 般的强大体验

- session: dbConnection
- 推荐理由: 强大的自动补全
- export/import: file/db
- data source: `⌘;`
- 执行当前语句: `⌘⏎`
- explain raw: `⌃⏎`
- 参数绑定: parameter `where id=:id` `where id=?`
- foreign key: 不建议使用

---

- datagrip
  - first step: connect query-result-export/import localDB MSSQL diagram shortcut
  - quickstart: install(toolbox) > clone: git(dump file) + docker(compose, db image) > attach dir with sql script > write code
- db coon `⌘;`
  - inner + JDBC/lib path/DDL
  - manage: view/group/color/filter copy/share/import/export
  - conf: jdbc option-keepalive/refresh mode-readonly ssh/ssl
  - session
- run `⌘⏎`
  - query bind `where id=:id` `where id=?`
  - query console `⌘↓`; history `⌥⌘E`; refresh `⌘R`
  - run: `⌃R` `⌃⌥R` / to file / procedure / run conf / tests / debug / transaction / migration
  - query result: in-editor / new tab / title / edit value / export file/clipboard / compare / relate `⌘B` `⌥⌘B` / pin / sort / table view / column / export / insert / log
- data
  - export: object(file/DDL/clipboard) / data(dump) / copy table `F5`
  - import: sql file(dialog) / csv / resotre dump
  - Data extractors set rules of how to copy or view your data in the editor
- design
  - table: create/drop / modify `⌘F6` / copy / filter `⌘F` / view data / compare
  - column row cell pk index fk schema user/role
- file
- code
- ide conf
- plugin
  - big data tool

## idea product

- ide - idea-base
- space - team
- kotlin
- ide - dot

## docs

> 以 goland 为例: <https://www.jetbrains.com/help/go>
> A-action P-perference(setting/config) S-search

- goland
  - first step: quick start > go setup > GOROOT GOPATH > go mod > go tool > search > debug > profile > shortcut > get help / FAQ
  - install: toolbox register(比验证码复制来复制去方便)
  - quick start
    - project
    - UI: project(view|tool windows) editor VCS-actions(view|toolbar) gutter(debug/run) scrollbar(code inspections) statusbar
    - code: refactor complete generate live-template inspection intention-action
    - run debug test benchmark
    - keyboard shortcut
    - work offline / accessibility / get help / FAQ
- ide config (level-global/project backup/restore/share)
  - appearance (editor / action indicator/list / navigation bar / status bar / tool window / context menu / popup menu / main window)
    - UI theme
    - menu & toolbar: customize / quick list(action|name of quick list)
    - status bar
    - tool window: project/structure/run/database open/hide bar/button/component arrange view-mode speed-search
    - view mode: presentation distraction-free full-screen zen
    - background image / touch bar support / linux naive menu
  - code style: set-from scope format .editorconfig wrap
  - color & font: color-scheme highlight color-scheme-font/console-font
  - file watcher: auto run compiler/formatter/linter when file change
  - shortcut: add/reset location conflict
  - plugin: setting install/remove/disable search-tips(`/`)
    - popular: tech keymap theme
  - file encoding / file type association / ignore file/folder / scope(a group of file/folder) / browser(Use a custom profile)
  - advanced config: GOLAND_VM_OPTIONS GOLAND_PROPERTIES GOLAND_JDK / Configuration directory
- project config (Whatever you do in GoLand, you do that in the context of a project)
  - project
    - `.idea` folder
    - new
      - go module: vgo / UML diagram / ENV(get info)
      - wasm
    - import: local/VCS
    - setting: global / current project / new project
    - Save projects as template
  - go conf
    - go tool: fmt goimports generate vet
    - go template
    - GOROOT & GOPATH
    - Build constraints: `// +build linux,amd64,go1.12,!cgo`
    - vendor buildin: `>=Go 1.14 RC`
  - content root (a collection of files)
    - type: resource excluded
    - attach/detach/mark/reset dictionary
  - favorite
  - index: exclude > file: `mark as plain text` / folder: `Mark Directory as | Excluded`
  - invalidate cache
- coding
  - action / search everywhere / new class/file/package/scratch / select code construct / select occurrence / tab/indent / copy/paste / create a type / line of code / code statement / code fragment / multiple caret / param hint / code folding / autosave
  - auto imports / exclude
  - editor basic: scrollbar/breadcrumb/gutter/tabs / navigation: panel/tool-window/terminal/layout / tabs(copy path) / split screen / quick popup: definition/doc/context/err/tooltip/usage/hector / editor conf / working with text
  - lightEdit mode
  - src navigation: caret / move caret / recent location / bookmark / declaration(package-level)/usage / implement / popup: select in / locate a file / err / structure view / method / lens mode / breadcrumb / line/column / file path / recent file
    - code hierarchy: type call
    - structure / file structure
  - search: find/replace caret/selection/file/project list/new-line/case/RE occur find-tool-window usage structural(db template) everywhere(`/` abbr) / db full-text search
  - reformat rearrange / comment leading space
  - completion: basic function smart statement hippie postifx reference hierarchy / fill struct
  - generate `⌘N`: action / live template / surround / paired character
    - action: constructor getter/setter implement test copyright
  - refactor `⌃T` (resolve conflict)
    - change signature: name param / default value / return value
    - extract: constant interface method var(introduce)
    - inline move/copy rename / safe delete
  - inspection: scope(order)/severity problem(`F2`) widget quick-fix inspection-profile
  - intention action `⌥⏎` `⌥Space`
  - live template: predefined var/function / emmet
  - compare file
  - reference: definition param inlay-hint quick-doc ext-doc type / copy tooltip
  - todo
  - Language injections let you work with pieces of code in other languages embedded in your code
  - scratch: Sometimes you may need to create temporary notes or draft up some code outside of the project context
  - Macros provide a convenient way to automate repetitive procedures you do frequently while writing code
  - copyright
  - proofreading: spellchecking grammar
- run: conf / log option
- debug
  - before: debugger conf
  - procedure: breakpoint > run in debug mode > use the debugger to get the information about the state of the program and how it changes during running
  - breakpoint `⌘F8`: Use breakpoints for debug printing / Set logging breakpoints more quickly / Add breakpoint descriptions / Group breakpoints
  - start debugger session: debug non-responding app / do more with pause /(`step`) / run before-launch task
  - examine suspended program: debug window `⌘5` / frame / var(copy/set`F2`/jump to source/evaluate expression/watch/inline) / Return to the current execution point`⌥F10`
  - step: over`F8`/out`⇧F8`/force step over`⌥⇧F8`/into`F7`/smart into`⇧F7`/run to cursor`⌥F9`/force run to cursor`⌥⇧F9`
  - go coredump / window minidumps / mozilla rr / profiler label / gops
- deploy: remote host / ssh
- test (gotest gocheck gobench)
  - new: Generate`⌘N` / Navigate | Test`⇧⌘T`
  - run: run`⌃⇧R`/stop`⌘F2` test / debug failed test
  - result: explore`⌘4` / sort&filter / track / manage / statistic / jump declaration / previous
  - code coverage
  - test RESTful: http client plugin
  - testify toolkit
- prifilling: an analysis of your program performance; cpu/mem/blocking/mutex
- integrated tools
  - VC`⌘9`: compare / resolve conflict / issue tracker / changeist / shalve change / patch / review change / anotation(date format) /
  - git: sync(fetch/pull/update) commit/push investigate(history/commit/merge/author) branch(merge/rebase/cherry-pick/separate) conflict undo(revert/drop/reset/previous) tag history(commit message/amend/squash/drop/rebase)
  - github: accout fork-PR-rebase gist
  - local history
  - database tool & SQL -> [datagrip](#datagrip)
  - docker: image container compose
  - terminal`⌥F12`
  - IDE scripting console: you can think of it as a lightweight alternative to a plugin
  - external tool
  - issue tracker / task / time tracking / context
  - textMate plugin: syntax highlight
  - command-line interface: toolbox / arg
  - feature trainer: `action|learn`
- no-go tech
  - json: json5 / json schema
  - html: new reference fragment doc preview
  - shell
  - k8s: completion inspection quick-fix live-template helm resource-define
  - xml: validata DTD doc reference
  - js ts css
- tutorial
- [reference](https://www.jetbrains.com/help/go/ui-reference.html)
- release note

![appearance](https://www.jetbrains.com/help/img/idea/2020.2/go_main_window.png)
![UI](https://www.jetbrains.com/help/img/idea/2020.2/go_QST_lookAroundThumb.png)
![setting for current project](https://www.jetbrains.com/help/img/idea/2020.2/sett-pref-dialog.png)
