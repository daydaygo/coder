# ai

<https://www.paddlepaddle.org.cn/>
PaddleHub的课程地址：<https://aistudio.baidu.com/aistudio/course/introduce/1070>
PaddleHub的教程地址：<https://aistudio.baidu.com/aistudio/personalcenter/thirdview/79927>
PaddleHub的模型地址：<https://github.com/PaddlePaddle/PaddleHub/tree/release/v1.6/demo>

<https://ai.baidu.com/easydl>

## todo

- 百度飞浆 PPDE: <https://aikuaichedao.wjx.cn/m/81859172.aspx>
- 深度学习工程师认证- 百度AI Studio - 人工智能学习与实训社区 <https://aistudio.baidu.com/aistudio/certification?_origin=learnmap>
- 2020百度之星开发者大赛：交通标识检测与场景匹配 - 百度AI Studio - 人工智能学习与实训社区 <https://aistudio.baidu.com/aistudio/competition/detail/39>
- 百度AI Studio课程_学习成就梦想，AI遇见未来_AI课程 - 百度AI Studio - 人工智能学习与实训社区 <https://aistudio.baidu.com/aistudio/education/group/info/1335>
- 精品】AI从入门到精通，一张图带你玩转ModelArts_ModelArts_EI企业智能-华为云论坛 <https://bbs.huaweicloud.com/forum/thread-62576-1-1.html>
- 解决 时间/地点/人物/事件 的自然语言处理
- 百度顶会论文复现营 <https://live.bilibili.com/21689802> <https://aistudio.baidu.com/aistudio/education/group/info/1340>

- ai.baidu.com aistudio.com <https://ai.baidu.com/productlist> <https://aistudio.baidu.com/aistudio/learnmap> 百度paddle
- <https://tianchi.aliyun.com/course>
  - 52个优质数据集: <https://mp.weixin.qq.com/s/Ms1EfPUnzynI-5dn213kSg>
  - 保姆级NLP学习路线: <https://mp.weixin.qq.com/s/LwhWcnoWw13Cpm2-c7kexQ>
  - 达摩院 PAI XDL MaxCompute MNN
    - <https://vision.aliyun.com/>
- ai.qq.com open.youtu.qq.com ai.aliyun.com
- 华为云
- google
  - <https://github.com/google/trax>
- 机器学习面试题: https://www.interviewquery.com/blog-machine-learning-interview-questions/
- [ddbourgin/numpy-ml: Machine learning, in numpy](https://github.com/ddbourgin/numpy-ml)

## 应用场景

- NLP
- image
- generates shareable social images and thumbnails for your articles: <https://thumbnail.ai/>

## 思维导图

- 传统
  - 回归: 线性回归(客户终生价值 货币基金流入流出 电影票房 销量/股票) 岭回归
  - 分类: 逻辑回归(垃圾邮件/故障预测) Sigmoid函数(预测肿瘤良性/恶性) k-近邻 朴素贝叶斯 决策树 支持向量机SVM 随机森林
  - 聚类: 用户画像 k-means DBSCAN
  - 关联分析 FPGrowth
  - 降维 主成分分析PCA
  - 协同过滤: 推荐系统 apriori SVD
- ai -> 机器学习 -> 神经网络 -> 深度学习
  - 全连接nn 卷积nn 循环nn
  - 流程: 数据准备(采集/标注/预处理) 特征工程 模型设计(问题分类/参数/评估函数) 数据预测
  - 一般过程
    - 模型 paddlehub: 建立模型->损失函数->参数学习
    - 策略: 模型选择 损失函数选择 无(关联分析/聚类) 有(分类/回归) 强化(状态/动作/奖励最大化 控制/视觉/NLP)
    - 算法: 学习参数 优化算法 方向传播算法
- NN
  - 技术: 激活函数(非线性 变换 Sigmoid Tanh ReLU LeakyReLU) 偏置值-阈值 损失函数(均方差/交叉熵)-调参 前向/后向传播-学习率 梯度消失(ReLU) 信号a-网络连接w
  - DNN: 分类/回归
  - CNN: 特征提取(conv-activation激励(激活 [0,1] 特征)-pool池化(max sum avg -> 压缩))->分类(全连接) LeNet-AlexNet-GoogLeNet(参数过多/过拟合)-ResNet(残差网络)
  - RNN: 输入(历史+当前) LSTM(forget/input/output Sigmoid)
  - GAN: 生成G/判别D(对抗)-无监督 视频分类
- NLP: 深层-词性标注/实体识别 循环-机器翻译/问答系统 递归-句子解析/情感分析 卷积-文本分类/语义提取
  - 情感分析-分词/停用词-正负面 机器翻译-词性/结构/语序 知识图谱-关联信息/搜索质量-实体识别/关系抽取/实体统一/指代消减
  - 应用: 分词/实体/词性 翻译 知识图谱-语义/情报 机器翻译/文本识别 循环/LSTM
- cv 机器视觉 图像: 处理-灰度化/对比度/旋度
  - 分类(主体/类别) 卷积
  - 检测 R-CNN->SVM
  - 跟踪 生成式 判别式
  - 语义分割 全CNN
  - 应用: 图像分类 物体检测 自动驾驶 人脸识别 卷积
- 语音 信号预处理/特征(模型)/模式匹配 前端/后端(在线)/自适应反馈
  - 对话系统 关键帧-波形/矩阵-NN-循环NN
  - 应用: 转换 合成 客服 直播字幕 循环/LSTM
- 知识图谱 实体/关系 数据采集预处理-设计-存储(GDB)-业务
- 机器人
- 其他
  - 机器学习分类预测 XGBoost LightGBM
- TF
  - 计算图(数据流图/有向图) in/run update/return
  - 基础: tensor->scalar/vector/matrix/cube/nTensor D维度/Shape tensor/var/placeholder/constant TensorBoard
  - 开发流程: Estimator(train eval predict) NN(多层感知 cost-optimizer; CNN conv2d-maxpool-全连接-relu-dropout) optimizer(gradient momentum adam)
- PAI: 传统/DL 拖拽/模型调用
  - 架构: 分布式框架->datawork(清洗/特征)-> AL/Studio(可视化)/DSW(交互 notebook) -> EAS
  - oss-图片/训练数据/代码 nas-文件系统
  - studio: 文本分类/金融风控/商品推荐
- 数据集 dataSet
  - MINIST: 新手入门数据集

## pytorch

```sh
conda create -n pytorch --clone base
pip install torch torchvision torchaudio matplotlib
```

## paddle

```sh
python -m pip install paddlepaddle -i https://mirror.baidu.com/pypi/simple
```

## mark

- 人工智能专家 JudeaPearl
