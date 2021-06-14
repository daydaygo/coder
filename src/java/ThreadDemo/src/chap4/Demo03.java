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
