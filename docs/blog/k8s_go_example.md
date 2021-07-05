# devops| k8s 上 go 服务实战: 扩容 发版更新 回滚 平滑重启

实践为主, 部分 内容/细节 略去, 详情请查看最后的资料

写在前面:

- go 稳坐 **云原生第一编程语言**
- 对概念的理解很重要, 这篇主要涉及 k8s 发布涉及到的 `deployment replicaSet pod` 3个概念
- 动手对学习真的很重要, 不仅要 `BB那么多, show me the code`, 还要 `show and run the code`
- 不要把时间浪费在 **没完没了** 的折腾工具上, 比如 本地安装 k8s, alibaba cloud toolkit 一键部署到 k8s, 用它, 是因为 **用起来爽呀**

## 环境装备

本地开发机使用的 mac, win 平台同理

mac 安装 k8s:

- [docker desktop 安装 - 使用阿里云加速](https://cr.console.aliyun.com/undefined/instances/mirrors "docker desktop 安装")
- [docker desktop 开启 k8s](https://github.com/AliyunContainerService/k8s-for-docker-desktop "docker desktop 开启 k8s")

![mac 安装 k8s]](https://s1.ax1x.com/2020/09/16/wgf4AK.png)

更多 k8s 开发环境准备:

[minikube](https://minikube.sigs.k8s.io/docs/ "minikube")
[kind](https://kind.sigs.k8s.io/ "kind")

继续之前, 请确保自己已经了解 k8s 的基础知识

推荐教程:

- 云原生技术公开课(https://developer.aliyun.com/learning/roadmap/cloudnative "aliyun x cloudnative")

## go 服务代码准备

go http server 入门代码, 开启 3000 端口, 并包含 `/` `/health_check` 2 个请求地址

```go
package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health check</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/health_check", check)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)
}
```

## go 镜像准备

号称 **云原生第一语言** 的实力开始显露出来了: 2 阶段构建, 一个镜像用来便来编译, 一个镜像只包含可执行文件

```dockerfile
FROM golang:alpine AS build-env
WORKDIR /app
ADD . /app
RUN cd /app && go build -o goapp

FROM alpine
WORKDIR /app
COPY --from=build-env /app/goapp /app/
EXPOSE 3000
ENTRYPOINT ./goapp
```

本地进行测试:

```sh
# 镜像构建
docker build -t test . # -t name:tag, 这里简单测试
# 容器运行
docker run -d --rm -p 3000:3000

# 测试
curl http://localhost:3000/
curl http://localhost:3000/health_check
```

## k8s deployment 准备

简单理一理 k8s 中 deployment 的概念:

- deployment: 管理应用的所有发布, 可使用 `kubectl get deploy` 查看
- ReplicaSet: 每次发布的都是应用的 `副本`, 不同版本会生成不同 `副本`, 副本可以进行水平伸缩, 使用 `kubectl get replicasets` 查看
- pod: 应用, 由应用需要的一组 container(容器) 组成, 本文示例比较简单, 只有一个包含简单 `go http server` 的容器

使用 [官方提供的 deployment 示例](https://k8s.io/examples/application/deployment.yaml "官方提供的 deployment 示例") 稍微改一改就行:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: go-app
  template:
    metadata:
      labels:
        app: go-app
    spec:
      containers:
        - name: go-app-container
          image: test # 上一步生成的镜像
          resources:
            limits:
              memory: "128Mi"
              cpu: "500m"
          ports:
            - containerPort: 3000
```

运行:

```sh
kubectl apply -f deployment.yaml
➜ kubectl get pod
NAME                                        READY   STATUS         RESTARTS   AGE
go-app-5757fcdcc5-pcw8l                     0/1     ErrImagePull   0          19s
```

嘛, 报错了, [胜败乃兵家常事](玩过仙剑的请举手 "胜败乃兵家常事"), 错误信息 `ErrImagePull`, 需要将镜像放到 docker hub 上

我们现在缺少的步骤:

- 镜像构建 `docker build`: `name:tag` 要想好
- 上传到 `docker hub`, 或者其他 docker 镜像管理平台, 获取可供使用的镜像地址
- 修改 `deployment.yaml`, 重新 `kubectl apply`

这一步, 敢不敢简单点? 敢!

## 使用 alibaba cloud toolkit 一键部署

[文档在此](https://help.aliyun.com/document_detail/162966.html "alibaba cloud toolkit")

![action > `deploy to kube` 快速打开](https://s1.ax1x.com/2020/09/16/wgjlWR.png)

![镜像构建上传](https://s1.ax1x.com/2020/09/16/wgjWkQ.png)

![部署到本地 k8s](https://s1.ax1x.com/2020/09/16/wgjq7F.png)

```sh
# 检查发布状态
➜ kubectl get deploy
NAME                        READY   UP-TO-DATE   AVAILABLE   AGE
go-app                      1/1     1            1           6h47m
➜ kubectl get replicasets
NAME                                  DESIRED   CURRENT   READY   AGE
go-app-6fd4487dd                      1         1         1       6h32m
➜  kubectl get pod
NAME                                        READY   STATUS             RESTARTS   AGE
go-app-6fd4487dd-d4rgv                      1/1     Running            0          6h31

# 暴露服务快速验证
kubectl expose deployment go-app --type=NodePort --name=go-app-svc --target-port=3000
➜ kubectl get svc
NAME                 TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
go-app-svc           NodePort    10.97.88.209    <none>        3000:32016/TCP   6h29m
```

PS: 仔细看一下 pod 的名字 `go-app-6fd4487dd-d4rgv`, 是不是理解了 `deployment replicaSet pod` 之间的联系?

## let's party with k8s

- 水平扩展/收缩

```sh
➜ kubectl get replicasets
NAME                DESIRED   CURRENT   READY   AGE
go-app-84f9f889c6   1         1         1       6s
➜ kubectl scale --replicas=3 deploy go-app --record
deployment.apps/go-app scaled
➜ kubectl get deploy
NAME     READY   UP-TO-DATE   AVAILABLE   AGE
go-app   3/3     3            3           31s
➜ kubectl get replicasets
NAME                DESIRED   CURRENT   READY   AGE
go-app-84f9f889c6   3         3         3       36s
➜ kubectl get pod
NAME                      READY   STATUS    RESTARTS   AGE
go-app-84f9f889c6-bx5cx   1/1     Running   0          40s
go-app-84f9f889c6-nlc4r   1/1     Running   0          13s
go-app-84f9f889c6-wd7k8   1/1     Running   0          13s
```

伸缩后, pod 新增了 2 个, replicaSet 名字没有变(水平伸缩还是使用当前 `版本`), deployment 名字也没有变

- 滚动更新

简单修改下代码, 使用 alibaba cloud toolkit 更新下

```go
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>test</h1>")
}
```

再次查看: deployment 没有变, replicaSet 新增, pod 更新为新的 3 个

```sh
➜ kubectl get deploy
NAME     READY   UP-TO-DATE   AVAILABLE   AGE
go-app   3/3     3            3           6m39s
➜ kubectl get replicasets
NAME                DESIRED   CURRENT   READY   AGE
go-app-5dd75d6f55   3         3         3       28s
go-app-84f9f889c6   0         0         0       6m46s
➜ kubectl get pod
NAME                      READY   STATUS    RESTARTS   AGE
go-app-5dd75d6f55-gkbp4   1/1     Running   0          31s
go-app-5dd75d6f55-m9w8d   1/1     Running   0          32s
go-app-5dd75d6f55-zl8g2   1/1     Running   0          33s
```

到这里, 也可以猜到, 如果回滚, 就是回到上一个 replicaSet

- 回滚: `kubectl rollout undo  deployment my-go-app --to-revision=1`

- 平滑重启

Deployment 会保证服务的连续性，确保滚动一定有 pod 可用, 从而保证服务可用, 这样就保证了服务能平滑更新

同时, 可以通过 deployment spec 来设置滚动更新策略:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-app
spec:
  replicas: 3
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 1
```

## 写在最后

我是 dayday, 读书写作敲代码, 永远在路上
