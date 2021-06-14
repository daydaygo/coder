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