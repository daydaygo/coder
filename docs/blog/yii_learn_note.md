# yii| 学习笔记

命名空间: 除了 composer 的 autoload, yii 自己还有一个 autoload

```php
// https://www.yiichina.com/doc/guide/2.0/concept-autoloading
Yii::setAlias('@console', dirname(dirname(__DIR__)) . '/console');
'controllerNamespace' => 'api\controllers',

$className = ltrim($this->controllerNamespace . '\\' . str_replace('/', '\\', $prefix)  . $className, '\\');
$classFile = Yii::getAlias('@' . str_replace('\\', '/', $className) . '.php');
```

## 分页 page

文档中给出的分页示例:

```php
public function actionIndex()
{
    $query = Country::find();

    $pagination = new Pagination([
        'defaultPageSize' => 5,
        'totalCount' => $query->count(),
    ]);

    $countries = $query->orderBy('name')
        ->offset($pagination->offset)
        ->limit($pagination->limit)
        ->all();

    return $this->render('index', [
        'countries' => $countries,
        'pagination' => $pagination,
    ]);
}
```

yii中的一些 **部件** 之所以好用, 在于和页面中的 widget 联动, 不过不写管理后台, 或者前后端分离, 或者 api 项目, 就没啥感觉了.

这里 `Pagination` 的关键在于:

- 接收分页参数, 然后转化成 `limit + offset`
- 通过 `count() + limit` 显示分页数

这里可以优化的地方:

- 优化查询: `limit x offset y` 为 `where id>z limit x`. 拓展知识, **limit offset 为什么慢**.
- 减少查询: `count()` 这次查询使用固定值代替. 拓展知识 **innoDB 下 count 优化**.

## 脚手架 gii

使用 gii 来生成代码很有必要:

- model 类对应着数据表的字段, 使用 `@property` 注释来方便 IDE 自动提示
- 推荐使用 `php yii` 命令行来使用 gii

gii 配置:

```php
'bootstrap' => ['gii'],
'modules' => [
    'gii' => [
        'class' => \yii\gii\Module::class,
        'allowedIPs' => ['127.0.0.1', '::1', '192.168.0.*', '192.168.178.20'], // 默认只允许本机访问, 推荐使用命令行
    ]
],
```

常用的几个:

```
// model
php yii gii/model --ns=app\\models --tableName=sms_ori_batch --modelClass=SmsOriBatch --db=db_platform
```

## 调试 debug

常用的一些选项

## 日志 log

使用阿里云日志服务
设置 prefix

跨域: https://www.yiichina.com/doc/guide/2.0/structure-filters#cors