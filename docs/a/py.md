# python 基础

- why: da ai
- env
  - pip
  - [anacoda](https://www.anaconda.com/download/): [清华镜像](https://mirrors.tuna.tsinghua.edu.cn/help/anaconda/) numpy+pandas+ipython+jupyter
  - debug: pdb cpython Cyberbrain
- syntax
  - type: dynamicType mutableType `None`
  - var: str list tuple(拆包/封包) dict set
  - expr: `: ⇥`
  - op: `a[s:e] -> [s:e)` 左闭右开; `and or not` `in`
  - control: `for-in-range` pass
  - fn: map(func, Iterable) reduce/filter lambda
  - oop: isinstance/type/dir/attr `class B(A)` `__init__(self)`
  - package: `from import as` `__init__.py` `csv.py`
  - error: raise TypeError(); `try-except-else-finally`
  - generator
  - decorator装饰器 contextManager
- fn
  - os
  - math random numpy Matplotlib Anaconda/Canopy
  - jupiter notebook: shortcut-command/edit; shell `!`; 变量监控/运行历史/debug; Magic `% %%`; help `??`
  - pdb
  - format: black yapf pep8
  - pytest unittest RobotFramework RedwoodHQ Jasmine selenium(js数据渲染)
  - Requests BeautifulSoup lxml untangle xmltodict
  - django Flask Tornado jinja2
  - ops/tool: mycli Ansible Puppet Jenkins Travis-CI
  - grpcio grpcio-tools
  - ipython cython Concurrent.futures
  - PyCrypto
  - GUI: tkinter
- project
  - `__init__.py`
  - dep: `pyproject.toml + poetry.lock`
  - doc: sphinx

```sh
# pip 加速: win ~/pip/pip.ini; linux ~/.pip/pip.conf
[global]
index-url = https://pypi.tuna.tsinghua.edu.cn/simple
trusted-host=pypi.tuna.tsinghua.edu.cn
[list]
format=columns

pip install virtualenv
virtualenv --no-site-packages py2

python -m SimpleHTTPServer 3000
landslide config.cfg # https://github.com/adamzap/landslide

# baidu aistudio
!mkdir /home/aistudio/external-libraries # 持久化
!pip install beautifulsoup4 -t /home/aistudio/external-libraries
import sys
sys.path.append('/home/aistudio/external-libraries')
```

```py
'''str'''
'foo' + 'bar'
''.join(map(str, range(20))) # [0, 20) 的连续字符串
'{foo}{bar}'.format(foo=foo, bar=bar)

[1,2,3] # list
a,b,*c = (1,2,3,4) # tuple
{'a': 1} # dict
set([1,2,3]) # set

g = (x * x for x in range(10)) # generator
next(g)
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

sorted()

json.loads(json_string)
json.dumps(dict_data)

# numpy
sum/mean/std/var/min/max/argmin/argmax/cumsum/cumprod

# cv2: opencv-python
```

## conda

```sh
CONDA_SUBDIR=osx-arm64
brew install miniconda
conda create -n native --clone base -c conda-forge # will get you a osx-arm64

conda clean -i # 清除索引缓存，保证用的是镜像站提供的索引
# ~/.condarc
channels:
  - defaults
show_channel_urls: true
report_errors: false
default_channels:
  - https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/main
  - https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/r
  - https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/msys2
custom_channels:
  conda-forge: https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud
  msys2: https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud
  bioconda: https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud
  menpo: https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud
  pytorch: https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud
  simpleitk: https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud
```

- notebook https://www.cnblogs.com/rangger/p/9520123.html
- 图解 numpy https://mp.weixin.qq.com/s/CgJ4gCsi8RGsj5fSQaW7dw
- Pandas
  - 图解: https://mp.weixin.qq.com/s/MFwdjMrPVQESMJjQw__rXw
- vaex(DataFrame): 0.052 秒打开 100GB 数据
- https://github.com/mwouts/jupytext
- [知乎周刊 - 编程小白学Python](https://www.zhihu.com/publications/weekly/19550511)
- [python 书单, 不将就](http://blog.csdn.net/turingbooks/article/details/46459349)
- [廖雪峰的Python3教程](http://www.liaoxuefeng.com/wiki/0014316089557264a6b348958f449949df42a6d3a2e542c000)
- [第三方库 awesome python](https://awesome-python.com/)
- [realpython - learn python programming by example](https://realpython.com/)
- [python course](https://www.python-course.eu/python3_course.php)
- 崔庆才_Python3爬虫入门到精通课程视频附软件与资料 34课
哲学理念: REPL(交互式解释器) -> import this
spaCy+Cython加速NLP
