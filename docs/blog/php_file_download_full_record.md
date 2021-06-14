# php 文件导出全记录

phper 几乎都接触过来自「管理后台」的需求吧, 其中几乎少不了「文件导出功能」. 做了三年 phper, 几乎接了三年这样的需求, 最近虽然在写服务器, 不过运营小哥既然「诚心诚意地求了」, 我当然要「大发慈悲地搭把手」, 于是又折腾了一次, 所以干脆汇总一下, 也帮助其他 phper 快速上手.

## 最简单的 demo

导出 csv 原理其实很简单, 设置 http header, 然后输出 csv 文件即可

- 读取 csv 文件, 然后导出

```php
// output csv from file
if (file_exists($csvFile)) {
    header('Content-type:text/csv');
    header('Content-disposition:attachment;filename='.$csvFile); // 需要注意文件名是否会乱码, 最简单的办法就是 「英文 + 日期」
    echo file_get_contents($csvFile);
}
```

- 直接将数据格式化为 csv, 然后导出

```php
// output csv directly
$arr = [['header1', 'header2'], ['body1', 'body2']];
$content = '';
foreach ($arr as $v) {
    $content .= join(',', $v) . PHP_EOL; // 换行符常量
}
```

- 对比一下, 输出图片

```php
// output img
if (file_exists($imgFile)) {
    header("Content-Type:image/jpg");
    echo file_get_contents($imgFile);
}
```

**当然还可以加上其他 http header 来调优**

## 站在巨人的肩膀上

这个句子被用得有点多, 显得有点「俗」了, 但「巨人」才足以表达我对开源贡献者的敬意.

phper 可以安装 package, 轻松实现文件导出: `composer require league/csv`

```php
require_once 'vendor/autoload.php';

use League\Csv\Writer;

$csv = Writer::createFromFileObject(new \SplTempFileObject()); // SPL, Standard PHP Library, 建议 phper 了解一下
$csv->insertOne([1, 2, 3]);
$csv->insertAll([['a1', 'a2', 'a3'], ['a1', 'a2', 'a3']]);
$csv->output('xxx' . date('Ymd') . '.csv');
die;
```

好了, 基本使用已经有了, 再来说说印象比较深刻的 2 个例子

## 案例: 礼包码

需求: 8 位, 数字 + 小写字母, 其实不区分大小写, 方便用户输入; 可能一次超过 10w 个

```php
$str = '0123456789abcdefghijklmnopqrstuvwxyz';
$code = '';
for ($i=0; $i < 8; $i++) {
    $code .= $str[mt_rand(0, 35)];
}
```

实现很简单, 列举几个注意点:

- 如果生成比较多, 使用异步队列处理或者 task(命令行脚本)
- `mt_rand()` 来生成随机数, 比 `rand()` 性能高 4 倍, `7.1.0 rand() has been made an alias of mt_rand()`
- 分批生成数据, 分批插入数据库, 比如 5k条/次, 否则可能 `爆内存` 或者 `max_execution_time`
- 关于使用随机数是否会重复, 我做过测试, 生成 100w 条, 并没有重复, 大家也可以写个 join 查询试试

## 案例: 快递单

因为发货量猛增, 2k+, 从开始的运营自己打包, 到请兼职(高峰期 10+), 然后接入电子单, 从后台管理系统导出快递信息, 导入到快递平台, 再从快递平台拿到快递单, 回填到后台管理系统. 回填这一步之前运营是一个一个手动填的.

```php
// 订单导出
$items = [];
$items[] = ['订单编号', '收件人', '手机', '地址', '发货信息', '备注', 'created_at'];
foreach ($rs as $k => $v) {
    $it['id'] = rand(1,9). sprintf('%011d', rand(0, 99999999999)). $v->id . "\t"; // 拼上固定 12 位防止订单重复, 这里是业务细节, 不用太关注
    $it['name'] = $v->address->recipent_name;
    $it['mobile'] = $v->address->recipent_mobile . "\t"; // 拼上特殊字符转为字符串, 方便 excel 打开查看
    $it['address'] = $v->address->recipent_address;
    $it['extra_info'] = $extra_info->ext_info ?? ""; // ?? 操作符, 这样就可以不用 isset() 了
    $it['products'] = $v->products . '|' . $it['extra_info']; // | 来作为分隔符
    $it['created_at'] = $v->created_at;
    foreach ($it as &$itv) {
        $itv = iconv("UTF-8", "GB2312//IGNORE", $itv); // 转码, //IGNORE 转码失败时不会终止程序
    }
    $items[] = $it;
}

$csv = Writer::createFromString('');
$csv->insertAll($items);

$csv->output($ags['dt'] . ".csv");
die;

// 读取 csv 文件, 自动匹配内部订单号和快递单号
$f = fopen($file, 'r');
while ($arr = fgetcsv($f)) { // 还有一个函数是 fputcsv()
    if (!$arr[0]) break; // csv 没有数据了
    $product_order_id = $arr[1];
    $product_order_id = substr($product_order_id, 12); // 除去拼上的随机数
    $express_no = explode(',', $arr[14]); // 订单号可能重复, 取最后一个
    $express_no = end($express_no);

    echo "order:$product_order_id, express: $express_no\n";
    Logger::info("doShipping | order:$product_order_id, express: $express_no");
    $this->doShipping($product_order_id, $express_no);
}
```
列举几个重要点:
- `iconv` 比 `mb_convert_encoding` 速度快很多
- 还是推荐 utf-8 编码, 之前在游戏公司遇到过, 少数名族姓名用 GBK 可能无法显示

## 其他值得 mark 的点

### 礼包码姊妹篇: 邀请码

先看一个案例, 某外包实现的「邀请码」:

```js
while (true) {
  const code = generateCode() // 只有 5 位数字, 随机生成
  const usersWithThisCode = await knex('user_meta') // 查询数据库, 看是否已经生成
    .where('name', INVITE_CODE_KEY)
    .where('string_1', code)
    .count('* as c')
  if (usersWithThisCode[0].c > 0) { // 已有则继续, 没有则更新数据
    continue
  }
  await knex('user_meta')
    .where('name', INVITE_CODE_KEY)
    .where('user_id', user.get('id'))
    .del()
  await knex('user_meta').insert({
    user_id: user.get('id'),
    name: INVITE_CODE_KEY,
    string_1: code,
  })
  return code
}
```

这样的结果是, 我们高配版的腾讯云数据库实例(25w IO, 6k 连接数), 直接被跑满.

当 mysql 出现性能问题时, 一个简单有效的办法:

```sql
SHOW PROCESSLIST -- mysql 上当前连接的进程
SELECT * FROM information_schema.`PROCESSLIST` -- information_schema 非常有用

kill id -- 通过上面查询出来的 id 杀掉进程
```

有 2 种思路来做邀请码:

- 使用用户 id 来保证去重, 然后拼接到固定位数, 比如 `sprintf('%08d', $uid)`
- 类似礼包码的, 预先生成, 需要时取一个即可

### 数据导出

数据有时候并不是那么简单获取到的, 需要进行复杂的 sql 查询, 这时候有以下几点需要注意:

- 是否需要使用「中间表」(或者「临时表」), 通过分步来简化查询, 如果是统计的话, 经常会用定时脚步来跑「大表」
- 加时间或者分页来控制数据量, 小心 `爆内存` 或者 `max_execution_time`
- 数据量大到一定程度, 可能就需要其他方案了, 比如 mysql 的 `select ... into outfile` / `mysqldump`

### 关于 phpExcel

csv 是如此的简单好用, 除非 **刚需**, 否则我基本不会用到 `phpExcel`, 推荐 **phpExcel 从入门到放弃**
