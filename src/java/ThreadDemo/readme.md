# thread demo

- theory理论 重排序与数据依赖性 asIfSerial/happensBefore 性质=可见性+有序性+原子性
  - 内存模型: 原子操作(lock unlock read load use assign store write)
- thread线程
  - API: interrupt中断 suspend/resume暂停恢复 yield wait/notify priority
  - standard标准: safety liveness reusability performance(throughput responsiveness latency capacity efficiency scalability degradation)
  - pool: ThreadPoolExecutor 4种常见 FutureTask
- lock锁
  - synchronized: 执行完/异常->unlock 底层基于lockFree队列
  - aqs AbstractQueuedSynchronizer = 单一资源state + FIFO
  - ReentranLock: 可重入 可响应中断 公平or抢占
  - ReentrantReadWriteLock: 锁降级 Condition等待队列 都多写少场景
  - lockSupport: suspend/resume Unsafe类实现 `unpark(thread)`指定唤醒线程
  - condition
    - wait->await notify->signal
    - 不响应中断 lock可以有多个condition await可设置超时时间 复用aqs的Node类
- coContainer并发容器
  - ConcurrentHashMap ConcurrentLinkedQueue CopyOnWriteArrayList ThreadLocal
  - BlockingQueue ArrayBlockingQueue LinkedBlockingQueue
- coTool并发工具
  - CountDownLatch/CyclicBarrier Semaphore/Exchanger

- chap1 basic
  - 1 `Thread.sleep(1000)`
  - 2 `extends Thread`
    - 3/4 rand随机性: `t.run()` `t.start()`
    - 6/7 成员变量/共享数据/线程安全
  - 5 `implements Runnable` java单继承?
  - 1,8-11 ThreadAPI `currentThread() isAlive() sleep() getId() stopThread() suspend/resume yield`
  - 12-18 `stopThread interrupt`
  - 19-22 `suspend/resume`
  - 23 `yield`
  - 24/25 `priority`
  - 26 `daemon`
- chap2 `synchronized`
  - 1局部变量 2成员变量 3/4方法锁=对象锁 5脏读 6锁重入/继承 7锁自动释放 8override
  - 9慢 10-15codeBlock 16/17非this对象/脏读 18/19/20静态方法=类锁 21String对象锁 22/23死锁 24/25改变内容
  - `volatile`变量多个线程可见 26
- chap3 线程间通信
  - wait&notify/notifyAll 1-3wait停notify继续 5wait+interrupt 6wait(long) 7通知过早 8wait条件变化 9-14producer-consumer
  - join等待线程销毁 15demo 16interrupte 17join(long) 18sleep(long)
  - ThreadLocal类=线程私有数据
- chap4 Lock
  - ReentranLock 1demo 2synchronized 3wait/notify 4Condition唤醒 5producer-consumer 6/7公平?
  - ReentrantReadWriteLock 8
- chap5
  - status: new runnable blocked waiting time_waiting terminated
  - 5 Callable接口
  - 6 线程池 ExecutorService ScheduleExecutorService ThreadPoolExecutor

```java
// Thread类 Runnable接口
Thread.sleep(ms, ns) Thread.interrupt
synchronized // 同步, 互斥
wait() notify() notifyAll() // 线程协调, wait set
```

- pattern
  - singleThreadedExecution(criticalSection region): 对临界区加以保护, 多线程互斥共享
  - immutable: 状态不会变化
  - guardedSuspension(spinLock guardedWait): 线程等待到适合的状态
  - balking: 状态不合适就直接退出
  - producerConsumer: channel
  - readWriteLock: 多读一写
  - threadPerMessage: 为了简化启动线程, 可以使用匿名内部类
  - workerThread(ThreadPool backgroundThread)
  - future: 任务委托给future参与者, 并将future当做返回值, 任务的处理结果事后设置给 future
  - twoPhaseTermination: 添加终止请求的标识, 让线程自己检测并退出
  - ThreadSpecificStorage: `java.lang.ThreadLocal`
  - activeObject(actor): 独立于client/server外, 异步处理耗时任务

![多线程程序设计的模式语言](http://qiniu.dayday.tech/java_multi_thread.png)
