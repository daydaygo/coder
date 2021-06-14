# coder| 重回 win 开发环境之基于 wsl2 的 PHP 开发环境

> 不推荐使用, phpstorm 加载文件非常慢

内容简介:

- 开启 wsl2
- 使用 wsl Ubuntu
- 使用 window terminal
- 使用 docker desktop 并开启 wsl2 支持
- 使用 vscode 并开启 wsl 支持
- wsl Ubuntu 配置 PHP 环境
- phpstorm 开启 wsl 支持
- mark: 总结

相关资源:

- win wsl2: <https://docs.microsoft.com/zh-cn/windows/wsl/wsl2-index>
- docker wsl2: <https://docs.docker.com/docker-for-windows/wsl/>
- win terminal: <https://docs.microsoft.com/en-us/windows/terminal/>

## 开启 wsl2

```powershell
# 开启 wsl
dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart

# 更新到 wsl2
dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart
```

**重启, 完成更新** (我在这里犯了一个 **低级错误**, 看文档不仔细, 没有重启, 浪费了很多时间)

## 使用 wsl Ubuntu

到 `Microsoft store` 搜索 wsl 即可, 推荐使用 Ubuntu, 使用人数最多, 遇到 bug 方面搜索文档来解决

进入到 wsl Ubuntu 中, 第一次需要设置用户名密码, 设置完就完成了 wsl Ubuntu 的安装

- wsl 常用命令

```powershell
wsl # 进入默认的 wsl, 或者输入 wsl 的名字进入

wsl -l -v # 查看
ws --set-default <name> # 设置默认
\\wsl$\<name>\home\<user> # 默认目录, 使用此路径可以访问到 wsl 中的文件
```

## 使用 window terminal

安装十分简单, `Microsoft store` 中搜索安装即可, 安装后根据文档修改了 主题(theme) 和 快捷键(key-binding), 参考如下

```json
// 只保留我修改的部分
{
    // 设置 wsl Ubuntu 为默认
    "defaultProfile": "{07b52e3e-de2c-5db4-bd2d-ba144ed6c273}",
    "profiles":
    {
        "defaults":
        {
            // 设置主题, 并对所有 profile 生效
            // Put settings here that you want to apply to all profiles.
            "acrylicOpacity" : 0.7,
            "colorScheme" : "Campbell",
            "cursorColor" : "#FFFFFD",
            "fontFace" : "Cascadia Code PL",
            "useAcrylic" : true
        },
        "list":
        [
            // 这里就是新增的 wsl Ubuntu
            {
                "guid": "{07b52e3e-de2c-5db4-bd2d-ba144ed6c273}",
                "hidden": false,
                "name": "Ubuntu-20.04",
                "source": "Windows.Terminal.Wsl"
            },
        ]
    },
    "keybindings":
    [
        // 添加了我常用的快捷键
        { "command": "closeTab", "keys": "ctrl+shift+w" },
        { "command": { "action": "switchToTab", "index": 0 }, "keys": "alt+1" },
        { "command": { "action": "switchToTab", "index": 1 }, "keys": "alt+2" },
        { "command": { "action": "switchToTab", "index": 2 }, "keys": "alt+3" },
        { "command": { "action": "switchToTab", "index": 3 }, "keys": "alt+4" },
        { "command": { "action": "switchToTab", "index": 4 }, "keys": "alt+5" },
        { "command": { "action": "switchToTab", "index": 5 }, "keys": "alt+6" },
        { "command": { "action": "switchToTab", "index": 6 }, "keys": "alt+7" },
        { "command": { "action": "switchToTab", "index": 7 }, "keys": "alt+8" },
        { "command": { "action": "switchToTab", "index": 8 }, "keys": "alt+9" }
    ]
}
```

window terminal 的其他优点就不提了, 用过 iterm2 之类 terminal 都不会觉得有啥, 大概是 cmd / powershell 实在是太令人失望了吧...

## 使用 docker desktop 并开启 wsl2 支持

参看此文档 `https://docs.docker.com/docker-for-windows/wsl/` 即可, 下载好 `docker desktop` 然后再配置中开始 wsl 支持即可

## 使用 vscode 并开启 wsl 支持

这个也简单, vscode 中搜索 `wsl` 插件安装即可, 之后就可以在 wsl 中使用 code 打开文件

注意: root 下没有 code 命令, 可以使用 `echo $PATH` 查看到差别, wsl 初始化时设置的用户会带上 win 下的几个 path, 所有可以使用 `explorer.exe .` 来打开目录, 而 root 下的 `PATH` 是没有这些可执行文件的

## wsl Ubuntu 配置 PHP 环境

- 修改 Ubuntu 镜像源:

> 清华大学镜像源 mirror: <https://mirror.tuna.tsinghua.edu.cn/help/ubuntu/>

- 安装 PHP 7.2 版本

> PS: PHP 要锁定到 7.2 / 7.2 这样的版本号, 否则可能导致 7.3 下依赖的包, 在 7.2 无法安装, 开发团队内部/各环境之间, 保持 PHP 版本统一

```bash
# 添加支持安装不同 php 版本的源
sudo apt-get install software-properties-common python-software-properties
sudo add-apt-repository ppa:ondrej/php && sudo apt-get update

# 包含 pecl 等工具, 方便安装扩展
sudo apt-get -y install php7.2-dev php7.2-xml php7.2-bcmath php7.2-curl php7.2-mbstring

# 验证
php -v
```

- 安装依赖的扩展 ext: redis mongodb rdkafka swoole

```bash
# 1.1 直接使用 pecl 安装
sudo pecl install redis

# 1.2 下载好文件后安装
wget http://pecl.php.net/get/redis-5.3.0.tgz
sudo pecl install redis

# 2. 添加到 ini 文件中
php --ini # 查看 ini 文件位置
sudo chmod 777 /etc/php/7.2/cli/php.ini # 添加编辑权限, 方便 vscode 打开
code /etc/php/7.2/cli/php.ini
# 添加以下内容
extension=redis.so

# mongodb 直接安装

# rdkafka
sudo apt install -y librdkafka-dev
sudo pecl install rdkafka

# swoole 使用编译安装, 方便修改编译参数
wget http://pecl.php.net/get/swoole-4.5.2.tgz
tar zxvf swoole-4.5.2.tgz
cd swoole-4.5.2
sudo phpize
./configure
make
sudo make install

# 添加 sdebug 支持
git clone https://github.com/swoole/sdebug.git
cd sdebug
sudo bash rebuild.sh
```

- php.ini 完整配置

```ini
extension=redis.so
extension=mongodb.so
extension=rdkafka.so
extension=swoole.so
swoole.use_shortname=0 # hyperf 需要使用

[xdebug]
zend_extension=xdebug.so
xdebug.remote_enable = On
xdebug.remote_handler = dbgp
xdebug.remote_host= localhost
xdebug.remote_port = 9000 # 默认使用 9000 端口, 可以更新其他端口
xdebug.idekey = PHPSTORM
```

- 安装 composer

```bash
mkdir tmp && cd tmp
# 安装稳定版
wget https://getcomposer.org/composer-stable.phar

# 安装最新版
wget https://getcomposer.org/composer.phar

# 添加到环境变量中
# 可以使用 `echo $PATH` 先检查
sudo cp composer.phar /usr/local/bin/composer
sudo chmod +x /usr/local/bin/composer

# 使用镜像源
composer config -g repo.packagist composer https://mirrors.aliyun.com/composer/
```

## phpstorm 开启 wsl 支持

最新版的 phpstorm 已经开始了 wsl 支持, 安装好 插件(plugin) 后, 就可以开始使用了

打开项目, 使用上面提到的路径:

`\\wsl$\<name>\home\<user>\<project path>`

配置 wsl:

`设置(setting) -> 语言与框架(language & framework) -> PHP -> cli interpreter -> 添加 -> 选择 wsl`

## mark 写在最后

目前还在初步使用阶段中, 能把 hyperf-skeleton 跑起来, 后续使用过程中遇到更多场景, 会陆续更新到新的 blog 里
