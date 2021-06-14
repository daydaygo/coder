# podman 快速入门

podman 上手简单攻略:

- 安装 podman
  - macos 可以通过 `multipass` 快速安装 `Ubuntu` 虚拟机(底层使用 `hyperkit`)
  - 云服务直接安装 Ubuntu, 20.10 直接包含 podman, 其他版本按照官方文档安装; **不要用 centos**, 会出现 podman 版本不一致的问题
- macos 安装 podman client & 配置好 podman connection
- 快速上手骚操作: `alias docker=podman`

## 概念
- OCI -> CRI / CNI
- buildah 镜像构建
- skopeo 镜像管理
- podman 容器管理

## macos 使用 podman

- podman 需要在 linux 下运行
- hyperkit: 轻量级虚拟机, 用来创建 linux

```sh
# 安装 hyperkit
# 方式一
brew cask install multipass
# 方式二: 新版的 docker desktop 自带

# 创建虚拟机
multipass launch -c 2 -d 10G -m 2G -n podman # -n name; -c CPU; -m mem; -d disk

# 查看
multipass list

# 进入
multipass shell podman

# 安装 podman
# ubuntu
. /etc/os-release
echo "deb https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/xUbuntu_${VERSION_ID}/ /" | sudo tee /etc/apt/sources.list.d/devel:kubic:libcontainers:stable.list
curl -L https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/xUbuntu_${VERSION_ID}/Release.key | sudo apt-key add -
sudo apt-get update
sudo apt-get -y upgrade 
sudo apt-get -y install podman

# config first connection
# enabl podman service: 依赖 systemd 的 socket activation 特性
sudo systemctl cat podman.socket
sudo systemctl cat podman.service
sudo systemctl enable podman.socket --now
# 确认 podman.socket 是否开启成功
podman info
# 加速
vim /etc/containers/registries.conf
[registries.search]
registries = ['c3ywro5t.mirror.aliyuncs.com','docker.io']

# podman client
brew install podman
podman system connection add ubuntu --identity ~/.ssh/id_rsa ssh://root@192.168.64.2/run/podman/podman.sock
podman system connection list

# 骚操作
alias podman=docker
```

## 写在最后

- 苦 docker desktop 久矣, MBP 出门开 docker desktop 就没超过 2h 过...