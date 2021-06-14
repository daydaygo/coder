# go| gokit

## code

```sh
➜  gokit git:(master) tree -L 1 --dirsfirst
.
├── auth
├── circuitbreaker
├── cmd/kitgen
├── endpoint
├── examples
├── log
├── metrics
├── ratelimit
├── sd
├── tracing
├── transport
├── util # coon
├── README.md
├── codecov.yml
├── docker-compose-integration.yml # etcd consul zk eureka
├── go.mod
├── lint # github.com/alecthomas/gometalinter
└── update_deps.bash
```

## note

- transport: http/grpc amqp awslambda httprp nats netrpc thrift
- endpoint
  - sd(consul dnssrv etcd eureka internal lb zk)
  - ratelimit(tokenBucket) auth(basic jwt casbin) circuitbreaker(handy hystrix)
  - tracing(ipencensus opentracing zipkin)
  - metrics(cloudwatch discard dogstatsd expvar generic grahite influx...)
  - log: level syslog/logrus/zap
- service-业务
- example
  - addsvc: cmd(addcli addsvc) pb thrift pkg(transport/endpoint/service)
  - apigateway
  - profilesvc - rest
  - shipping - ms
  - stringsvc1-4
