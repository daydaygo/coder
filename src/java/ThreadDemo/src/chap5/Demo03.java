package chap5;

public class Demo03 {
    public static void main(String[] args) throws InterruptedException {
        Object lock = new Object();
        Thread t = new Demo03Thread(lock);
        t.start();
        Thread.sleep(1000);
        System.out.println("线程在main方法中的状态是：" + t.getState());
    }
}

class Demo03Thread extends Thread{
    private Object lock;

    public Demo03Thread(Object lock){
        this.lock = lock;
    }

    @Override
    public void run() {
        try {
            synchronized (lock){
                lock.wait();
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
