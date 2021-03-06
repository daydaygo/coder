# DIP, digital image process, 数字图像处理

数字图像处理（python） - 有你的文章 - 知乎
https://zhuanlan.zhihu.com/p/345533329

- 数据集
- 资料

我现在给你们发了一个数据集，包含131张飞机飞行图像。你们这段时间可以学习数字图像处理的知识，以这131张图像为样本，进行特征学习。（1）要识别出飞机在图像中的位置，并圈出来。（2）要有详细的处理过程。

数据处理 https://blog.csdn.net/sinat_36458870/article/details/78825571

- 开题ppt，plan 开题论文与老师项目
- mem sjtu paper
- 数字图像处理 ed3 英文版

- 从图像工程的角度认识图像处理
- 图像处理基础
- 典型图像变换理论
- 图像视觉质量提升
- 图像复原与超分辨率重建
- 图像压缩编码
- 彩色和多光谱图像处理
- 图像形态学处理
- 图像处理编程基础及应用实例

- 绪论
  - 图像处理
  - 图像基本概念
  - 图像处理分类
  - 数字图像处理
- 图像与视觉系统
  - 视觉过程
  - 光度学
  - 采样与量化
  - 图像类型
- 像素空间关系
  - 基本关系
  - 距离
  - 几何变换
  - 几何失真校验
- 空域变换增强
  - 算术运算
  - 逻辑运算
  - 直方图处理
  - 灰度变换
- 空域滤波增强
  - 卷积
  - 线性平滑
  - 非线性平滑
  - 线性锐化
  - 非线性锐化
- 图像变换
  - 一维离散
  - 二维离散
  - 傅里叶
  - 离散余弦
- 频域图像增强
  - 低通
  - 高通
  - 频域空域关系

- 数字图像处理与彩色数字图像处理
- 空间滤波与频域滤波
- 图像特征提取
- 图像压缩与图像小波变换

- 获取图像
- 人类视觉
- 打印与存储
- 修正成像缺陷
- 空间域图像增强
- 频率空间中的图像处理
- 分割与阈值处理
- 二值图像处理
- 全局图像测量
- 特定特征的测量
- 形状表征
- 特征识别与分类
- 层析成像
- 三维视图
- 表面成像

- 数字图像处理基础
• 概述
✓ 概念：图像、数字图像、像素
✓ 数字图像处理的起源
✓ 数字图像处理的应用领域
✓ 图像处理系统的部件
• 基础知识
✓ 图像的采样和量化
✓ 数字图像的表示
✓ 数字图像的质量
✓ 像素间的一些基本关系
- 图像增强
  - 处理方法
    - 空域: 点/模板
    - 频域
  - 处理策略: 全局/局部
  - 处理对象: 灰度/彩色
- 彩色图像处理
 彩色基础知识
 彩色空间
 伪彩色处理
 全彩色图像处理
 彩色变换
 彩色图像平滑和尖锐化
- 基于内容的图像检索 CBIR QBIC
 为什么需要基于内容的图像检索?
 查询方式，查询demo，现有系统简介
 具体内容
 特征提取
 相似度匹配
 相关反馈
 索引结构
 MPEG-7介绍：性能评价等
 思考的几个问题？
- 傅里叶变换
 傅里叶变换及其反变换 DFT
 傅里叶变换的性质
 快速傅里叶变换（FFT）
- 频率域图像增强
 频率域滤波
 频率域平滑（低通）滤波器
 频率域锐化（高通）滤波器
- 图像复原
 图像退化/复原过程的模型
 噪声模型
 空间域滤波复原（唯一退化是噪声）
 频率域滤波复原（削减周期噪声）
- 图像压缩
 基本概念
 图像压缩模型
 信息论基础
 无损压缩
 有损压缩
 图像压缩标准
 视频压缩标准
- 形态学图像处理
 概述
 集合论基础知识
 膨胀和腐蚀：产生滤波器作用
 开操作和闭操作：产生滤波器作用
 击中或击不中变换
 形态学的主要应用
- 图像分割
 概述
 间断检测
 边缘连接和边界检测
 阈值处理
 基于区域的分割
 分割中运动的应用
- 表示与描述
 概述
 表示方法
 边界描述子
 关系描述子

- 绪论
- 图像变换
- 图像增强与恢复技术
- 图像编码与压缩
- 图像分割
- 特征提取
- 图像目标检测与识别
- 影像匹配与镶嵌
- 图像分割新技术

## image 图像

- 表示: 二值.黑白 灰度.256(0.纯黑 255.纯白) 彩色(三基色.红绿蓝RGB 光学.波长/纯度/明度 视觉/心理学.色调/饱和度/亮度)
- ROI regionOfInterest 感兴趣区域

## mark

- 数字图像处理与python实现; 岳亚伟 2020.1
- 数字图像处理--原理与实现; 黄进 李剑波 2020.1 以matlab为实验平台
- opencv轻松入门: 面向python; 李立宗 2019.5
- 数字图像处理原理与实践; 秦志远 2017.12
- the image processing handbook 数字图像处理, ed6; Russ J.C. 2014.8 经典教材
  - 作者主页: <https://www.drjohnruss.com/>
- digital image processing, ed3; Rafael C. Gonzalez(冈萨雷斯) 2017.1 经典教材
  - [数字图像处理-电子科技大学-李庆嵘(冈萨雷斯版)](https://www.bilibili.com/video/av795664551)
  - [百度文库: 北京大学数字图像处理(冈萨雷斯)课件](https://wenku.baidu.com/view/11f7d420b0717fd5360cdc8e.html)
- [学堂在线: whu 图像处理分析](https://www.xuetangx.com/course/WHU08121000520/5882473)
- 学习opencv的建议: <https://www.zhihu.com/question/26881367>
