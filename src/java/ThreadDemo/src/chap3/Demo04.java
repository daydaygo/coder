package chap3;

public class Demo04 {
    public static void main(String[] args) throws InterruptedException {
        Object lock = new Object();
        Thread t1 = new Demo04Thread(lock);
        t1.start();
        Thread.sleep(2000);
        t1.interrupt();
    }
}

class Demo04Service{
    public void foo(Object lock){
        try{
            synchronized (lock){
                System.out.println("准备开始等待");
                lock.wait();
                System.out.println("结束等待");
            }
        }catch (InterruptedException e){
            System.out.println("出现异常，因为wait状态的线程被interrupt了");
            e.printStackTrace();
        }
    }
}

class Demo04Thread extends Thread{
    private Object lock;
    public Demo04Thread(Object lock){
        this.lock = lock;
    }

    @Override
    public void run() {
        Demo04Service service = new Demo04Service();
        service.foo(lock);
    }
}