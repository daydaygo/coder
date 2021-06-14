# java| 从 spring 中学习笔记

- spring modernJava
  - version(snapshot alpha beta release GA)前世今生 arch架构设计
  - core: IoC/DI`@bean` AOP context
  - module: core jdbc+DAO+ORM+事务 web+MVC
- spring-boot build
  - 约定优于配置conf autoConf env: Annotation properties xml YAML
  - spring-boot-starter
    - web
    - data-jpa=hibernate mybatis Flyway=migration
    - data-redis session-redis mongodb cache
    - amqp=mq
    - mail=attachment+template+staticResource ldap
    - security
    - test: xxxTests `MockMvc`
    - actuator=监控=应用配置(配置刷新)+度量指标+操作控制 log cli
    - devtools=hotload
    - admin
    - swagger=apidoc
    - cli StateMachine(state->event)
    - task: 同步 异步/异步回调 线程池 优雅关闭(平滑重启) Future->Runnable/Callable->任务取消/是否完成/获取结果->阻塞
- spring-cloud=微服务 coordinate 分布式系统基础设施
  - [子模块](https://spring.io/projects/spring-cloud)
  - [版本记录](https://github.com/spring-cloud/spring-cloud-release/releases)
  - [开发工具](https://spring.io/tools)
- dataFlow=大数据 connect
- native=云原生=graalVM+graalVMNativeImage
- eco
  - github: spring-projects spring-native 程序员DDdyc87112
  - [start](https://start.spring.io) [Start Your Dev Journey on Aliyun](https://start.aliyun.com)
  - [探秘 spring AOP](https://www.imooc.com/learn/869) [google guice](http://www.imooc.com/learn/901)

```sh
mvn spring-boot:run
mvn install & java -jar
java -jar target/xxx.jar --spring.profiles.active=prod --server.port=8888 # 命令行设置属性 env
spring.profiles.active=test # 多环境: dev/test/prod
application-dev.properties

com.didispace.blog.desc=${com.didispace.blog.name}正在努力写《${com.didispace.blog.title}》 # 参数引用
com.didispace.blog.value=${random.value} # 随机数

spring.application.name=trace-1
server.port=9101
# 服务注册中心
eureka.client.serviceUrl.defaultZone=http://eureka.didispace.com/eureka/
# 服务消费者
ribbon.eager-load.enabled=true
ribbon.eager-load.clients=hello-service, user-service
# mq 生产
spring.cloud.stream.bindings.input.group=Service-A # 分组
spring.cloud.stream.bindings.input.destination=greetings
spring.cloud.stream.bindings.input.consumer.partitioned=true # 分区
spring.cloud.stream.instanceCount=2
spring.cloud.stream.instanceIndex=0
# mq 消费
spring.cloud.stream.bindings.output.destination=greetings # 分组
spring.cloud.stream.bindings.output.producer.partitionKeyExpression=payload # 分区
spring.cloud.stream.bindings.output.producer.partitionCount=2
# 随机端口
server.port=0 # spring自动随机分配
eureka.instance.instance-id=${spring.application.name}:${random.int}
server.port=${random.int[10000,19999]} # 直接使用random指定

# 监控
/autoconfig // 配置信息
/beans
/configprops
/env
/mappings
/info
/metrics // 度量指标
/health
/dump
/trace
/shutdown // 操作控制
```

```java
// 从注解中获取属性
@Value("${com.didispace.blog.name}")
private String name;

// web
@RestController() // 不需要 @ResponseBody
@Controller()
@ResponseBody()
@RequestMapping()
@PathVariable
@RequestParam
@ModelAttribute

// Swagger2
@ApiOperation(value="更新用户详细信息", notes="根据url的id来指定更新对象，并根据传过来的user信息来更新用户详细信息")
@ApiImplicitParams({
    @ApiImplicitParam(name = "id", value = "用户ID", required = true, dataType = "Long"),
    @ApiImplicitParam(name = "user", value = "用户详细实体user", required = true, dataType = "User")
})

// Spring-data-jpa
@Query("from User u where u.name=:name")
User findUser(@Param("name") String name);

// mybatis-spring-boot-starter
@Select("SELECT * FROM USER WHERE NAME = #{name}")
User findByName(@Param("name") String name);
@Insert("INSERT INTO USER(NAME, AGE) VALUES(#{name}, #{age})")
int insert(@Param("name") String name, @Param("age") Integer age);

// 数据库事务
@Transactional(isolation = Isolation.DEFAULT, propagation = Propagation.REQUIRED)
@Rollback

@EnableCaching // application
@Cacheable // repository
@CachePut // 缓存更新

// task
@EnableScheduling
@Scheduled(fixedRate = 5000) // 5s
@Scheduled(fixedDelay = 5000)
@Scheduled(initialDelay=1000, fixedRate=5000)
@Scheduled(cron="*/5 * * * * *")
@EnableAsync
@Async

// 等待任务全部执行完
while(true) {
    if(task1.isDone() && task2.isDone() && task3.isDone()) {
        // 三个任务都调用完成，退出循环等待
        break;
    }
    Thread.sleep(1000);
}

// task 线程池
@Async("taskExecutor")
@Bean("taskExecutor")
public Executor taskExecutor() {
    ThreadPoolTaskExecutor executor = new ThreadPoolTaskExecutor();
    executor.setCorePoolSize(10);
    executor.setMaxPoolSize(20);
    executor.setQueueCapacity(200);
    executor.setKeepAliveSeconds(60);
    executor.setThreadNamePrefix("taskExecutor-");
    executor.setRejectedExecutionHandler(new ThreadPoolExecutor.CallerRunsPolicy());
    // 优雅关闭
    executor.setWaitForTasksToCompleteOnShutdown(true);
    executor.setAwaitTerminationSeconds(60);
    return executor;
}

@Test
@Component // 实现命名替换

// mq
@StreamListener(Sink.INPUT)
@Input(Sink.INPUT)
SubscribableChannel input();
```
