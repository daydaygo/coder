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