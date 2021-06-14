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
