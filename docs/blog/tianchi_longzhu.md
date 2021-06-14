# ai| 2021阿里云天池龙珠计划 AI训练营

- python sql docker 机器学习(分类预测) 深度学习(人脸识别) 强化学习(PPO算法)
- [数据分析-美国大选](https://tianchi.aliyun.com/notebook-ai/home#notebookLabId=149877&notebookType=ALL&isHelp=false&operaType=5)

## mysql

- [Task06：综合练习题-10道经典题目-天池龙珠计划SQL训练营-天池技术圈-天池技术讨论区](https://tianchi.aliyun.com/forum/postDetail?postId=170359)
- [SQL训练题习题答案汇总 - SQL_Answers.pdf](chrome-extension://oikmahiipjniocckomdccmplodldodja/pdf-viewer/web/viewer.html?file=http%3A%2F%2Ftianchi-media.oss-cn-beijing.aliyuncs.com%2Fdragonball%2FSQL%2Fother%2FSQL_Answers.pdf)

```mysql
create table Addressbook (
    regist_no int auto_increment primary key comment '注册编号',
    name varchar(128) not null comment '姓名',
    address varchar(256) not null comment '住址',
    tel_no char(10) comment '电话号码',
    mall_address char(20) comment '邮箱地址'
) comment '地址簿';
alter table Addressbook add column postal_code char(8) not null comment '邮政编码';
drop table Addressbook;
-- 删除的表无法恢复
```
