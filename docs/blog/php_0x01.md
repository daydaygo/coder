# PHP| 开发拾遗 0x01

记录 PHP 开发中的二三事

提纲:

- 请使用 const 常量
- 再说 `= == ===`
- 原生函数 `json_encode()/array_merge()/preg_match_all()`
- 一次「压平」 if 的踩坑记录
- 简单的「频次限制」, 常见场景比如重复点击重复请求

## 请使用 const 常量

其实开始我是「拒绝」的, 理由是增加了一层 **映射**, 就是变相的增加了一层 **复杂度**, 比如最常见的常量应用场景, 表示各种状态:

```php
const STATUS_UNDO = 'undo';
const STATUS_DOING = 'doing';
const STATUS_SUCCESS = 'success';
const STATUS_FAIL = 'fail';
```

之前一直抱有的观点是, 记忆一次 `undo/doing/success/fail` 就够了, 比没要 **常量** 再来一层, 用的地方保持就好了.

但最近几次密集的使用常量, 让我改变了这个想法 -- 常量是可以 **IDE** 提示的. 由于编程是一件 **精确** 的活, 原来的方式需要 **精确记忆细节** 或者 **复制粘贴**, 有了 IDE 提示后就简单多了.

同时再想想下面的场景:

- 如果字符串不是这么简单, 有点长(这种情况太多了), 经常在 array key 类似的场景用到
- 如果这个字段是数据表中映射出来的, 有十多个类似字段
- 如果这个表示状态用的数字, 比如 0-undo, 1-doing, 2-success, 3-fail, 就需要在使用的使用带上注释了

```php
if (3 == $status) { // 失败状态
    // fail case
} else if (2 == $status) {
    // success case
}
```

综上, 常量其实一件 **轻松** 的事儿

## 再说 `= == ===`

上面的代码其实已经示范了一个例子, `=` 与 `==` 是初学很容易遇到的 **困惑**, 这里再简单重申一下定义:

- `=`: 赋值语句, 给变量赋值
- `==`: 判断是否 **相等**
- `===`: 判断是否 **全等**, 区别与 `==` 的是要求**数据类型一致**

理解清楚定义, 然后再看 2 个场景:

```php
if ($status = 1) { // 如果这里把 == 少写了
    //
} else {
    //
}
```

上面的错误基本每个人都犯过吧, 尤其是是使用 `if ($var = xxx)` 确实有另外一个用途:

```php
$var = 'xxx'; // 给 $var 赋值
if ($var) {
    //
}

if ($var = 'xxx') { // 常见的缩写方式

}
```

有效避免这种错误的方式:

```php
if (1 == $status) { // 如果少写了 =, IDE会自动提示
    //
} else {
    //
}
```

推荐这样的写法, 因为最近一次 bug 就是这个问题导致的, 指不定哪个夜黑风高的晚上, 又写出这种 bug 出来.

再来说说 `==` 和 `===`:

```php
if (strpos('abc', 'a')) { // 判断字符串是否存在
    echo 'yes';
} else {
    echo 'no';
}
```

这里明显是个错误的例子, 因为 `strpos()` 函数返回的是匹配到的 **起始位置**, 是 `int 0`, 不匹配是返回 `bool false`, 正确的做法应该是:

```php
if (strpos('abc', 'a') !== false) {
    echo 'yes';
} else {
    echo 'no';
}
```

`==` 和 `===` 的关键点就在于数据类型上, 弱类型是 **对人友好**, 强类型是 **对机器友好**.

## 原生函数 `json_encode()/array_merge()/preg_match_all()`

接着上面的 `strpos()` 继续聊几个原生函数.

因为 json 的**大行其道**, `json_encode()/json_encode()` 就会经常使用到了, 直接说几个要点:

- 需要 `ext-json` 扩展支持, 不过 PHP 默认是开启这个扩展的
- json 数据类型: bool int string array object, 因为 PHP 的弱类型, 带来几个需要注意的类型转换的问题:

```php
json_encode(1); // int <-> string
json_encode('1');

json_encode(true); // bool <-> string
json_encode('true');

echo json_encode(['a' => '']); // 空字符串: {"a":""}
echo json_encode(['a' => []]); // 空数组: {"a":[]}
echo json_encode(['a' => new \Stdclass()]); // 空对象: {"a":{}}
```

尤其要注意这里的 `new \Stdclass()`, 毕竟 PHP 编程中, 经常是只使用 array 这一种数据结构.

这也是为什么要使用 `json_decode($str, true)` 的原因, 默认返回 `Stdclass` 类型, 带 `true` 参数才是 array 类型

- Unicode转义

```php
echo json_encode('中国'); // "\u4e2d\u56fd"
echo json_encode('中国', JSON_UNESCAPED_UNICODE); // "中国"
```

这样就不用拿到之后还要 `decode` 一下了, 而且实际使用的字符量减少了, 具体可以参考[鸟哥的 blog](http://www.laruence.com/)

`array_merge()` 的坑不知道有多少人踩过, 在 PHP manual 上是有说的: **不能递归合并**, 所以很多框架都提供了辅助函数来处理:

```php
// 比如 yii 框架的 \yii\helpers\BaseArrayHelper::merge()
public static function merge($a, $b)
{
    $args = func_get_args();
    $res = array_shift($args);
    while (!empty($args)) {
        $next = array_shift($args);
        foreach ($next as $k => $v) {
            if ($v instanceof UnsetArrayValue) {
                unset($res[$k]);
            } elseif ($v instanceof ReplaceArrayValue) {
                $res[$k] = $v->value;
            } elseif (is_int($k)) {
                if (isset($res[$k])) {
                    $res[] = $v;
                } else {
                    $res[$k] = $v;
                }
            } elseif (is_array($v) && isset($res[$k]) && is_array($res[$k])) {
                $res[$k] = self::merge($res[$k], $v);
            } else {
                $res[$k] = $v;
            }
        }
    }

    return $res;
}
```

`preg_match_all()` 是正则匹配, 比较适合日常使用了, 这里简单mark一下:

```php
preg_match_all("/href='(.*?)'/", $str, $output);
// $output[0] 返回所有满足整段正则的字符串
// $output[1] 开始以此返回 () 中匹配到的值, 类似 perl 中的 $1
```

## 一次「压平」 if 的踩坑记录

首先是一个简单风格的对比:

```php
// if-else
if ('a' == $a) {
    //
} else {
    //
}

// 压平一点
$a = 'xxx';
if ('a' == $a) {
    //
}
```

个人倾向于后一种, 这样可以只有考虑一次 `if`, 当然具体情况要具体分析.

来看看具体的采坑记:

```php
function checkStatus() { // 读取配置请检查条件是否满足
    $a = getConfig(); // 获取配置

    // 基础条件: 任一一个不满足就返回 false
    if ($a['base']) {
        foreach ($a['base'] as $v) {
            if (!$v) {
                return false;
            }
        }
    }

    // 附加条件: 满足基础条件的情况下, 还需要满足附加中的一项
    if ($a['option']) {
        foreach ($a['option'] as $v) {
            if (!$v) {
                return true;
            }
        }
    }

    return false;
}
```

这是初版的代码, **基础条件** 和 **附加调价** 都是可配置的, 如果没有配置 **附加条件**, 就出现了 2 个问题:

- 没有对变量值进行检测, 尤其是 array 的 key, 这是 PHP 中 **非常常见** 的错误
- 只是简单的改为 `if (isset($a['option']})`, 恭喜你, **逻辑错误**, 这也是 **根据报错改代码** 容易遇到的问题

正确的版本:

```php
function checkStatus() { // 读取配置请检查条件是否满足
    $a = getConfig(); // 获取配置

    // 基础条件: 任一一个不满足就返回 false
    if ($a['base']) {
        foreach ($a['base'] as $v) {
            if (!$v) {
                return false;
            }
        }
    }

    // 附加条件: 满足基础条件的情况下, 还需要满足附加中的一项
    if (empty($a['option'])) {
        return true; // 通过了基础条件, 到这里就需要返回 true
    }
    foreach ($a['option'] as $v) { // 判断附加条件
        if (!$v) {
            return true;
        }
    }

    return false;
}
```

这里提醒 2 点:

- 使用 `isset() / empty()` 来进行变量检测
- 尽管大部分情况下, **写业务** 看起来就是写 `if-else`, 但是请务必小心, 随着复杂度提升, 很容易出错的

## 简单的「频次限制」

常见场景: 防止页面重复点击后端重复处理, 加入 60s 点击限制

先来看最终结果:

```php
if (1 == MyRedis::incr(MyRedis::CLICK_ITEM_A)) {
    MyRedis::expire(MyRedis::CLICK_ITEM_A, 60); // 60s 过期时间
    // 业务逻辑
}
```

`MyRedis()` 类是使用 facade 设计模式, 对  `exe-redis` 扩展的封装, 这样业务不用关心 redis client 初始化的相关的细节:

```php
$redis = new \Reids();
$redis->connct('127.0.0.1');
$redis->auth('password')
$redis->select(1);

// 方法参数和返回值 和 ext-redis 扩展保持一致
$redis->incr($key);
MyRedis::incr(MyRedis::CLICK_ITEM_A);
```

而且自己封装的 `MyRedis()` 类还可以使用常量, 有效标识出 **具体业务**

- 关于 facade 设计模式, 可以参考我之前的 [blog - laravel源码解读](https://www.jianshu.com/p/b7ea3f2a55f6)
- rate limit 的更多应用, 可以参考 [redis 官方文档 - incr](https://redis.io/commands/incr)

## 写在最后

- 细节出魔鬼
- practice make perfect
