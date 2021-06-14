# hyperf| hyperf/metric 上手指南

这期又开始聊微服务的基础设施之一: 实时监控. 更准确的说法, 基于 prometheus 的实时监控. 关于技术选型这里就不多啰嗦啦, 很多时候「从众」或者「用脚投票」往往是最有效的

> 真的猛士, 敢于走少有人走的路. 我们选择 prometheus, 这条猛士已经走过的路.

## prometheus 实战

之所以先把实战放上来, 在于「看得见」往往不如「摸得着」来得 ~~刺激~~ 深刻. 有了一套环境在那放着, 变学变实践, 才不会停留在 **花了时间学, 用的时候还是不会** 这种低效的循环里. 至于环境问题怎么解决, 又到了我们的老朋友 docker 了.

```yml
version: '3.1'
services:
    hyperf:
        image: hyperf/hyperf
        volumes:
            - ./:/data
        ports:
            - "9501:9501"
        tty: true
    # https://prometheus.io/docs/prometheus/latest/installation/
    prometheus:
        image: prom/prometheus
        ports:
            - "9090:9090"
        volumes:
            - ./config/prometheus.yml:/etc/prometheus/prometheus.yml
    # https://grafana.com/docs/installation/docker/
    grafana:
        image: grafana/grafana
        ports:
            - "3000:3000"
```

docker 知识不熟悉的小伙伴, 可以先学着把自己本地使用的环境 docker 化.

> 把 docker 当做一个工具, 就会发现这家伙是的真的简单. 原理可能很复杂, 但是我只是使用呀, 这和 CRUD 能有多大差别?

### hyperf/metric 配置

首先是 hyperf 里面使用 prometheus, 这里参考官网文档 [hyperf/metric](https://doc.hyperf.io/#/zh/metric) 即可. 当然官网的 metric 包比较强大(这不是吹水, 官方通常要考虑更 **通用**, 监控选项的支持会更多一些), 这里我们只关注 prometheus:

- 安装

```
# 安装
composer require hyperf/metric

# 添加配置
php bin/hyperf vendor:publish hyperf/metric
```

- 配置

```php
return [
    'default' => env('METRIC_DRIVER', 'prometheus'),
    'use_standalone_process' => env('METRIC_USE_STANDALONE_PROCESS', true),
    'enable_default_metric' => env('METRIC_ENABLE_DEFAULT_METRIC', true),
    'default_metric_interval' => env('DEFAULT_METRIC_INTERVAL', 5),
    'metric' => [
        'prometheus' => [
            'driver' => Hyperf\Metric\Adapter\Prometheus\MetricFactory::class,
            'mode' => Constants::SCRAPE_MODE,
            'namespace' => env('APP_NAME', 'skeleton'),
            'scrape_host' => env('PROMETHEUS_SCRAPE_HOST', '0.0.0.0'),
            'scrape_port' => env('PROMETHEUS_SCRAPE_PORT', '9502'),
            'scrape_path' => env('PROMETHEUS_SCRAPE_PATH', '/metrics'),
            'push_host' => env('PROMETHEUS_PUSH_HOST', '0.0.0.0'),
            'push_port' => env('PROMETHEUS_PUSH_PORT', '9091'),
            'push_interval' => env('PROMETHEUS_PUSH_INTERVAL', 5),
        ],
        ...
```

使用默认配置就好, 默认配置就是 prometheus, 之后访问 `http://localhost:9502/metrics` 就可以查看组件默认配置的 metric

### prometheus 配置

重复一句, **你只是使用 prometheus**, 根据「学习思维三部曲: what->how->why」, 只是 `what` 其实很简单(简单不代表低级, 命名很简单却还没有完成的问题比比皆是), 这里都使用默认配置运行起来即可

参考文件, 即上面 docker-compose 中配置的:

```yml
    prometheus:
        image: prom/prometheus
        ports:
            - "9090:9090"
        volumes:
            - ./config/prometheus.yml:/etc/prometheus/prometheus.yml
```

对应其实只需要修改:

```yml
# my global config
global:
  scrape_interval:     15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.
  evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.
  # scrape_timeout is set to the global default (10s).

# Alertmanager configuration
alerting:
  alertmanagers:
  - static_configs:
    - targets:
      # - alertmanager:9093

# Load rules once and periodically evaluate them according to the global 'evaluation_interval'.
rule_files:
  # - "first_rules.yml"
  # - "second_rules.yml"

# A scrape configuration containing exactly one endpoint to scrape:
# Here it's Prometheus itself.
scrape_configs:
  # The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.
  - job_name: 'prometheus'

    # metrics_path defaults to '/metrics'
    # scheme defaults to 'http'.

    static_configs:
    - targets: ['localhost:9090']
  - job_name: 'hyperf'
    static_configs:
    - targets: ['ms:9502']
  - job_name: 'grafana'
    static_configs:
    - targets: ['grafana:3000']
```

只需要添加需要监控的 job 即可, 这里我添加 `grafana` 和 `hyperf` 作为示例

### grafana 配置

grafana 纯 webUI 交互, 启动后按照页面提示, 就可以添加 prometheus 作为数据源, 然后可以选择一个 `grafana` 的 UI 模板, 就可以看到采集的效果了

同理, hyperf 官方团队也提供了 UI 模板, 在 grafana theme 中搜索即可, 替换相应的参数, 就可以看到 hyperf 默认提供的 metric 参数

### 实践小结

**实践是检验真理的唯一标准**, 这句话在面对新技术, 面对未知时尤其有用, 学习一个新知识, 通常面对的知识 `what` 一级的内容, 很多时候天然具备 **简单** 属性, 这个时候的摸爬滚打几乎是 **最公平** 的 -- 只需要像小孩子一样, 保持好奇心, 多试几次.

## prometheus 的典型生产实践

项目中使用 prometheus 监控最常见的 2 个场景:

- api 监控 -> 扩展在所有的业务服务监控
- db 监控 -> 扩展到所有基础服务监控

到了这里, 我们需要补充一点 prometheus 的基础知识:

> prometheus docs: https://prometheus.io/docs/introduction/overview/

prometheus 文档并不长, 阅读大概需要 15min

PS: 阅读英文有障碍, 其实主要是心理作用, 安装一个网页翻译查词插件, 哪里不会点哪里, 常用的计算机词汇熟悉后, 几乎就不会影响阅读体验了

摘抄一下需要掌握的基础知识:
- 基础概念
    - metric: 监控指标, 比如 pv_total
    - label: 相当于 tag, 附在 metric 上, 可以做细化的查询分析
- 指标类型
    - count: 只增计数器, 比如 pv
    - gauge: 可增减计数器, 比如当前 qps
    - histogrm: 直方图, 比如 95值
- query: PromQL

熟悉 MySQL 记忆有一些代码量的读者, 我相信 `show me the code` 绝对是绝佳的学习方式:

```bash
# api 访问次数
sum(rate(app_api_query_time_count{api_name="self"}[1m]))
# service 对应 request_uri
sum(rate(app_api_query_time_count{api_name="self"}[1m])) by (service)

# api 访问耗时95值
# https://prometheus.io/docs/concepts/metric_types/#histogram
histogram_quantile(0.95, sum(rate(app_api_query_time_bucket{api_name="self"}[1m])) by (le))
histogram_quantile(0.95, sum(rate(app_api_query_time_bucket{api_name="self"}[1m])) by (le, service))

# 表查询次数
sum(irate(app_db_query_time_count{db=~"mysql:.*"}[1m])) by (db,table)
# 表访问次数
# type: curd
sum(rate(app_db_query_time_count{db=~"mysql:.*"}[30s])) by (type)
# db 响应时间95值
histogram_quantile(0.95, sum(irate(app_db_query_time_bucket{db=~"mysql:.*"}[1m])) by (le))
# db 平均查询时间
sum(rate(app_db_query_time_sum{db=~"mysql:.*"}[30s]))/sum(rate(app_db_query_time_count{db=~"mysql:.*"}[30s]))

# api qps
sum(rate(app_api_query_time_count{api_name!="self"}[1m])) by (api_name)
# api 平均响应时间
sum(rate(app_api_query_time_sum{api_name!="self"}[1m])) by (api_name)/sum(rate(app_api_query_time_count{api_name!="self"}[1m])) by (api_name)
```

PS: 关于95值, 熟悉 web 性能指标对 95值 就不陌生了, 直译是 `95%的请求在多久时间内返回`, 95值 在监控服务指标上使用广泛, 如果需要更好的性能, 可以使用 97值 甚至 99值, 作为衡量标准

## 一次 PHP微服务 prometheus 落地实践之旅

监控由微服务基础设施团队下的服务保障组来统一维护, 落实好 技术执行(业务开发) / 技术决策(架构师/服务 owner) / 技术支撑(基础设施/服务保障) 3大角色的分工合作, 上线 prometheus 选择了对业务方几乎零打扰的方案:

- prometheus 使用的单独的 redis 进行存储

```php
<?php

namespace Mt\Metric;

use Hyperf\Redis\RedisFactory;
use Hyperf\Metric\Adapter\Prometheus\Redis;

class RedisStorageFactory
{
    public function __invoke()
    {
        $redis = container(RedisFactory::class)->get('metric');
        Redis::setPrefix(config('app_name'));
        return Redis::fromExistingConnection($redis);
    }
}
```

- 添加默认 metric 配置

```php
'metric' => [
        'enable' => true,
        'default' => 'prometheus',
        // 不适用单独进程
        'use_standalone_process' => false,
        // 不使用默认 metric
        'enable_default_metric' => false,
        // 5s 统计周期
        'default_metric_interval' => 5,
        'metric' => [
            'prometheus' => [
                'driver' => \Hyperf\Metric\Adapter\Prometheus\MetricFactory::class,
                // 自定义模式
                'mode' => \Hyperf\Metric\Adapter\Prometheus\Constants::CUSTOM_MODE,
                // hyperf/metric 已修复这个问题, 使用 _ 作为分隔符
                'namespace' => \Hyperf\Utils\Str::camel(env('APP_NAME', 'app')),
            ],
        ],
    ],
```

- 统一监控 api/db

```php
<?php

namespace Mt\Metric;

use Hyperf\Metric\Metric;

class Prometheus
{
    public static function isEnable()
    {
        return config('metric.enable') === true;
    }

    public static function apiQuery($api, $http_code, $use_time, $method = '')
    {
        if (!self::isEnable()) {
            return;
        }
        $name = 'api_query_time';
        $labels = [
            'hostname' => php_uname('n'),
            'service' => config('app_name', 'app'),
            'api_name' => $api,
            'code' => (string)$http_code,
            'method' => $method,
        ];
        Metric::put($name, (float)($use_time * 1000), $labels);
    }

    public static function dbQuery($type, $db, $table, $use_time, $result = 'ok')
    {
        if (!self::isEnable()) {
            return;
        }
        $name = 'db_query_time';
        $labels = [
            'hostname' => php_uname('n'),
            'service' => config('app_name', 'app'),
            'type' => $type,
            'db' => $db,
            'table' => $table,
            'result' => $result,
        ];
        Metric::put($name, (float)$use_time, $labels);
    }

    public static function responseCode($api, $response_code)
    {
        if (!self::isEnable()) {
            return;
        }
        $name = 'api_response_code';
        $labels = [
            'hostname' => php_uname('n'),
            'service' => config('app_name', 'app'),
            'api' => $api,
            'response_code' => (string)$response_code,
        ];
        Metric::count($name, 1, $labels);
    }
}
```

- 最后一步, 由一个基础设施的 tools 服务中暴露路由给 prometheus 服务即可

```php
// metrics
$renderer = new RenderTextFormat();
$prometheus_collector = new PrometheusCollector();
$micro_services = ['app1', 'app2', 'app3'];
foreach ($micro_services as $service_name) {
    Router::get('/' . $service_name . '/metrics', function() use($service_name, $prometheus_collector, $renderer) {
        $prometheus_collector->setPrefix($service_name);
        return $renderer->render($prometheus_collector->collect());
    }, ['middleware' => [MetricAuthMiddleware::class]]);
}
```

## 写在最后

prometheus 的监控之旅还没有走完, 不断暴露问题, 不断提升监控的数量和质量, 对了, 用一句经典的话语:

> 用发展的眼光来看问题