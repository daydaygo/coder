# pff, Python for finance Python金融大数据分析

## env

```bash
# 完整版
brew cask install ananconda
# 配置 PATH(fish)
set -gx fish_user_paths /usr/local/anaconda3/bin $fish_user_paths

# miniconda
brew cask install miniconda
# 配置 PATH(fish)
set -gx fish_user_paths ~/opt/miniconda3/bin $fish_user_paths
```

## py 语言基础

```py
# 10^100
10 ** 100

# 浮点数精度
import decimal
from decimal import Decimal
decimal.getcontext() # 查看默认配置
decimal.getcontext.prec = 4 # 小于默认精度
e = Decimal(1) / Decimal(11)

# keyword
import keyword
keyword.kwlist

# 函数式编程: map() filter() reduce()
def even(x):
    return x%2 == 0
list(map(even, range(10)))
list(map(lambda x:x**2, range(10)))

# OO
```

## NumPy

```py
# 避免引用指针
from copy import deepcopy

# python array 类型
import array

import numpy as np
np.where()
```

## pandas

```py
import pandas as pd
df = pd.DataFrame()

df.info()
df.describe()

# 基础可视化
from pylab import plt,mpl
plt.style.use('seaborn')
mpl.rcParams['font.family'] = 'serif'
%matplitlib inline
# 专为 DataFrame 对象设计的 matplotlib 包装器: plot()
df.cumsum().plot(lw=2.0, figsize=(10,6))

# 柱状图
df.plot.bar(figsize=(10,6), rot=15)
df.plot(kind='bar', figsize=(10,6))

df.groupby('Quarter')

s = pd.Series(np.linspace(0, 15, 7), name='series')
```

## 可视化

```py
import matplotlib as mpl
import matplotlib.pyplot as plt
plt.sytle.use('seaborn')
mpl.rcParams['font.family'] = 'serif'
%matplotlib inline
```

## 通用函数

```py
import numpy as np
np.mean() # 均值
np.min()
np.std() # 标准差
np.median() # 中位数
np.max()

import pandas as pd
df = pd.DataFrame()
df.diff().head() # 2个索引之间的绝对变化
df.pct_change().round(3).head() # 2个索引值之间的变化率
```

## IO

```py
f = open('a.csv', 'w') # r b
f.close()

import csv
with open('a.csv', 'r') as f:
    csv_reader = csv.reader(f)
    # csv_reader = csv.DictReader(f)
    lines = [line for line in csv_reader]

import sqlite3 as sql
conn = sql.connect('a.db')
query = 'select * from t'
conn.execute(query)
conn.commit()
q = conn.execute # 定义别名
rows = q(query).fetchall()
res = np.array(q(query).featchall()).round(3) # 转换成 ndarray 对象
conn.close()

np.save('a.npy', data)
np.load('a.npy')

# pd 原生支持 csv/sql/xls/json/html
pd.read_sql(query) 

h5s = pd.HDFStore('a.h5s', 'w)
h5s['data'] = data
h5s.close()

data.to_csv('a.csv')
df = pd.read_csv('a.csv')

data[:100].to_excel('a.xls')
df = pd.read_excel('a.xls')

# pytables
import tables as tb
h5 = tb.open_file('a.h5', 'w')
tb.Filters(complevel=5, complib='blosc') # 压缩表
out = h5.create_earray() # 内存外支持
expr = tb.Expr('3 * sin(ear) + sqrt(abs(ear)))') # 数学公式支持

# tstables
import tstables as tstab
h5 = tb.open_file('a.h5', 'w')
ts = h5.create_ts('/', 'ts', ts_desc)
rows = ts.read_range()
rows.info() # pd
h5.close()
```

## 性能

- 风格与范型: 替代循环
- 库: NumPy.ndarray pandas.DataFrame
- 编译: 静态-Cython 动态-Numba
- 并行化

## 数学公式

- 逼近法
- 凸优化
- 积分
- 符号数学 SymPy

```py
import scipy.integrate as sci
x = np.linespace(0, 10)
y = f(x)
a = 0.5 # 积分左界
b = 9.5 # 积分右界
Ix = np.linspace(a, b) # 积分区间值
Iy = f(Ix) # 积分函数值

import sympy as sy
x = sy.Symbol('x') # 定义要处理的符号
y = sy.Symbol('y')
sy.sqrt(x) # 对一个符号应用函数
3 + sy.sqrt(x) - 4 ** 2 # 在符号上定义数值表达式
f = x ** 2 + 3 + 0.5 * x ** 2 + 3 / 2 # 符号化定义的函数
sy.simplify(f) # 简化的函数表达式
sy.pretty(sy.sqrt(x) + 0.5) # sy 渲染器: LaTeX Unicode ASCII
sy.solve(x**2 - 1) # 解方程
```

## 推断统计学
## 统计学

# 算法交易
## FXCM交易平台
## 交易策略
## 自动化交易
## 衍生品分析
## 估值框架
## 金融模型的模拟
## 衍生品估值
## 投资组合估值
## 基于市场的估值