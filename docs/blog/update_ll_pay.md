# coder| 升级连连支付

> 我一直持有一个观点：如果投入时间还没有把事情做好，一定要好好反思

## 1 背景

接到一个升级 *连连支付* 的任务，但是却拖了很长时间，事后想想，很多地方的处理都可以优化

## 2 单个支付方式概览

1. 支付（有用户交互） = 发起支付 + （同步 + 异步 + 查询）
2. 发起支付：构建自提交表单
3. 同步：从支付渠道回到商家页面（**最好**使用查询确认支付结果）
4. 异步：支付渠道异步通知商家支付家国
5. 查询：做好**轮询**策略

技术关键字：http + sign

### 2.1 http

本质：http 请求
关键：https、method（get、post）、params（type、length、format，care money）、sign、pay result、header、autosubmit form、response

### 2.2 签名

1. 签名方式：对称加密与非对称加密；md5 算法与 rsa 算法；PEM format 密钥格式
2. 构造签名：需要签名的参数、拼接签名串、加密、验签

## 3 接入流程

自己总结了一份流程，希望对之后接入支付有帮助：

1. 理解需求很重要：大部分来自 PM（这次要支持单笔 1w 以上），过一过测试 case，以及隐藏的需求，别把之前的功能整挂了
2. 对接群中多提问
3. 获取开发指南和 demo，并开发指南为准，并谨记，2者可能都有坑，要多问
4. 接通每个接口（发起、同步、异步、查询），完成自测
5. 对接测试环境和**上线**相关
6. 提交测试并继续跟进，可能会有隐藏（或者奇葩）的 bug

## 4 简单示例

### 4.1 以下代码基于 php 实现，其他语言同理

关于 rsa 密钥的格式，可以去搜索相关文章了解下，各语言相关函数可能会有限制，注意查看文档

```php
// openssl_get_privatekey() 等函数需要使用 PEM format 格式密钥
function l_pem_format($key, $type){
	$tmp = $type ? 'RSA PRIVATE':'PUBLIC';
	$str = "-----BEGIN $tmp KEY-----\n";
	$len = strlen($key);
	$i = 0;
	while ($i <= $len) {
		$str .= substr($key, $i, 64)."\n";
		$i += 64;
	}
	$str .= "-----END $tmp KEY-----\n";

	return $str;
}

// 默认 json encode 是会转中文等为 unicode(\u2345 这种类型)
// 这样会导致2个问题：有些请求不允许传入 \ 这样特殊字符；增加数据量
json_encode($str, JSON_UNESCAPED_UNICODE);

// 处理金额，是否为整数，是否需要保留 2 位小数
$money_order = rtrim(rtrim(sprintf('%.2f', $money_order  / 100), '0'), '.');

// name_goods 字段有长度限制
$i = 40>>1;
while (strlen($name_goods)>=40) {
    $name_goods = mb_substr($name_goods,0,$i,'utf-8');
    $i >>=1;
}

// 处理特殊字符
str_replace('\\', '', $str); // 可以传入数组，替换多个字符

// 构建自提交表单，需要注意：是否制定 header、input 带 hidden
$sHtml = '<html><head><meta http-equiv="Content-Type" content="text/html; charset=utf-8"></head><body>';
$sHtml .= "<form id='llpaysubmit' name='llpaysubmit' action='" . self::llpay_gateway_new . "' method='POST'>";
foreach ($para_sort as $k => $v) {
    $sHtml .= "<input type='hidden' name='{$k}' value='{$v}'/>";
}
$sHtml .= "<script>document.forms['llpaysubmit'].submit();</script>";
$sHtml .= '</body></html>';
echo $sHtml;

// 获取接口数据
$data = file_get_contents("php://input"); // 返回原生数据，支持 get/post
$data = $_POST; // 获取 post 数据，会自动解析到数组中，推荐使用

// 简单打日志，建议封装日志类，封装上时间戳
file_put_contents('log.txt','noti-'.$data."\n\n", FILE_APPEND);

// 使用 php header() 实现重定向
header('location: '. $url);die;
```

### 4.2 mark 一个以前接入米大师（应用宝）的例子：

说明：

1. 此处代码为 Job 类方法并实现异步队列接口，Controller 接收请求数据，传入并调用此方法
2. 此方法包含 测试、正式环境切换；使用 https 并添加 cookie；轮询；日志

```
// 指定 cookie
$cookie_get = 'session_id=openid;session_type=kp_actoken;org_loc=%25mpay%25get_balance_m'; // 对方文档错误，双方沟通才确定
// 查询余额接口
$url = 'https://ysdk.qq.com/mpay/get_balance_m'; // 现网
$secret = '';
// $url = 'https://ysdktest.qq.com/mpay/get_balance_m'; // 沙箱
// $secret = '';
$arr = $input; // 此处框架层处理输入数据
unset($arr['billno']);
$arr['sig'] = QqSnsSigCheck::makeSig('get', '/v3/r/mpay/get_balance_m', $arr, $secret); // 此处使用了 demo 的方法，但是 demo 签名写错了，最后读文档发现的
$url = $url. '?'. http_build_query($arr);

$opts = [
    'ssl' => ["verify_peer"=>false, "verify_peer_name"=>false],
    'http' => ['header' => "Cookie: {$cookie_get}\r\n"]
];
$context = stream_context_create($opts);
\Log::info('sdk qq-query: '. $url);

// 2分钟之内间隔15秒多次调用，直到查到当前充值已到账
$time = 0;
while ($time<120) {
    $data = file_get_contents($url, false, $context);
    $data = json_decode($data, true);
    // 查询到游戏币
    if($data['ret']==0){
        \Log::info('sdk qq-query: ', $data);
        if($data['balance']>0) break;
    }
    \Log::info('sdk qq-unpaid: ', $data);
    DB::table('orders')->where('id', $order_id)->update(['status'=>7, 'updated_at'=>date('Y-m-d H:i:s')]);
    sleep(15);
    $time +=15;
}
if($time>=120){
    \Log::info('sdk qq-query-timeout: ', $data);
    DB::table('orders')->where('id', $order_id)->update(['status'=>7, 'updated_at'=>date('Y-m-d H:i:s')]);
    return;
}

// 后面是扣款接口，和上面类似
```

## 5 入坑指南

1. 产品说连连那边的人说只要改一下 url 和换一下签名商户号等配置就 ok 了，最重要的是，我也信了
2. 当发现上面的路走不通，就想使用 demo，然后采用原始的方法：对比文件差异，覆盖并修改配置
3. 做了上面 2 个无用功（唯一的用途可能就是多阅读了一下 demo 的代码），遇到问题就到讨论组里面问
4. 第三步涉及到 2 个公司之间沟通，沟通成本高，一个问题可能很久才能回复，甚至跨天，最后终于定位到：demo 有问题
5. 忍无可忍之下根据文档重写 **发起支付** 接口 并顺利调试通，然后被告知要走上线流程
6. 按照对方要求填好文档之后，被告知 **同步调用**、**异步通知** 接口不通
7. 继续阅读文档重写这 2 个接口，同步走上线流程
8. 加急下成功上线，但是导致了连连没有走正常上线流程：只验证支付成功，没有验证**同步调用**、**异步通知** 接口
9. 测试发现 **同步调用**、**异步通知** 报错，只能推迟上线时间
10. 战线太长（前后 2 个星期，并采用直接重写的方式）以及双开（同时接连连 wap 和 web），忽略了**同步调用**、**异步通知** 2 个接口以及 2 个不同版本的差异，导致编码出错
11. 加班按照上面的 **接入流程** 进行二次 coding，终于成功
12. 测试继续测试发现隐藏 bug，同步修复

总结：

1. 流程很重要，流程是为了减少犯错的几率，开发最容易犯的流程错误就是开发完认为万事大吉，不自测就直接提交给测试
2. 再次重申：流程很重要。可以看出上面做了很多无用功，最开始以为改改配置就好，不行就用 demo 替换，结果联调下来发现 demo 问题很多，最后根据文档重写，如果一开始按照 **接入流程** 来做，完全可以顺利提测
3. 后台开发要用好 **日志**，我这次犯的错误是不熟悉框架的日志方法而不打日志
4. 还是尽量减少双开这样的情况，我犯的错误是以为 wap 版和 web 版基本类似，同步和异步接口类似，导致很多改变的地方没有调整，最后出错
5. 其他小坑：沟通之后才知道要走上线流程；走上线流程才知道风控参数很变态；测试发现有一个参数有长度限制，以前没有处理；测试使用的密钥是 string 而不是 PEM format，定位了一圈才知道 *还有 PEM format 这东东*

最后，提供一些参考：

- php 中 `$_POST，$HTTP_RAW_POST_DATA 和 php://input` 的区别：http://blog.wpjam.com/m/post-http_raw_post_data-php-input/。注意`$_POST`，和 http post 提交的数据格式有关
