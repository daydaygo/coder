# devops| k8s 上 go 微服务实战: go 实现 istio bookinfo 微服务

在完成 `k8s 上快速部署 go 服务` 和 `k8s: istio 入门` 后, 继续 **膨胀**, 使用 go 来实现 istio 提供的 bookinfo 微服务 demo

快速回顾之前的 blog:

- [k8s 上 go 服务实战: 扩容 发版更新 回滚 平滑重启](https://zhuanlan.zhihu.com/p/248866126 "k8s 上 go 服务实战: 扩容 发版更新 回滚 平滑重启")
- [k8s 上 go 服务实战: 使用 helm 快速构建云原生应用](https://zhuanlan.zhihu.com/p/251261134 "k8s 上 go 服务实战: 使用 helm 快速构建云原生应用")
- [k8s: istio 入门实践](https://zhuanlan.zhihu.com/p/258104456 "k8s: istio 入门实践")

涉及到的问题:

- [istio bookinfo demo](https://istio.io/latest/zh/docs/examples/bookinfo/#start-the-application-services "istio bookinfo demo")
- [istio task](https://istio.io/latest/zh/docs/tasks/traffic-management/request-routing/ "istio task")

简单实践步骤:

- 使用 go 重写 bookinfo 微服务并部署到 k8s 中
- 基于 go 版 bookinfo 微服务, 验证 istio task

## 使用 go 重写 bookinfo 微服务并部署到 k8s 中

先回顾一下 bookinfo 微服务应用的端到端架构:

![](https://pic3.zhimg.com/80/v2-688a343d587cff9686ab57accbf42218_720w.jpg)

包含 4 个微服务:

- rating: 书籍评分
- review: 书籍评价, 有 v1/v2/v3 三个版本, 其中 v2/v3 需要调用 `rating` 服务
- detail: 书籍详情
- productpage: 书籍产品页面, 需要调用 `review + detail` 服务

### 实现 rating 服务

可以参考 [k8s 上 go 服务实战: 使用 helm 快速构建云原生应用](https://zhuanlan.zhihu.com/p/251261134 "k8s 上 go 服务实战: 使用 helm 快速构建云原生应用") 快速部署 rating 服务

- go

```go
// rating/main.go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        fmt.Fprintf(writer, "rating")
    })
    http.ListenAndServe(":80", nil)
}
```

- dockerfile

```dockerfile
FROM golang:alpine as builder
WORKDIR /
COPY main.go .
RUN go build -o app main.go

FROM alpine
WORKDIR /
COPY --from=builder /app /app
ENTRYPOINT /app
EXPOSE 80
```

- 使用 alibaba cloudtookit 快速打包上传镜像

![](https://s1.ax1x.com/2020/09/27/0AGkCT.png)

- helm 快速部署

```sh
helm create rating

# 修改 values.yaml
repository: registry.cn-shanghai.aliyuncs.com/daydaygo/istio_bookinfo_rating
# 修改 chart.yml
appVersion: 0.1.1

helm lint --strict rating
helm install rating rating # 部署到 k8s 中
kubectl get pod # 查看 pod
kubectl port-forward $POD_NAME 8081:80 # 开启 port-forward 测试, 本地端口:pod端口
➜ curl localhost:8081
rating⏎
```

同理, 实现 `productpage` `detail` 服务

### 实现 review 服务并包含 3 个版本

- 实现三个版本的 review 应用

| 应用名称  | 镜像版本 |
| --------- | -------- |
| review-v1 | 0.1.0    |
| review-v2 | 0.2.0    |
| Review-v3 | 0.3.0    |

- 直接使用 yaml 文件部署( helm 单应用多版本下次继续折腾)

```yaml
##################################################################################################
# Reviews service
##################################################################################################
apiVersion: v1
kind: Service
metadata:
  name: review
  labels:
    app: review
    service: review
spec:
  ports:
    - port: 80
      name: http
  selector:
    app: review
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: bookinfo-review
  labels:
    account: review
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: review-v1
  labels:
    app: review
    version: v1
spec:
  replicas: 1
  selector:
    matchLabels:
      app: review
      version: v1
  template:
    metadata:
      labels:
        app: review
        version: v1
    spec:
      serviceAccountName: bookinfo-review
      containers:
        - name: review
          image: registry.cn-shanghai.aliyuncs.com/daydaygo/istio_bookinfo_review:0.1.0
          imagePullPolicy: IfNotPresent
          env:
            - name: LOG_DIR
              value: "/tmp/logs"
          ports:
            - containerPort: 80
          volumeMounts:
            - name: tmp
              mountPath: /tmp
            - name: wlp-output
              mountPath: /opt/ibm/wlp/output
      volumes:
        - name: wlp-output
          emptyDir: {}
        - name: tmp
          emptyDir: {}
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: review-v2
  labels:
    app: review
    version: v2
spec:
  replicas: 1
  selector:
    matchLabels:
      app: review
      version: v2
  template:
    metadata:
      labels:
        app: review
        version: v2
    spec:
      serviceAccountName: bookinfo-review
      containers:
        - name: review
          image: registry.cn-shanghai.aliyuncs.com/daydaygo/istio_bookinfo_review:0.2.0
          imagePullPolicy: IfNotPresent
          env:
            - name: LOG_DIR
              value: "/tmp/logs"
          ports:
            - containerPort: 80
          volumeMounts:
            - name: tmp
              mountPath: /tmp
            - name: wlp-output
              mountPath: /opt/ibm/wlp/output
      volumes:
        - name: wlp-output
          emptyDir: {}
        - name: tmp
          emptyDir: {}
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: review-v3
  labels:
    app: review
    version: v3
spec:
  replicas: 1
  selector:
    matchLabels:
      app: review
      version: v3
  template:
    metadata:
      labels:
        app: review
        version: v3
    spec:
      serviceAccountName: bookinfo-review
      containers:
        - name: review
          image: registry.cn-shanghai.aliyuncs.com/daydaygo/istio_bookinfo_review:0.3.0
          imagePullPolicy: IfNotPresent
          env:
            - name: LOG_DIR
              value: "/tmp/logs"
          ports:
            - containerPort: 80
          volumeMounts:
            - name: tmp
              mountPath: /tmp
            - name: wlp-output
              mountPath: /opt/ibm/wlp/output
      volumes:
        - name: wlp-output
          emptyDir: {}
        - name: tmp
          emptyDir: {}
---
```

- 部署到 k8s 中

```sh
kubectl apply -f review.yaml
kubectl get pod
# 重复验证各个 pod
kubectl port-forward $POD_NAME 8081:80 # 开启 port-forward 测试, 本地端口:pod端口
➜ curl localhost:8081
review-v1⏎
```

## 基于 go 版 bookinfo 微服务, 验证 istio task

- istio task 文档: https://github.com/istio/istio.io/content/en/docs/tasks
- task 都可以在 istio sample 找到例子: https://github.com/istio/istio/samples

```sh
# istio.io/content/en/docs/tasks
➜  tasks git:(master) tree -d
.
├── observability
│   ├── distributed-tracing
│   │   ├── configurability
│   │   ├── jaeger
│   │   ├── lightstep
│   │   ├── overview
│   │   └── zipkin
│   ├── gateways
│   ├── kiali
│   ├── logs
│   │   └── access-log
│   └── metrics
│       ├── classify-metrics
│       ├── customize-metrics
│       ├── querying-metrics
│       ├── tcp-metrics
│       └── using-istio-dashboard
├── security
│   ├── authentication
│   │   ├── authn-policy
│   │   └── mtls-migration
│   ├── authorization
│   │   ├── authz-deny
│   │   ├── authz-http
│   │   ├── authz-ingress
│   │   ├── authz-jwt
│   │   ├── authz-tcp
│   │   └── authz-td-migration
│   └── cert-management
│       ├── dns-cert
│       └── plugin-ca-cert
└── traffic-management
    ├── circuit-breaking
    ├── egress
    │   ├── egress-control
    │   ├── egress-gateway
    │   ├── egress-gateway-tls-origination
    │   ├── egress-gateway-tls-origination-sds
    │   ├── egress-kubernetes-services
    │   ├── egress-tls-origination
    │   ├── http-proxy
    │   └── wildcard-egress-hosts
    ├── fault-injection
    ├── ingress
    │   ├── ingress-control
    │   ├── ingress-sni-passthrough
    │   ├── kubernetes-ingress
    │   └── secure-ingress
    ├── mirroring
    ├── request-routing
    ├── request-timeouts
    ├── tcp-traffic-shifting
    └── traffic-shifting

53 directories
```

## 写在最后

istio 几乎涵盖了 服务治理/流量控制 的方方面面, 作为服务治理层的基础设施 **完全够用**, 问题开始从 **行不行**, 转向 **用哪些**, 让 业务层/devops工作流/k8s基础设施 用起来更爽

还需要解决的问题:
- helm 对单应用多版本的支持, 比如 review 应用是通一个 srv, 多个 deploy
- k8s/helm 发布对服务间依赖的支持, 比如 A 服务必须依赖 B 服务更新后才能更新