package chap3;

public class Demo06 {
    public static void main(String[] args) throws InterruptedException {
        Object lock = new Object();
        Thread t1 = new Demo06ThreadA(lock);
        t1.start();
        Thread t3 = new Demo06ThreadA(lock);
        t3.start();
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
                System.out.println(Thread.currentThread().getName() + "进入同步代码块于" + System.currentTimeMillis());
                lock.wait(3000);
                System.out.println(Thread.currentThread().getName() + "结束同步代码块于" + System.currentTimeMillis());
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