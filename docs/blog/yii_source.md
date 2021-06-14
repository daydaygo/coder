# yii| 源码解读

- [yii| 源码解读](https://www.jianshu.com/p/fd85383783eb)

本篇博客阅读指南:

- php & 代码提示: 工欲善其事必先利其器
- yii 源码阅读指南: 整体上全貌上进行了解
- 之后的章节: 细节入手, 没错, 都是知识点

写完上篇 [yii 框架解读] 后, 发现干货有点少, 写来写去还是 **底层是服务容器** 这样的老生常谈. 虽然这个真的很重要, 我认为理解服务容器的 php 程序员, 算是 **境界提升**(至少不用自嘲「码畜」了吧). 这篇就实实在在的阅读 yii 框架的源码, 希望可以给大家带来更多干货.

**备注**: 因为 yii 有一个惯用套路, 框架层实现使用 `BaseXXX`, 具体使用需要用 `Xxx` 来继承而不是直接使用 `BaseXxx` 类, 而最底层的基类 `BaseObject` 使用这种方式后的类 `Object`, 在 [php7.2](http://www.php.net/ChangeLog-7.php#7.2.0) 中被添加为关键字.

## 请使用 phpstorm

详细介绍一个 IDE 怎么用不太现实, 各种黑科技还是自己体会, 我比较喜欢凭自己落笔时的印象来判断 -- 第一时间想到的, 往往是最熟悉的.

- 错误提示: 单词拼写错误, 低级语法错误, 这些开发过程中最常见的问题
- 代码提示: 函数以及函数的参数和返回; 类以及类的属性和方法; 等等等等
- 跳转: 方便的跳转对阅读代码有多重要就不多说了, 而且可以跳转 php 内部函数和类, 明显减少 php manual 的使用

还有其他很多高级功能, 比如 重构/db连接/版本控制, 这些都不是重点, 或者说锦上添花, 用 phpstorm 的理由非常简单:

> 明显提高开发效率. 崇尚极简也同样适用, 只关心编辑功能也会发现效率提升.

**友情提示**: 开发机请使用 16G 内存. 更多使用小技巧可以参考我的 [wiki - tools - ide](http://www.php.net/ChangeLog-7.php#7.2.0)

## 代码提示

phpstorm 之所以会让人感觉很 **智能**, 很多地方都来自于完善的 **代码提示**. 当然现实是很多人写代码, 不写注释.

> 我 TM 代码都写不完, 你还要我写注释?!

不就这个话题展开, 但是可以给一个关于开源代码选用的标准, 如果你打算使用的开源代码注释和文档不完善, 建议你最好不要选用. 否则, 一定要确认可以对接的人(同样适用于接手维护旧代码).

这里八卦一下, 之前一直有人 *骂* swoole 文档烂, 到处是 *坑*. 我这里说句公道话, swoole 的 wiki 里写到有 2 个开源项目提供代码提示(关于代码提示, 可以参考之前的 [blog - 聊一聊 php 代码提示](http://www.jianshu.com/p/b3daadb3c4c5), 一种使用 php Reflection(反射) 实现, 一种是提取代码注释然后手动完善. 并且 swoole 的 wiki 1400+ 页, 下面的评论往往也是干货满满. **在你骂别人文档烂, 坑多的时候, 你凭什么这样说?**

> 能力足够, 可以参加核心开发组; 文档不够完善, 但是它可直接编辑; 使用发现代码提示不够好, 代码提示的开源项目可以参与. 最后是坑多的说法, 有没有想过更多是经验不够, 所以没用好工具.

发这个牢骚不是想 **探究人性** 之类的, 只是面对有些现实, 其实明明可以往 **好一点** 的方向前进一小步. 当然, 我可不敢公然和 **喷子** 叫板.

得益于 php 语言的简单, 代码提示在这里也非常简单, 而且 yii 框架的代码提示做得非常好, 几乎任何输入的地方, 都会有 IDE 的自动提示.

注释的语法很简单: **指令(@开头) + 指令内容**. 基本都是只要看到就能理解什么意思:

```
// 描述函数参数, 格式: @param type var define
@param string $name the property name

// 描述函数返回值, 格式和上面类似: @return type define
@return mixed the property value or the value of a behavior's property

// 变量提示, 格式也类似: @var type define
@var array the attached event handlers (event name => handlers)
```

当然还有一些其他提示, 各有作用, 使用较少, 就不一一列举了.

**备注**: 使用 swoole 的过程中也被回调函数难以代码提示烦恼过, 所以参与了 [swoole-ide-helper](https://github.com/swoole/ide-helper) 项目, 提交 pr, 来一起改善 swoole 的编程体验.

> 道理很多依旧过不好人生什么的, 原因可能并没有那么复杂, 真的只是因为太懒了一点.

## yii 源码阅读指南

yii 框架的源码很简单, 层次很清晰:

- `yii\base\BaseObject`: 基类, 几乎所有类都继承自这个类, 使用 `__get()/__set()` 等魔术方法, 方便操作类属性等

```php
class BaseObject implements Configurable
{
    public static function className()
    {
        return get_called_class(); // 等同于 static::CLASS, 区别与 get_class()
    }
    public function __get($name){};
    public function __set($name, $value)
}
```

这里实现了 `Configurable` 接口, 给框架了带来了 **基于配置** 的超强灵活性, 后面会有具体代码讲到

- `yii\base\Component`: 组件, 继承自 `BaseObject`, yii 框架提供的所有功能, 几乎都是 `component`, 这样就可以 `\Yii:$db` 这样的形式来调用

```php
class Component extends BaseObject
{
    private $_events = [];
    private $_behaviors;
}
```

`Component` 扩展了 `BaseObject`, 并为所有组件定义了特性: *property*, *event* and *behavior*

- `\yii\di\Container`: 容器, 这个概念就不再啰嗦

- `\BaseYii`: yii 框架主体, 定义了部分框架运行的功能, *log*, *profile* 等

- `\Yii`: 实例化 `BaseYii`, 这种方式 yii 中随处可见, 基本定义基础功能, 具体使用时继承基类并自己按需扩展. `\Yii` 会同时启动一个 `Container`.

```php
class Yii extends \yii\BaseYii
{
}

spl_autoload_register(['Yii', 'autoload'], true, true);
Yii::$classMap = require __DIR__ . '/classes.php';
Yii::$container = new yii\di\Container();
```

这里使用 classMap` 的方式来注册框架核心类, 性能会比 composer 的 psr-4 稍高, 但是也导致了你有 2 种方式来管理依赖, 这点我是持 **消极** 态度的.

- `web\index.php`: 入口脚本, 加载配置和 `Yii`, 实例化 `application`, 来完成请求

```php
defined('YII_DEBUG') or define('YII_DEBUG', true);
defined('YII_ENV') or define('YII_ENV', 'dev');

require(__DIR__ . '/../vendor/autoload.php');
require(__DIR__ . '/../../vendor/yiisoft/yii2/Yii.php');

$config = require(__DIR__ . '/../config/web.php');

(new yii\web\Application($config))->run();
```

得益于 `BaseObject` 和 `Component`, 几乎所有特性, 都可以通过这里的 `$config` 进行配置.

## 关于链式调用

这次阅读源码的过程中, 在使用 `yii\widgets\DetailView` 卡了一小会, 被自己之前关于链式调用的理解给绕进去了. 首先看第一种方式 `$this`:

```php
class a {
    public $b = 0;
    function aa() {
        $this->b += 1;
        return $this;
    }
    function bb() {
        $this->b += 2;
        return $this;
    }
}

$a = new a();
$a->aa()->bb()->aa();
echo $a->b;
```

通过在类方法中返回 `$this`, 从而实现链式调用, 这样的写法, 可以参考 `yii\db\Query` 的源码, 使用链式调用来构建 sql 语句.

因为对这种方式 **印象太深**, 导致忽略了下一种更常见的方式:

```php
class A {
    public function b()
    {
        $b = new b();
        return $b; // 返回 b 对象
    }
}

class B {
    public function c() {
        echo 'czl';
    }
}

$a = new A();
$a->b()->c();
```

使用 **其他对象作为自己的属性或者函数函数返回**, 这是更常见的链式调用, 而在 yii 中, 这种方式更是随处可见, 这里用 `\yii\widgets\DetailView` 中使用 `\yii\i18n\Formatter` 来展示一下 **基于配置** 的超强灵活性:

```php
DetailView::widget([
    'model' => $model, // 和 Model 类无缝配合
    'attributes' => [
        'id',
        'title',
        'content:ntext',
        'tags:ntext',
        'create_time:datetime',
        'update_time:datetime',
        [
            'attribute' => 'author_id',
            'value' => $model->author->nickname,
        ],
    ],
    'template' => '<tr><th width="120px">{label}</th><td{contentOptions}>{value}</td></tr>',
    'formatter' => [
        'class' => \yii\i18n\Formatter::class,
        'datetimeFormat' => 'short',
    ]
]);
```

查看 [api 文档](http://www.yiichina.com/doc/api/2.0/yii-widgets-detailview#$attributes-detail), 会发现这里的 attribute 非常的强大:

- 这里的 attribute 属性, 可以和 Model 中的 attribute 对应
- 这里的 attribute 属性, 可以使用 `attribute:format:label` 格式, 其中的 format 就是对应的
`\yii\i18n\Formatter`, 大部分常用的格式化方法, 这里都有定义, 比如这里的 `create_time:datetime` 表示使用 `\yii\i18n\Formatter` 中的 `asDatetime()` 进行格式化

你以为到这里就结束了么:

- `template`: 直接可以配置页面的 html
- `formatter`: 不止可以用 `\yii\i18n\Formatter`, 还可以配置

还没完, 我们再全局也是可以配置的 `config/web.php`:

```php
$config = [
    'id' => 'myYii',
    ...
    'components' => [
        'formatter' => [
            'datetimeFormat' => 'Y-m-d H:i:s',
        ]
    ],
];
```

当然, 全局的配置, 会被这里具体使用的地方给覆盖掉.

另外还有 `\yii\widgets\ActiveForm` 和 `\yii\widgets\ActiveField`

> 非常推荐大家阅读一下这块的代码, 尝试动手改改, 只要这里理解清楚了, 对框架的整体理解基本没问题了.

**PS**: 我之前表达过观点, 前后端分离是大势, phper 应该更关注 **后端**, 关注写出更好的 api. 但是 yii 这种前后端无缝对接高可配置的方式, 还是把我惊艳到了. 但是我的观点还是没有变, phper 还是应该更关注后端, 我倾向于把 yii 应用到不需要 **设计** 的场合, 比如管理后台.

## 关于 db

很多初级 phper 会感觉 db 这块的内容 **很多**, 一方面是数据相关的基础知识就很多(基础的增删改查并不是难度好不好), 然后 php 和数据库联动的过程, 又增加了一层抽象. 我之前的 [blog - hyperframework WebClient 源码解读](http://www.jianshu.com/p/cf39804b7c04) 也提过这样一个观点:

> 层出不穷的工具, 目的就是对现有问题作出更 **易用** 的抽象. 但是伴随抽象的不断增多, 基础部分的更加不可见, 导致越来越容易 **摸不着头脑**. 所以我希望我写的东西, 能在一开始就给大家划定出一个核心的范围, 而不是有一个工具的堆砌.

先来聊 db 的第一个话题, php 使用 db 的三种方式.

### 3 种 db 访问方式

数据库作为一个服务, 其实 php 是作为 client 端来访问. 数据库的架构通常是分层结构, 最外层的和我们平时写的 **接口** **网关** 其实是一样的 -- 通过暴露 api 来提供服务. 只是我们最终提供的 web 服务, 走的是 http 协议, 而数据库走的数据库的协议, 比如和 mysql 通信需要实现 mysql 协议. 嗯, 这个比较底层了, 协议的细节被抽象掉了, 最终暴露给我们的, 其实就是 **sql**.

> 这就是我划定的核心范围, 说是 3 种方式, 本质还是就是执行 sql 语句而已.

```php
// 直接执行 sql 语句
$postStatus = \Yii::$app->db->createCommand('SELECT id,`name` FROM poststatus')->queryAll();
$postStatus = array_column($postStatus, 'name', 'id');

// 使用查询构造器
$postStatus = (new \yii\db\Query())
->select(['name', 'id'])
->from('poststatus')
->indexBy('id')
->column();

// 使用 ActiveRecord
$postStatus = \app\models\Poststatus::find()
->select(['name', 'id'])
->indexBy('id')
->column();
```

三种方式的关系也很简单:

- ActiveRecord 调用 `find()` 后, `@return ActiveQuery the newly created [[ActiveQuery]] instance`, 其实就是返回一个拼上表名的 ActiveQuery 实例
- ActiveQuery 通过链式调用, 拼接出一个完整的 sql
- 最终和 `\Yii::$app->db->createCommand()` 执行没啥区别, 只是 ActiveQuery 又提供了一些方法, 对查询到的结果集做一些处理

这也是目前大部分框架采用的方式 -- 提供三种方式给大家使用. 这里还是发表一下我个人的观点, 我们的 [hyperframework](http://hyperframework.com/cn/manual) 中是不提供 ActiveQuery 这样的实现的, 因为我们相信, 大部分情况下, sql 是最好的选择.

- 实现一个 `ActiveQuery` 类并不难, 用起来也不难, 但是 sql 是需要要掌握的, 掌握了 sql 之后其实就可以用第一种方法解决问题了.
- `ActiveQuery` 在复杂 sql 下面非常难写, 甚至不能 -- 来自游戏数据统计的血泪史

当然, `ActiveQuery` 也有优点和合适的场景, 比如代码提示和条件查询:

```php
$query = $db->select('xxx');
if (!empty($search['a'])) {
    $query = $query->where('a', $search['a']);
}
```

### 关联查询

上一节只是 **浅尝辄止** 的提到 `ActiveRecord`, 这里详细讲讲, 然后再深入一点. 先提个醒: 设计出 `ActiveRecord` 这样的抽象, 真的非常厉害.

`ActiveRecord`, 中文翻译为活动记录, 对应于 MVC 中的 Model 这一层, 但是它是和数据库结合最紧密的地方. 一个 `ActiveRecord` 类, 用来对应数据库里的一张表, 一个 `ActiveRecord` 实例化对象, 用来对应这张表里面的一条记录, 进而通过对象的 新建/属性修改/方法调用, 实现数据库的增删改查.

```php
// 增
$post = new Post();
$post->title = 'daydaygo';
$post->save();
// 查
$post = Post::find(1);
// 删
$post->delete();
// 改
$post->title = 'czl';
$post->save();
```

你看这样的代码, 是不是感受不到 sql 的存在, 但是你却轻松实现了需要的功能. 这就是我认为 **厉害的地方**.

再来看更厉害的 -- 关联查询:

```php
// Post 中定义和 author 的关联
public function getAuthor()
{
    return $this->hasOne(Adminuser::className(), ['id' => 'author_id']);
}

// 这样访问 author 就简单了
$post->author;
```

这里先解释一下, `$post->author` 会去寻找 Post 中的 `getAuthor()` 方法, 然后根据这里定义的关联关系, 执行查询, 并将查询到 author 记录, 赋值给 `$post->author` 属性. 这里有 2 个细节:

- author -> getAuthor() 其实是通过 `yii\db\BaseActiveRecord` 中的 `__get()` 魔术方法实现的, 这也是 yii 核心的设计理念之一, 通过实现 `__get()` 等魔术方法, 让 *类* 更好用
- Post 的注释中有这样一句 `@property Adminuser $author`, 这样使用 `$post->author` 就有酸爽的代码提示了

关于关联查询, 这里还有 2 个细节:

- 查询缓存, 这也是 yii 为什么性能这么高的原因. 一点题外话, 在看源码的过程中, 有函数被标记不推荐使用, 点进入发现是使用缓存的姿势不够优雅, 强耦合

```php
// 关联查询
$user = User::findOne();
$orders = $user->orders; // 执行关联查询, 结果被缓存
unset($user->orders); // 清楚缓存, 重新查询
$orders2 = $user->orders;
```

- 多对多的查询, 需要注意查询上的优化:

```php
// 多次查询
$users = User::find()->all(); // 查询 user
foreach($users as $user){
    $oders = $user->orders; // 查询 order
}

$users = User::find()->with('orders')->all(); // 2次查询, 一次 user, 一次 order
foreach($users as $user){
    $oders = $user->orders; // 此处不会执行数据库查询
}
```

## 关于锁

基础稍差的话, 可能对锁的概念会有些陌生. 简单的解释是: 在多进程或者多线程编程的情况下, 同时访问同一个资源导致程序的最终结果不可控.

首选需要区分 2 个概念: 并发 vs 并行

- 并发 Concurrent: 多线程多进程场景下, 微观上 cpu 切换看, 快到人类无法直观感知极限(0.1s), 所以宏观上看起来是 *同时* 运行
- 并行 parallel: 真正的 **同时** 运行, 必须要都 cpu 支持

再来一个概念: 竞态资源

- 在某个资源上产生了并发访问, 导致程序执行后没有达到预期, 那么这个资源就是竞态资源

套用一下数据事务的例子: 2 个账户间转账, 必须加事务, 只有一个账户上钱扣了, 另一个账户上钱增加了, 才算完成, 这时候去取到的 2 个账户的余额才是准确的.

好了, 前戏差不多了, 这里来讲讲 yii 中用到的 2 个锁.

### mutex 互斥锁

yii 中特地添加了 `yii\mutex\Mutex`, 并且提供了不同驱动下(file, 不同 db)的实现

互斥锁的理念非常简单: 保证当前只有一个进程(或线程)访问当前资源

实现也非常简单, 就 2 个方法:

- acquire(): 使用前请求锁, 请求成功就执行业务逻辑, 失败就退出
- release(): 使用后释放锁

```php
function lock($lockName = NULL) {
    if (empty($lockName)) {
        $backtrace = debug_backtrace(null, 2);
        $class = $backtrace[1]['class'];
        $func = $backtrace[1]['function'];
        $args = implode('_', $backtrace[1]['args']);
        $lockName = base64_encode($class . $func . $args);
    }

    $lock = \Yii::$app->mutex->acquire( $lockName ); // 请求锁
    if (!$lock) {
        $err = "cannot get lock {$lockName}.";
        throw new \Exception($err);
    }

    register_shutdown_function(function() use($lockName) {
        return \Yii::$app->mutex->release($lockName); // 释放锁
    });

    return TRUE;
}
```

### db optimisticLock() 乐观锁

这个就隐藏的比较深了. 因为已经养成数据库中使用自动更新的 `create_time / update_time` 字段, 所以深入 ActiveRecord 的 update() 源码进去, 然后就发现了这家伙. 详细的解释可以看这里 [百度百科 - 乐观锁](https://baike.baidu.com/item/%E4%B9%90%E8%A7%82%E9%94%81)

```php
/**
 * @see update()
 * @param array $attributes attributes to update
 * @return int|false the number of rows affected, or false if [[beforeSave()]] stops the updating process.
 * @throws StaleObjectException
 */
protected function updateInternal($attributes = null)
{
    if (!$this->beforeSave(false)) {
        return false;
    }
    $values = $this->getDirtyAttributes($attributes);
    if (empty($values)) {
        $this->afterSave(false, $values);
        return 0;
    }
    $condition = $this->getOldPrimaryKey(true);
    $lock = $this->optimisticLock();
    if ($lock !== null) {
        $values[$lock] = $this->$lock + 1;
        $condition[$lock] = $this->$lock;
    }
    // We do not check the return value of updateAll() because it's possible
    // that the UPDATE statement doesn't change anything and thus returns 0.
    $rows = static::updateAll($values, $condition);

    if ($lock !== null && !$rows) {
        throw new StaleObjectException('The object being updated is outdated.');
    }

    if (isset($values[$lock])) {
        $this->$lock = $values[$lock];
    }

    $changedAttributes = [];
    foreach ($values as $name => $value) {
        $changedAttributes[$name] = isset($this->_oldAttributes[$name]) ? $this->_oldAttributes[$name] : null;
        $this->_oldAttributes[$name] = $value;
    }
    $this->afterSave(false, $changedAttributes);

    return $rows;
}
```

## 关于 log & error handler

写代码到一定程度, 就会开始意识到 log & error handler 的重要性, 然而在小白程序员升级打怪的过程中, 一直在写业务, 这 2 块关注太少以致有些 **苍白**. 并且这块也是我比较薄弱的地方, 几个月前在在添加 Exception 的时候卡住了.

> 知道短处, 补补就好了.

在聊这 2 块之前, 先补一下关于 **回调** 的基础知识:

- [call_user_func_array()](http://php.net/manual/en/function.call-user-func-array.php)
- [Callbacks / Callables](http://php.net/manual/en/language.types.callable.php)

平时大家可能这样写代码的情况不多, 不过如果接触过 [swoole](http://wiki.swoole.com), 写过一段时间的 **异步编程**, 这个知识点就再熟悉不过了, 在 swoole 的 wiki 中也特意提到过, 里面列举了 4 种, 官方文档这里列举了 5 种.

### log 模块

先看整体结构:

```sh
├── Logger.php
├── Dispatcher.php
└── Target.php
    ├── DbTarget.php
    ├── EmailTarget.php
    ├── FileTarget.php
    └── SyslogTarget.php
```

由 `Logger - Dispatcher - Target` 的 3 层结构:

- Logger: 日志 **入口**(生产者)
- Dispatcher: 日志的 **分发**(通道)
- Target: 日志 **处理**(消费者)

其实日志系统的设计已经相当成熟了, 几乎都采用 **消息队列** 的设计模式:

> 生产者 - 消费者 模型.

这里看一点代码细节:

- yii 框架中的 profile 功能, 可能大家有用过, 也是通过 `Logger` 实现的

```php
\Yii::beginProfile('block1');
// some code to be profiled
    \Yii::beginProfile('block2');
    // some other code to be profiled
    \Yii::endProfile('block2');
\Yii::endProfile('block1');

// yii\log\Logger
const LEVEL_PROFILE_BEGIN = 0x50
const LEVEL_PROFILE_END = 0x60
```

- 使用 `flush`: 先 **缓存** 一下, 然后再一起落地, 性能要比直接写直接落地高一些

```php
public function log($message, $level, $category = 'application')
{
    $time = microtime(true);
    $traces = [];
    // ...
    $this->messages[] = [$message, $level, $category, $time, $traces, memory_get_usage()]; // 暂时缓存到这里
    if ($this->flushInterval > 0 && count($this->messages) >= $this->flushInterval) {
        $this->flush();
    }
}
```

- 回调终于登场了, `register_shutdown_function()` 也不常见, 但是下面还会提到

```php
public function init()
{
    parent::init();
    register_shutdown_function(function () {
        // make regular flush before other shutdown functions, which allows session data collection and so on
        $this->flush();
        // make sure log entries written by shutdown functions are also flushed
        // ensure "flush()" is called last when there are multiple shutdown functions
        register_shutdown_function([$this, 'flush'], true);
    });
}
```

日志模块的代码还是很简单的. 实现日志模块其实并不难, 但是新手想用好日志却感觉有点 **经验积累** 的意思, 特别是遇到的问题的时候发现没有日志辅助定位问题. 我的建议也很简单:

> 多打日志, 多用日志.

## error handler 模块

如果说日志大部分时候只用 `Logger::info()` 这样调用一下就好了, `Exception` 天生就要复杂一点了, 因为完整的过程是这样的:

```php
try {
    // do something
    throw new \Exception("Error Processing Request", 1);

} catch (\Exception $e) {
    // handle error
}
```

但是, 其实只要记住这个基本 **骨架**, 任何地方都是同样的. 如果这块比较薄弱, 可以 [参考官方手册 - Exception](http://php.net/manual/en/language.exceptions.php) 多看一看.

yii 框架中 Exception 的使用很多, 所以看起来会比较凌乱, 但其实层次很清晰:

首先是 base, 这基本确定了 Exception 的分类:

- `\yii\base\Exception`: 异常基类, 下面会着重讲一下 `register()` 方法
- `\yii\base\ErrorException`: 处理未捕获的 php 错误和异常, 下面会着重讲一下 `register()` 方法
- `\yii\base\UserException`: 用户可见异常基类, 这个很重要, 添加了一个明显分类
- `\yii\base\XxxException`: 其他通用异常

然后就是根据应用不同:

- `\yii\web\XxxException`: web 应用下的异常
- `\yii\console\XxxException`: console 应用下的异常

好了, 再来看点源码:

- `\yii\base\ErrorException` 中的 `register()` 方法: 注册函数回调; 兼容 HHVM

```php
public function register()
{
    ini_set('display_errors', false);
    set_exception_handler([$this, 'handleException']);
    if (defined('HHVM_VERSION')) {
        set_error_handler([$this, 'handleHhvmError']);
    } else {
        set_error_handler([$this, 'handleError']);
    }
    if ($this->memoryReserveSize > 0) {
        $this->_memoryReserve = str_repeat('x', $this->memoryReserveSize);
    }
    register_shutdown_function([$this, 'handleFatalError']);
}
```

- `\yii\base\UserException`

```php
/**
 * UserException is the base class for exceptions that are meant to be shown to end users.
 * Such exceptions are often caused by mistakes of end users.
 */
class UserException extends Exception
{
}
```

一个明显的场景, 就是 http 的 4xx 错误:

```php
class HttpException extends UserException
{
    /**
     * @var int HTTP status code, such as 403, 404, 500, etc.
     */
    public $statusCode;
}
```

- 还有一个常用的方式(**套路**), 将应用整个包在 `try-catch` 中, 统一捕获异常

```php
// 入口脚本: web/index.php
require(__DIR__ . '/../../vendor/yiisoft/yii2/Yii.php');
$config = require(__DIR__ . '/../config/web.php');
(new yii\web\Application($config))->run();

// \yii\base\Application
public function run()
{
    try {
        $this->state = self::STATE_BEFORE_REQUEST;
        $this->trigger(self::EVENT_BEFORE_REQUEST);

        $this->state = self::STATE_HANDLING_REQUEST;
        $response = $this->handleRequest($this->getRequest());

        $this->state = self::STATE_AFTER_REQUEST;
        $this->trigger(self::EVENT_AFTER_REQUEST);

        $this->state = self::STATE_SENDING_RESPONSE;
        $response->send();

        $this->state = self::STATE_END;

        return $response->exitStatus;
    } catch (ExitException $e) {
        $this->end($e->statusCode, isset($response) ? $response : null);
        return $e->statusCode;
    }
}
```

## 写在最后

聊了这么多, 内容多了之后, 也会有些 **杂乱**, 而且也无法深入到太多的细节. 我比较满意的是, **在一开始我就计划好使用脑图, 尝试整体的理解架构**, 那些记下的细节, 反而有点像 **意外之喜**.

> 大型开源项目的源码是一座宝矿.

不得不说很多人还停留在 **我 TM 业务代码都写不完, 你让我看这个?** 的阶段, 所以业界也容易流传开一些类似 **10x程序员** 或者 **我们从来不认为月薪 2w 以下的是程序员** 这样的话题, 究其原因:

> 编程也是一项技艺, 如同江湖中对武功的崇拜一样, 程序员也会对自己的一技之长产生骄傲.

也许确实没有大段的时间去阅读源码, 但是使用方法时, 多点进去看看, 也经常会有所收获, 比如 cache 相关的方法:

```php
// 平时使用
Yii::$app->cache->set('key', 'value');

// 进入会发现, 可以设置 过期时间 + 缓存依赖
public function set($key, $value, $duration = null, $dependency = null)
```

