# LaTeX 数学公式

- [latex在线: $吴文中^\circledR$数学公式编辑器](https://latex.91maths.com/)
- 工具: vscode + ext: markdown all in one(写) + ext: markdown preview enhance(读)
  - markdown all in one -> mdmath(markdown-math) -> katex
    - [katex table](https://katex.org/docs/support_table.html) -> 搜索
    - [katex function](https://katex.org/docs/supported.html) -> 知识地图
    - [语雀数学公式举例](https://www.yuque.com/yuque/help/brzicb)
- 知识地图
  - Accents 重音 $\overrightharpoon{ac}$
  - Delimiters 分隔符 $\lang\rang$
  - Environments
    - matrix $$\begin{matrix} a & b \\ c & d \end{matrix}$$
  - HTML
  - Letters and Unicode
    - Greek Letters $\alpha$ $\Alpha$
    - Other Letters $\imath$ $\jmath$
    - Unicode
  - Layout
    - Annotation $$\tag{hi} x+y^{2x}$$ $$\tag*{hi} x+y^{2x}$$
    - Line Breaks
    - Vertical Layout $x_n$ $e^x$
    - Overlap and Spacing
  - Logic and Set Theory $\to$
  - Macros
  - Operators
    - Big Operators $\sum_{0}^{n}a_i$
    - Binary Operators
    - Fractions and Binomials $\dfrac{a}{b}$
    - Math Operators $\sqrt[3]{x}$
  - Relations
    - Negated Relations
    - Arrows
  - Special Notation
  - Style, Color, Size, and Font
  - Symbols and Punctuation $\infty$ $\circledR$
  - Units

## 示例

- 行内公式: `$x$`
- 行间公式: `$$x$$`
- 普通字母 直接使用: $x$
- 角标 `_` `^`
  - 下标: $x_y$ $x_{yz}$
  - 上标: $x^y$ $x^{yz}$
- 乘法 直接使用: $ab$ $xy$
- 分式: `\frac` `\dfrac`
  - $\frac{x+y}{a+b}$ $\dfrac{x+y}{a+b}$
- 根式 `\sqrt[n]{}`
  - 二次根: $\sqrt{a}$
  - n次根: $\sqrt[3]{a}$
- 对数 $\log_x$
  - $\ln{x} = \log_ex$
  - $\lg{x} = \log_{10}x$
- 求积/求和 `\sum` `\int`
  - 求和: $\sum_{k=1}^nx$
  - 求积: $\sum_{k=1}^{\infty}\frac{x^n}{n!} = \int_0^{\infty}e^x$
  - 改变上线限位置: `\limits` `\nolimits` $\sum\limits_{k=1}^nx$

## 数学公式

- [三角函数公式](https://baike.baidu.com/item/三角函数公式) $\sin{A} = \dfrac{a}{c}$
- 手指计数基本法则 $1+1=2$ 计数法/十进制
- 勾股定理/毕达哥拉斯定理 $a^2+b^2=c^2$ 万物皆数/有理数
- 阿基米德杠杆原理 $f_1x_1=f_2x_2$ 给我一个支点, 我可以撬动地球
  - 数学之神: 阿基米德 牛顿 高斯
- 纳皮尔指数/对数关系公式 $e^{\ln{N}} = N$ $e=2.71828……$ 奇妙的对数定律说明书
  - 伽利略: 给我时间、空间和对数，我可以创造出一个宇宙来
- 牛顿万有引力定律 $F=\dfrac{Gm_1m_2}{r^2}$ G引力常量 $m_1m_2$两个物体质量 r距离
- 麦克斯韦电磁方程组
- 爱因斯坦质能关系式 $E=mc^2$ E能量 m质量 c光速
  - 物理中的数学美: 简洁美与对称美
- 德布罗意公式/波粒二象性 $\lambda=h/mv$ $\lambda$波长 h普朗克常量 mv粒子的动量
- 玻尔兹曼公式/熵增法则 $S=k*logW$ S熵 k玻尔兹曼常数 W微观态数
- 齐奥尔科夫斯基公式/宇航理论/火箭技术 $V=V_e\ln\dfrac{M_0}{M_i}$ V火箭速度增量 $V_e$喷流相对火箭速度 $M_0$火箭开启时质量 $M_i$火箭关闭时质量
- 最美的数学公式-圆周率和自然对数的关系 $i^2 = 1 \Leftrightarrow e^{\pi i} + 1 = 0$
- 最小二乘法 $f(\boldsymbol x) = {1\over 2}(\mathbf{w}^T\boldsymbol{x}-y)^2+\sum_{i=1}^K w_i^2$

## 资料

- 在线编辑公式: https://latex.vimsky.com/
- 改变数学的十个公式: https://mp.weixin.qq.com/s/UeEi9pT5jC_M5sxITOm41g
- 语雀help 如何插入数学公式: https://www.yuque.com/yuque/help/math
- latex入门: https://book.douban.com/subject/24703731/