# istio

## ms

- app
  - 容错: 无状态 幂等 超时 重试 缓存 实时监控度量
  - auth
  - 配置中心
- net -> istio
  - 服务注册发现 限流 降级 熔断 LB
  - log/err tracing metric monitor

## install

```sh
brew install istioctl
istioctl manifest apply # 安装 istio 到 k8s
kubectl get svc -n istio-system # 校验安装是否成功
kubectl get pods -n istio-system

kubectl label namespace default istio-injection=enabled # default 命名空间下开启 istio
istioctl kube-inject -f register-service.yaml | kubectl apply -f - # pod 注入

# kiali dashboard https://istio.io/latest/docs/ops/integrations/kiali/#installation
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.8/samples/addons/kiali.yaml
istioctl dashboard kiali
```

## istio samples

- <https://github.com/istio/istio>

```sh
# istio/samples
➜  samples git:(master) tree -L 2
.
├── README.md
├── addons
│   ├── README.md
│   ├── extras
│   ├── grafana.yaml
│   ├── jaeger.yaml
│   ├── kiali.yaml
│   └── prometheus.yaml
├── bookinfo
│   ├── README.md
│   ├── build_push_update_images.sh
│   ├── networking
│   ├── platform
│   ├── policy
│   ├── src
│   └── swagger.yaml
├── certs
│   ├── README.md
│   ├── ca-cert.pem
│   ├── ca-key.pem
│   ├── cert-chain.pem
│   ├── generate-workload.sh
│   ├── root-cert.pem
│   ├── workload-bar-cert.pem
│   ├── workload-bar-key.pem
│   ├── workload-foo-cert.pem
│   └── workload-foo-key.pem
├── custom-bootstrap
│   ├── README.md
│   ├── custom-bootstrap.yaml
│   └── example-app.yaml
├── external
│   ├── README.md
│   ├── aptget.yaml
│   ├── github.yaml
│   └── pypi.yaml
├── health-check
│   ├── liveness-command.yaml
│   ├── liveness-http-same-port.yaml
│   └── server.go
├── helloworld
│   ├── README.md
│   ├── helloworld-gateway.yaml
│   ├── helloworld.yaml
│   ├── loadgen.sh
│   └── src
├── httpbin
│   ├── README.md
│   ├── httpbin-gateway.yaml
│   ├── httpbin-nodeport.yaml
│   ├── httpbin-vault.yaml
│   ├── httpbin.yaml
│   └── sample-client
├── https
│   ├── default.conf
│   └── nginx-app.yaml
├── kubernetes-blog
│   ├── bookinfo-ratings.yaml
│   ├── bookinfo-reviews-v2.yaml
│   └── bookinfo-v1.yaml
├── multicluster
│   ├── README.md
│   ├── expose-istiod.yaml
│   ├── expose-services.yaml
│   └── gen-eastwest-gateway.sh
├── rawvm
│   ├── Makefile
│   ├── README.md
│   ├── demo.sh
│   ├── k8cli.yaml.in
│   └── k8services.yaml.in
├── security
│   └── psp
├── sleep
│   ├── README.md
│   ├── sleep-vault.yaml
│   └── sleep.yaml
├── tcp-echo
│   ├── README.md
│   ├── src
│   ├── tcp-echo-20-v2.yaml
│   ├── tcp-echo-all-v1.yaml
│   ├── tcp-echo-services.yaml
│   └── tcp-echo.yaml
└── websockets
    ├── README.md
    ├── app.yaml
    └── route.yaml

25 directories, 63 files
```

## istio task

- <https://github.com/istio/istio.io>

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
