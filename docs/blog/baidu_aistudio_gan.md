# AI| 百度顶会论文复现营

- 课程地址: https://aistudio.baidu.com/aistudio/course/introduce/1340

## 资料

- GAN
1. LARGE SCALE GAN TRAINING FOR HIGH FIDELITY NATURAL IMAGE SYNTHESIS 
https://github.com/sxhxliang/BigGAN-pytorch
2. Few-shot Video-toVideo Synthesis
https://github.com /NVlabs/few-shotvid2vid
3. First Order Motion Model for Image Animation
https://github.com/AliaksandrSiarohin/first-order-model
4. StarGAN v2: Diverse Image Synthesis for Multiple Domains
https://github.com/clovaai/stargan-v2
5. U-GAT-IT: Unsupervised Generative Attentional Networks with Adaptive Layer-Instance Normalization for Image-to-Image Translation
https://github.com/znxlwm/UGATIT-pytorch

- 视频分类
1. ECO: Efficient Convolutional Network for Online Video Understanding
网址：https://github.com/mzolfaghari/ECO-pytorch
2. Temporal Pyramid Network for Action Recognition
网址：https://github.com/decisionforce/TPN
3. Learning Spatio-Temporal Features with 3D Residual Networks For Action Recognition
网址：https://github.com/kenshohara/3D-ResNets-PyTorch
4. Representation Flow for Action Recognition
网址：https://github.com/piergiaj/representation-flow-cvpr19

## py 基础
```py
l = [1,2,3] # list
t = (1,2,3) # tuple
d = {'a': 1, 'b': 2} # dict
s = set([1,2,3]) # set

# yield
def foo():
    print("starting...")
    while True:
        res = yield 4
        print("res:",res)
g = foo()
print(next(g))
print("*"*20)
print(next(g))
# print(g.send(7))

for n in range(1000): print(n)
```

## numpy
- sum 和
- mean 算数平均数
- std 标准差; var 方差
- min/max argmin/argmax 索引
- cumsum 累加; cumprod 累积

## notebook 基础
> AI Studio基本操作-Notebook篇: https://aistudio.baidu.com/aistudio/projectdetail/671885
> AI Studio基本操作(二) Debug篇: https://aistudio.baidu.com/aistudio/projectdetail/69987

- 文件 
  - 支持 .py, .json, .txt, .log等格式, 部分图片格式
  - >30M => 使用数据集功能
  - 上传 notebook(.ipynb)
  - 下载
  - git 同步代码
- 快捷键: 熟悉 vim 熟悉这个起来很轻松
- 使用 shell 命令, 如 `!pwd`
  - pip 持久化安装

```ipynb
!mkdir /home/aistudio/external-libraries
!pip install beautifulsoup4 -t /home/aistudio/external-libraries

import sys
sys.path.append('/home/aistudio/external-libraries')
```

- magic 命令
  - % - 作用于单行 %% - 作用于单元格
  - %lsmagic 列出全部可用 magic
  - %%timeit 统计运行时长
  - %matplotlib inline
  - %env：设置环境变量
  - %%writefile and %pycat: 导出cell内容/显示外部脚本的内容

- 查看帮助信息: ? ??(addl详细信息), 如 `?pe.fluid.layers.conv3d`

- 变量监控
  - 修改内核选项 `ast_note_interactivity`

```ipynb
from IPython.core.interactiveshell import InteractiveShell
InteractiveShell.ast_node_interactivity = "all"
```

- 调试 pdb
  - ENTER (重复上次命令)
  - c (继续)
  - l (查找当前位于哪里)
  - s (进入子程序,如果当前有一个函数调用，那么 s 会进入被调用的函数体)
  - n(ext) 让程序运行下一行，如果当前语句有一个函数调用，用 n 是不会进入被调用的函数体中的
  - r (运行直到子程序结束)
  - `!<python 命令>`
  - h (帮助)
  - a(rgs) 打印当前函数的参数
  - j(ump) 让程序跳转到指定的行数
  - l(ist) 可以列出当前将要运行的代码块
  - p(rint) 最有用的命令之一，打印某个变量
  - q(uit) 退出调试
  - r(eturn) 继续执行，直到函数体返回
  - ipdb 彩色输出: `!pip install ipdb -i https://pypi.tuna.tsinghua.edu.cn/simple`

## conda 多环境管理
> https://www.jianshu.com/p/edaa744ea47d

```
conda env list
conda create -n <name> python=3.7
conda remove -n <name> --all
conda activate <name>
conda deactivate <name>
```

## paddlehub

- porn_detection_lstm

```
# TODO: 使用 conda 设置的新环境安装后, 还缺少大量的包, 安装后才跑起来
pip install paddlehub

# 情感分析
hub run senta_lstm --input_text "这家餐厅很好吃"

# 口罩检测
hub run pyramidbox_lite_mobile_mask --input_path='test.jpg'

# 人像抠图
hub run deeplabv3p_xception65_humanseg --input_path='test.jpg'

# 风格迁移
hub run stylepro_artistic
```

- 命令式编程 

```py
import paddle.fluid as fluid
from paddle.fluid.dygraph.base import to_variable

# 开启动态图模式
with fluid.dygraph.guard():
    # 动态图模式下，将numpy的ndarray类型的数据转换为Variable类型
    x = fluid.dygraph.to_variable(data)
    print('In DyGraph mode, after calling dygraph.to_variable, x = ', x)
    # 动态图模式下，对Variable类型的数据执行x=x+10操作
    x += 10
    # 动态图模式下，调用Variable的numpy函数将Variable类型的数据转换为numpy的ndarray类型的数据
    print('In DyGraph mode, data after run:', x.numpy())
```