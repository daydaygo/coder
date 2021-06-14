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
