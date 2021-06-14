# win

- chrome shortcut: https://support.google.com/chrome/answer/157179?hl=en

```cmd
ftype /? # 查看帮助

assoc .md=markdown
ftype markdown="C:\bin\vscode\bin\Code.exe" %1
```

## 软件列表

win10下载: https://msdn.itellyou.cn/
win10激活工具: http://www.tudoupe.com/win10/win10jihuo/
window 常用运行库
常用软件下载 腾讯软件中心: https://pc.qq.com/
snipaste(截图工具)
utools: https://u.tools/
everything(查找文件)
total-command: 文件浏览器
foxmail
potplayer
软媒时间
谷歌访问助手: http://www.ggfwzs.com/
工具: bandicam(录屏) camtasia(录屏) carnac(按键屏幕显示) 万彩办公大师
pdf: Acrobat 给 pdf 添加表单、直接修改 pdf 可能导致样式问题
Rufus: Create bootable USB drives the easy way

win-r 运行按照环境变量来找可运行程序, 包括打开目录
sysdm.cpl -> 设置环境变量
netplwiz -> 设置开机不用输入密码
共享: 运行或者文件夹地址栏输入 `\\192.168.0.107`
录音有杂音: 声音图标 -> 右键 -> 录音设备, 然后一步一步设置就行了
翻页: pageup / pagedown 或者 左右方向键
新建隐藏文件: 末尾加一个点
文件被占用无法删除: 资源管理器->cpu->关联句柄->搜索文件名, 就可以看到是什么程序在使用文件了
window 快捷键冲突：尝试了一圈**小软件**，在 Win10 下都不好用，最后只好修改热键；先用 qq 检测热键是否被占用，然后修改
开机自启动：C:\ProgramData\Microsoft\Windows\Start Menu\Programs\StartUp
安装 msi 报 error2303: `admin cmd ->msiexec -i $file`
清理 windows.old: 磁盘 - 属性 - 磁盘清理 - 系统文件
放大镜: win-+ 打开 win-esc 退出 设置view->停靠

win10 常用快捷键: http://www.pconline.com.cn/win10/680/6807260_all.html
potplayer 好用的视频播放器: z正常倍速x减c加 d上一帧f下一帧 添加链接播放(flv/m3u8)
[少数派](https://sspai.com/) [小众软件](https://www.appinn.com/)
[善用佳软](http://xbeta.info/): [大鱼号](https://id.tudou.com/i/UMTMzNjg5MzkyMA==) [TC学堂](https://xbeta.info/studytc/index.htm)
[total commander](http://blog.sina.com.cn/s/blog_631af5fc0102wjpq.html)
[snipaste - 截图工具](https://www.v2ex.com/t/295433)
bandicam: 注册机破解 设置-全屏/保存到桌面/视频格式avi/音频格式pcm f12开始结束 S-f12暂停
camtasia studio: 屏幕录制工具, 设置视频格式avi, 音频立体声 f9录制 f10结束录制
[朗读女](http://www.443w.com/tts): 文本朗读工具
[carnac](http://code52.org/carnac/): 屏幕显示键盘按键
倍速播放: 腾讯视频(2x, 广告); 网易公开课 web [github](https://github.com/theFool32/163OpenCourse_playbackRate) flvcd+m3u8视频流; app离线+dice player(没有中文字幕)
mac 与 pc 键盘布局对比: http://7xksek.com1.z0.glb.clouddn.com/daydaygomac&pc-keymap.jpg
许 ppt 视频教程 [7个网站+4个app](https://post.smzdm.com/p/730480)

## fiddler抓包
> https://imooc.com/learn/37

设置代理:
1. tools - options - connection: 设置 port
2. 设置手机wifi - 手动代码 - fiddler 所在机器 ip:port
3. 电脑先测试一下 ip:port 是否ok

请求重放: 右键 - replay - R(request)/E(edit 记得改 host 部分)

## powershell & cmd
> 真爱生命, 别玩 PS/cmd
- https://channel9.msdn.com/Series/GetStartedPowerShell3
- 《Learn Windows PowerShell 3 in a Month of Lunches》
- 遇到一个很好的讲 cmd 的教程: http://www.crifan.com/summary_usage_of_win7_cmd/

## excel
- 批量输入相同内容: 选中 输入 cmd-enter
- 批量选中
- 减少鼠标/键盘来回切换
- 批量添加序列
- 数据
  - 原数据: 编辑栏
  - format: 使用 format 而不是直接修改原数据
  - 数据验证
  - 导入: ocr/web/csv/db
- 分析
  - 智能表格
  - 数据透视表

# photoshop

- 慕课网 Oeasy: http://www.imooc.com/learn/139
- 慕课网 祁连山: http://www.imooc.com/view/159
- 慕课网 爱米 http://www.imooc.com/learn/506
