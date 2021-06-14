# coder| lbs LocationBasedServices 基于位置的服务

- 高德
  - web=js: 地图 数据可视化 地图组件 uri 室内 地铁图
    - [高德地图 JSAPI 2.0 (GL+) 正式版发布 | 高德地图API](https://lbs.amap.com/news/jsapi-v2/)
      - ![zooms[2,20] 最大级别放开至20级，满足街区级、室内级精细数据展示](https://a.amap.com/jsapi/static/doc/images/world.png)
  - Android/ios: 地图 轻量版地图 定位 导航 猎鹰=轨迹 室内 室内定位 商家入驻
  - web服务: api 猎鹰
    - api: ip定位 地理编码-逆地理编码(addr-geo) 输入提示(keyword->addr) 关键词搜索(keyword+poi) 周边(geo+keyword) 多边形 行政区域查询
      - [相关下载-Web服务 API | 高德地图API](https://lbs.amap.com/api/webservice/download/): POI分类编码 城市编码表
  - other: wxapp 高德app=手机+车机 flutter
    - [概述-Flutter插件 | 高德地图API](https://lbs.amap.com/api/flutter/summary) 定位基础能力 地图基础能力

- [全国行政区划信息查询平台](http://xzqh.mca.gov.cn/map)
- Coordinate坐标 geo: gps mapbar baidu
- poi PointOfInterest 兴趣点: 医疗保健服务-综合医院-三级甲等医院 政府机构与社会团体 交通设施服务-机场相关/火车站 地点地址信息-普通/自然/交通/门牌 事件活动-突发事件-公共卫生事件
- adcode城市编码表

使用MongoDB存储地理位置信息
MongoDB原生支持地理位置索引，可以直接用于位置距离计算和查询。
另外，它也是如今最流行的NoSQL数据库之一，除了能够很好地支持地理位置计算之外，还拥有诸如面向集合存储、模式自由、高性能、支持复杂查询、支持完全索引等等特性。
对于我们的需求，在MongoDB只需一个命令即可得到所需要的结果：
db.runCommand( { geoNear: "places", near: [ 121.4905, 31.2646 ], num:100 })
查询结果默认将会由近到远排序，而且查询结果也包含目标点对象、距离目标点的距离等信息。
由于geoNear是MongoDB原生支持的查询函数，所以性能上也做到了高度的优化，完全可以应付生产环境的压力

MongoDB地理位置索引常用的有两种

- 2d 平面坐标索引，适用于基于平面的坐标计算。也支持球面距离计算，不过官方推荐使用2dsphere索引
- 2dsphere 几何球体索引，适用于球面几何运算

```js
// MongoDB地理位置索引
db.places.ensureIndex({'coordinate':'2d'})
db.places.ensureIndex({'coordinate':'2dsphere'})
```

查询方式分三种情况：

- Inclusion。范围查询，如百度地图“视野内搜索”。
- Inetersection。交集查询。不常用。
- Proximity。周边查询，如“附近500内的餐厅”。

```js
// 查询坐标对(经纬度), 单位: 弧度/度
db.<collection>.find({
  <location field>: {
    $nearSphere: [<x>, <y>],
    $maxDistance: <distance in radians>
}});

// 查询 GeoJson 距离, 单位: 米
db.<collection>.find({
  <location field>: {
    $nearSphere: {
      $geometry: {
        type: "Point", coordinates: [<longitude>, <latitude>]},
        $maxDistance: <distance in meters>
}}});
```

使用场景:

- 附近的人: 默认由近到远排列
- 区域内搜索(inclusion): $box 矩形; $center 圆形; $centerSphere 球面; $polygon 多边形
- 附近POI: 增加 POI 信息, 比如机场/火车站/三甲医院等, 可以获得更精准的搜索结果

```js
// 附近的人
db.places.find({'coordinate':{$near: [121.4905, 31.2646], $maxDistance:2}}).limit(2)
db.places.find({'coordinate':{$nearSphere: [121.4905, 31.2646]}})

// 区域内搜索
// 矩形
db.places.find({
  coordinate: {
    $geoWithin: {
      $box: [[121.44, 31.25], [121.5005, 31.2846]]
}}});
// 圆形
db.places.find({'coordinate':{$geoWithin:{$centerSphere:[[121.4905, 31.2646] , 0.6/111] }}})
// 多边形
db.places.find({coordinate : {$geoWithin : {$polygon :[
  [121.45183 , 31.243816],
  [121.533181, 31.24344],
  [121.535049, 31.208983],
  [121.448955, 31.214913],
  [121.440619, 31.228748]
]}}})

// 附近POI, 比如三甲医院
db.runCommand({ geoNear: "三甲医院", near: [121.4905, 31.2646], spherical: true, maxDistance:1/6371, num:2 });
```
