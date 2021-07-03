# java| Java多线程编程入门

## 1. 多线程编程基础

### 1.1 进程、线程

1.1.1 进程

狭义：进程是正在运行的程序的实例。

广义：进程是一个具有一定独立功能的程序，关于某个数据集合的一次运行活动。

进程是操作系统动态执行的基本单元，在传统的操作系统中， 进程即是基本的分配单元，也是基本的执行单元。

1.1.2 线程

线程是操作系统能够进行运算调试的最小单位。它被包含在进程中，是进程中的实际动作单位。一个线程指的是进程中的一个单一顺序的控制流，一个进程中可以并发多个线程，每个线程执行不同的任务。

1.1.3 多线程的优点

1. 可以把占据时间较长的任务放到后台去处理
2. 程序的运行速度加快

### 1.2 使用多线程

1.2.1 继承Thread

步骤：

1. 创建一个类，这个类需要继承Thread
2. 重写Thread类的run方法（业务代码）
3. 实例化创建好的线程类
4. 调用实例化对象的start方法启动线程

```java
package chap1;

public class Demo02 {
    public static void main(String[] args) {
        Thread t = new Demo02Thread();
        t.start();  // 启动线程
        System.out.println("运行了main方法");
    }
}

class Demo02Thread extends Thread{
    @Override
    public void run() {
        System.out.println("运行了run方法");
    }
}
```

线程运行具有以下特点

1. 随机性

```java
package chap1;

public class Demo03 {
    public static void main(String[] args) {
        Thread t = new Demo03Thread();
        t.start();
        try {
            for (int i = 0; i < 10; i++) {
                System.out.println("运行main方法");
                Thread.sleep(100);
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo03Thread extends Thread{
    @Override
    public void run() {
        try {
            for (int i = 0; i < 10; i++) {
                System.out.println("运行run方法");
                Thread.sleep(100);
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
```

在多线程编程中，代码的执行结果与代码的执行顺序或调用顺序是无关的。线程是一个子任务，CPU以不确定的方式或者是以随机的时间来调用线程中的run方法。

如果直接调用线程对象的run方法，不是启动线程，而是由main主线程来调用run方法。

2. start的执行顺序与线程的启动顺序不一致

```java
package chap1;

public class Demo04 {
    public static void main(String[] args) {
        Thread t1 = new Demo04Thread(1);
        Thread t2 = new Demo04Thread(2);
        Thread t3 = new Demo04Thread(3);
        Thread t4 = new Demo04Thread(4);
        Thread t5 = new Demo04Thread(5);
        t1.start();
        t2.start();
        t3.start();
        t4.start();
        t5.start();
    }
}

class Demo04Thread extends Thread{
    private int val;
    public Demo04Thread(int val){
        this.val = val;
    }

    @Override
    public void run() {
        System.out.println("val=" + val);
    }
}
```

1.2.2 实现Runnable接口

步骤：

1. 创建一个类，这个类需要实现Runnable接口
2. 重写Runnable接口的run方法
3. 实例化创建的这个类
4. 实例化一个Thread对象，并把第3步创建的对象通过Thread的构造方法进行传递
5. 调用Thread对象的start方法

```java
package chap1;

public class Demo05 {
    public static void main(String[] args) {
        Runnable r = new Demo05Thread();
        Thread t = new Thread(r);
        t.start();
        System.out.println("运行了main方法");
    }
}

class Demo05Thread implements Runnable{
    @Override
    public void run() {
        System.out.println("运行了run方法");
    }
}
```

使用Thread继承的方式开发多线程应用程序是有局限的，Java是单继承，为了改变这种限抽，可以使用Runnable接口方式来实现多线程。

1.2.3 成员变量与线程安全

自定义线程类中的成员变量针对其它线程可以分为共享与不共享，这个多线程之间交互时很重要的一个技术点

1. 不共享数据

```java
package chap1;

public class Demo06 {
    public static void main(String[] args) {
        Thread t1 = new Demo06Thread();
        Thread t2 = new Demo06Thread();
        Thread t3 = new Demo06Thread();
        t1.start();
        t2.start();
        t3.start();
    }
}

class Demo06Thread extends Thread{
    private int count = 5;

    @Override
    public void run() {
        while(count > 0) {
            count--;
            System.out.println(Thread.currentThread().getName() + ": count=" + count);
        }
    }
}
```

每个线程都有各自的count变量，自己减少自己的count变量的值，这种情况就是不共享

2. 共享数据

```java
package chap1;

public class Demo07 {
    public static void main(String[] args) {
        Thread t = new Demo07Thread();
        Thread t1 = new Thread(t);
        Thread t2 = new Thread(t);
        Thread t3 = new Thread(t);
        Thread t4 = new Thread(t);
        Thread t5 = new Thread(t);
        t1.start();
        t2.start();
        t3.start();
        t4.start();
        t5.start();
    }
}

class Demo07Thread extends Thread {
    private int count = 5;

    @Override
    public void run() {
        count--;
        System.out.println(Thread.currentThread().getName() + ", count=" + count);
    }
}
```

共享数据有概率出现不同线程使用相同的count的值，产生了『非线程安全问题』。

### 1.3 线程常用API

1.3.1 currentThread方法

返回代码正在那个线程调用的详细信息

```java
package chap1;

public class Demo08 {
    public static void main(String[] args) {
        Thread t = new Demo08Thread();
        t.start();
        System.out.println("main方法：" + Thread.currentThread().getName());
    }
}

class Demo08Thread extends Thread{
    public Demo08Thread(){
        System.out.println("构造方法：" + Thread.currentThread().getName());
    }

    @Override
    public void run() {
        System.out.println("run方法" + Thread.currentThread().getName());
    }
}
```

main方法被命名为main的线程调用，线程类的构造方法是被main线程调用，run方法是被Thread-0的线程调用，run方法是自动被调用的方法。

1.3.2 isAlive

是判断当前的线程是否处于活动状态

```java
package chap1;

public class Demo09 {
    public static void main(String[] args) {
        Thread t = new Demo09Thread();
        System.out.println("准备开始启动线程：" + t.isAlive());
        t.start();  // 启动线程
        System.out.println("启动线程后：" + t.isAlive());
    }
}

class Demo09Thread extends Thread{
    @Override
    public void run() {
        System.out.println("run方法运行状态：" + this.isAlive());
    }
}
```

活动状态就是线程已经启动且运行没有结束。线程处于正在运行或准备开始运行的状态，就认为线程是『存活』的状态。

1.3.3 sleep方法

作用是在指定的毫秒数内让当前『正在执行的线程』暂停执行。

```java
package chap1;

import java.sql.SQLOutput;

public class Demo10 {
    public static void main(String[] args) {
        Thread t = new Demo10Thread();
        System.out.println("开始于" + System.currentTimeMillis());
        t.start();
        System.out.println("结束于" + System.currentTimeMillis());
    }
}

class Demo10Thread extends Thread{
    @Override
    public void run() {
        try {
            System.out.println("线程开始于" + System.currentTimeMillis());
            Thread.sleep(2000);
            System.out.println("线程结束于" + System.currentTimeMillis());
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
```

1.3.4 getId方法

作用获取当前线程的唯一标识

```java
package chap1;

public class Demo11 {
    public static void main(String[] args) {
        Thread t = Thread.currentThread();
        System.out.println(t.getName() + ", " + t.getId());
    }
}
```

### 1.4 停止线程

停止一个线程意味着在线程处理完成任务之前结束正在执行的操作。

1. 使用退出标志，使线程正常退出，也就是当run方法完成后线程终止。

```java
package chap1;

public class Demo12 {
    public static void main(String[] args) throws InterruptedException {
        Demo12Thread t = new Demo12Thread();
        t.start();
        Thread.sleep(2000);
        t.stopThread();
    }
}

class Demo12Thread extends Thread{
    private boolean flag = true;

    @Override
    public void run() {
        try {
            while (flag) {
                System.out.println("time=" + System.currentTimeMillis());
                Thread.sleep(1000);
            }
            System.out.println("线程执行结束");
        }catch(InterruptedException e){
            e.printStackTrace();
        }
    }

    public void stopThread(){
        flag = false;
    }
}
```

2. stop方法强制结束线程

```java
package chap1;

public class Demo13 {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Demo13Thread();
        t.start();
        Thread.sleep(2000);
        t.stop();
    }
}

class Demo13Thread extends Thread{
    @Override
    public void run() {
        try {
            while (true) {
                System.out.println("time=" + System.currentTimeMillis());
                Thread.sleep(1000);
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }catch (ThreadDeath e){
            System.out.println("进入catch块");
            e.printStackTrace();
        }
    }
}
```

stop方法已经被作废，如果强制让线程停止有可能使一些清理性的工作得不到完成。另外一个原因是对锁定的对象进行『解锁』，导致数据得不同同步的处理，出现数据不一致性的问题。

```java
package chap1;

public class Demo14 {
    public static void main(String[] args) throws InterruptedException {
        Demo14User user = new Demo14User();
        Thread t = new Demo14Thread(user);
        t.start();
        Thread.sleep(500);
        t.stop();
        System.out.println("username=" + user.getUsername() + ", password=" + user.getPassword());
    }
}

class Demo14Thread extends Thread{
    private Demo14User user;
    public Demo14Thread(Demo14User user){
        this.user = user;
    }

    @Override
    public void run() {
        user.updateUsernameAndPassword("b", "bb");
    }
}

class Demo14User {
    private String username = "a";
    private String password = "aa";
    public void setUsername(String username){
        this.username = username;
    }
    public String getUsername(){
        return username;
    }
    public void setPassword(String passpword){
        this.password = passpword;
    }
    public String getPassword(){
        return password;
    }
    synchronized public void updateUsernameAndPassword(String username, String password){
        try {
            this.username = username;
            Thread.sleep(10000);
            this.password = password;
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
```

3. 使用interrupt方法中断线程

```java
package chap1;

public class Demo15 {
    public static void main(String[] args) {
        Thread t = new Demo15Thread();
        t.start();
        t.interrupt();
    }
}

class Demo15Thread extends Thread{
    @Override
    public void run() {
        for (int i = 0; i < 10000; i++) {
            System.out.println("i=" + i);
        }
    }
}
```

调用interrupt方法不会真正的结束线程，在当前线程中打上一个停止的标记。

Thread类提供了interrupted方法测试当前线程是否中断，isInterrupted方法测试线程是否已经中断。

```java
package chap1;

public class Demo15 {
    public static void main(String[] args) {
        Thread t = new Demo15Thread();
        t.start();
        t.interrupt();
        System.out.println("是否停止1？" + t.isInterrupted());
        System.out.println("是否停止2？" + Thread.interrupted());
    }
}

class Demo15Thread extends Thread{
    @Override
    public void run() {
        for (int i = 0; i < 10000; i++) {
            System.out.println("i=" + i);
        }
    }
}
```

t.isInterrupted方法是检查Demo15Thread是否被打上停止的标记，Thread.interrupted方法是检查主线程是否被打上停止的标记。

```java
package chap1;

public class Demo16 {
    public static void main(String[] args) {
        Thread.currentThread().interrupt();
        System.out.println("是否停止1？" + Thread.interrupted());
        System.out.println("是否停止2？" + Thread.interrupted());
    }
}
```

测试当前线程是否已经中断，线程的中断状态由方法清除。如果 连接两次调用该方法，第二次调用将返回false。

```java
package chap1;

public class Demo17 {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Demo17Thread();
        t.start();
//        Thread.sleep(10);
        t.interrupt();
    }
}

class Demo17Thread extends Thread{
    @Override
    public void run() {
        try {
            for (int i = 0; i < 10000; i++) {
                if (this.isInterrupted()) {
                    System.out.println("已经是停止状态了！我有退出");
//                break;
                    throw new InterruptedException();
                }
                System.out.println("i=" + i);
            }
            System.out.println("这里是结束循环后的代码");
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
```

```java
package chap1;

public class Demo18 {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Demo18Thread();
        t.start();
        Thread.sleep(10);
        t.interrupt();
    }
}

class Demo18Thread extends Thread{
    @Override
    public void run() {
        try {
            for (int i = 0; i < Integer.MAX_VALUE; i++) {
                String s = new String();
            }
            System.out.println("开始线程");
            Thread.sleep(20000);
            System.out.println("结束线程");
        }catch (InterruptedException e){
            System.out.println("异常进入到catch代码块");
            e.printStackTrace();
        }
    }
}
```

如果程中有sleep代码，不管是否进入到sleep的状态，如果调用了interrupt方法都会产生异常信息。

### 1.5 暂停线程

暂停线程使用suspend，重启暂停线程使用resume方法

```java
package chap1;

public class Demo19 {
    public static void main(String[] args) throws InterruptedException {
        Demo19Thread t = new Demo19Thread();
        t.start();
        Thread.sleep(1000);

        t.suspend();    // 暂停线程
        System.out.println("A=" + System.currentTimeMillis() + ", i=" + t.getI());
        Thread.sleep(1000);
        System.out.println("A=" + System.currentTimeMillis() + ", i=" + t.getI());

        t.resume();     // 恢复暂停线程运行
        Thread.sleep(1000);

        t.suspend();
        System.out.println("B=" + System.currentTimeMillis() + ", i=" + t.getI());
        Thread.sleep(1000);
        System.out.println("B=" + System.currentTimeMillis() + ", i=" + t.getI());
    }
}

class Demo19Thread extends Thread{
    private long i = 0;

    public long getI(){
        return i;
    }

    public void setI(long i){
        this.i = i;
    }

    @Override
    public void run() {
        while (true){
            i++;
        }
    }
}
```

suspend如果独占公共的同步对象，使其它线程无法访问公共同步对象

```java
package chap1;

public class Demo20 {
    public static void main(String[] args) throws InterruptedException {
        Demo20Service service = new Demo20Service();
        Thread t1 = new Thread(){
            @Override
            public void run() {
                service.printString();
            }
        };
        t1.setName("A");    // setName修改线程的名称，这里就是把Thread-0修改为A
        t1.start();
        Thread.sleep(10);

        Thread t2 = new Thread(){
            @Override
            public void run() {
                service.printString();
            }
        };
        t2.start();
    }
}

class Demo20Service {
    synchronized public void printString(){
        System.out.println("线程开始");
        if ("A".equals(Thread.currentThread().getName())){
            System.out.println("A线程永远suspend了");
            Thread.currentThread().suspend();
        }
        System.out.println("线程结束");
    }
}
```

suspend下println方法不会释放同步锁

```java
package chap1;

public class Demo21 {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Demo21Thread();
        t.start();
        Thread.sleep(100);
        t.suspend();
        System.out.println("main结束");
    }
}

class Demo21Thread extends Thread{
    private long i = 0;

    @Override
    public void run() {
        while(true){
            i++;
            System.out.println("i=" + i);
        }
    }
}
```

程序运行到println方法内部时，同步锁没有被释放，导致当前的PrintStream对象的println方法一直呈『暂停』状态。

suspend会造成共享对象数据不同步

```java
package chap1;

public class Demo22 {
    public static void main(String[] args) throws InterruptedException {
        Demo22User user = new Demo22User();
        Thread t1 = new Thread(){
            @Override
            public void run() {
                user.updateUsernameAndPassword("b", "bb");
            }
        };
        t1.setName("A");
        t1.start();

        Thread.sleep(10);

        Thread t2 = new Thread(){
            @Override
            public void run() {
                user.printUseruserAndPassword();
            }
        };
        t2.start();
    }
}

class Demo22User {
    private String username = "a";
    private String password = "aa";

    public void updateUsernameAndPassword(String username, String password){
        this.username = username;
        if ("A".equals(Thread.currentThread().getName())){
            System.out.println("停止A线程");
            Thread.currentThread().suspend();
        }
        this.password = password;
    }

    public void printUseruserAndPassword(){
        System.out.println("username=" + username + ", password=" +password);
    }
}
```

### 1.6 yield方法

yield方法的作用是放弃当前的CPU资源，将资源让给其它的任务去占用CPU执行。但放弃时间不确定，有可能刚刚放弃，马上又获取CPU时间片。

```java
package chap1;

public class Demo23 {
    public static void main(String[] args) {
        Thread t1 = new Demo23Thread();
        t1.start();
    }
}

class Demo23Thread extends Thread{
    @Override
    public void run() {
        int count = 0;
        long start = System.currentTimeMillis();
        for (int i = 0; i < 500000; i++) {
            Thread.yield();
            count  += i + 1;
        }
        long end = System.currentTimeMillis();
        System.out.println("用时：" + (end - start) + "毫秒");
    }
}
```

### 1.7 线程的优先级

在操作系统中，线程可以划分优先级，优先级较高的线程得到更多的CPU资源，也就CPU会优先执行优先级较高的线程对象中的任务。设置线程优先有助于帮助『线程调度器』确定在下一次选择哪个线程优先执行。

设置线程的优先级使用setPriority方法，优级分为1~10个级别，如果设置优先级小于1或大于10，JDK抛出IllegalArgumentException。JDK默认设置3个优先级常量，MIN_PRIORITY=1(最小值)，NORM_PRIORITY=5(中间值，默认)，MAX_PRIORITY=10(最大值)。

获取线程的优先级使用getPriority方法。

```java
package chap1;

public class Demo24 {
    public static void main(String[] args) {
        System.out.println("主线程的运行优先级是：" + Thread.currentThread().getPriority());
        Thread.currentThread().setPriority(1);
        System.out.println("主线程的运行优先级是：" + Thread.currentThread().getPriority());
        Thread t = new Demo24Thread();
        t.start();
    }
}

class Demo24Thread extends Thread{
    @Override
    public void run() {
        System.out.println("线程的运行优先级是：" + this.getPriority());
    }
}
```

线程的优先级具有继承性，比如线程A启动线程B，线程B的优先级与线程A是一样的。

```java
package chap1;

import java.util.Random;

public class Demo25 {
    public static void main(String[] args) {
        for (int i = 0; i < 5; i++) {
            Thread t1 = new Demo25Thread();
            t1.setName("A" + i);
            t1.setPriority(Thread.MAX_PRIORITY);
            t1.start();

            Thread t2 = new Demo25Thread();
            t2.setName("B" + i);
            t2.setPriority(Thread.MIN_PRIORITY);
            t2.start();
        }
    }
}

class Demo25Thread extends Thread{
    @Override
    public void run() {
        long start = System.currentTimeMillis();
        long count = 0;
        for (int i = 0; i < 10; i++) {
            for (int j = 0; j < 50000; j++) {
                Random r = new Random();
                count += i * j + r.nextInt();;
            }
        }
        long end = System.currentTimeMillis();
        System.out.println("线程" + this.getName() + "执行完成使用了" + (end  - start) + "毫秒");
    }
}
```

高优先级的线程总是大部分先执行完，但不代表高优先级的线程全部执行完。当线程优先级的等级差距很大时，谁先执行完和代码的调用顺序无关。

线程的优先还有『随机性』，也就是说优先级高的线不一定每一次都先执行完成。

### 1.8 守护线程

在Java线程中有两种线程，一种是用户线程，另一种是守护线程。

守护线程是一种特殊的线程，特殊指的是当进程中不存在用户线程时，守护线程会自动销毁。典型的守护线程的例子就是垃圾回收线程，当进程中没有用户线程，垃圾回收线程就没有存在的必要了，会自动销毁。

```java
package chap1;

public class Demo26 {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Demo26Thread();
        t.setDaemon(true);  // 把指定的线程设置为守护线程
        t.start();
        Thread.sleep(3000);
        System.out.println("主线程结束");
    }
}

class Demo26Thread extends Thread{
    @Override
    public void run() {
        try {
            while (true) {
                System.out.println("time=" + System.currentTimeMillis());
                Thread.sleep(1000);
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
```

设置守护线程必须要在调用start方法之前设置，否则JDK会产生IllegalThreadStateException

## 2. 线程的同步机制

Java多线程中的同步，如何在Java语言中开发出线程安全的程序，如何在Java语文中解决非线程安全所带来的相关问题。

### 2.1 synchronized同步方法

『线程安全』与『非线程安全』是学习多线程技术时一定会遇到的经典问题。『非线程安全』其实当多个线程访问同一个对象中的成员变量时产生的，产生的后果就是『脏读』，就是取到的数据其实是被更改过的。而『线程安全』就是以获取的成员变量的值是经过同步处理的，不会出现脏读的现象。

1. 局部变量是线程安全的

```java
package chap2;

public class Demo01 {
    public static void main(String[] args) {
        Demo01Service service = new Demo01Service();
        Thread t1 = new Demo01ThreadA(service);
        t1.start();
        Thread t2 = new Demo01ThreadB(service);
        t2.start();
    }
}

class Demo01Service{
    public void add(String username){
        int num = 0;
        if ("a".equals(username)){
            num = 100;
            System.out.println("a set over");
            // 等待另外一个线程修改num的值
            try {
                Thread.sleep(2000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }else {
            num = 200;
            System.out.println("b set over");
        }
        System.out.println("username=" + username + ", num=" + num);
    }
}

class Demo01ThreadA extends Thread{
    private Demo01Service service;
    public Demo01ThreadA(Demo01Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.add("a");
    }
}
class Demo01ThreadB extends Thread{
    private Demo01Service service;
    public Demo01ThreadB(Demo01Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.add("b");
    }
}
```

局部变量不存在非线程安全问题，永远都是线程安全的。这是由局部变量是私有的特征所造成的。

2. 成员变量不是线程安全

```java
package chap2;

public class Demo02 {
    public static void main(String[] args) {
        Demo02Service service = new Demo02Service();
        Thread t1 = new Demo02ThreadA(service);
        Thread t2 = new Demo02ThreadB(service);
        t1.start();
        t2.start();
    }
}

class Demo02Service {
    private int num;
    synchronized public void add(String username){
        if ("a".equals(username)){
            num = 100;
            System.out.println("a is set");
            try {
                Thread.sleep(2000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }else{
            num = 200;
            System.out.println("b is set");
        }
        System.out.println("username=" + username + ", num=" + num);
    }
}

class Demo02ThreadA extends Thread{
    private Demo02Service service;
    public Demo02ThreadA(Demo02Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.add("a");
    }
}

class Demo02ThreadB extends Thread{
    private Demo02Service service;
    public Demo02ThreadB(Demo02Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.add("b");
    }
}
```

如果有两个线程同时操作业务对象中的成员变量，可能会产生『非线程安全』问题，需要在方法前使用关键字synchronized进行修饰。

3. 多个对象使用多个对象锁

```java
package chap2;

public class Demo03 {
    public static void main(String[] args) {
        Demo03Service service1 = new Demo03Service();
        Demo03Service service2 = new Demo03Service();

        Thread t1 = new Demo03ThreadA(service1);
        Thread t2 = new Demo03ThreadB(service2);
        t1.start();
        t2.start();
    }
}

class Demo03Service {
    private int num;
    synchronized public void add(String username){
        if ("a".equals(username)){
            num = 100;
            System.out.println("a is set");
            try {
                Thread.sleep(2000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }else{
            num = 200;
            System.out.println("b is set");
        }
        System.out.println("username=" + username + ", num=" + num);
    }
}

class Demo03ThreadA extends Thread{
    private Demo03Service service;
    public Demo03ThreadA(Demo03Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.add("a");
    }
}

class Demo03ThreadB extends Thread{
    private Demo03Service service;
    public Demo03ThreadB(Demo03Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.add("b");
    }
}
```

synchronized取得的锁都是对象锁，而不是把一段代码或方法作为锁，所以那个线程先执行带synchronized关键字修饰的方法，哪个方法就持有该方法所属对象的锁，其它线程只能呈等待状态，前提是多个线程访问的是同一个对象。如果多个线程访问多个对象，JVM会创建出多个对象锁。

4. synchronize方法锁的是整个对象

```java
package chap2;

public class Demo04 {
    public static void main(String[] args) {
        Demo04Service service = new Demo04Service();
        Thread t1 = new Demo04ThreadA(service);
        Thread t2 = new Demo04ThreadB(service);
        t1.start();
        t2.start();
    }
}

class Demo04Service {
    synchronized public void foo1(){
        System.out.println("开始运行foo1方法，threadName：" + Thread.currentThread().getName());
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println("foo1方法运行结束");
    }

    synchronized public void foo2(){
        System.out.println("开始运行foo2方法，threadName：" + Thread.currentThread().getName());
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println("foo2方法运行结束");
    }
}

class Demo04ThreadA extends Thread{
    private Demo04Service service;
    public Demo04ThreadA(Demo04Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo1();
    }
}

class Demo04ThreadB extends Thread{
    private Demo04Service service;
    public Demo04ThreadB(Demo04Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo2();
    }
}
```

A线程先持有Object对象的对象锁，B线程就不可以异步方式调用Object对象使用Synchronize修饰的方法，线程B只能等线程A的方法执行完成释放对象锁才能够执行，也就是同步执行。B线程可以以异步的方式调用Object对象没有使用synchronize修改的方法。

5. 脏读

```java
package chap2;

public class Demo05 {
    public static void main(String[] args) throws InterruptedException {
        Demo05User user = new Demo05User();
        Thread t = new Demo05Thread(user);
        t.start();
        Thread.sleep(200);
        user.getValue();
    }
}

class Demo05User{
    private String username = "a";
    private String password = "aa";

    synchronized public void setUsernameAndPassword(String username, String password){
        this.username = username;
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        this.password = password;
        System.out.println("setUsernameAndPassword方法，线程名称：" + Thread.currentThread().getName() +
                "，username=" +username + ", password=" + password);
    }

    synchronized public void getValue(){
        System.out.println("getValue方法，线程名称" + Thread.currentThread().getName() +
                ", username=" + username + ", password=" + password );
    }
}

class Demo05Thread extends Thread{
    private Demo05User user;
    public Demo05Thread(Demo05User user){
        this.user = user;
    }

    @Override
    public void run() {
        user.setUsernameAndPassword("B", "BB");
    }
}
```

当A线程调用使用synchronize关键字修饰的方法时，A线程就获取到一个方法锁，准确来说是获取到对象锁，所以其它线程必须等A线程执行完毕后才可以调用其它使用synchrnoize修改的方法。这时A线程已经执行完一个完整任务，也就是说username和password两个成员变量已经被同时赋值，所以不存在脏读的。

6. 锁重入

关键字synchronized拥有锁重入的功能，就是说在使用synchronized时，当一个线程得到一个对象锁后，再次请求些对象锁时是可以再次得到该对象锁。

```java
package chap2;

public class Demo06 {
    public static void main(String[] args) {
        Thread t = new Demo06Thread();
        t.start();
    }
}

class Demo06Service{
    synchronized public void foo1(){
        System.out.println("foo1方法");
        foo2();
    }
    synchronized public void foo2(){
        System.out.println("foo2方法");
        foo3();
    }
    synchronized public void foo3(){
        System.out.println("foo3");
    }
}

class Demo06Thread extends Thread{
    @Override
    public void run() {
        Demo06Service service = new Demo06Service();
        service.foo1();
    }
}
```

可重入锁也支持父子类继承的环境中。

```java
package chap2;

public class Demo06 {
    public static void main(String[] args) {
        Thread t = new Demo06Thread();
        t.start();
    }
}

class Demo06Service{
    synchronized public void foo1(){
        System.out.println("foo1方法");
        foo2();
    }
    synchronized public void foo2(){
        System.out.println("foo2方法");
        foo3();
    }
    synchronized public void foo3(){
        System.out.println("foo3");
    }
}

class Demo06ServiceB extends Demo06Service{
    synchronized public void foo4(){
        System.out.println("foo4方法");
        super.foo1();
    }
}

class Demo06Thread extends Thread{
    @Override
    public void run() {
//        Demo06Service service = new Demo06Service();
//        service.foo1();
        Demo06ServiceB serviceB = new Demo06ServiceB();
        serviceB.foo4();
    }
}
```

7. 锁的自动释放

当一个线程执行的代码出现了异常，其所持有的锁会自动释放。

```java
package chap2;

public class Demo07 {
    public static void main(String[] args) {
        Demo07Service service = new Demo07Service();
        Thread t1 = new Demo07Thread(service);
        t1.setName("A");
        t1.start();
        Thread t2 = new Demo07Thread(service);
        t2.start();
    }
}

class Demo07Service {
    synchronized public void foo(){
        if ("A".equals(Thread.currentThread().getName())){
            System.out.println("线程A开始于" + System.currentTimeMillis());
            while(true){
                if(("" + Math.random()).substring(0, 8).equals("0.123456")){
                    System.out.println("线程A结束于" + System.currentTimeMillis());
                    Integer.parseInt("A");
                }
            }
        }else{
            System.out.println("线程B开始于" + System.currentTimeMillis());
        }
    }
}

class Demo07Thread extends Thread{
    private Demo07Service service;
    public Demo07Thread(Demo07Service service){
        this.service = service;
    }
    @Override
    public void run() {
        service.foo();
    }
}
```

8. 同步不具有继承性

```java
package chap2;

public class Demo08 {
    public static void main(String[] args) {
        Demo08ServiceB serviceB = new Demo08ServiceB();
        Thread t1 = new Demo08Thread(serviceB);
        t1.setName("A");
        t1.start();

        Thread t2 = new Demo08Thread(serviceB);
        t2.setName("B");
        t2.start();
    }
}

class Demo08ServiceA{
    synchronized public void foo(){
        try {
            System.out.println("父类：" + Thread.currentThread().getName() + "，开始于" + System.currentTimeMillis());
            Thread.sleep(5000);
            System.out.println("父类：" + Thread.currentThread().getName() + "，结束于" + System.currentTimeMillis());
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo08ServiceB extends Demo08ServiceA{
    public void foo(){
        try {
            System.out.println("子类：" + Thread.currentThread().getName() + "，开始于" + System.currentTimeMillis());
            Thread.sleep(5000);
            System.out.println("子类：" + Thread.currentThread().getName() + "，结束于" + System.currentTimeMillis());
            super.foo();
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo08Thread extends Thread{
    Demo08ServiceB service;
    public Demo08Thread(Demo08ServiceB service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo();
    }
}
```

### 2.2 synchronized同步语句块

1. synchronized方法的缺点

```java
package chap2;

public class Demo09 {
    public static void main(String[] args) throws InterruptedException {
        Demo09Service service = new Demo09Service();
        Thread t1 = new Demo09ThreadA(service);
        t1.setName("A");
        Thread t2 = new Demo09ThreadB(service);
        t2.setName("B");

        t1.start();
        t2.start();

        Thread.sleep(11000);

        long start = Demo09Utils.start1 > Demo09Utils.start2 ? Demo09Utils.start2 : Demo09Utils.start1;
        long end = Demo09Utils.end1 > Demo09Utils.end2 ? Demo09Utils.end1 : Demo09Utils.end2;
        System.out.println("总耗时：" + (end - start) / 1000 + "秒");
    }
}

class Demo09Utils{
    static long start1;
    static long start2;
    static long end1;
    static long end2;
}

class Demo09Service{
    synchronized public void foo(){
        try{
            System.out.println("开始任务");
            Thread.sleep(5000);
            System.out.println("长时任务处理完成，线程" + Thread.currentThread().getName());
            System.out.println("结束任务");
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo09ThreadA extends Thread{
    public Demo09Service service;
    public Demo09ThreadA(Demo09Service service){
        this.service = service;
    }

    @Override
    public void run() {
        Demo09Utils.start1 = System.currentTimeMillis();
        service.foo();
        Demo09Utils.end1 = System.currentTimeMillis();
    }
}

class Demo09ThreadB extends Thread{
    private Demo09Service service;
    public Demo09ThreadB(Demo09Service service){
        this.service = service;
    }

    @Override
    public void run() {
        Demo09Utils.start2 = System.currentTimeMillis();
        service.foo();
        Demo09Utils.end2 = System.currentTimeMillis();
    }
}
```

2. synchronized同步代码块使用

```java
package chap2;

public class Demo10 {
    public static void main(String[] args) {
        Demo10Service service = new Demo10Service();
        Thread t1 = new Demo10Thread(service);
        t1.setName("A");
        t1.start();
        Thread t2 = new Demo10Thread(service);
        t2.setName("B");
        t2.start();
    }
}

class Demo10Service{
    public void foo(){
        try {
            synchronized (this) {
                System.out.println(Thread.currentThread().getName() + "开始于" + System.currentTimeMillis());
                Thread.sleep(2000);
                System.out.println(Thread.currentThread().getName() + "结束于" + System.currentTimeMillis());
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
class Demo10Thread extends Thread{
    private Demo10Service service;
    public Demo10Thread(Demo10Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo();
    }
}
```

3. 使用同步代码块解决同块方法的问题

```java
package chap2;

public class Demo11 {
    public static void main(String[] args) throws InterruptedException {
        Demo11Service service = new Demo11Service();
        Thread t1 = new Demo11ThreadA(service);
        t1.setName("A");
        t1.start();

        Thread t2 = new Demo11ThreadB(service);
        t2.setName("B");
        t2.start();

        Thread.sleep(10000);
        long start = Demo11Utils.start1 > Demo11Utils.start2 ? Demo11Utils.start2 : Demo11Utils.start1;
        long end = Demo11Utils.end1 > Demo11Utils.end2 ? Demo11Utils.end1 : Demo11Utils.end2;
        System.out.println("耗时：" + (end - start) /1000 + "秒");

    }
}

class Demo11Utils{
    static long start1;
    static long start2;
    static long end1;
    static long end2;
}

class Demo11Service{
    /*synchronized*/ public void foo(){
        try {
            System.out.println(Thread.currentThread().getName() + "开始任务");
            Thread.sleep(3000);
            synchronized (this) {
                System.out.println(Thread.currentThread().getName() + "处理计算结果");
            }
            System.out.println(Thread.currentThread().getName() + "结束任务");
        }catch(InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo11ThreadA extends Thread{
    private Demo11Service service;
    public Demo11ThreadA(Demo11Service service){
        this.service = service;
    }

    @Override
    public void run() {
        Demo11Utils.start1 = System.currentTimeMillis();
        service.foo();
        Demo11Utils.end1 = System.currentTimeMillis();
    }
}

class Demo11ThreadB extends Thread{
    private Demo11Service service;
    public Demo11ThreadB(Demo11Service service){
        this.service = service;
    }

    @Override
    public void run() {
        Demo11Utils.start2 = System.currentTimeMillis();
        service.foo();
        Demo11Utils.end2 = System.currentTimeMillis();
    }
}
```

4. 半异步半同步

```java
package chap2;

public class Demo12 {
    public static void main(String[] args) {
        Demo12Service service = new Demo12Service();
        Thread t1 = new Demo12Thread(service);
        t1.setName("A");
        t1.start();

        Thread t2 = new Demo12Thread(service);
        t2.setName("B");
        t2.start();
    }
}

class Demo12Service {
    public void foo(){
        try{
            for (int i = 0; i < 100; i++) {
                System.out.println("非同步线程" + Thread.currentThread().getName() + ", i=" + i);
                Thread.sleep(10);
            }
            System.out.println();
            synchronized (this){
                for (int i = 0; i < 100; i++) {
                    System.out.println("同步线程" + Thread.currentThread().getName() + ", i=" + i);
                    Thread.sleep(10);
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo12Thread extends Thread{
    private Demo12Service service;
    public Demo12Thread(Demo12Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo();
    }
}
```

不在synchronized块中的就是异步执行，在synchroinzed块中的代码就是同步执行。

5. synchronized代码块间的同步性

```java
package chap2;

public class Demo13 {
    public static void main(String[] args) throws InterruptedException {
        Demo13Service service = new Demo13Service();
        Thread t1 = new Demo13ThreadA(service);
        t1.start();
        Thread.sleep(10);
        Thread t2 = new Demo13ThreadB(service);
        t2.start();
    }
}

class Demo13Service{
    public void foo1(){
        try {
            synchronized (this) {
                System.out.println("foo1开始于" + System.currentTimeMillis());
                Thread.sleep(2000);
                System.out.println("foo1结束于" + System.currentTimeMillis());
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }

    public void foo2(){
        synchronized (this){
            System.out.println("foo2开始于" + System.currentTimeMillis());
            System.out.println("foo2结束于" + System.currentTimeMillis());
        }
    }
}

class Demo13ThreadA extends Thread{
    private Demo13Service service;
    public Demo13ThreadA(Demo13Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo1();
    }
}

class Demo13ThreadB extends Thread{
    private Demo13Service service;
    public Demo13ThreadB(Demo13Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo2();
    }
}
```

使用同步synchronize(this)代码时需要注意，当一个线程访问object的一个synchronized(this)同步代码块时，其它线程对这个object的其它syncrhonized(this)同步的访问会被阻塞，说明synchronized使用的对象锁是同一个。

6. synchroinze(this)代码块是锁定当前对象

```java
package chap2;

public class Demo14 {
    public static void main(String[] args) throws InterruptedException {
        Demo14Service service = new Demo14Service();
        Thread t1 = new Demo14ThreadA(service);
        Thread t2 = new Demo14ThreadB(service);
        t2.start();
        Thread.sleep(10);
        t1.start();

    }
}

class Demo14Service{
    synchronized public void foo1(){
        System.out.println("foo1正在运行");
    }

    public void foo2(){
        try {
            synchronized (this) {
                System.out.println("foo2开始");
                Thread.sleep(2000);
                System.out.println("foo2结束");
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo14ThreadA extends Thread{
    private Demo14Service service;
    public Demo14ThreadA(Demo14Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo1();
    }
}

class Demo14ThreadB extends Thread{
    private Demo14Service service;
    public Demo14ThreadB(Demo14Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo2();
    }
}
```

synchronized(this)代码块与synchronized方法一样，synchronized(this)代码块也是锁定当前对象的。

7. 使用任意对象作为对象锁

除了可以使用syncrhonized(this)来同步代码块，Java还支持『任意对象』作为『对象锁』来实现同步的功能。这个『任意对象』大数是成员变量或方法的参数，使用格式synchronized(非this对象)。

1) 在多个线程持有『对象锁』为同一个对象的情况下，同一时间只有一个线程可以执行synchronized(非this对象)同步代码块中的代码。如果使用不是同一个对象锁，运行的结果就是异步调用，会交叉运行。

```java
package chap2;

public class Demo15 {
    public static void main(String[] args) {
        Demo15Service service = new Demo15Service();
        Thread t1 = new Demo15Thread(service);
        t1.setName("A");
        t1.start();
        Thread t2 = new Demo15Thread(service);
        t2.setName("B");
        t2.start();
    }
}

class Demo15Service{
    private Object lockObject = new Object();

    public void foo(){
        try {
            synchronized (lockObject) {
                System.out.println(Thread.currentThread().getName() + "开始于" + System.currentTimeMillis());
                Thread.sleep(2000);
                System.out.println(Thread.currentThread().getName() + "结束于" + System.currentTimeMillis());
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo15Thread extends Thread{
    private Demo15Service service;
    public Demo15Thread(Demo15Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo();
    }
}
```

2）非this对象具有的优点

一个类有很多个synchrnoized方法，虽然可以实现同步，受到阻塞，所以影响运行效率。如果使用同步代码块非this对象，则synchronized非this对象的代码块中的程序与同步方法是奶的，不与其它的锁(this)争抢this锁，大大提高运行的效率。

```java
package chap2;

public class Demo16 {
    public static void main(String[] args) throws InterruptedException {
        Demo16Service service = new Demo16Service();
        Thread t1 = new Demo16ThreadA(service);
        t1.start();

        Thread.sleep(10);
        Thread t2 = new Demo16ThreadB(service);
        t2.start();
    }
}

class Demo16Service{
    private Object lockObject = new Object();
    public void foo(){
        try {
            synchronized (lockObject) {
                System.out.println("foo方法开始于" + System.currentTimeMillis());
                Thread.sleep(2000);
                System.out.println("foot方法结束于" + System.currentTimeMillis());
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
    synchronized public void foo2(){
        System.out.println("foo2方法开始" + System.currentTimeMillis());
        System.out.println("foo2方法结束" + System.currentTimeMillis());
    }
}

class Demo16ThreadA extends Thread{
    private Demo16Service service;
    public Demo16ThreadA(Demo16Service service){
        this.service  =service;
    }

    @Override
    public void run() {
        service.foo();
    }
}

class Demo16ThreadB extends Thread{
    private Demo16Service service;
    public Demo16ThreadB(Demo16Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo2();
    }
}
```

3) synchronized(非this对象)解决脏读问题

```java
package chap2;

import java.util.ArrayList;
import java.util.List;

public class Demo17 {
    public static void main(String[] args) throws InterruptedException {
        Demo17List list = new Demo17List();
        Thread t1 = new Demo17ThreadA(list);
        t1.start();
        Thread t2 = new Demo17ThreadB(list);
        t2.start();
        Thread.sleep(5000);
        System.out.println("list size is " + list.size());
    }
}

class Demo17List{
    private List list = new ArrayList();
    synchronized public void add(Object obj){
        list.add(obj);
    }
    synchronized public int size(){
        return list.size();
    }
}

class Demo17Service{
    private Object lockObject = new Object();
    public void add(Demo17List list, Object obj){
        try {
            synchronized (list) {
                if (list.size() < 1) {
                    Thread.sleep(2000); // 模拟数据的获取
                    list.add(obj);
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo17ThreadA extends Thread{
    private Demo17List list;
    public Demo17ThreadA(Demo17List list){
        this.list = list;
    }

    @Override
    public void run() {
        Demo17Service service = new Demo17Service();
        service.add(list, "a");
    }
}

class Demo17ThreadB extends Thread{
    private Demo17List list;
    public Demo17ThreadB(Demo17List list){
        this.list = list;
    }

    @Override
    public void run() {
        Demo17Service service = new Demo17Service();
        service.add(list, "b");
    }
}
```

synchronized(非this对象x)格式的写法是将x对象本身作为『对象锁』，这样得出以下的结论：

a. 当多个线程同时执行syncrhonized(x)同步代码块时呈同步效果；

b. 当其它 线程执行x对象中的synchronized同步方法时呈同步效果；

c. 当其它线程执行x对象方法里的synchronize(this)代码块时也呈同步效果

注意：如果其它 线程调用不加synchronized关键字的方法时，还是异步调用

8. 静态同步synchronized方法与synchronized(class)代码块

关键字synchronized还可以应用在静态方法上，如果 这样就是对象当前的*.java文件对应的Class进行加锁

```java
package chap2;

public class Demo18 {
    public static void main(String[] args) {
        Thread t1 = new Demo18ThreadA();
        t1.setName("A");
        t1.start();

        Thread t2 = new Demo18ThreadB();
        t2.setName("B");
        t2.start();
    }
}

class Demo18Service{
    synchronized public static void foo1(){
        System.out.println(Thread.currentThread().getName() + "进入方法foo1在" + System.currentTimeMillis());
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println(Thread.currentThread().getName() + "结束方法foo1在" + System.currentTimeMillis());
    }
    synchronized public static void foo2(){
        System.out.println(Thread.currentThread().getName() + "进入方法foo2在" + System.currentTimeMillis());
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println(Thread.currentThread().getName() + "结束方法foo2在" + System.currentTimeMillis());
    }
}

class Demo18ThreadA extends Thread{
    @Override
    public void run() {
        Demo18Service.foo1();
    }
}

class Demo18ThreadB extends Thread{
    @Override
    public void run() {
        Demo18Service.foo2();
    }
}
```

从运行结果来年在静态方法上使用synchronized修饰与在非静态方法上使用的结果是一到致的，都是同步运行。从本质来讲它们是有区别的，在静态方法上使用synchronized是给Class类上锁，而在非静态方法上使用syncrhonized是给对象上锁。

```java
package chap2;

public class Demo19 {
    public static void main(String[] args) {
        Demo19Service service = new Demo19Service();
        Thread t1 = new Demo19ThreadA(service);
        t1.setName("A");
        t1.start();
        Thread t2 = new Demo19ThreadB(service);
        t2.setName("B");
        t2.start();
    }
}

class Demo19Service{
    synchronized public static void foo1(){
        System.out.println(Thread.currentThread().getName() + "进入foo1方法在" + System.currentTimeMillis());
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println(Thread.currentThread().getName() + "结束foo1方法在" + System.currentTimeMillis());
    }

    synchronized public void foo2(){
        System.out.println(Thread.currentThread().getName() + "进入foo2方法在" + System.currentTimeMillis());
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println(Thread.currentThread().getName() + "结束foo2方法在" + System.currentTimeMillis());
    }
}

class Demo19ThreadA extends Thread{
    private Demo19Service service;
    public Demo19ThreadA(Demo19Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo2();
    }
}

class Demo19ThreadB extends Thread{
    private Demo19Service service;
    public Demo19ThreadB(Demo19Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo1();
    }
}
```

以异步方式运行的原因持有锁是不一样的，非静态方法持有提对象锁，而静态方法持有的是Class锁，Class锁可以对类的所有对象实例起作用。

同步synchronized(class)代码块的作用其实和synchronized static方法的作用是一样的

```java
package chap2;

public class Demo20 {
    public static void main(String[] args) {
        Demo20Service service = new Demo20Service();
        Thread t1 = new Demo20ThreadA(service);
        t1.setName("A");
        t1.start();

        Thread t2 = new Demo20ThreadA(service);
        t2.setName("B");
        t2.start();
    }
}

class Demo20Service{
    synchronized public static void foo1(){
        System.out.println(Thread.currentThread().getName() + "进入foo1方法在" + System.currentTimeMillis());
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println(Thread.currentThread().getName() + "结束foo1方法在" + System.currentTimeMillis());
    }
    public static void foo2(){
        synchronized (Demo20Service.class) {
            System.out.println(Thread.currentThread().getName() + "进入foo1方法在" + System.currentTimeMillis());
            try {
                Thread.sleep(2000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            System.out.println(Thread.currentThread().getName() + "结束foo1方法在" + System.currentTimeMillis());
        }
    }
}

class Demo20ThreadA extends Thread{
    private Demo20Service service;
    public Demo20ThreadA(Demo20Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo1();
    }
}

class Demo20ThreadB extends Thread{
    private Demo20Service service;
    public Demo20ThreadB(Demo20Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo2();
    }
}
```

9. synchronized(String)

```java
package chap2;

public class Demo21 {
    public static void main(String[] args) {
        Thread t1 = new Demo21ThreadA();
        t1.setName("A");
        t1.start();

        Thread t2 = new Demo21ThreadB();
        t2.setName("B");
        t2.start();
    }
}

class Demo21Service{
    public static void foo1(String lockObject){
        try {
            synchronized (lockObject) {
                while (true) {
                    System.out.println("线程" + Thread.currentThread().getName());
                    Thread.sleep(1000);
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
    public static void foo2(Object lockObject){
        try {
            synchronized (lockObject) {
                while (true) {
                    System.out.println("线程" + Thread.currentThread().getName());
                    Thread.sleep(1000);
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo21ThreadA extends Thread{
    @Override
    public void run() {
        Demo21Service.foo1("AA");
    }
}

class Demo21ThreadB extends Thread{
    @Override
    public void run() {
        Demo21Service.foo2("AA");
    }
}
```

出现 只有线程A在运行的原因是，两个线程的对象锁都是使用AA，两个线程持有的相同的锁，所以造成线程B能运行。这就是String常量所带来的问题。不应该使用字符串类型作为对象锁，而应该使用其它类型，例如new Object()，对象就不会放入缓存 。

10. 解决synchronized方法无限等待问题

```java
package chap2;

public class Demo22 {
    public static void main(String[] args) {
        Demo22Service service = new Demo22Service();
        Thread t1 = new Demo22ThreadA(service);
        t1.start();
        Thread t2 = new Demo22ThreadB(service);
        t2.start();
    }
}

class Demo22Service{
    private Object lockObject1 = new Object();
    /*synchronized */public void foo1(){
        synchronized (lockObject1) {
            System.out.println("foo1方法开始执行");
            boolean isContinue = true;
            while (isContinue) {

            }
            System.out.println("foo2方法执行结束");
        }
    }
    private Object lockObject2 = new Object();
    /*synchronized */public void foo2(){
        synchronized (lockObject2) {
            System.out.println("foo2方法开始执行");
            System.out.println("foo2方法执行结束");
        }
    }
}

class Demo22ThreadA extends Thread{
    private Demo22Service service;
    public Demo22ThreadA(Demo22Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo1();
    }
}

class Demo22ThreadB extends Thread{
    private Demo22Service service;
    public Demo22ThreadB(Demo22Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo2();
    }
}
```

11. 死锁

所谓死锁，是指多个线程在运行过程中因争夺资源而造成的一种僵局，当线程处于这种僵持的状态时，如果没有外力作用，这些线程都无法再继续运行。

```java
package chap2;

public class Demo23 {
    public static void main(String[] args) throws InterruptedException {
        Demo23Thread t = new Demo23Thread();
        t.setFlag("a");

        Thread t1 = new Thread(t);
        t1.start();
        Thread.sleep(10);

        t.setFlag("b");
        Thread t2 = new Thread(t);
        t2.start();
    }
}

class Demo23Thread extends Thread{
    private String flag;    // 标志，控制代码以什么样的方式运行
    private Object lock1 = new Object();
    private Object lock2 = new Object();

    public void setFlag(String flag){
        this.flag = flag;
    }

    @Override
    public void run() {
        try {
            if ("a".equals(flag)) {
                synchronized (lock1) {
                    System.out.println("flag=" + flag);
                    Thread.sleep(3000);
                    synchronized (lock2) {
                        System.out.println("按lock1=>lock2的顺序执行");
                    }
                }
            } else {
                synchronized (lock2){
                    System.out.println("flag=" + flag);
                    Thread.sleep(3000);
                    synchronized (lock1){
                        System.out.println("按lock2->lock1的顺序执行");
                    }
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
```

11. 锁对象的内容改变

```java
package chap2;

public class Demo24 {
    public static void main(String[] args) throws InterruptedException {
        Demo24Service service = new Demo24Service();
        Thread t1 = new Demo24Thread(service);
        t1.setName("A");
        t1.start();

        Thread.sleep(50);

        Thread t2 = new Demo24Thread(service);
        t2.setName("B");
        t2.start();
    }
}

class Demo24Service{
    private String lockObject = "123";

    public void foo(){
        try {
            synchronized (lockObject) {
                System.out.println(Thread.currentThread().getName() + "开始于" + System.currentTimeMillis());
                lockObject = "456";
                Thread.sleep(2000);
                System.out.println(Thread.currentThread().getName() + "结束于" + System.currentTimeMillis());
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo24Thread extends Thread{
    private Demo24Service service;
    public Demo24Thread(Demo24Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo();
    }
}
```

线程A的锁是123，50毫秒后线程B得到的锁是456，所以钱程之间就是异步。如果没有Thread.sleep(50)，线程A和线程B的锁都是123，虽然代码里将锁改成456，但结果还是同步，因为线程A和线程B争抢的锁是123.

```java
package chap2;

public class Demo25 {
    public static void main(String[] args) throws InterruptedException {
        Demo25User user = new Demo25User();
        Demo25Service service = new Demo25Service();

        Thread t1 = new Demo25Thread(service, user);
        t1.setName("A");
        t1.start();

        Thread.sleep(50);

        Thread t2 = new Demo25Thread(service, user);
        t2.setName("B");
        t2.start();
    }
}

class Demo25User{
    public String username;
    public String password;
}

class Demo25Service{
    public void foo(Demo25User user){
        try {
            synchronized (user) {
                System.out.println(Thread.currentThread().getName() + "开始" + System.currentTimeMillis());
                user.username = "b";
                user.password = "bb";
                Thread.sleep(2000);
                System.out.println(Thread.currentThread().getName() + "结束" + System.currentTimeMillis());
            }

        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo25Thread extends Thread{
    private Demo25Service service;
    private Demo25User user;

    public Demo25Thread(Demo25Service service, Demo25User user){
        this.service = service;
        this.user = user;
    }

    @Override
    public void run() {
        service.foo(user);
    }
}
```

只要对象不变，即使对象的属性被改变，运行的结果 还是同步。

### 2.3 volatile关键字

volatile主要作用是使变量可以在多个线程间见。

```java
package chap2;

public class Demo26 {
    public static void main(String[] args) {
        Demo26Service service = new Demo26Service();
        service.foo();
        System.out.println(Thread.currentThread().getName() + "准备停止foo方法的循环");
        service.flag = false;
    }
}

class Demo26Service{
    // 标志，控制循环
    public boolean flag = true;

    public void foo(){
        System.out.println("foo开始运行");
        while(flag) {

        }
        System.out.println("foo运行结束");
    }
}
```

程序开始运行后停不下来，主要原因是main线程一直在处理while循环，导致程序不能继续 执行后的代码。

```java
package chap2;

public class Demo26 {
    public static void main(String[] args) throws InterruptedException {
        Demo26Service service = new Demo26Service();
//        service.foo();
        service.start();
        Thread.sleep(100);
        System.out.println(Thread.currentThread().getName() + "准备停止foo方法的循环");
        service.flag = false;
    }
}

class Demo26Service extends Thread{
    // 标志，控制循环
    public boolean flag = true;

    public void run(){
        foo();
    }

    public void foo(){
        System.out.println("foo开始运行");
        while(flag) {

        }
        System.out.println("foo运行结束");
    }
}
```

如果上面的代码运行在服务器模式(JDK8及以上版本默认就是使用服务器械、JDK7默认使用的客户端模式)中的JVM时就会出现死循环。

```java
package chap2;

public class Demo26 {
    public static void main(String[] args) throws InterruptedException {
        Demo26Service service = new Demo26Service();
//        service.foo();
        service.start();
        Thread.sleep(100);
        System.out.println(Thread.currentThread().getName() + "准备停止foo方法的循环");
        service.flag = false;
    }
}

class Demo26Service extends Thread{
    // 标志，控制循环
    volatile public boolean flag = true;

    public void run(){
        foo();
    }

    public void foo(){
        System.out.println("foo开始运行");
        while(flag) {

        }
        System.out.println("foo运行结束");
    }
}
```

在启动线程是，变量public boolean flag=true存在在公共堆栈及线程的私有堆栈中。当使用服务器模式时为了线程运行的效率，线程一直在私有堆栈中取得flag的值一直都是true。当执行service.flag=false时，虽然被执行，更新的却是在公共堆栈中的flag，所以一直是死循环的状态。使用volatile修饰成员变量后，强制虚拟机从公共堆栈中获取变量的值。

synchronize和volatile区别：

1) volatile是线程同步的轻易级实现，它的性能比synchronized要好，并且volatile只能修饰变量。而synchronized可以修饰方法及代码块。随着JDK的版本更新，synchronized在执行效率也得到很大的提升，在开发中synchronized的使用率还是较高

2) 多线程访问volatile不会发生阻塞，而synhcronized会出现阻塞

3) volatile能保证数据的可见性，不能保证原子性，而synchronized可以保证原子性，也可以间接保证可见性，因为它会将私有内存和公共内存中的数据做同步。

4) volatile解决的是变量在多个线程之见的可见性，而synchronized是解决多个线程之间访问资源的同步性。

## 3. 线程间的通信

线程是程序中独立的个体，但这些个体如果不经过处理就能成为一个整体。线程间的通信就是成为整体的必胜方案之一，可以说使线程间通讯后，线程之间的交互性会更强大，大大提高CPU复用率的同时还会使程序员对各线程任务在处理的过程中进行有效的把控与监督。

### 3.1 wait与notify

1. 不使用wait与notifty实现线程间通讯有什么问题

```java
package chap3;

import java.util.ArrayList;
import java.util.List;

public class Demo01 {
    public static void main(String[] args) {
        List list = new ArrayList();
        Thread t1 = new Demo01ThreadA(list);
        t1.start();
        Thread t2 = new Demo01ThreadB(list);
        t2.start();
    }
}

class Demo01ThreadA extends Thread{
    private List list;
    public Demo01ThreadA(List list){
        this.list = list;
    }

    @Override
    public void run() {
        try {
            for (int i = 0; i < 10; i++) {
                list.add(i);
                System.out.println("添加了" + (i + 1) + "个元素");
                Thread.sleep(1000);
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo01ThreadB extends Thread{
    // 强制线程从公共堆栈中获取数据
    volatile private List list;
    public Demo01ThreadB(List list){
        this.list = list;
    }

    @Override
    public void run() {
        try {
            while (true) {
                if (list.size() == 5) {
                    System.out.println("list已经有5个数据了，线程B要退出");
                    throw new InterruptedException();
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
```

虽然两个线程间通信实现了，但是有一个弊端，线程B需要不停地通过while语句轮询来检测一个条件，这样做会浪费CPU资源。

如果轮询的时间间很小，更浪费CPU资源，如果轮询的时间间隔很大，有可能会取不到想要到的数据。所以就需要有一种机制 来实现减少CPU的资源浪费，而且还可以实现在多个线程间通信，它就是『wait/notify』机制。

2. 什么是等待和通知机制

厨师与服务中上菜就是上等待和通知机制。

前面的例子线程之间也可以实现通信，原因就是多个线程共同访问同一个变量，但是这种机制 并不是『等待/通知』机制，两个线程完全是主动式地读取一个共享变量，在花费读取时间的基础上，读取到的是不是我们想要的，并不完全确定。所以需要使用『等待/通知』机制满足开发的需求。

3. wait/notifty机制的实现

wait方法的作用是使当前正在执行的线程进入等待状态，wait方法是Object类的方法，该方法用来将当前线程放入到『预执行队列』中，并且在wait所在的代码行进行停止执行，直到接到通知或被中断为止。在调用wait方法之前，线程必须获取得该对象的对象锁，也就是说只能在同步方法或同步代码块中调用wait方法。如果在执行wait方法后，当前线程锁会自动释放。当wait方法返回线程与其它线程重新竞争获得锁。

```java
package chap3;

public class Demo02 {
    public static void main(String[] args) {
        try {
            /*String str = new String();
            // Exception in thread "main" java.lang.IllegalMonitorStateException
            str.wait();*/

            String str = new String();
            System.out.println("同步代码块之前");
            synchronized (str){
                System.out.println("同步代码块的第一行");
                str.wait();
                System.out.println("wait之后的第一行代码");
            }

        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}
```

notify方法也要在同步方法或同步代码块中使用，在调用前线程必须获得该对象的对象锁，如果没有获取重适当的锁也会抛出IllegalMonitorStateException。这个方法是用来通知那些可能等待锁对象的其它线程，如果有多个线程等待，由线程调试器随机挑选一个在wait状态的线程，向其发出通知，并使等待获取该对象的对象锁。

在执行notify方法后，当前线程不会马上释放该对象锁，wait状态的线程也不能马上获取取该对象锁，要等到执行notify方法的线程将任务执行完成后，也就是退出synchronize代码块后，当前线程才会释放锁，wait状态线程才可以获取到锁。

当第一个获得该对象锁的wait线程运行完成后，它会释放掉该对象锁，如果 该对象没有再次使用nofity语句，则对象处理空闲状态，其它 wait状态的线程由于没有得到通知，还会继续处理阻塞的wait状态，直到这个对象发现次发出通知。

```java
package chap3;

public class Demo03 {
    public static void main(String[] args) throws InterruptedException {
        Object lock = new Object();
        Thread t1 = new Demo03ThreadA(lock);
        t1.start();
        Thread.sleep(2000);
        Thread t2 = new Demo03ThreadB(lock);
        t2.start();
    }
}

class Demo03ThreadA extends Thread{
    private Object lock;
    public Demo03ThreadA(Object lock){
        this.lock = lock;
    }

    @Override
    public void run() {
        try {
            synchronized (lock) {
                System.out.println("线程A开始等待在" + System.currentTimeMillis());
                lock.wait();
                System.out.println("线程A结束等待在" + System.currentTimeMillis());
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
class Demo03ThreadB extends Thread{
    private Object lock;
    public Demo03ThreadB(Object lock){
        this.lock = lock;
    }

    @Override
    public void run() {
        try {
            synchronized (lock) {
                System.out.println("线程B准备发出通知在" + System.currentTimeMillis());
                lock.notify();
                System.out.println("线程B结束发出通知在" + System.currentTimeMillis());
                Thread.sleep(1000);
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
```

wait就使线程停止运行，notify使停止的线程继续运行。

4. wait方法自动释放锁与notify方法不会释放锁

notify方法必须执行会同步代码后才会释放锁。

5. wait方法遇到intterrup方法

```java
package chap3;

public class Demo04 {
    public static void main(String[] args) throws InterruptedException {
        Object lock = new Object();
        Thread t1 = new Demo04Thread(lock);
        t1.start();
        Thread.sleep(2000);
        t1.interrupt();
    }
}

class Demo04Service{
    public void foo(Object lock){
        try{
            synchronized (lock){
                System.out.println("准备开始等待");
                lock.wait();
                System.out.println("结束等待");
            }
        }catch (InterruptedException e){
            System.out.println("出现异常，因为wait状态的线程被interrupt了");
            e.printStackTrace();
        }
    }
}

class Demo04Thread extends Thread{
    private Object lock;
    public Demo04Thread(Object lock){
        this.lock = lock;
    }

    @Override
    public void run() {
        Demo04Service service = new Demo04Service();
        service.foo(lock);
    }
}
```

当线程呈wait状态时，调用线程对象的interrupt方法会产生InterruptedException异常。

通知前面的例了总结 出：

1) 执行完同步代码块就会释放对象锁；

2) 在执行同步代码块的过程中，遇到异常而导致线程终止，锁也会被释放；

3) 在执行同步代码块的过程中，执行了锁所属对象的wait 方法，这个线程会释放对象锁，而此线程对象会进行线程等待池中等待被唤醒。

6. 只唤醒一个线程或所有线程

调用notify方法一次只随机通知一个线程进行唤醒

```java
package chap3;

public class Demo05 {
    public static void main(String[] args) throws InterruptedException {
        Object lock = new Object();
        Thread t1 = new Demo05ThreadA(lock);
        t1.setName("A");
        t1.start();
        Thread t2 = new Demo05ThreadA(lock);
        t2.setName("B");
        t2.start();
        Thread t3 = new Demo05ThreadA(lock);
        t3.setName("C");
        t3.start();
        Thread.sleep(1000);
        Thread t4 = new Demo05ThreadB(lock);
        t4.start();
    }
}
class Demo05Service{
    public void foo(Object lock){
        try{
            synchronized (lock){
                System.out.println(Thread.currentThread().getName() + "进入了foo方法，准备执行wait方法");
                lock.wait();
                System.out.println(Thread.currentThread().getName() + "结束了foo方法");
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
class Demo05ThreadA extends Thread{
    private Object lock;
    public Demo05ThreadA(Object lock){
        this.lock = lock;
    }

    @Override
    public void run() {
        Demo05Service service = new Demo05Service();
        service.foo(lock);
    }
}
class Demo05ThreadB extends Thread{
    private Object lock;
    public Demo05ThreadB(Object lock){
        this.lock = lock;
    }

    @Override
    public void run() {
        synchronized (lock){
            // notify仅随机唤醒一个线程
//            lock.notify();
            // 如果要唤醒多个线程，需要多次调用notify方法
//            lock.notify();
//            lock.notify();
//            lock.notify();
//            lock.notify();
            lock.notifyAll();
        }
    }
}
```

如果有多个线程，可以调用多次notify方法唤醒所有的线程，不能保证系统中确定有多少个线程，也就是说如果notify方法的调用次数小于线程数量时，会出现有部分线程无法被唤醒。为了唤醒全部的线程，可以使用notifyAll方法。

7. wait(long)的使用

带一个long参数的方法的作用等待某一时间内是否有线程对象锁进行唤醒，如果超过这个等待时间线程会自动唤醒。

```java
package chap3;

public class Demo06 {
    public static void main(String[] args) throws InterruptedException {
        Object lock = new Object();
        Thread t1 = new Demo06ThreadA(lock);
        t1.start();
        Thread.sleep(5000);
        Thread t2 = new Demo06ThreadB(lock);
        t2.start();
    }
}

class Demo06ThreadA extends Thread{
    private Object lock;
    public Demo06ThreadA(Object lock){
        this.lock = lock;
    }

    @Override
    public void run() {
        try{
            synchronized (lock){
                System.out.println("进入同步代码块于" + System.currentTimeMillis());
                lock.wait(3000);
                System.out.println("结束同步代码块于" + System.currentTimeMillis());
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo06ThreadB extends Thread{
    private Object lock;
    public Demo06ThreadB(Object lock){
        this.lock = lock;
    }

    @Override
    public void run() {
        synchronized (lock){
            System.out.println("开始唤醒线程在" + System.currentTimeMillis());
            lock.notify();
            System.out.println("结束唤醒线程在" + System.currentTimeMillis());
        }
    }
}
```

sleep(long)与wait(long)非常想象，都是在指定的时间后线程会自动唤醒，区别在于sleep是不会释放对象锁，而wait方法可以释放对象锁。

8. 通知过早

如果通知过早，就会打乱程序正常运行逻辑。

```java
package chap3;

public class Demo07 {
    public static void main(String[] args) throws InterruptedException {
        Object lock = new Object();
        Demo07Object canRun = new Demo07Object();
        Thread t1 = new Demo07ThreadA(lock, canRun);
        Thread t2 = new Demo07ThreadB(lock, canRun);
        // 正常执行顺序
//        t1.start();
//        t2.start();
        // 增加了同步代码块是否可以运行的判断
//        t2.start();
//        Thread.sleep(100);
//        t1.start();
        t1.start();
        Thread.sleep(100);
        t2.start();
    }
}

class Demo07Object{
    boolean canRun = true;
}

class Demo07ThreadA extends Thread{
    private Object lock;
    private Demo07Object canRun;

    public Demo07ThreadA(Object lock, Demo07Object canRun){
        this.lock = lock;
        this.canRun = canRun;
    }

    @Override
    public void run() {
        try{
            synchronized (lock){
                while(canRun.canRun) {
                    System.out.println("准备进入等待状态");
                    lock.wait();
                    System.out.println("结束等待状态");
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo07ThreadB extends Thread{
    private Object lock;
    private Demo07Object canRun;

    public Demo07ThreadB(Object lock, Demo07Object canRun){
        this.lock = lock;
        this.canRun = canRun;
    }

    @Override
    public void run() {
        synchronized (lock){
            System.out.println("准备执行唤醒方法");
            lock.notify();
            System.out.println("结束唤醒方法");
            canRun.canRun = false;
        }
    }
}
```

9. wait的条件发生变化

在使用wait/notify模式时，需要注意另外一种情况，就是wait等待的条件发生了变化，很容易会造成程序逻辑的混乱。

```java
package chap3;

import java.util.ArrayList;
import java.util.List;

public class Demo08 {
    public static void main(String[] args) throws InterruptedException {
        Demo08Service service = new Demo08Service();
        Thread t1 = new Demo08ThreadB(service);
        t1.start();
        Thread t2 = new Demo08ThreadB(service);
        t2.start();
        Thread.sleep(1000);
        Thread t3 = new Demo08ThreadA(service);
        t3.start();
    }
}

class Demo08Service{
    private List list = new ArrayList();
    private Object lock = new Object();

    public void add(){
        synchronized (lock){
            list.add("a");
            lock.notifyAll();
        }
    }
    public void subtrac(){
        try {
            synchronized (lock) {
                if (list.size() == 0) {
                    System.out.println(Thread.currentThread().getName() + "开始等待数据");
                    lock.wait();
                    System.out.println(Thread.currentThread().getName() + "结束获取数据等待");
                }
                if (list.size() > 0) {
                    list.remove(0);
                }
                System.out.println(Thread.currentThread().getName() + "：list的大小是" + list.size());
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
class Demo08ThreadA extends Thread{
    private Demo08Service service;
    public Demo08ThreadA(Demo08Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.add();
    }
}
class Demo08ThreadB extends Thread{
    private Demo08Service service;
    public Demo08ThreadB(Demo08Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.subtrac();
    }
}
```

10. 生产者/消费者模式实现

1. 一个生产者与一个消费者(操作者)

```java
package chap3;

public class Demo09 {
    public static void main(String[] args) {
        Object lock = new Object();
        Thread t1 = new Demo09Producer(lock);
        t1.start();
        Thread t2 = new Demo09Consumer(lock);
        t2.start();
    }
}

class Demo09VO{
    public static String value = "";
}

// 生产者
class Demo09Producer extends Thread {
    private Object lock;
    public Demo09Producer(Object lock){
        this.lock = lock;
    }

    @Override
    public void run() {
        try{
            while(true) {
                synchronized (lock) {
                    if (!"".equals(Demo09VO.value)){
                        lock.wait();
                    }
                    String value = System.currentTimeMillis() + "_" + System.nanoTime();
                    System.out.println("Set的值是：" + value);
                    Demo09VO.value = value;
                    lock.notify();
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo09Consumer extends Thread{
    private Object lock;
    public Demo09Consumer(Object lock){
        this.lock = lock;
    }

    @Override
    public void run() {
        try{
            while(true){
                synchronized (lock){
                    if ("".equals(Demo09VO.value)){
                        lock.wait();
                    }
                    System.out.println("Get的值是：" + Demo09VO.value);
                    Demo09VO.value = "";
                    lock.notify();
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
```

如果在这个代码的基础上，设计出多生产者与多消费者，在运行过程很有可能会出现 『假死』的情况，也就是所有的线程都是WAIT状态。

2. 多生产与多消费者(操作值)

```java
package chap3;

public class Demo10 {
    public static void main(String[] args) {
        Object lock = new Object();
        int size = 2;
        Thread[] producers = new Thread[size];
        Thread[] consumers = new Thread[size];
        for (int i = 0; i < size; i++) {
            char c = (char)('A' + i);
            producers[i] = new Demo10Producer(lock);
            producers[i].setName("生产者" + c);

            consumers[i] = new Demo10Consumer(lock);
            consumers[i].setName("消费者" + c);

            producers[i].start();
            consumers[i].start();
        }
    }
}

class Demo10VO{
    public static String value = "";
}

// 生产者
class Demo10Producer extends Thread {
    private Object lock;
    public Demo10Producer(Object lock){
        this.lock = lock;
    }

    @Override
    public void run() {
        try{
            while(true) {
                synchronized (lock) {
                    if (!"".equals(Demo10VO.value)){
                        System.out.println(Thread.currentThread().getName() + "等待中……");
                        lock.wait();
                    }
                    System.out.println(Thread.currentThread().getName() + "生产中……");
                    String value = System.currentTimeMillis() + "_" + System.nanoTime();
                    Demo10VO.value = value;
                    //lock.notify();
                    lock.notifyAll();
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo10Consumer extends Thread{
    private Object lock;
    public Demo10Consumer(Object lock){
        this.lock = lock;
    }

    @Override
    public void run() {
        try{
            while(true){
                synchronized (lock){
                    if ("".equals(Demo10VO.value)){
                        System.out.println(Thread.currentThread().getName() + "等待中……");
                        lock.wait();
                    }
                    System.out.println(Thread.currentThread().getName() + "消费中……");
                    Demo10VO.value = "";
//                    lock.notify();
                    lock.notifyAll();
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
```

假设生产A它在生产数据，其它3个线程(生产者B、消费者A、消费者B)都是呈等待状态，当生产者A生产完成后，随机唤醒1个线程，刚好唤醒了生产B，生产者B发现ValueObject中有数据所以进行到等待状态(生产者A竞争锁，生产者B、消费者A、消费者B等待状态)，A又重新获得锁但是它发现创建出来的值没还有被消费，所以它又进入等待状态，结果就是4个线程都是在待状态。怎样样解决这个问题，使用notifyAll方法把有的线程都唤醒，保证生产出来的值一定会被消费掉。

3. 一生产者与一消费者(操作集合)

```java
package chap3;

import java.util.ArrayList;
import java.util.List;

public class Demo11 {
    public static void main(String[] args) {
        Demo11VO vo = new Demo11VO();
        Thread producer = new Demo11Producer(vo);
        producer.setName("生产者");
        producer.start();
        Thread consumer = new Demo11Consumer(vo);
        consumer.setName("消费者");
        consumer.start();
    }
}

class Demo11VO{
    private List<String> list = new ArrayList<>();

    synchronized public void push(String val){
        try {
            if (list.size() == 1) {
                System.out.println(Thread.currentThread().getName() + "等待中");
                this.wait();
            }
            list.add(val);
            System.out.println(Thread.currentThread().getName() + "：添加数据" + val);
            System.out.println(Thread.currentThread().getName() + "：还有" +  list.size() + "个数据");
            this.notify();
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }

    synchronized public String pop(){
        String returnValue = null;
        try {
            if (list.size() == 0) {
                System.out.println(Thread.currentThread().getName() + "等待中");
                this.wait();
            }
            returnValue = list.get(0);
            list.remove(0);
            System.out.println(Thread.currentThread().getName() + "：消费数据" + returnValue);
            System.out.println(Thread.currentThread().getName() + "：还有" + list.size() + "个数据");
            this.notify();
        }catch (InterruptedException e){
            e.printStackTrace();
        }
        return returnValue;
    }
}

class Demo11Producer extends Thread{
    private Demo11VO vo;
    public Demo11Producer(Demo11VO vo){
        this.vo = vo;
    }

    @Override
    public void run() {
        while(true) {
            vo.push(Math.random() + "");
        }
    }
}

class Demo11Consumer extends Thread{
    private Demo11VO vo;
    public Demo11Consumer(Demo11VO vo){
        this.vo = vo;
    }

    @Override
    public void run() {
        while(true){
            vo.pop();
        }
    }
}
```

4. 一生产者多消费者(操作集合)

```java
package chap3;

import java.util.ArrayList;
import java.util.List;

public class Demo12 {
    public static void main(String[] args) {
        Demo12VO vo = new Demo12VO();
        Thread producer = new Demo12Producer(vo);
        producer.setName("生产者");
        producer.start();

        Thread[] consumers = new Thread[5];
        for (int i = 0; i < consumers.length; i++) {
            consumers[i] = new Demo12Consumer(vo);
            consumers[i].setName("消费者" + (char)('A' + i));
            consumers[i].start();
        }
    }
}

class Demo12VO{
    private List<String> list = new ArrayList<>();

    synchronized public void push(String val){
        try {
            if (list.size() == 1) {
                System.out.println(Thread.currentThread().getName() + "等待中");
                this.wait();
            }
            list.add(val);
            System.out.println(Thread.currentThread().getName() + "：添加数据" + val);
            System.out.println(Thread.currentThread().getName() + "：还有" +  list.size() + "个数据");
            this.notify();
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }

    synchronized public String pop(){
        String returnValue = null;
        try {
            while (list.size() == 0) {
                System.out.println(Thread.currentThread().getName() + "等待中");
                this.wait();
            }
            returnValue = list.get(0);
            list.remove(0);
            System.out.println(Thread.currentThread().getName() + "：消费数据" + returnValue);
            System.out.println(Thread.currentThread().getName() + "：还有" + list.size() + "个数据");
//            this.notify();
            this.notifyAll();
        }catch (InterruptedException e){
            e.printStackTrace();
        }
        return returnValue;
    }
}

class Demo12Producer extends Thread{
    private Demo12VO vo;
    public Demo12Producer(Demo12VO vo){
        this.vo = vo;
    }

    @Override
    public void run() {
        while(true) {
            vo.push(Math.random() + "");
        }
    }
}

class Demo12Consumer extends Thread{
    private Demo12VO vo;
    public Demo12Consumer(Demo12VO vo){
        this.vo = vo;
    }

    @Override
    public void run() {
        while(true){
            vo.pop();
        }
    }
}
```

注意：

1. 消费者在判断集合是否有数据时不能使用if，有可能重复唤醒的还是消费，需要使用while语句判断，保证集合中一定是有数据的。
2. 消费者唤醒线程不使用notify，notify是随机唤醒一个线程，万一重复唤醒消费者就没有办法重新生产新的数据。需要使用notifyAll方法。

5. 多生产者与一消费(操作集合)

```java
package chap3;

import java.util.ArrayList;
import java.util.List;

public class Demo13 {
    public static void main(String[] args) {
        Demo13VO vo = new Demo13VO();
        Thread[] producers = new Thread[5];
        for (int i = 0; i < producers.length; i++) {
            producers[i] = new Demo13Producer(vo);
            producers[i].setName("生产者" + (char)('A' + i));
            producers[i].start();
        }


        Thread consumer = new Demo13Consumer(vo);
        consumer.setName("消费者");
        consumer.start();
    }
}

class Demo13VO{
    private List<String> list = new ArrayList<>();

    synchronized public void push(String val){
        try {
            while (list.size() == 1) {
                System.out.println(Thread.currentThread().getName() + "等待中");
                this.wait();
            }
            list.add(val);
            System.out.println(Thread.currentThread().getName() + "：添加数据" + val);
            System.out.println(Thread.currentThread().getName() + "：还有" +  list.size() + "个数据");
            this.notifyAll();
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }

    synchronized public String pop(){
        String returnValue = null;
        try {
            if (list.size() == 0) {
                System.out.println(Thread.currentThread().getName() + "等待中");
                this.wait();
            }
            returnValue = list.get(0);
            list.remove(0);
            System.out.println(Thread.currentThread().getName() + "：消费数据" + returnValue);
            System.out.println(Thread.currentThread().getName() + "：还有" + list.size() + "个数据");
            this.notify();
        }catch (InterruptedException e){
            e.printStackTrace();
        }
        return returnValue;
    }
}

class Demo13Producer extends Thread{
    private Demo13VO vo;
    public Demo13Producer(Demo13VO vo){
        this.vo = vo;
    }

    @Override
    public void run() {
        while(true) {
            vo.push(Math.random() + "");
        }
    }
}

class Demo13Consumer extends Thread{
    private Demo13VO vo;
    public Demo13Consumer(Demo13VO vo){
        this.vo = vo;
    }

    @Override
    public void run() {
        while(true){
            vo.pop();
        }
    }
}
```

注意：

1. 需要不断的判断集合中是否有数据，如果有就重新等待，直到没有数据 时再做添加
2. 生产者不能使用notify唤醒其它线程，有可能造成死锁，应该使用notifyAll方法把所有的线程都唤醒。

7. 多生产者与多消费者(操作集合)

```java
package chap3;

import java.util.ArrayList;
import java.util.List;

public class Demo14 {
    public static void main(String[] args) {
        Demo14VO vo = new Demo14VO();

        int size = 5;
        Thread[] producers = new Thread[size];
        Thread[] consumers = new Thread[size];

        for (int i = 0; i < size; i++) {
            char c = (char)('A' + i);
            producers[i] = new Demo14Producer(vo);
            producers[i].setName("生产者" + c);
            producers[i].start();

            consumers[i] = new Demo14Consumer(vo);
            consumers[i].setName("消费者" + c);
            consumers[i].start();
        }
    }
}

class Demo14VO{
    private List<String> list = new ArrayList<>();

    synchronized public void push(String val){
        try {
            while (list.size() == 1) {
                System.out.println(Thread.currentThread().getName() + "等待中");
                this.wait();
            }
            list.add(val);
            System.out.println(Thread.currentThread().getName() + "：添加数据" + val);
            System.out.println(Thread.currentThread().getName() + "：还有" +  list.size() + "个数据");
            this.notifyAll();
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }

    synchronized public String pop(){
        String returnValue = null;
        try {
            while (list.size() == 0) {
                System.out.println(Thread.currentThread().getName() + "等待中");
                this.wait();
            }
            returnValue = list.get(0);
            list.remove(0);
            System.out.println(Thread.currentThread().getName() + "：消费数据" + returnValue);
            System.out.println(Thread.currentThread().getName() + "：还有" + list.size() + "个数据");
            this.notifyAll();
        }catch (InterruptedException e){
            e.printStackTrace();
        }
        return returnValue;
    }
}

class Demo14Producer extends Thread{
    private Demo14VO vo;
    public Demo14Producer(Demo14VO vo){
        this.vo = vo;
    }

    @Override
    public void run() {
        while(true) {
            vo.push(Math.random() + "");
        }
    }
}

class Demo14Consumer extends Thread{
    private Demo14VO vo;
    public Demo14Consumer(Demo14VO vo){
        this.vo = vo;
    }

    @Override
    public void run() {
        while(true){
            vo.pop();
        }
    }
}
```

### 3.2 join方法

在很情况下，都由主线程创建并运行子线程，如果子线程中需要进行大量的耗时运行，主线程往往将早于子线程结束。这时，如果主线程想等待子线程执行完成后再线路，比如子线处理一个数据，主线程需要取得这个数据中的值，可以使用join方法。join方法的作用是等待线程对象销毁。

1. join方法的使用

```java
package chap3;

public class Demo15 {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Demo15Thread();
        t.start();
        t.join();
        System.out.println("子线程执行完成以后再执行");
    }
}

class Demo15Thread extends Thread{
    @Override
    public void run() {
        try {
            int value = (int) (Math.random() * 10000);
            System.out.println("需要等待" + value + "毫秒");
            Thread.sleep(value);
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
```

方法join的作用就是使所属的线程对象x(子线程)执行run方法中的任务，而使得当前线程z(主线程)进行无限期的阻塞，等待线程x销毁后再继续 执行线程z后面的代码。方法join具有线程排队运行的作用，有些类似于同步代码运行的效果。

join与synchroinzed区别：join的内部使用wait方法进行等待，而synchronized关键字使用的是『对象锁』的机制作为同步。

2. join方法与异常

```java
package chap3;

public class Demo16 {
    public static void main(String[] args) throws InterruptedException {
        Demo16ThreadB t = new Demo16ThreadB();
        t.start();
        Thread.sleep(10);
        Thread t2 = new Demo16ThreadC(t);
        t2.start();
    }
}

class Demo16ThreadA extends Thread{
    @Override
    public void run() {
        for (int i = 0; i < Integer.MAX_VALUE; i++) {
            // 耗时操作
            String s = new String();
            Math.random();
        }
    }
}
class Demo16ThreadB extends Thread{
    @Override
    public void run() {
        try {
            Thread t = new Demo16ThreadA();
            t.start();
            t.join();
            System.out.println("B线程正常结束");
        }catch (InterruptedException e){
            System.out.println("B线程异常结束");
            e.printStackTrace();
        }
    }
}
class Demo16ThreadC extends Thread{
    private Demo16ThreadB t;
    public Demo16ThreadC(Demo16ThreadB t){
        this.t = t;
    }

    @Override
    public void run() {
        t.interrupt();
    }
}
```

join与interrupte方法如果彼此遇到，则会出现异常，但进程并没有结束，原因是ThreadA还在继续运行，线程A并没有出现 异常，是正常状态下继续 执行。

3. join(long)方法使用

方法join(long)中的参数是设定等待的时间

```java
package chap3;

public class Demo17 {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Demo17Thread();
        t.start();
        t.join(2000);
        System.out.println("主线程结束于" + System.currentTimeMillis());
    }
}

class Demo17Thread extends Thread{
    @Override
    public void run() {
        try {
            System.out.println("子线程开始于" + System.currentTimeMillis());
            Thread.sleep(5000);
            System.out.println("子线程结束于" + System.currentTimeMillis());
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
```

如果 代码改成使用sleep(2000)，运行的效果还是等待2000毫秒，那使用join(2000)与sleep(2000)有什么区别？运行效果没有任何区别，但是同步的处理上不一样，sleep不释放锁，而join底层是使用wait来实现所以会释放锁。

4. join(long)与sleep(long)的区别

```java
package chap3;

public class Demo18 {
    public static void main(String[] args) {
        Demo18ThreadA t1 = new Demo18ThreadA();
        Demo18ThreadB t2 = new Demo18ThreadB(t1);
        t2.start();
        Demo18ThreadC t3 = new Demo18ThreadC(t1);
        t3.start();
    }
}
class Demo18ThreadA extends Thread{
    @Override
    public void run() {
        try {
            System.out.println("线程A开始于" + System.currentTimeMillis());
            Thread.sleep(5000);
            System.out.println("线程A结束于" + System.currentTimeMillis());
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }

    synchronized public void foo(){
        System.out.println("方法执行时间" + System.currentTimeMillis());
    }
}
class Demo18ThreadB extends Thread{
    private Demo18ThreadA t;
    public Demo18ThreadB(Demo18ThreadA t){
        this.t = t;
    }

    @Override
    public void run() {
        synchronized (t) {
            try {
                t.start();
//                Thread.sleep(6000);
                t.join();
                for (int i = 0; i < Integer.MAX_VALUE; i++) {
                    String s = new String();
                    Math.random();
                }
            }catch (InterruptedException e){
                e.printStackTrace();
            }
        }
    }
}
class Demo18ThreadC extends Thread{
    private Demo18ThreadA t;
    public Demo18ThreadC(Demo18ThreadA t){
        this.t = t;
    }

    @Override
    public void run() {
        t.foo();
    }
}
```

方法join(long)的功能在内部使用wait(long)来实现同步，所以joing(long)方法具有释放同步锁的特点。

### 3.3 ThreadLocal类

变量值共享可以使用静态变量形式，所有的线程都可以使用同一个静态变量。如果想实现每一个线程都有自己的共享变量如何解决？Java提供ThreadLocal类来解决这个问题。

ThreadLocal类主要解决每个线程绑定自己的值，可以将ThreadLocal类比喻成全局存储数据的盒子，盒子中可以存储每个线程的私有数据 。

```java
package chap3;

public class Demo19 {
    public static void main(String[] args) {
        ThreadLocal tl = new ThreadLocal();
        if (tl.get() == null){ // 获取当前线程存储的数据
            System.out.println("从未放过值");
            tl.set("a");        // 保存当前线程的数据
        }
        System.out.println(tl.get());
        System.out.println(tl.get());
    }
}
```

ThreadLocal，使用get方法获取当前线程保存的数据，使用set方法将数据保存当前的线程中。

```java
package chap3;

public class Demo20 {
    public static void main(String[] args) {
        ThreadLocal tl = new ThreadLocal();
        Thread t1 = new Demo20ThreadA(tl);
        t1.setName("A");
        t1.start();

        Thread t2 = new Demo20ThreadA(tl);
        t2.setName("B");
        t2.start();
    }
}

class Demo20ThreadA extends Thread{
    private ThreadLocal tl;
    public Demo20ThreadA(ThreadLocal tl){
        this.tl = tl;
    }

    @Override
    public void run() {
        try {
            for (int i = 0; i < 100; i++) {
                tl.set(Thread.currentThread().getName() + (i + 1));
                System.out.println(Thread.currentThread().getName() + " get value = " + tl.get());
                Thread.sleep(100);
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
```

## 4. Lock

### 4.1 ReentranLock

在Java多线程中，可以使用synchronized关键字来实现线程之间的同步互斥，但在JDK1.5中新增加ReentranLock类也可以达到同样的效果，并且在扩展功能上更加强大，比如具有嗅探锁定、多路分支通知等，而且在使用上也比synchronized更加灵活。

1. 使用ReentranLock

```java
package chap4;

import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class Demo01 {
    public static void main(String[] args) {
        Lock lock = new ReentrantLock();
        Thread t1 = new Demo01Thread(lock);
        Thread t2 = new Demo01Thread(lock);
        Thread t3 = new Demo01Thread(lock);
        t1.start();
        t2.start();
        t3.start();
    }
}

class Demo01Thread extends Thread{
    private Lock lock;
    public Demo01Thread(Lock lock){
        this.lock = lock;
    }

    @Override
    public void run() {
        lock.lock();    // 加上同步锁
        for (int i=0; i< 5; i++){
            System.out.println(Thread.currentThread().getName() + ", " + (i + 1));
        }
        lock.unlock();  // 解开同步锁
    }
}
```

从运行结果来看，当前线程打印完毕之后将锁进行释放，其它线程才可以继续 打印。线程打印的数组是分级打印，因为当前线程已经持有锁，但线程之间打印的顺序是随机的。

```java
package chap4;

import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class Demo02 {
    public static void main(String[] args) {
        Demo02Service service = new Demo02Service();
        Thread t1 = new Demo02ThreadA(service);
        t1.setName("A");
        t1.start();

        Thread t2 = new Demo02ThreadB(service);
        t2.setName("B");
        t2.start();

    }
}

class Demo02Service{
    private Lock lock = new ReentrantLock();
    /*synchronized*/ public void foo1(){
        try {
            lock.lock();
            System.out.println(Thread.currentThread().getName() + "开始执行foo1方法" + System.currentTimeMillis());
            Thread.sleep(2000);
            System.out.println(Thread.currentThread().getName() + "结束执行foo1方法" + System.currentTimeMillis());
            lock.unlock();
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
    /*synchronized*/ public void foo2(){
        try{
            lock.lock();
            System.out.println(Thread.currentThread().getName() + "开始执行foo2方法" + System.currentTimeMillis());
            Thread.sleep(2000);
            System.out.println(Thread.currentThread().getName() + "结束执行foo2方法" + System.currentTimeMillis());
            lock.unlock();
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
class Demo02ThreadA extends Thread{
    private Demo02Service service;
    public Demo02ThreadA(Demo02Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo1();
    }
}
class Demo02ThreadB extends Thread{
    private Demo02Service service;
    public Demo02ThreadB(Demo02Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo2();
    }
}
```

调用lock方法的线程就会持有对象锁，其它的线程只能等待锁被释放(调用unlock方法)才可以再次争抢锁。效果和使用synchronized关键字一样，线程之间还是按照顺序执行。

2. ReentrantLock实现wait和notifty

关键字synchronized与wait/notify或notifyAll方法想结合可以实现等待/通知模式，类RentrantLock也可以实现同样的功能，但需要借助Condition对象。Condition类是在JDK5中出现的技术，使用它有更好的灵活性，比如可以实现多路通知功能，也就是在一个Lock对象里面可以创建多个Condition实例，线程对象可以注册在指定的Condition中，从而可以有选择性地进行线程通知，在调试线程上更加灵活。

在使用notify/notifyAll方法进行通知时，被通知线程是由JVM随机选择的，但使用ReentrantLock结合Condition类可以实现选择性通知，这个功能是非常重要的，而且在Condition在中是默认提供的。

而synchronized就相当于整个Lock对象中只有一个单一Condition对象，所有的线程都是注册在它一个对象的身上。线程开始notifyAll时需要通知所有正在等待的线程，没有选择权，会出现相当大效率问题。

```java
package chap4;

import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class Demo03 {
    public static void main(String[] args) throws InterruptedException {
        Demo03Service service = new Demo03Service();
        Thread t = new Demo03Thread(service);
        t.start();
        Thread.sleep(2000);
        service.signal();
    }
}

class Demo03Service {
    private Lock lock = new ReentrantLock();
    private Condition condition = lock.newCondition();

    public void await(){
        try {
            lock.lock();
            System.out.println("await方法开始于" + System.currentTimeMillis());
            condition.await();  // 需要在同步代码中调用，否则会抛出IllegalMonitorStateException
            System.out.println("await方法结束于" + System.currentTimeMillis());
            lock.unlock();
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
    public void signal(){
        lock.lock();
        System.out.println("signal方法开始于" + System.currentTimeMillis());
        condition.signal();
        lock.unlock();
    }
}
class Demo03Thread extends Thread{
    private Demo03Service service;
    public Demo03Thread(Demo03Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.await();
    }
}
```

3. 使用Condition唤醒不同的线程

```java
package chap4;

import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class Demo04 {
    public static void main(String[] args) throws InterruptedException {
        Demo04Service service = new Demo04Service();
        Thread t1 = new Demo04ThreadA(service);
        t1.setName("A");
        t1.start();

        Thread t2 = new Demo04ThreadB(service);
        t2.setName("B");
        t2.start();

        Thread.sleep(2000);
        service.signalAll_A();

    }
}

class Demo04Service{
    private Lock lock = new ReentrantLock();
    private Condition conditionA = lock.newCondition();
    private Condition conditionB = lock.newCondition();

    public void awaitA(){
        try{
            lock.lock();
            System.out.println(Thread.currentThread().getName() + "开始执行awaitA方法" + System.currentTimeMillis());
            conditionA.await();
            System.out.println(Thread.currentThread().getName() + "结束执行awaitA方法" + System.currentTimeMillis());
        }catch (InterruptedException e){
            e.printStackTrace();
        } finally {
            lock.unlock();
        }
    }

    public void awaitB(){
        try{
            lock.lock();
            System.out.println(Thread.currentThread().getName() + "开始执行awaitB方法" + System.currentTimeMillis());
            conditionB.await();
            System.out.println(Thread.currentThread().getName() + "结束执行awaitB方法" + System.currentTimeMillis());
        }catch (InterruptedException e){
            e.printStackTrace();
        } finally {
            lock.unlock();
        }
    }

    public void signalAll_A(){
        try{
            lock.lock();
            System.out.println(Thread.currentThread().getName() + "唤醒所有的线程在" + System.currentTimeMillis());
            conditionA.signalAll();
        }finally {
            lock.unlock();
        }
    }

    public void signalAll_B(){
        try{
            lock.lock();
            System.out.println(Thread.currentThread().getName() + "唤醒所有的线程在" + System.currentTimeMillis());
            conditionB.signalAll();
        }finally {
            lock.unlock();
        }
    }
}

class Demo04ThreadA extends Thread{
    private Demo04Service service;
    public Demo04ThreadA(Demo04Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.awaitA();
    }
}

class Demo04ThreadB extends Thread{
    private Demo04Service service;
    public Demo04ThreadB(Demo04Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.awaitB();
    }
}
```

分别唤醒不同的线程，就需要使用多个Condition对象，也就是Condition对象可以唤醒部分指定的线程，有助于提升程序的运行效率。

5. 使用Lock实现生产者与消费者

```java
package chap4;

import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class Demo05 {
    public static void main(String[] args) {
        Demo05Service service = new Demo05Service();
        // 一生产一消费
//        Thread producer = new Demo05ProducerThread(service);
//        producer.start();
//
//        Thread consumer = new Demo05ConsumerThread(service);
//        consumer.start();

        // 多生产多消费
        int size = 2;
        Thread[] producers = new Thread[size];
        Thread[] consumers = new Thread[size];

        for (int i = 0; i < size; i++) {
            char c = (char)('A' + i);
            producers[i] = new Demo05ProducerThread(service);
            producers[i].setName("生产者" + c);
            producers[i].start();

            consumers[i] = new Demo05ConsumerThread(service);
            consumers[i].setName("消费者" + c);
            consumers[i].start();
        }
    }
}

class Demo05Service{
    private Lock lock = new ReentrantLock();
    private Condition condition = lock.newCondition();
    private String val = "";

    public void set(){
        try{
            lock.lock();
            while(!"".equals(val)){
                System.out.println(Thread.currentThread().getName() + "开始等待");
                condition.await();
            }
            val = System.currentTimeMillis() + "-" + System.nanoTime();
            System.out.println(Thread.currentThread().getName() + "生产值：" + val);
            //condition.signal();
            condition.signalAll();
        }catch (InterruptedException e){
            e.printStackTrace();
        } finally {
            lock.unlock();
        }
    }

    public void get(){
        try{
            lock.lock();
            while("".equals(val)){
                System.out.println(Thread.currentThread().getName() + "开始等待");
                condition.await();
            }
            System.out.println(Thread.currentThread().getName() + "消费值：" + val);
            val = "";
//            condition.signal();
            condition.signalAll();
        }catch (InterruptedException e){
            e.printStackTrace();
        } finally {
            lock.unlock();
        }
    }
}

class Demo05ProducerThread extends Thread{
    private Demo05Service service;
    public Demo05ProducerThread(Demo05Service service){
        this.service = service;
    }

    @Override
    public void run() {
        while(true){
            service.set();
        }
    }
}

class Demo05ConsumerThread extends Thread{
    private Demo05Service service;
    public Demo05ConsumerThread(Demo05Service service){
        this.service = service;
    }

    @Override
    public void run() {
        while(true){
            service.get();
        }
    }
}
```

6. 公平锁与非公平锁

Lock分为公平锁与非公平锁两种。公平锁表示线程获取锁的顺序是按照加锁的顺序来分配的，也就是先来先得。而非公平锁就是一种获取锁的抢占机制，是随机获得锁，与公平锁不一样的是先来的不一定先得到锁，这种方式可能会造成某些线程一直都拿不到锁，结果也就是不公平的。

```java
package chap4;

import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class Demo06 {
    public static void main(String[] args) {
//        Demo06Service service = new Demo06Service(true);        // 公平锁
        Demo06Service service = new Demo06Service(false);        // 非公平锁

        Thread[] threads = new Thread[50];
        for (int i = 0; i < threads.length; i++) {
            threads[i] = new Demo06Thread(service);
        }
        for (int i = 0; i < threads.length; i++) {
            threads[i].start();
        }
    }
}

class Demo06Service{
    private Lock lock;
    public Demo06Service(boolean isFair){
        // 无参的ReentrantLock是非公平锁，通过boolean参数控制锁的类型，如果是true会使用公平锁，否则就是非公平锁
        lock = new ReentrantLock(isFair);
    }
    public void foo(){
        try{
            lock.lock();
            System.out.println(Thread.currentThread().getName() + "获得锁定");
        }finally {

            lock.unlock();
        }

    }
}

class Demo06Thread extends Thread{
    private Demo06Service service;
    public Demo06Thread(Demo06Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.foo();
    }
}
```

7. 实现线程的顺序执行

```java
package chap4;

import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class Demo07 {
    private static Lock lock = new ReentrantLock();
    private static Condition conditionA = lock.newCondition();
    private static Condition conditionB = lock.newCondition();
    private static Condition conditionC = lock.newCondition();
    volatile private static int nextPrintWho = 1;

    public static void main(String[] args) {
        Thread t1 = new Thread(){
            @Override
            public void run() {
                try{
                    lock.lock();
                    while (nextPrintWho != 1){
                        conditionA.await();
                    }
                    for (int i = 0; i < 3; i++) {
                        System.out.println("ThreadA打印第" + (i + 1) + "次");
                    }
                    nextPrintWho = 2;
                    conditionB.signalAll();
                }catch (InterruptedException e){
                    e.printStackTrace();
                }finally {
                    lock.unlock();
                }

            }
        };
        Thread t2 = new Thread(){
            @Override
            public void run() {
                try{
                    lock.lock();
                    while (nextPrintWho != 2){
                        conditionB.await();
                    }
                    for (int i = 0; i < 3; i++) {
                        System.out.println("ThreadB打印第" + (i + 1) + "次");
                    }
                    nextPrintWho = 3;
                    conditionC.signalAll();
                }catch (InterruptedException e){
                    e.printStackTrace();
                }finally {
                    lock.unlock();
                }

            }
        };
        Thread t3 = new Thread(){
            @Override
            public void run() {
                try{
                    lock.lock();
                    while (nextPrintWho != 3){
                        conditionC.await();
                    }
                    for (int i = 0; i < 3; i++) {
                        System.out.println("ThreadC打印第" + (i + 1) + "次");
                    }
                    nextPrintWho = 1;
                    conditionA.signalAll();
                }catch (InterruptedException e){
                    e.printStackTrace();
                }finally {
                    lock.unlock();
                }

            }
        };

        Thread[] threadsA = new Thread[5];
        Thread[] threadsB = new Thread[5];
        Thread[] threadsC = new Thread[5];
        for (int i = 0; i < threadsA.length; i++) {
            threadsA[i] = new Thread(t1);
            threadsB[i] = new Thread(t2);
            threadsC[i] = new Thread(t3);

            threadsA[i].start();
            threadsB[i].start();
            threadsC[i].start();
        }
    }
}
```

### 4.2 ReentrantReadWriteLock

类ReentrantLock具有完全互斥排他锁，也就是同一时间内只有一个线程可以在执行Reentrant.lock方法后面的任务。这样虽然可以保证实例变量的线程安全性，但是效率却是非常的低下。所以在JDK中提供了一种读写锁ReentrantReadWriteLock类，使用它可以加快运行的效率，在某些不需要操作实例变量的方法中，完全可以使用读写锁来提升方法的代码运行效率。

读写锁表示有两个锁，一个是读操作也称作共享锁，另一个是写操作也叫排他锁。也就是多个读锁之间不互斥，读锁与写锁是互斥，写锁与写锁之间也是互斥的。在没有线程进入定入操作时，进行读取操作的多个线程都可以获取 读取锁，而进行写入操作时只有在获取写锁后才能进行写入操作。即可以有多个线程同时进行读取操作，但同一时内只能有一个线程进行写操作。

```java
package chap4;

import java.util.concurrent.locks.ReentrantReadWriteLock;

public class Demo08 {
    public static void main(String[] args) {
        Demo08Service service = new Demo08Service();
//        Thread t1 = new Demo08ThreadA(service);
//        t1.setName("读线程A");
//        t1.start();
//        Thread t2 = new Demo08ThreadA(service);
//        t2.setName("读线程B");
//        t2.start();
        Thread t3 = new Demo08ThreadB(service);
        t3.setName("写线程A");
        t3.start();
        Thread t4 = new Demo08ThreadB(service);
        t4.setName("写线程B");
        t4.start();
    }
}

class Demo08Service{
    private ReentrantReadWriteLock lock = new ReentrantReadWriteLock();

    public void read(){
        try{
            lock.readLock().lock();
            System.out.println(Thread.currentThread().getName() + "获得读锁于" + System.currentTimeMillis());
            Thread.sleep(1000);
            System.out.println(Thread.currentThread().getName() + "解除读锁于" + System.currentTimeMillis());
        }catch (InterruptedException e){
            e.printStackTrace();
        }finally {
            lock.readLock().unlock();
        }
    }

    public void write(){
        try{
            lock.writeLock().lock();
            System.out.println(Thread.currentThread().getName() + "获得写锁于" + System.currentTimeMillis());
            Thread.sleep(1000);
            System.out.println(Thread.currentThread().getName() + "解除写锁于" + System.currentTimeMillis());
        }catch (InterruptedException e){
            e.printStackTrace();
        }finally {
            lock.writeLock().unlock();
        }
    }
}


class Demo08ThreadA extends Thread{
    private Demo08Service service;
    public Demo08ThreadA(Demo08Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.read();
    }
}

class Demo08ThreadB extends Thread{
    private Demo08Service service;
    public Demo08ThreadB(Demo08Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.write();
    }
}
```

读写或写读操作都是互斥的，只要出现 写操作的过程，就是互斥的。读读是共享的。

## 5. 拾遗增补

### 5.1 线程的状态

| 状态         | 描述                                                     |
| ------------ | -------------------------------------------------------- |
| NEW          | 尚未启动的线程处于这种状态                               |
| RUNNABLE     | 在Java虚拟机中正在执行的线程处于此状态                   |
| BLOCKED      | 补阻塞等待对象锁的线程处于此状态                         |
| WAITING      | 正在等待另一个线程执行特定动作的线程处于h此x状态         |
| TIME_WAITING | 正在等待另一个线程执行达到指定的等待时间的线程处于此状态 |
| TERMINATED   | 已经退出运行的线程处于此状态                             |

### 5.2 Callable接口

Runnable是执行工作的独立任务，但是它不返回任何值。如果希望任务在完成的同时能够返回一个值，可以通过实现Callable接口。在JDK5.0中引入的Callable接口是一种具有类型参数的泛型，它的类型参数表示从方法call中返回的值的类型。

```java
package chap5;

import java.util.concurrent.Callable;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.FutureTask;

public class Demo05 {
    public static void main(String[] args) throws ExecutionException, InterruptedException {
        Callable<Integer> callable = new Demo05Callable();
        FutureTask<Integer> task = new FutureTask<>(callable);
        Thread t1 = new Thread(task);
        t1.start();
        System.out.println("线程返回的值是：" + task.get());
    }
}

class Demo05Callable implements Callable<Integer>{
    @Override
    public Integer call() throws Exception {
        System.out.println(Thread.currentThread().getName() + "调用了callable接口的实现类");
        int val = (int)(Math.random() * 10);
        System.out.println("准备返回的值是：" + val);
        return val;
    }
}
```

### 5.3 线程池

线程池的作用：

就是限制在系统中可以执行的线程的数量。根据系统环境的情况，自动或手动设置线程数量，以达到运行的最佳效果。

用线程池控制线程的数量，多余的线程必须排队等候，一个任务执行完毕，再从队列中取最前面的任务开始执行。如果队列中没有等待的线程，线程池的资源就处于等待的状态。当一个新任务需要运行时，如果线程池中有等待的工作线程，就歌词开始运行，否则进入等待队列。

为什么需要使用线程池

1. 减少创建和销毁线程的次数，每个工作线程都可以被重复复用，可执行多个任务。
2. 可以根据系统的随能力，调整线程池中工作线程的数量，防止因为消耗过多的内存，导致服务器崩溃。

Java里面线程池顶级接口是Executor，但是从严格意义上来讲Executor并不是一个线池，只是一个执行线程的工具。

比较重要的几个类：

ExecutorService：真正的线程池接口

ScheduleExecutorService：解决那些需要任务重复执行的问题

ThreadPoolExecutor：ExecutorSrevice的默认实现

ScheduleThreadPoolExecutor：继承ThreadPoolExecutor的ScheduledExecutorService接口的实现，周期性任务调试的类实现。

Executors类里提供了一些静态工厂，生成一些常用的线程池。

| 方法                                                         | 解释                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| newCachedThreadPool(ThreadPoolExecutor)                      | 创建一个可缓存的线程池。如果线程池的大小超过了处理任务所需要的线程，就会回收部分空闲的线程(60秒没有执行的线程)，当任务数量增加时，以可以智能的添加新的线程来处理任务。这个线程池不会对线程池的大小进行限制，线程池的大小完全依赖操作系统(JVM)能够创建的最大线程大小。 |
| newFixedThreadPool(ThreadPoolExecutor)                       | 创建固定大小的线程池。每次提交一个任务就创建一个线程，直到线程池达到最大大小。线程池的大小一旦达到最大值就会保持不变，如果某个线程因为执行而异常结束，那么线程池会补充一个新的线程。 |
| newSginleThreadExecutor(ThreadPoolExecutor)                  | 创建一个单线程的线程池。这个线程池只有一个工作线程，也就是相当于单线程串行执行所有任务。如果 这个唯一的线程因为异常结束，那么会有一个新的线程来替代它。些线程池保证所有任务的执行顺序 会按照 任务的提交顺序 执行。 |
| newScheduledThreadPool(ScheduledThreadPoolExecutor)          | 创建一个大小无限的线程池。些线程池运行定时以及周期性执行任务的需求。 |
| newSingleThreadScheduledExecutor(ScheduledThreadPoolExecutor) | 创建一个单线程用于定时以及周期性执行任务的需求。             |

```java
package chap5;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;

public class Demo06 {
    public static void main(String[] args) {
        Runnable r = new Demo06Runnable();
//        ExecutorService es = Executors.newCachedThreadPool();
//        ExecutorService es = Executors.newFixedThreadPool(2);
        /*ExecutorService es = Executors.newSingleThreadExecutor();
        for (int i = 0; i < 10; i++) {
            // execute就是线程代码
            es.execute(r);
        }*/

//        ScheduledExecutorService es = Executors.newScheduledThreadPool(2);
        ScheduledExecutorService es = Executors.newSingleThreadScheduledExecutor();
        System.out.println("准备执行：" + System.currentTimeMillis());
        // 延迟执行
//        es.schedule(r, 3, TimeUnit.SECONDS);
        // 定时执行
        // scheduleAtFixedRate(需要执行的线程, 第1个需要执行的线程延迟多长时间行, 两个线程间相隔的时间, 使用到时间单位)
        es.scheduleAtFixedRate(r, 2, 5, TimeUnit.SECONDS);

        // 关闭线程池
//        es.shutdown();
    }
}

class Demo06Runnable implements Runnable{
    static long waitTime = 1000;
    @Override
    public void run() {
        try {
            System.out.println(Thread.currentThread().getName() + "开始于" + System.currentTimeMillis());
            synchronized (this) {
                waitTime += 100;
            }
            Thread.sleep(waitTime);
            System.out.println(Thread.currentThread().getName() + "结束于" + System.currentTimeMillis());
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
```
