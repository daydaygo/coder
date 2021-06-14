# https://www.paddlepaddle.org.cn/install/quick

import paddle

print(paddle.__version__)

# 加载内置数据集
from paddle.vision.transforms import ToTensor

train_dataset = paddle.vision.datasets.MNIST(mode='train', transform=ToTensor())
val_dataset = paddle.vision.datasets.MNIST(mode='test', transform=ToTensor())

# 模型搭建
mnist = paddle.nn.Sequential(  # 一层一层网络结构组件起来
    paddle.nn.Flatten(),  # [1,28,28] -> [1,784]
    paddle.nn.Linear(784, 512),
    paddle.nn.ReLU(),
    paddle.nn.Dropout(0.2),
    paddle.nn.Linear(512, 10),
)

# 模型训练
model = paddle.Model(mnist)
model.prepare(
    paddle.optimizer.Adam(parameters=model.parameters()),  # 优化器
    paddle.nn.CrossEntropyLoss(),  # 损失计算方法
    paddle.metric.Accuracy(),  # 精度计算方法
)
model.fit(train_dataset, epochs=10, batch_size=64, verbose=1)

# 模型评估
model.evaluate(val_dataset, verbose=1)
