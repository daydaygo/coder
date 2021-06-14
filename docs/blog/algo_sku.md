# algo| 业务: 电商 sku & spu

业务上碰到的关于电商系统中sku与spu的一个难题

> [algo| 关于电商系统中sku与spu的一个难题](https://www.jianshu.com/p/edc5b28e2842)

最近业务上碰到的一个难题, 分享出来给大家瞧瞧.

PS: 本文使用 php 演示, 实现思路不局限于语言

## 需求

在电商领域, 有 sku 和 spu 这 2 个概念:

- sku 就是表示一个商品, 用户最后加到购物车中的那个
- spu 是一组商品, 比如一件衣服有 `L XL XXL` 等不同型号, 那么每个数据就是一个 sku, 这组 sku 的集合就是 spu

假如我们有一个sku 为 4483094 的商品, 获取到的商品 spu 信息信息如下:

![spu信息](https://upload-images.jianshu.io/upload_images/567399-2e01d197174d8509.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

这里提供了 json 版的数据, 方便感兴趣的朋友自己尝试

```json
{"颜色":{"曜石黑":[4483094],"草木绿":[3893503],"钻雕蓝":[3893493],"钻雕金":[4736669,4483110,4483074],"陶瓷白":[4483100]},"版本":{"128GB":[4483100, 4483074],"64GB":[3893493,4736669,3893503,4483094,4483110]},"普通版":{"普通版":[3893493,4483100,3893503,4483094,4483110,4483074],"移动4G+版":[4736669]}}
```

spu 信息的数据层级很清晰:

- 第一层表示的规格
- 第二层表示的规格包含的属性
- 第三层就是这个属性下有哪些 sku

接着, 来看看我们的需求:

1. 获取 sku 有哪些属性
2. 用户选择了规格属性时, 判断其他规格的属性是否可选
3. 隐藏需求: 数据怎么存以及 CRUD

## 需求1: sku 有哪些属性呢

一次简单的遍历即可

```php
// $spuArr: spu信息格式化后的数组 $specArr['颜色']['曜石黑']
// $spuSet = []; // spu下有多少 sku
$skuSpec = []; // sku的规格属性信息
foreach ($spuArr as $spec => $specArr) {
    foreach ($specArr as $attr => $skuArr) {
        // $spuSet = $spuSet + $skuArr; // 数组并集
        foreach ($skuArr as $sku) {
            $skuSpec[$sku][] = $attr;
        }
    }
}
// spu下有多少sku
// var_dump($spuSet);
// var_dump(array_keys($skuSpec));
// sku的属性信息
var_dump($skuSpec);
// 指定 sku 的属性信息
var_dump(join(' ', $skuSpec[4483094]));
```

最后获取到 sku 为 4483094 的商品的属性信息: `曜石黑 64GB 普通版`

PS: 我之前在求 `spu 下有哪些 sku` 犯了一个失误, **通过遍历整个 $spuArr, 当做集合问题来求解**(见上面注释掉的代码), 但是请仔细观察数据 -- **每一层规格都包含了此 spu 下的所有 sku**

> 数据有意思~ 多观察, 多体会~

## 2. 用户选择了规格属性时, 判断其他规格的属性是否可选

其实是很简单的交集就可以解决

```php
// 初始状态, 用户还没有选择规格属性, 所有规格属性都可选
// 用户选择了规格属性时, 判断其他规格的属性是否可选
$attrSelect = ['颜色' => '曜石黑']; // 用户选择了的规格属性
$attrSku = []; // 拥有这些属性的 sku
foreach ($attrSelect as $spec => $attr) {
    if (!$attrSku) {
        $attrSku = $spuArr[$spec][$attr];
    } else {
        $attrSku = array_intersect($attrSku, $spuArr[$spec][$attr]); // 交集
    }
}
// var_dump($attrSku);
// 判断其他规格的属性是否可选
foreach ($spuArr as $spec => $specArr) {
    if (!isset($attrSelect[$spec])) { // 此规格下用户没有选择属性, 判断次规格下的属性是否可选
        foreach ($specArr as $attr => $skuArr) {
            $spuArr[$spec][$attr] = array_intersect($attrSku, $skuArr); // 交集, 如果为空, 则此属性不可选
        }
    }
}
var_dump($spuArr);
```

## 3. 隐藏需求: 数据怎么存呢? 要考虑到 CRUD

需求1 的存储比较简单, 可以给 sku表添加一个字段 `attr`, 用来存商品的所有属性

需求2 的存储怎么设计呢? 直观的感受是, 我们会有这些表:

- spu表, 自增ID
- sku表: 自增ID + spu_id(1 : n) + attr(需求1 中格式化后的属性信息)
- spec表(规格表): 自增ID
- spec_attr表(规格属性表): 自增ID + spec_id
- attr_sku表(属性与sku关系对应表): 自增ID + sku_id + spec_attr_id, 其实是上面第二层数据和第三层数据的映射

5个ID, 你还 hold 住么?

- 解决需求2的问题时, 我们需要从 5 张表里取出数据并格式化成 `$spuArr`
- 新增 spu 数据的时候, 我们需要更新 5 张表
- 如果 spu 数据有变化, 比如某个属性下的商品没货了, 嗯, 又是 5 张表

> 就问你怕不怕?

通过格式化的 `$spuArr` 数据, 我们已经可以解决我们的需求, 也就是说, 无论我们的数据怎么存, 最后我们都会格式化出 `$spuArr` 数据来解决问题, 为什么不可以直接这样存储呢:

- sku表: 自增ID + spu_id + attr(需求1 中格式化后的属性信息), spu_id 只用来标识一组 sku, 可以用 id 生成器生成
- spu_info表: spu_id + `json_encode($spuArr)`

**只用 2 张表**, 就完成了我们的需求

`json_encode($spuArr)` 这样的文档型数据, 可以考虑使用 **mongoDB**

## 写在最后

- **商品搜索系统** 可以录入 sku表的 attr, 实现搜索商品属性的功能
- 原来的系统没有 spu 的概念, 只有 goods 表(商品), 流程依靠 goods 表都走通了, 怎么进行 sku 和 spu 的拆分呢? -- **其实 sku 表就相当于系统现有的 goods 表, 按照需求3 的方式录入数据即可**
- **大道至易**: 不要过度设计 不要过度设计 不要过度设计
- **算法 - 并查集**: 有多个不相交的集合, 怎么快速找到一个数属于哪个集合呢?