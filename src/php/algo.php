<?php
// arr
$arr = [1, 2, 3];
$arr[0]; // 查
end($arr);
$arr[3] = 4; // 增/改
$arr[4] = isset($arr[4]) ? $arr + 1 : 1; // $arr[4]++ 会报错
foreach ($arr as $k => $v) { // 遍历
    // code...
}
for ($i = 0; $i < count($arr); $i++) {
    // code
}
array_map();
array_slice($arr, 1, 0, [4]); // 插入
array_slice($arr, 1, 1); // 删除
unset($arr[1]);
array_merge(); // 并集
$arr + $arr;
array_intersect(); // 交集
array_diff(); // 差集
array_filter(); // 过滤
array_unique(); // 去重

// stack
// array_push($arr, 1); array_pop($arr);
// queue
// array_push($arr, $val); array_unshift($arr);

// str
$s = "far";
$s[-1]; // 查
for ($i = 0; $i < strlen($s); $i++) $s[$i]; // 遍历
str_repeat($s, 1); // 改
$s[1] = 'b';
strripos(); // i 大小写敏感; r 最后一次出现
strrpos();
strripos();
strrpos();

// SPL 数据结构: https://www.php.net/manual/zh/spl.datastructures.php
$stack = new SplStack();
$stack->pop();
$stack->push();
$stack->top();
$stack->isEmpty();

$queue = new SplQueue();
$queue->push();
$queue->pop();
$queue->top();