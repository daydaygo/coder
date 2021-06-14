# OAM

- <https://oam.dev>
  - <https://kubevela.io>
  - IaaS <https://www.terraform.io/>

- 钉钉群
- OAM 核心依赖库项目：<https://github.com/crossplane/oam-kubernetes-runtime>
- Crossplane 项目：<https://github.com/crossplane/crossplane>
- OAM 系列文章: <https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzUzNzYxNjAzMg==&action=getalbum&album_id=1416643008195608577&scene=173#wechat_redirect>

![](https://github.com/crossplane/oam-kubernetes-runtime/blob/master/assets/arch.png)

- compoents 应用中的组件 dev: 应用本身 应用依赖(如 db) deployment function
- 运维侧
  - traits 运维特征: scalling rollout route cert traffic
  - work load
  - app scope
- appconfig 组装component和trait

OAM 的核心思想是，业务研发人员的工作以从编写源代码开始，到构建完容器镜像结束

kubeVela

应用的亲密性 函数/ipc/rpc
我们可以根据业务的特点，逐步地把一些业务代码拆分成一个个 RPC 或者 IPC 服务，这样它们就可以独立的发布和运维了
IaC（Infrastructure as Code）和 GitOps -- oam
可运行、可观测、可治理、可变更 -- 云效
定义抽象的模板使用了 AWS 自己的 Cloud Formation (KubeVela 目前支持的是 Google 开源的 CUELang 模板语言）
本地开发与测试 tilt kt-connect nocalhost
