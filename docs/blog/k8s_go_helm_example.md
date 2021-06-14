# devops| k8s 上 go 服务实战: 使用 helm 快速构建云原生应用
> [k8s 上 go 服务实战: 使用 helm 快速构建云原生应用](https://zhuanlan.zhihu.com/p/251261134 "k8s 上 go 服务实战: 使用 helm 快速构建云原生应用")

上一篇折腾了 [k8s 如何跑 go 服务](https://zhuanlan.zhihu.com/p/248866126 k8s运行go服务), 并且对 k8s 基础概念 deployment / replicaSet / pod 进行了详细的讲解

实践过的小伙伴可能会继续追问: 好多好多步骤呢, 尤其是写 `deployment.yml` 文件时, 还挺折腾的, 敢不敢更简单点?

一如既往: 敢!

> [helm: The package manager for Kubernetes](https://helm.sh/ "helm")

## 服务镜像准备

首先, 还是老三样, 准备好我们的镜像

- go 服务代码, 还是简单的 hello 为例

这里使用了 **环境变量**, 方便后续演示: 可以通过环境变量这种方式来控制应用

```go
package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "80"
    }
    username := os.Getenv("USERNAME")
    if username == "" {
        username = "world"
    }
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        fmt.Fprintf(writer, "hello %s\n\n", username)
    })
    http.ListenAndServe(":" + port, nil)
}
```

- Dockerfile 构建镜像, 依旧是 **两阶段构建**

```Dockerfile
FROM golang:alpine as builder
WORKDIR /app
COPY main.go .
RUN go build -o hello main.go

FROM alpine
WORKDIR /app
ARG PORT=80
COPY --from=builder /app/hello /app/hello
ENTRYPOINT ./hello
EXPOSE 80
```

- alibaba cloud toolkit 上传镜像, 依旧是一件搞定

![alibaba cloud toolkit 上传镜像](https://s1.ax1x.com/2020/09/17/wfd559.png)

## helm: 主角登场

介绍一下概念, 方便理解:
- helm: k8s 的包管理工具, 能把 k8s 的应用/服务 打包好
- chart: helm 中对 k8s 中 应用/服务 的抽象, 简单理解 `chart = k8s 应用`

不废话, 直接一顿操作猛如虎:

```sh
brew info helm # 安装 helm

helm init hello # 初始化一个 helm chart 工程
➜ tree -L 2
.
├── Chart.yaml # chart 的基础新
├── charts
├── templates # k8s 的标准 yaml 文件
│   ├── NOTES.txt
│   ├── _helpers.tpl
│   ├── deployment.yaml
│   ├── hpa.yaml
│   ├── ingress.yaml
│   ├── service.yaml
│   ├── serviceaccount.yaml
│   └── tests
└── values.yaml # 主要需要修改的文件, 值都在这里定期, 供 templates 下的文件使用
```

只需要修改 2 个文件:

- `deployment.yaml`: 把环境变量加上

```yaml
      containers:
        - name: {{ .Chart.Name }}
          securityContext:
            {{- toYaml .Values.securityContext | nindent 12 }}
          image: "{{ .Values.image.repository }}:{{ .Values.image.tag | default .Chart.AppVersion }}"
          imagePullPolicy: {{ .Values.image.pullPolicy }}
          env: # 添加环境变量
            - name: USERNAME
              value: {{ .Values.Username }}
```

- `values.yaml`: 修改镜像地址 + 配置环境变量

```yaml
image:
  repository: registry.cn-shanghai.aliyuncs.com/daydaygo/open
  pullPolicy: IfNotPresent
  # Overrides the image tag whose default is the chart appVersion.
  tag: "20200917224128"

Username: dayday
```

继续 helm 操作:

```sh
helm lint --strict hello # lint: 校验

➜ helm package hello # package: 打包
Successfully packaged chart and saved it to: /Users/dayday/coder_at_work/docker/k8s/helm_test/hello-0.1.0.tgz

➜ helm install hello hello-0.1.0.tgz # install: 安装
NAME: hello
LAST DEPLOYED: Thu Sep 17 22:54:54 2020
NAMESPACE: default
STATUS: deployed
REVISION: 1
NOTES:
1. Get the application URL by running these commands:
  export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=hello,app.kubernetes.io/instance=hello" -o jsonpath="{.items[0].metadata.name}")
  echo "Visit http://127.0.0.1:8080 to use your application"
  kubectl --namespace default port-forward $POD_NAME 8080:80

➜ kubectl port-forward hello-ffbd5b4d7-bhwtp 3000:80 # port-forward: 端口转发, 方便本地测试
Forwarding from 127.0.0.1:3000 -> 80
Forwarding from [::1]:3000 -> 80

➜ curl localhost:3000 # 大功告成
hello dayday
```

## 写在最后

很简单有木有, 赶紧也动手试试

更多例子: [cloudnativeapp/handbook](https://github.com/cloudnativeapp/handbook "Cloud Native App Handbook")

伴随着云原生生态的不断发展与壮大, 标准/工具链越发的成熟, 服务的部署与运维会越来越 easy, 一起期待更美好的明天~

我是 dayday, 读书写作敲代码, 永远在路上