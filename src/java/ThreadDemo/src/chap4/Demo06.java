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