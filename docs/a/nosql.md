# nosql

- nosql: 70`no sql` 80`know sql` 00`no sql!` 05`not only sql` 13`no, sql!`
- tool: datagrip

## mongo

- https://docs.mongodb.com/manual/core/document
- chunk块.最小64M -> ReplicationSet复制集/sharding分片->cluster 数据均衡 master=primary slave=secondary journaling日志
- db
- collection集合: metadata
- document文档: bson relation=include包含+point引用
  - field index cursor
  - aggregate聚合
- gridfs: 存储16M+文件->分布式网盘

```js
// Mongo的shell运行在JavaScript之上
// 指令: help exit
db.help()
show dbs/collections
use db_xxx
db.dropDatabase()
db.createCollection('users') db.users.drop()

$set $unset $inc UpSert upAll
db.users.insert()/save()/update()/remove()

$or $gt $exists $in
db.unicorns.find({_id: ObjectId("TheObjectId")})
db.users.find().pretty().sort().limit().skip().count()
db.unicorns.find({str: /far/)}) // regex
db.unicorns.ensureIndex({name: 1})/dropIndex({name: 1})/getIndexs()
db.crcsCashHitRule.distinct('rule_set_name')
db.cashEventReport.group({ key:{event_id:1, utm_source:1}, cond:{event_id:"1000109"}, reduce: function(curr, res){ res.cnt++; }, initial: {cnt:0} })
db.users.aggregate([{$group:{_id:"$name", user:{$sum:"$user_id"}}}])
db.users.aggregate([{$match:{user_id:{$gt:0,$lte:2}}},{$group:{_id:"user",count:{$sum:1}}}]) // aggreagte with pipe
db.users.find({gender:"M"},{user_name:1,_id:0}).hint({gender:1,user_name:1}).explain() // hint() force to use the input as index
db.getLastError() // safeMode
// geo
var map = [{"gis":["x":185,"y":150]},{"gis":["x":70,"y":180}]
db.map.ensureIndex({"gis":"2d"},{min:-1,max:201}) // 默认会建立一个[-180,180]之间的2d索引
db.map.find({"gis":{$near:[70,180]}},{gis:1,_id:0}).limit(3) // 查询点(70,180)最近的3个点
db.map.find({gis:{$within:{$box:[[50,50],[190,190]]}}},{_id:0,gis:1}) // 查询以点(50,50)和点(190,190)为对角线的正方形的所有的点
db.map.find({gis:{$within:{$center:[[56,80],50]}}},{_id:0,gis:1}) // 查询出以圆心(56,80),半径为50的圆中的点
// MapReduce
// sharding copy
db.stats() db.unicorns.stats()
db.setProfilingLevel(2)
db.system.profile.find()
db.system.profile.find().sort({$natrual:-1}).limit(3)

// admin cpu使用率高 https://help.aliyun.com/document_detail/62224.html
db.currentOp({"waitingForLock": true, $or: [{"op": {"$in": ["insert", "update", "remove"]}}, {"query.findandmodify": {$exists: true }}]})
db.killOp(opid)
```

```sh
mongodump --db learn --out backup
mongorestore --collection unicorns backup/learn/unicorns.bson
mongoexport --db learn -collection unicorns # json
mongoexport --db learn -collection unicorns --csv -fields name,weight,vampires
```

## hdfs

## Cassandra
