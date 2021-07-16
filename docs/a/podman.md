# podman

- https://multipass.run/docs
- http://docs.podman.io/
  - https://podman.io/getting-started/installation

- dockerless https://mkdev.me/en/posts/dockerless-part-3-moving-development-environment-to-containers-with-podman

相关概念:

- OCI: image(container) -> buddle(fs+config) -> runtime(runc)
- OCI(runc) -> CRI / CNI
- buildah 镜像构建
- skopeo 镜像管理
  - harbor权威指南
- podman 容器管理

```sh
# Ubuntu20.10 虽然包含 podman, 但是版本不是最新
brew cask install multipass
multipass find # image -> Ubuntu20.10 = groovy, 默认 Ubuntu20.04
multipass launch -c 2 -d 10G -m 2G -n podman

# 安装最新版本
. /etc/os-release
echo "deb https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/xUbuntu_${VERSION_ID}/ /" | sudo tee /etc/apt/sources.list.d/devel:kubic:libcontainers:stable.list
curl -L https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/xUbuntu_${VERSION_ID}/Release.key | sudo apt-key add -
sudo apt-get update
sudo apt-get -y upgrade
sudo apt-get -y install podman
podman info # 确认安装形成
# enabl podman service: 依赖 systemd 的 socket activation 特性
# sudo systemctl cat podman.socket
# sudo systemctl cat podman.service
sudo systemctl enable podman.socket --now
# 加速
vim /etc/containers/registries.conf
[registries.search]
registries = ['c3ywro5t.mirror.aliyuncs.com','docker.io']

# podman client
brew install podman
podman system connection add ubuntu --identity ~/.ssh/id_rsa ssh://root@192.168.64.2/run/podman/podman.sock
podman system connection list

# podman 使用
alias docker=podman
podman run -d -p 6379:6379 redis:alpine
podman run -d -p 3306:3306 -e TZ='Asia/Shanghai' -e MYSQL_ROOT_PASSWORD=root mysql
podman run -d -p 2379:2379 -e ETCD_ADVERTISE_CLIENT_URLS='http://0.0.0.0:2379' -e ETCD_LISTEN_CLIENT_URLS='http://0.0.0.0:2379' -e ETCDCTL_API='3' quay.io/coreos/etcd
podman run -d -p 5672:5672 -p 15672:15672 -e RABBITMQ_DEFAULT_USER='dayday' -e RABBITMQ_DEFAULT_PASS='dayday' --hostname aliyun rabbitmq:management-alpine
podman run -d --name=dev-consul -e CONSUL_BIND_INTERFACE=eth0 consul

podman generate kube xxx # 容器 -> kube pod yaml
```

## buildah

> https://buildah.io/

- define image
  - Dockerfile
  - buildah: from run config mount commit rm

```sh
# install
. /etc/os-release
sudo sh -c "echo 'deb http://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/x${ID^}_${VERSION_ID}/ /' > /etc/apt/sources.list.d/devel:kubic:libcontainers:stable.list"
wget -nv https://download.opensuse.org/repositories/devel:kubic:libcontainers:stable/x${ID^}_${VERSION_ID}/Release.key -O Release.key
sudo apt-key add - < Release.key
sudo apt-get update -qq
sudo apt-get -qq -y install buildah

# use
buildah bud -t docker.io/mkdevme/mattermost:5.8.0 .
buildah login --username=1252409767@qq.com registry.cn-shanghai.aliyuncs.com
```

- 镜像大小: 轻量化基础镜像 多阶段构建

```dockerfile
# old: 1.16G
FROM node:10
WORKDIR /app
COPY app /app
RUN npm install -g webserver.local
RUN npm install && npm run build
EXPOSE 3000
CMD webserver.local -d ./build

# new: 22.4M
FROM node:10-alpine AS build
WORKDIR /app
COPY app /app
RUN npm install && npm run build
FROM nginx:stable-alpine
COPY --from=build /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

# alpine
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai
```

## minikube

> https://minikube.sigs.k8s.io/

```sh
# install
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube_latest_amd64.deb
sudo dpkg -i minikube_latest_amd64.deb

minikube start --driver=podman --container-runtime=cri-o
minikube config set driver podman
```
