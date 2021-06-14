# DP DesignPattern 设计模式

- 实现了面向对象的语言都可以应用设计模式
- 设计模式实际是特定场景下的更加优良的解决方案

- principle原则
  - SingleResponsibility单一职责 接口一定要做到 复杂度/变更引起的风险↓ 可读性/可维护性↑ 订单下单->售后一系列解耦
  - LiskovSubstitution里氏替换 所有引用基类的地方必须能透明地使用其子类的对象 互联网解决实际问题
  - InterfaceSegregation接口隔离
  - DependenceInversion依赖倒置 高层模块不应该依赖底层模块, 两者都应该依赖其抽象; 抽象不依赖细节; 细节依赖抽象; 第三方api封装
  - OpenClosed开闭 开放=扩展 修改=关闭 变化=逻辑+子模块+可见视图 通过扩展而非修改来实现变化 多设计源配置
  - LawOfDemeter迪米特=LeastKnowledge最小知识 接口设计
  - CompositeAggregateReuse合成复用 聚合复用场景

- 创建型
  - singleton单例 thread饿汉.枚举.container
  - simpleFactory简单工厂 产品创建细节
  - factory工厂 产品扩展
  - abstractFactory抽象工厂 产品扩展.dbPool
  - builder建造者 sqlBuilder
  - prototype原型 序列化
- 结构型
  - adapter适配器 第三方登录.代码优化
  - decorator装饰 煎饼.logExt
  - proxy代理 动=数据源切换.静=三层
  - bridge桥接 复杂消息
  - component组合 透明组合=课程目录.安全组合=无限级fs
  - facaed门面/外观 整合api
  - flyweight享元 pool(资源/db).状态(内/外)
- 行为型
  - templateMethod模板方法 hook
  - strategy策略 促销.优惠.模板
  - chainOfResponsibility责任链 权限热插拔.builder
  - state状态 登录态切换.订单状态流转.责任链.策略
  - delegate委派 task
  - observer观察者 notify.鼠标事件
  - Visitor访问者 KPI考核.分派(动/静/伪动)
  - memento备忘录 草稿箱
  - mediator中介者 群聊
  - iterator迭代器 set集合
  - interpreter解释器 mathExpr
  - command命令 播放器控制条

- <https://www.w3cschool.cn/uml_tutorial>
- 设计模式的应用

## Singleton 单例

- 定义: Ensure a class has only one instance, and provide a global point of access to it.（确保某一个类只有一个实例，而且自行实例化并向整个系统提供这个实例。）
- 使用场景: 唯一id; 共享访问点或共享数据; 创建一个对象需要消耗的资源过多, 如io
  - java: Runtime Calendar enum实现线程安全单例

```java
// 线程安全
public class Singleton {
    private static final Singleton singleton = new Singleton();

    //限制产生多个对象
    private Singleton() {
    }

    //通过该方法获得实例对象
    public static Singleton getSingleton() {
        return singleton;
    }

    //类中其他方法，尽量是static
    public static void doSomething() {
    }
}

// 线程不安全
public class Singleton {
    private static Singleton singleton = null;

    //限制产生多个对象
    private Singleton() {
    }

    //通过该方法获得实例对象
    public static Singleton getSingleton() {
        if (singleton == null) {
            singleton = new Singleton();
        }
        return singleton;
    }
}
```

## Factory 工厂模式

- 定义: Define an interface for creating an object,but let subclasses decide which class to instantiate.Factory Method lets a class defer instantiation to subclasses.（定义一个用于创建对象的接口，让子类决定实例化哪一个类。工厂方法使一个类的实例化延迟到其子类。）
- 简单工厂模式: 一个模块仅需要一个工厂类，没有必要把它产生出来，使用静态的方法
- 多个工厂类: 每个人种（具体的产品类）都对应了一个创建者，每个创建者独立负责创建对应的产品对象，非常符合单一职责原则
- 代替单例模式: 单例模式的核心要求就是在内存中只有一个对象，通过工厂方法模式也可以只在内存中生产一个对象
- 延迟初始化: ProductFactory负责产品类对象的创建工作，并且通过prMap变量产生一个缓存，对需要再次被重用的对象保留
- 使用场景：jdbc连接数据库，硬件访问，降低对象的产生和销毁
  - java: Boolean等不可变类 Logback&dubbo

```java
public class ConcreteCreator extends Creator {
    public <T extends Product> T createProduct(Class<T> c) {
        Product product = null;
        try {
            product = (Product) Class.forName(c.getName()).newInstance();
        } catch (Exception e) { //异常处理
        }
        return (T) product;
    }
}
```

## simpleFactory 简单工程

- java: Logback&jdk

## AbstractFactory 抽象工厂模式

- 定义: Provide an interface for creating families of related or dependent objects without specifying their concrete classes.（为创建一组相关或相互依赖的对象提供一个接口，而且无须指定它们的具体类。）
- 使用场景: 一个对象族（或是一组没有任何关系的对象）都有相同的约束。涉及不同操作系统的时候，都可以考虑使用抽象工厂模式
  - spring

```java
public abstract class AbstractCreator {
    //创建A产品家族
    public abstract AbstractProductA createProductA();

    //创建B产品家族
    public abstract AbstractProductB createProductB();
}
```

## TemplateMethod 模板方法模式

- 定义: Define the skeleton of an algorithm in an operation,deferring some steps to subclasses.Template Method lets subclasses redefine certain steps of an algorithm without changing the algorithm's structure.（定义一个操作中的算法的框架，而将一些步骤延迟到子类中。使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。）
- AbstractClass 抽象模板
  - 基本方法: 基本方法也叫做基本操作，是由子类实现的方法，并且在模板方法被调用。
  - 模板方法: 可以有一个或几个，一般是一个具体方法，也就是一个框架，实现对基本方法的调度，完成固定的逻辑。注意：为了防止恶意的操作，一般模板方法都加上final关键字，不允许被覆写。
- ConcreteClass 具体模板
- 使用场景
  - 多个子类有公有的方法，并且逻辑基本相同时。
  - 重要、复杂的算法，可以把核心算法设计为模板方法，周边的相关细节功能则由各个子类实现。
  - 重构时，模板方法模式是一个经常使用的模式，把相同的代码抽取到父类中，然后通过钩子函数（见“模板方法模式的扩展”）约束其行为。

## Builder 建造者模式

- 定义: Separate the construction of a complex object from its representation so that the same construction process can create different representations.（将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示。）
- 使用场景
  - 相同的方法，不同的执行顺序，产生不同的事件结果时，可以采用建造者模式。
  - 多个部件或零件，都可以装配到一个对象中，但是产生的运行结果又不相同时，则可以使用该模式。
  - 产品类非常复杂，或者产品类中的调用顺序不同产生了不同的效能，这个时候使用建造者模式非常合适。
- 建造者=组装顺序 工厂方法=创建零件

## Proxy 代理模式

- 定义: Provide a surrogate or placeholder for another object to control access to it.（为其他对象提供一种代理以控制对这个对象的访问。）
- 普通代理=需要知道代理存在 强制代理=代理直接调用真实角色 动态代理

## Prototype 原型模式

- 定义: Specify the kinds of objects to create using a prototypical instance,and create new objects by copying this prototype.（用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象。）
- 原型模式实际上就是实现Cloneable接口，重写clone()方法
  - 性能优良: 原型模式是在内存二进制流的拷贝，要比直接new一个对象性能好很多，特别是要在一个循环体内产生大量的对象时，原型模式可以更好地体现其优点。
  - 逃避构造函数的约束
- 使用场景：**
- 资源优化场景: 类初始化需要消化非常多的资源，这个资源包括数据、硬件资源等。
- 性能和安全要求的场景: 通过new产生一个对象需要非常繁琐的数据准备或访问权限，则可以使用原型模式。
- 一个对象多个修改者的场景: 一个对象需要提供给其他对象访问，而且各个调用者可能都需要修改其值时，可以考虑使用原型模式拷贝多个对象供调用者使用。

```java
public class PrototypeClass implements Cloneable {
    //覆写父类Object方法
    @Override
    public PrototypeClass clone() {
        PrototypeClass prototypeClass = null;
        try {
            prototypeClass = (PrototypeClass) super.clone();
        } catch (CloneNotSupportedException e) {//异常处理
        }
        return prototypeClass;
    }
}
```

## 中介者模式

- 定义: Define an object that encapsulates how a set of objects interact.Mediator promotes loose coupling by keeping objects from referring to each other explicitly,and it lets you vary their interaction independently.（用一个中介对象封装一系列的对象交互，中介者使各对象不需要显示地相互作用，从而使其耦合松散，而且可以独立地改变它们之间的交互。）
- Mediator 抽象中介者角色 抽象中介者角色定义统一的接口，用于各同事角色之间的通信
- ConcreteMediator 具体中介者角色 具体中介者角色通过协调各同事角色实现协作行为，因此它必须依赖于各个同事角色。
- Colleague 同事角色 depMethod依赖方法=依赖中介者 selfMethod自发行为
- 使用场景: 中介者模式适用于多个对象之间紧密耦合的情况，紧密耦合的标准是：在类图中出现了蜘蛛网状结构，即每个类都与其他的类有直接的联系

```java
public abstract class Mediator {
    //定义同事类
    protected ConcreteColleague1 c1;
    protected ConcreteColleague2 c2;

    //通过getter/setter方法把同事类注入进来
    public ConcreteColleague1 getC1() {
        return c1;
    }

    public void setC1(ConcreteColleague1 c1) {
        this.c1 = c1;
    }

    public ConcreteColleague2 getC2() {
        return c2;
    }

    public void setC2(ConcreteColleague2 c2) {
        this.c2 = c2;
    }

    //中介者模式的业务逻辑
    public abstract void doSomething1();

    public abstract void doSomething2();
}
```

## 命令模式

- 定义: Encapsulate a request as an object,thereby letting you parameterize clients with different requests,queue or log requests,and support undoable operations.（将一个请求封装成一个对象，从而让你使用不同的请求把客户端参数化，对请求排队或者记录请求日志，可以提供命令的撤销和恢复功能。）
- Receive接收者角色 Command命令角色 Invoker调用者角色
- 使用场景: 认为是命令的地方就可以采用命令模式, 如GUI按钮点击, 模拟dos命令

## 责任链模式

- 定义: Avoid coupling the sender of a request to its receiver by giving more than one object a chance to handle the request.Chain the receiving objects and pass the request along the chain until an object handles it.（使多个对象都有机会处理请求，从而避免了请求的发送者和接受者之间的耦合关系。将这些对象连成一条链，并沿着这条链传递该请求，直到有对象处理它为止。）
- handleMessage请求的处理方法(唯一对外开放) setNext编排方法 level能够处理的级别+resp任务处理
- 注意事项: 避免超长链, 如handler设置最大节点setNext中判断

```java
public abstract class Handler {
    private Handler nextHandler;

    //每个处理者都必须对请求做出处理
    public final Response handleMessage(Request request) {
        Response response = null;
        //判断是否是自己的处理级别
        if (this.getHandlerLevel().equals(request.getRequestLevel())) {
            response = this.echo(request);
        } else { //不属于自己的处理级别
            //判断是否有下一个处理者
            if (this.nextHandler != null) {
                response = this.nextHandler.handleMessage(request);
            } else { //没有适当的处理者，业务自行处理
            }
        }
        return response;
    }

    //设置下一个处理者是谁
    public void setNext(Handler _handler) {
        this.nextHandler = _handler;
    }

    //每个处理者都有一个处理级别
    protected abstract Level getHandlerLevel();

    //每个处理者都必须实现处理任务
    protected abstract Response echo(Request request);
}
```

## Decorator 装饰模式

- 定义: Attach additional responsibilities to an object dynamically keeping the same interface.Decorators provide a flexible alternative to subclassing for extending functionality.（动态地给一个对象添加一些额外的职责。就增加功能来说，装饰模式相比生成子类更为灵活。）
- Component抽象构件 ConcreteComponent具体构件 Decorator装饰角色 ConcreteDecorator具体装饰角色
- 使用场景
  - 需要扩展一个类的功能，或给一个类增加附加功能
  - 需要动态地给一个对象增加功能，这些功能可以再动态地撤销
  - 需要为一批的兄弟类进行改装或加装功能，当然是首选装饰模式
  - java: IO(Buffered BufferedReader BufferedWriter)

## Strategy 策略模式

- 定义: Define a family of algorithms,encapsulate each one,and make them interchangeable.（定义一组算法，将每个算法都封装起来，并且使它们之间可以互换。）
- Context封装角色 Strategy抽象策略(algo 算法)角色 ConcreteStrategy具体策略角色(超过4个考虑混合模式)
- 使用场景
  - 多个类只有在算法或行为上稍有不同的场景。
  - 算法需要自由切换的场景。
  - 需要屏蔽算法规则的场景。

```java
// 策略模式扩展: 策略枚举
// 需要暴露所有策略并由客户端决定使用
public enum Calculator {
    //加法运算
    ADD("+") {
        public int exec(int a, int b) {
            return a + b;
        }
    },
    //减法运算
    SUB("-") {
        public int exec(int a, int b) {
            return a - b;
        }
    };
    String value = "";

    //定义成员值类型
    private Calculator(String _value) {
        this.value = _value;
    }

    //获得枚举成员的值
    public String getValue() {
        return this.value;
    }

    //声明一个抽象函数
    public abstract int exec(int a, int b);
}
```

## Adapter 适配器模式

- 定义: Convert the interface of a class into another interface clients expect.Adapter lets classes work together that couldn't otherwise because of incompatible interfaces.（将一个类的接口变换成客户端所期待的另一种接口，从而使原本因接口不匹配而无法在一起工作的两个类能够在一起工作。）
- Target目标角色 Adaptee源角色 Adapter适配器角色
  - 类适配器=继承 对象适配器=对象合成
- 使用场景: 修改一个已经投产中的接口时=扩展应用=设计阶段不用考虑

## Iterator 迭代器模式

- 定义：Provide a way to access the elements of an aggregate object sequentially without exposing its underlying representation.（它提供一种方法访问一个容器对象中各个元素，而又不需暴露该对象的内部细节。）
- Iterator抽象迭代器(first() next() isDone()) ConcreteIterator具体迭代器 Aggregate抽象容器(createIterator()) ConcreteAggregate具体容器
  - iterator -> collection

## Composite 组合模式

- 定义: Compose objects into tree structures to represent part-whole hierarchies.Composite lets clients treat individual objects and compositions of objects uniformly.（将对象组合成树形结构以表示“部分-整体”的层次结构，使得用户对单个对象和组合对象的使用具有一致性。）
- Component抽象构件角色 Leaf叶子构件 Composite树枝构件
- 使用场景
  - 维护和展示部分-整体关系的场景，如树形菜单、文件和文件夹管理
  - 从一个整体中能够独立出部分模块或功能的场景
  - 树形结构就可以考虑组合模式

```java
public class Composite extends Component {
    //构件容器
    private ArrayList<Component> componentArrayList = new ArrayList<Component>();

    //增加一个叶子构件或树枝构件
    public void add(Component component) {
        this.componentArrayList.add(component);
    }

    //删除一个叶子构件或树枝构件
    public void remove(Component component) {
        this.componentArrayList.remove(component);
    }

    //获得分支下的所有叶子构件和树枝构件
    public ArrayList<Component> getChildren() {
        return this.componentArrayList;
    }
}
```

## Observer 观察者模式

- 定义: Define a one-to-many dependency between objects so that when one object changes state,all its dependents are notified and updated automatically.（定义对象间一种一对多的依赖关系，使得每当一个对象改变状态，则所有依赖于它的对象都会得到通知并被自动更新。）
- Subject被观察者(动态增删观察者) ConcreteSubject具体的被观察者 Observer观察者 ConcreteObserver具体的观察者
- 使用场景
  - 关联行为场景。需要注意的是，关联行为是可拆分的，而不是“组合”关系
  - 事件多级触发场景
  - 跨系统的消息交换场景，如消息队列的处理机制
  - java: Swing等事件监听
- 注意: 广播链的问题=一个对象既是观察者也是被观察者 异步处理问题

```java
public abstract class Subject {
    //定义一个观察者数组
    private Vector<Observer> obsVector = new Vector<Observer>();

    //增加一个观察者
    public void addObserver(Observer o) {
        this.obsVector.add(o);
    }

    //删除一个观察者
    public void delObserver(Observer o) {
        this.obsVector.remove(o);
    }

    //通知所有观察者
    public void notifyObservers() {
        for (Observer o : this.obsVector) {
            o.update();
        }
    }
}
```

## Facade 门面模式

- 定义: Provide a unified interface to a set of interfaces in a subsystem.Facade defines a higher-level interface that makes the subsystem easier to use.（要求一个子系统的外部与其内部的通信必须通过一个统一的对象进行。门面模式提供一个高层次的接口，使得子系统更易于使用。）
- Facade门面角色 subsystem子系统角色
- 使用场景
  - 为一个复杂的模块或子系统提供一个供外界访问的接口
  - 子系统相对独立——外界对子系统的访问只要黑箱操作即可
  - 预防低水平人员带来的风险扩散
- 注意: 一个子系统可以有多个门面; 门面不参与子系统内的业务逻辑

## Memento 备忘录模式

- 定义: Without violating encapsulation,capture and externalize an object's internal state so that the object can be restored to this state later.（在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。这样以后就可将该对象恢复到原先保存的状态。）
- 使用场景
  - 需要保存和恢复数据的相关状态场景。
  - 提供一个可回滚（rollback）的操作。
  - 需要监控的副本场景中。
  - 数据库连接的事务管理就是用的备忘录模式。
- 注意: 备忘录的生命期 备忘录的性能=不要在频繁建立备份的场景中使用备忘录模式
- Originator发起人角色 Memento备忘录角色 Caretaker备忘录管理员角色
- clone方式备忘录: 发起人角色融合了发起人角色和备忘录角色，具有双重功效
- 多状态的备忘录
- 多备份的备忘录

```java

// 多状态的备忘录: 增加BeanUtils类
public class BeanUtils {
    //把bean的所有属性及数值放入到Hashmap中
    public static HashMap<String, Object> backupProp(Object bean) {
        HashMap<String, Object> result = new HashMap<String, Object>();
        try {
            beanInfo = Introspector.getBeanInfo(bean.getClass()); //获得Bean描述 BeanInfo
            descriptors = beanInfo.getPropertyDescriptors(); //获得属性描述 PropertyDescriptor[]
            for (PropertyDescriptor des : descriptors) { //遍历所有属性
                String fieldName = des.getName(); //属性名称
                Method getter = des.getReadMethod(); //读取属性的方法
                Object fieldValue = getter.invoke(bean, new Object[]{}); //读取属性值
                if (!fieldName.equalsIgnoreCase("class")) {
                    result.put(fieldName, fieldValue);
                }
            }
        } catch (Exception e) {//异常处理
        }
        return result;
    }

    //把HashMap的值返回到bean中
    public static void restoreProp(Object bean, HashMap<String, Object> propMap) {
        try {
            BeanInfo beanInfo = Introspector.getBeanInfo(bean.getClass()); //获得Bean描述
            PropertyDescriptor[] descriptors = beanInfo.getPropertyDescriptors(); //获得属性描述
            for (PropertyDescriptor des : descriptors) { //遍历所有属性
                String fieldName = des.getName(); //属性名称
                if (propMap.containsKey(fieldName)) { //如果有这个属性
                    Method setter = des.getWriteMethod(); //写属性的方法
                    setter.invoke(bean, new Object[]{propMap.get(fieldName)});
                }
            }
        } catch (Exception e) { //异常处理
            System.out.println("shit");
            e.printStackTrace();
        }
    }
}
```

## Visitor 访问者模式

- 定义: Represent an operation to be performed on the elements of an object structure. Visitor lets you define a new operation without changing the classes of the elements on which it operates. （封装一些作用于某种数据结构中的各元素的操作，它可以在不改变数据结构的前提下定义作用于这些元素的新的操作。）
- Visitor抽象访问者(visit() 可以访问哪些元素) ConcreteVisitor具体访问者 Element抽象元素 ConcreteElement具体元素
- 使用场景
  - 一个对象结构包含很多类对象，它们有不同的接口，而你想对这些对象实施一些依赖于其具体类的操作，也就说是用迭代器模式已经不能胜任的情景
  - 需要对一个对象结构中的对象进行很多不同并且不相关的操作，而你想避免让这些操作“污染”这些对象的类

## 状态模式

- 定义: Allow an object to alter its behavior when its internal state changes.The object will appear to change its class.（当一个对象内在状态改变时允许其改变行为，这个对象看起来像改变了其类。）
- State抽象状态角色 ConcreteState具体状态角色 Context环境角色
接口或抽象类，负责对象状态定义，并且封装环境角色以实现状态切换。
- 使用场景
  - 行为随状态改变而改变的场景, 如权限设计
  - 条件、分支判断语句的替代者
- 注意: 行为=状态 状态<=5

## Interpreter 解释器模式(少用)

- 定义: Given a language, define a representation for its grammar along with an interpreter that uses the representation to interpret sentences in the language.（给定一门语言，定义它的文法的一种表示，并定义一个解释器，该解释器使用该表示来解释语言中的句子。）
- AbstractExpression抽象解释器 TerminalExpression终结符表达式 NonterminalExpression非终结符表达式 Context环境角色
- 使用场景: 重复发生的问题可以使用解释器模式 一个简单语法需要解释的场景
- 注意: 尽量不要在重要的模块中使用解释器模式 使用脚本语言(shell jruby groovy)代替

## Flyweight 享元模式

- 定义: Use sharing to support large numbers of fine-grained objects efficiently.（使用共享对象可有效地支持大量的细粒度的对象。）
- 对象信息=intrinsic内部状态(可共享)+extrinsic外部状态(不可共享=随env改变=对外依赖)
- Flyweight抽象享元角色 ConcreteFlyweight具体享元角色 unsharedConcreteFlyweight不可共享的享元角色 FlyweightFactory享元工厂
- 使用场景
  - 系统中存在大量的相似对象
  - 细粒度的对象都具备较接近的外部状态，而且内部状态与环境无关，也就是说对象没有特定身份
  - 需要缓冲池的场景
- 注意
  - 享元模式是线程不安全的，只有依靠经验，在需要的地方考虑一下线程安全，在大部分场景下不用考虑。对象池中的享元对象尽量多，多到足够满足为止
  - 性能安全：外部状态最好以java的基本类型作为标志，如String，int，可以提高效率。

```java
// 享元工厂的代码
public class FlyweightFactory {
    //定义一个池容器
    private static HashMap<String, Flyweight> pool = new HashMap<String, Flyweight>();

    //享元工厂
    public static Flyweight getFlyweight(String Extrinsic) {
        Flyweight flyweight = null; //需要返回的对象
        if (pool.containsKey(Extrinsic)) { //在池中没有该对象
            flyweight = pool.get(Extrinsic);
        } else { //根据外部状态创建享元对象
            flyweight = new ConcreteFlyweight1(Extrinsic);
            pool.put(Extrinsic, flyweight); //放置到池中
        }
        return flyweight;
    }
}
```

## Bridge 桥梁模式

- 定义: Decouple an abstraction from its implementation so that the two can vary independently.（将抽象和实现解耦，使得两者可以独立地变化。）
- Abstraction抽象化角色 Implementor实现化角色 RefinedAbstraction修正抽象化角色 ConcreteImplementor具体实现化角色
- 使用场景: 不希望或不适用使用继承的场景 接口或抽象类不稳定的场景 重用性要求较高的场景
- 注意: 发现类的继承有N层时，可以考虑使用桥梁模式。桥梁模式主要考虑如何拆分抽象和实现
