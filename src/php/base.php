<?php

use Hyperf\Utils\Str;

require_once __DIR__ . '/vendor/autoload.php';

// php
declare(ticks=1); // ticks
register_tick_function();
declare(encoding='ISO-8859-1'); // encoding
// 内存溢出
ini_set('memory_limit', '128m'); // 调整内存使用
memory_get_usage(); // 获取当前使用
memory_get_peak_usage(); // 获取内存使用峰值

// var
isset($a); // empty()
var_dump(); // die() print_r() var_export()
defined('YII_DEBUG') or define('YII_DEBUG', true); // 定义常量常见用法

// text/string
sprintf('%08d', 123); // fix length
preg_match();
preg_replace_callback();
iconv("UTF-8", "ISO-8859-1//IGNORE", '中国');
mb_substr();
mb_convert_encoding('中国', 'gbk');
echo Str::snake('They think they are very powerful');

// array
array_merge(); // array_merge_recursive();
array_chunk();
sort();

// json_encode JsonSerializable
json_encode([]);
json_encode((object)[]);
json_encode(new stdClass());
json_encode('中国', JSON_UNESCAPED_UNICODE); // \r 换行

// date
// https://www.php.net/manual/zh/datetime.formats.relative.php
// https://www.php.net/manual/zh/dateinterval.construct.php
date('Y-m-d H:i:s', time());
strtotime('+1 day', time());
$d = new DateTime('2021-06-07');
$d->add(new DateInterval('P1D'));
echo $d->format('Y-m-d');

// fs
mime_content_type();
flock(fopen('t.txt'), LOCK_EX | LOCK_NB);
function readdir_r($path) {
    $fd = opendir($path);
    while (($f = readdir($fd)) !== false) {
        if ($f == '.' || $f == '..') continue;
        $t = $path . '/' . $f;
        if (is_dir($t)) { // dir
            readdir_r($t);
            continue; // be care
        }
        // file
    }
}
function scandir_r($path) {
    $a = scandir($path);
    foreach ($a as $v) {
        if ($v == '.' || $v == '..') continue;
        $t = "$path/$v";
        if (is_dir($t)) { // dir
            scandir_r($t);
            continue; // be care
        }
        // file
    }
}
function scan_dir($dir)
{
    $a = scandir($dir);
    foreach ($a as $v) {
        if ($v=='.' || $v=='..') continue;
        $f = $dir.'/'.$v;
        if (is_dir($f)) scan_dir($f);
        // do file
    }
}
function composer_dep()
{
    $links = []; // ecCharts json data
    $node = [];
    $hc = [];

    $d = '/Users/dayday/hub/hyperf/hyperf/src';
    $r = scandir($d);
    foreach ($r as $v) {
        if ($v=='.'||$v=='..') continue;
        $json = file_get_contents($d. '/'. $v.'/composer.json');
        $deps = json_decode($json, true)['require'] ?? [];
        foreach ($deps as $k=>$ver) {
            $node[$k] = 1;
            $links[] = [
                'source' => 'hyperf/'.$v, // hyperf component
                'target' => $k, // dep of component
            ];
            $hc[] = ['hyperf/'.$v, $k];
        }
    }

    $n = [];
    foreach ($node as $k=>$v) {
        $n[] = ['id'=>$k];
    }
    echo json_encode($hc), "\n";
// echo json_encode($links);
}

// web
header('Content-type:text/csv'); // csv download
header('Content-disposition:attachment;filename='.date('YmdHis').'.csv'); // filename: ASCII, 避免乱码
echo file_get_contents('t.csv');
header("Content-Type:image/jpg"); // output image
echo file_get_contents('t.jpg');
http_build_query();
parse_url();
pathinfo();

// net
pack();
unpack('H6', '中'); // utf8

// crypto
hash('md5', 'xxx'); // md5('xxx)
$pw_hash = password_hash('xxx', PASSWORD_DEFAULT); // password_verify('xxx', $pw_hash)
// openssl
// key format: RFC4716(ssh-keygen default key format)
$pri_key = openssl_get_privatekey('pem_format_pri_key'); // ssh-keygen -e -f xxx -m pem
// key format: pkcs12
openssl_pkcs12_read(file_get_contents('xxx-pri.pfx'), $pkfArr, 'pri_key_passwd'); // pri: $pkfArr['pkey']
openssl_pkey_get_details(openssl_get_publickey(file_get_contents('xxx-pub.cer')))['key']; // pub
$sign_str = openssl_sign('xxx', $pri_key); // openssl_verify('xxx', $sign_str, $pub_key);
openssl_get_cipher_methods();
openssl_encrypt('$str', 'DES-ECB', '$enKey'); // openssl_decrypt($str, 'DES-ECB', $enKey)

// sys
shell_exec();

// backtrace
$source = debug_backtrace(DEBUG_BACKTRACE_IGNORE_ARGS, 3)[2]['class'];
preg_match('#\w+#', $source, $arr);
die($arr[0]);

// db
// 进程长时间运行时的长连接
if ($this->_linkr && mysqli_ping($this->_linkr)) {
    $this->_link = $this->_linkr;
    return true;
}
$this->_linkr = $this->_connect($host);
// pdo_ping
function pdo_ping($dbconn){
    try{
        $dbconn->getAttribute(PDO::ATTR_SERVER_INFO);
    } catch (PDOException $e) {
        if(strpos($e->getMessage(), 'MySQL server has gone away')!==false){
            return false;
        }
    }
    return true;
}
// 大delete -> 小delete
while (1) {
    mysqli_query("delete from logs where log_date<='2020-01-01' limit 1000");
    if (!mysqli_affected_rows()) {
        break;
    }
    usleep(50000);
}
