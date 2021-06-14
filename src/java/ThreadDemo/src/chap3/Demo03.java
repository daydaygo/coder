package chap3;

public class Demo03 {
    public static void main(String[] args) throws InterruptedException {
        Object lock = new Object();
        Thread t1 = new Demo03ThreadA(lock);
        t1.setName("A");
        t1.start();
        Thread t3 = new Demo03ThreadA(lock);
        t3.setName("C");
        t3.start();
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
                System.out.println("线程" + Thread.currentThread().getName() + "开始等待在" + System.currentTimeMillis());
                lock.wait();
                System.out.println("线程" + Thread.currentThread().getName() + "结束等待在" + System.currentTimeMillis());
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
                Thread.sleep(2000);
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}