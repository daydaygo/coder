# 程序员修炼之道

摘录几则书评, 作为是否读这本书的参考:

- 这样的赞美一直不绝于耳: 通过撰写一本书来推动整个行业
- 如果想让自己写的软件既领先与时代又易于维护, 手边当常备此书
- 它完全改变了我的职业轨迹!
- 二十年前的第一版, 彻底颠覆了我的技术生涯

## 通往务实的最高境界

> ed2 新增内容

- 务实的哲学
  - (add)人生是你的, 你有权选择 -- 改变组织/换个组织
- 务实的方法
  - (add)ETC, easy to change, 优秀的设计比糟糕的设计更易变更
- 基础工具
  - (add)工程日记
- (change)拧折不弯 Bend, or Break
  - 解耦: ETC 链式调用 全局数据
  - 在现实世界抛球杂耍: 有限状态机 观察者模式 发布/订阅 响应式编程
  - 变换式编程: 编程讲的是代码, 而程序谈的是数据; 不要囤积状态, 传递下去
  - 继承税: 接口表达多态 mixin共享功能
- (add)并发
  - 共享状态是不正确的状态
  - 角色与进程
- 当你编程时
  - (add)倾听蜥蜴脑: 相信直觉 啊哈时刻
  - (add)出门在外注意安全: 安全性的基本原则
    - 将攻击面最小化
    - 最小特权原则
    - 安全的默认值
    - 敏感数据要加密
    - 维护安全更新
- 项目开始之前
  - (change)敏捷的本质: 形容词 - 如何做的价值观, 而非名词 - 这么干就是敏捷
- 务实的项目
  - (change)务实的团队
    - 维持小而稳定的团队
    - 排上日程以待其成
    - 组织全功能团队
    - 椰子派不上用场: 起作用, 而非赶时髦
    - 在用户需要时交付
- (add)跋
  - 勿先伤害: 犯错在所难免, 确保犯错后没人因此受难
  - 不要助纣为虐: 因为这样做你也有编程纣王的风险

## 从小工到专家

1. 关心你的技艺 Care About Your Craft 如果你不在乎能否漂亮地开发出软件，你又为何要耗费生命去开发软件呢？
2. 思考！你的工作 ThinklAbout Your Work 关掉自动驾驶仪，接管操作。不断地批评和评估你的工作。
3. 提供各种选择，不要找蹩脚的借口 要提供各种选择，而不是找借口。不要说事情做不到；说明能够做什么。
4. 不要容忍破窗户 Don't Live with Broken Windows 当你看到糟糕的设计、错误的决策和糟糕的代码时，修正它们。
5. 做变化的催化剂 Be a Catalyst for Change 你不能强迫人们改变。相反，要向他们展示未来可能会怎样，并帮助他们参与对未来的创造。
6. 记住大图景 Remember the Big Picture 不要太过专注于细节，以致忘了查看你周围正在发生什么。
7. 使质量成为需求问题 Make Quality a Requirements Issue 让你的用户参与确定项目真正的质量需求。
8. 定期为你的知识资产投资 Invest Regularly in Your Knowledge Portfolio 让学习成为习惯。
9. 批判地分析你读到的和听到的 Critically Analyze What You Read and Hear 不要被供应商、媒体炒作、或教条左右。要依照你自己的看法和你的项目的情况去对信息进行分析。
10. 你说什么和你怎么说同样重要 如果你不能有效地向他人传达你的了不起的想法，这些想法就毫无用处。
11. 不要重复你自己 DRY-Don't repeat Vourself 系统中的每一项知识都必须具有单一、无歧义、权威的表示。
12. 让复用变得容易 Make It Easy to Reuse 如果复用很容易，人们就会去复用。创造一个支持复用的环境。
13. 消除无关事物之间的影响 Eliminate Effects Between Unrelated Things 设计自足、独立、并具有单一、良好定义的目的的组件。
14. 不存在最终决策 There Are No Final Decisions 没有决策是浇铸在石头上的。相反，要把每项决策都视为是写在沙滩上的，并为变化做好计划。
15. 用曳光弹找到目标 Use Tracer Bullets to Rnd the Target 曳光弹能通过试验各种事物并检查它们离目标有多远来让你追踪目标。
16. 为了学习而制作原型 Prototype to Learn 原型制作是一种学习经验。其价值并不在于所产生的代码，而在于所学到的经验教训。
17. 靠近问题领域编程 Program Close to the Problem domain用你的用户的语言进行设计和编码。
18. 估算，以避免发生窻外 Estimate to Avoid Surprises 在着手之前先进行估算。你将提前发现潜在的问题。
19. 通过代码对进度表进行迭代 Iterate the Schedule with the Code 用你在进行实现时获得的经验提炼项目的时间标度„
20. 用纯文本保存知识 Keep Knowledge in Plain Text 纯文本不会过时。它能够帮助你有效利用你的工作，并简化调试和测试。
21. 利用命令shell的力量 Use the Power of Command Shells 当图形用户界面无能为力时使用shell。
22. 用好一种编辑器 Use a Single Editor Well 编辑器应该是你的手的延伸；确保你的编辑器是可配置、可扩展和可编程的。
23. 总是使用源码控制 Always Use Source Code Control 源码控制是你的工作的时间机器一^你能够回到过去。
24. 要修正问题，而不是发出指责 Fix the Problem,Not the Blame bug是你的过锗还是别人的过错，并不是真的很有关系一它仍然是你的问题，它仍然需要修正。
25. 不要恐慌 Don't Panic When Debuging 做一次深呼吸，思考什么可能是bug的原因。
26. “Select”没有问题 'Select"lsn't Broken 在OS或编译器、甚或是第三方产品或库中很少发现bug。bug很可能在应用中。
27. 不要假定，要证明 Don't Assume It-Prove It 在实际环境中一^使用真正的数据和边界条件一证明你的假定。
28. 学习一种文本操纵语言 Learn a Text Manipulation Language 你用每天的很大一部分时间处理文本，为什么不让计算机替你完成部分工作昵？
29. 编写能编写代码的代码 Write Code That Writes Code 代码生成器能提高你的生产率，并有助于避免重复。
30. 你不可能写出完美的软件 You Can't Write Perfect Software 软件不可能完美。保护你的代码和用户，使它（他）们免于能够预见的错误。
31. 通过合约进行设计 Design with Contracts 使用合约建立文档，并检验代码所做的事情正好是它声明要做的。
32. 早崩溃 Crash Early 死程序造成的危害通常比有问题的程序要小得多。
33. 用断言避免不可能发生的事情 Use Assertions to Prevent the Impossible 断言验证你的各种假定。在一个不确定的世界里，用断言保护你的代码。
34. 将异常用于异常的问题 Use Exceptions for Exceptional Problems 异常可能会遭受经典的意大利面条式代码的所有可读性和可维护性问题的折磨。将异常保留给异常的事物。
35. 要有姶有终 Finish What You Start 只要可能，分配某资源的例程或对象也应该负贲解除其分配。
36. 使模块之间的耦合减至最少 Minimize Coupling Between Modules 通过编写“羞怯的”代码并应用得墨忒耳法则来避免耦合。
37. 要配置，不要集成 Configure.Don't Integrate 要将应用的各种技术选择实现为配置选项，而不是通过集成或工程方法实现。
38. 将抽象放进代码，细节放进元数据 Put Abstractions in Code,Details in Metadata 为一般情况编程，将细节放在被编译的代码库之外。
39. 分析工作流，以改善并发性 Analyze Workflow to Improve Concurrency利用你的用户的工作流中的并发性。
40. 用服务进行设计 Design Using Services 根据服务——独立的、在良好定义、一致的接口之后的并发对象——进行设计。
41. 总是为并发进行设计 Always Design for Concurrency 容许并发，你将会设计出更整洁、具有更少假定的接口。
42. 使视图与模型分离 Separate Views from Models 要根据模型和视图设计你的应用，从而以低廉的代码获取灵活性。
43. 用黑板协调工作流 Use Blackboards to Coordinate Workflow 用黑板协调完全不同的事实和因索，同时又使各参与方保持独立和隔离。
44. 不要靠巧合编程 Don't Program by Coincidence 只依靠可靠的事物。注意偶发的复杂性，不要把幸运的巧合与有目的的计划混为一谈。
45. 估算你的算法的阶 Estimate the Order of Your Algorithms 在你编写代码之前，先大致估算事情需要多长时间。
46. 测试你的估算 Test Your Estimates 对算法的数学分析并不会告诉你每一件事情。在你的代码的目标环境中测定它的速度。
47. 早重构，常重构 Refactor Early,Refactor Often 就和你会在花园里除草、并重新布置一样，在需要时对代码进行重写、重做和重新架构。要铲除问题的根源。
48. 为测试而设计 Design to Test 在你还没有编写代码时就开始思考测试问题。
49. 测试你的软件，否则你的用户就得测试 Test Your Software,or Your Users Will 无情地测试。不要让你的用户为你查找bug。
50. 不要使用你不理解的向导代码 Don't Use Wizard Code You Don't Understand 向导可以生成大置代码。在你把它们合并进你的项目之前，确保你理解全部这些代码。
51. 不要搜集需求一挖振它们 Don't Gather Requirements-Dig for Them 需求很少存在于表面上。它们深深地埋藏在层层假定、误解和政治手段的下面。
52. 与用户一同工作，以像用户一样思考 Work with a User to Think Like a User 要了解系统实际上将如何被使用，这是最好的方法。
53. 抽象比细节活得更长久 Abstractions Live Longer than Details ‘投资”于抽象，而不是实现。抽象能在来自不同的实现和新技术的变化的“攻击”之下存活下去。
54. 使用项目词汇表 Use a Project Glossary 创建并维护项目中使用的专用术语和词汇的单一信息源。
55. 不要在盒子外面思考一要找到盒子 Don't Think Outside the Box-Find the Box 在遇到不可能解决的问题时，要确定真正的约束。问问你自己：“它必须以这种方式完成吗？它真的必须完成吗？
56. 等你准备好再开始 Start When You're Ready 你的一生都在积累经验。不要忽视反复出现的疑虑。
57. 对有些事情“做”胜于“描述” Some Things Are Better Done than Described 不要掉进规范的螺旋——在某个时刻，你需要开姶编码。
58. 不要做形式方法的奴隶 Don't Be a Slave to Formal Methods 如果你没有把某项技术放进你的开发实践和能力的语境中，不要盲目地采用它。
59. 昂贵的工具不一定能制作出更好的设计 Costly Tools Don't Produce Better Designs 小心供应商的炒作，行业教条、以及价格标签的诱惑。要根据工具的价值判断它们。
60. 围绕功能组织团队 Organize Teams Around Functionality 不要把设计师与编码员分开，也不要把测试员与数据建模员分开。按照你构建代码的方式构建团队。
61. 不要使用手工流程 Don't Use Manual Procedures shell脚本或批文件会一次次地以同一顺序执行同样的指令。
62. 早测试，常测试，自动测试。 Test Early.Test Often.Test Automatically 与呆在书架上的测试计划相比，每次构建时运行的测试要有效得多。
63. 要到通过全部测试，编码才算完成 Coding Ain't Done'Til All the Tests Run就是这样。
64. "蓄意破坏”测试你的测试19 Use Saboteurs to Test Your Testing 在单独的软件副本上故意引入bug, 以检验测试能够抓住它们。
65. 测试状态覆盖，而不是代码覆盖 Test State Coverage,Not Code Coverage 确定并测试重要的程序状态。只是测试代码行是不够的。
66. —个bug只抓一次 Find Bugs Once 一旦测试员找到一个bug, 这应该是测试员最后一次找到它。此后自动测试应该对其进行检查。
67. 英语就是一种编程语言 English is Just a Programming Language 像你编写代码一样编写文档：遵守DflWS则、使用元数据、MVC、自动生成，等等。
68. 把文档建在里面，不要拴在外面 Build Documentation In,Don't Bolt It On 与代码分离的文档不太可能被修正和更新。
69. 温和地超出用户的期望 Gently Exceed Your Users'Expectations 要理解你的用户的期望，然后给他们的东西要多那么一点。
70. 在你的作品上签名 Sign Your Work 过去时代的手艺人为能在他们的作品上签名而自豪。你也应该如此。

检查清单
71. 要学习的语言 它们每一种都有不同的能力和不同的“风味”。用其中的一种或多种语言在家里开发一个小项目。 厌倦了C、C++和JAVA? 试试CLOS、Dylan、Eiffel、Objective C、Prolog、Smalltalk或TOM
72. WISDOM离合诗
    1. What do you want them to learn? 你想让他们学到什么？
    2. What is their interest in what you've got to say? 他们对你讲的什么感兴趣？
    3. How sophisticated are they? 他们有多富有经验？
    4. How much detail do they want? 他们想要多少细节？
    5. 想要让谁拥有这些信息？
    6. 你如何促使他们听你说话？
73. 怎样维持正交性
    - 设计独立、良好定义的组件。
    - 使你的代码保持解藕。
    - 避免使用全局数据。
    - 參重构相似的函数。
74. 应制作原型的事物
    - 架构
    - 已有系统中的新功能
    - 外部数据的结构或内容
    - 第三方工具或组件
    - 性能问题
    - 用户界面设计
75. 架构问题
    - 贵任是否得到了良好定义？
    - 协作是否得到了良好定义？
    - 參耦合是否得以最小化？
    - 你能否确定潜在的重复？
    - 接口定义和各项约束是否可接受？
    - 模块能否在需要时访问所需数据？
76. 调试检直滑单
    - 正在报告的问题是底层bug的直接结果，还是只是症状？
    - bug真的在编译器里？在OS里？或者是在你的代码里？
    - 如果你向同事详细解释这个问题，你会说什么？
    - 如果可疑代码通过了单元测试，测试是否足够完整？如果你用该数据运行单元测试，会发生什么？
    - 造成这个bug的条件是否存在于系统中的其他任何地方？
77. 函数的得墨忒耳法则 某个对象的方法应该只调用属于以下情形的方法:
    - 它自身
    - 传入的任何参数
    - 參它创建的对象
    - 组件对象
78. 怎样深思熟虑地编程
    - 总是意识到你在做什么。
    - 不要盲目地编程。
    - 按照计划行事
    - 依靠可靠的事物。
    - 为你的假定建立文档
    - 不要只是测试你的代码，还要测试你的假定。
    - 为你的工作划分优先级。
    - 不要做历史的奴隶。
79. 何时进行重构
    - 你发现了对DflVIS则的违反。
    - 你发现事物可以更为正交。
    - 你的知识扩展了。
    - 需求演变了。
    - 你需要改善性能。
80. 劈开戈尔迪斯结 在解决不可能解决的问题时，问问你自己:
    - 有更容易的方法吗？
    - 我是在解决正确的问题吗？
    - 这件事情为什么是一个问题？
    - 是什么使它如此难以解决？
    - 它必须以这种方式完成吗？
    - 它真的必须完成吗？
81. 测试的各个方面
    - 单元测试
    - 集成测试
    - 验证和校验
    - 资源耗尽、锗误及恢复
    - 性能测试
    - 可用性测试
    - 对测试自身进行测试
