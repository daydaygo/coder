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