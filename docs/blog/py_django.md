# py| 初探 django

- [python| 初探 django](https://www.jianshu.com/p/f200c1acd1f3)

遇到使用 django 的项目, 用自己的方式熟悉起来

**更新**: 一直局限于 pycharm 要本地有 python 环境, 才可以智能识别代码, 现在才发现可以使用 **project interpreter** 进行设置, `local/ssh/docker` 都可以, 又可以折腾起来了~

## 教程

- [django初体验](https://www.imooc.com/view/458)
- [django入门与实践](https://www.imooc.com/view/790)

## 开发环境配置

不多说, 直接上 dockerfile:

```dockerfile
# FROM ubuntu:18.04
FROM rastasheep/ubuntu-sshd:18.04
LABEL maintainer="1252409767@qq.com"

RUN sed -i 's/archive.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list && \
    apt update
RUN apt-get update && apt-get install -y locales && \
    # rm -rf /var/lib/apt/lists/* && \
    localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8
ENV LANG en_US.utf8

RUN apt install -y fish vim less curl telnet net-tools ipython3 python3-pip
RUN echo -e "[global]\nindex-url=https://mirrors.aliyun.com/pypi/simple/\nformat=columns" > /etc/pip.conf && \
    pip3 install --upgrade pip && \
    pip3 install django

WORKDIR /var/www/
```

还有 docker-compose:

```yaml
version: '3'
services:
    u:
        build:
            context: linux
            dockerfile: ubuntu.Dockerfile # https://hub.docker.com/_/ubuntu/
        volumes:
            - ../:/var/www
        ports: # https://github.com/rastasheep/ubuntu-sshd
            - "8038:80"
            - "1138:22"
        tty: true
```

简单说明几个有意思的点:

- 为啥使用 docker? 不提那些大道理, 可以足够 **任性** -- 想建就建, 想删就删
- 为啥使用 Ubuntu? 教程中使用的 Ubuntu, 避免可能 **冒出的问题**, 保持 OS 一致可以有效降低风险
- 为啥都使用的最新版? 无论是 Ubuntu 还是 python, 都选择了最新版本, **享受技术迭代带来的乐趣**
- 为啥不是官方的 Ubuntu 镜像? 因为我希望使用 ssh, 还是喜欢 xshell 下的 terminal. 另外想增加 ssh 非常简单, 先谢开源 [rastasheep/ubuntu-sshd](https://github.com/rastasheep/ubuntu-sshd)

还有一些细节 -- **Ubuntu使用国内源** **pip使用国内源** **pip设置** 等, 都是一点一点积累的.

> 正是这些积累, 尝试新事物 django 的时候, 不会有一个陡峭的学习曲线, 可以快速 get 一个新技能.

## django: 运行第一个应用

安装好 django 后命令行就可以 `django-admin`, 可以通过模板来创建项目

```
# 通过模板来创建项目
django-admin startproject startproject

# 查看
root@9e16d0821126 /v/w/c/p/d/startproject# tree
.
├── db.sqlite3
├── manage.py
└── startproject
    ├── __init__.py
    ├── __pycache__
    │   ├── __init__.cpython-36.pyc
    │   ├── settings.cpython-36.pyc
    │   ├── urls.cpython-36.pyc
    │   └── wsgi.cpython-36.pyc
    ├── settings.py
    ├── urls.py
    └── wsgi.py

2 directories, 10 files

# 启动进行测试
cd startproject
python3 manage.py runserver
```

启动后, 看到如下输入:

```
Run 'python manage.py migrate' to apply them.

August 29, 2018 - 14:13:27
Django version 2.1, using settings 'startproject.settings'
Starting development server at http://127.0.0.1:8000/
Quit the server with CONTROL-C.
```

**为啥把这段贴出来呢?** 是要传递一个很简单的观点:

> 英文很简单也很重要, 经常给我们很多有用信息

- 看了这段话才知道要运行 `python manage.py migrate`
- 服务器运行在 `127.0.0.1:8000` 上

我们来测试一下:

```
telnet 127.0.0.1 8000 # 测试端口是否通
netstat -alnp|grep 8000 # 查看网络服务的信息
curl 127.0.0.1:8000 # 测试 http, 也可以用 wget
```

## django: 深入一点示例项目

上一节已经使用 `tree` 来查看 **项目目录结构**, 这节我们来深入几个地方

- `python3 manage.py`: 项目管理入口, 是不是和 `php yii` 或者 `php artisan` 类似? 是的, **很多地方, 语言是想通的, 框架设计也是**
- `settings.py`: 应用配置文件, 基本通过 **变量名 + 注释** 就可以理解
- `urls.py`: 相当于其他框架中的 **路由**
- `wsgi.py`: Web Server Gateway Interface, web服务的通用抽象, 类比一下 php 中的 `fcgi`

```bash
# 项目管理入口
python3 manage.py

# 换ip:port, 这样可以在容器外访问
python3 manage.py runserver 0.0.0.0:80

# 调试, 可以配合之前的 ipython 使用
python3 manage.py shell
```

## django: project & app

简单看完 django 的 project 概念后, 我们来看到 project 下的一个新概念 -- `app`

```
# 新建一个名叫 blog 的 app
python3 manage.py startapp blog

# 目录结构
root@9e16d0821126 /v/w/c/p/d/startproject# tree blog/
blog/
├── __init__.py
├── __pycache__
│   ├── __init__.cpython-36.pyc
│   ├── admin.cpython-36.pyc
│   ├── models.cpython-36.pyc
│   └── views.cpython-36.pyc
├── admin.py
├── apps.py
├── migrations
│   ├── __init__.py
│   └── __pycache__
│       └── __init__.cpython-36.pyc
├── models.py
├── tests.py
└── views.py

3 directories, 12 files
```

着重来看几个:

- `views.py`: 视图文件, 熟悉 MVC 的应该都懂, 当然, 有视图也少不了 `template(模板)`
- `models.py`: MVC 中的 model, 我们要开始和数据库打交道了
- `migration`: 数据迁移, yii/laravel 这样的框架也都有
- `admin.py`: 相关管理配置
- `test.py`: 测试对代码质量有多重要, 就不多说了

还有一些常见的:

- project 中的 `admin` 模块, 进行一些常规的管理操作
- project 中的 `urls` 支持使用 `include` 加载 app 中的 `urls`
- project 中一些常用的配置: `i18n l10n` 等
- app 中 `views` 使用 `render` 来加载 `template`

```
# model + migration
python manage.py makemigrations blog # 通过 model 生成 migration 文件
python manage.py sqlmigrate blog 0001 # 查看 migration 对应的 sql 语句
python manage.py migrate blog # 执行 migrate

# admin模块
python manage.py createsuperuser
```

## 写在最后

为什么我们要使用框架呢? 因为框架了封装了很多 **通用的功能**, 有些是我们知道的, 有些是 **我们不知道的**, 了解框架, 既是学习技术上的 **架构设计**, 也是补充相应技术的 **领域知识**, 比如上面没有提到的 **安全**, 框架往往比我们做得更多.