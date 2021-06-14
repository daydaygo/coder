package chap3;

public class Demo09 {
    public static void main(String[] args) {
        Object lock = new Object();
        Thread t1 = new Demo09Producer(lock);
        t1.start();
        Thread t2 = new Demo09Consumer(lock);
        t2.start();
    }
}

class Demo09VO{
    public static String value = "";
}

// 生产者
class Demo09Producer extends Thread {
    private Object lock;
    public Demo09Producer(Object lock){
        this.lock = lock;
    }

    @Override
    public void run() {
        try{
            while(true) {
                synchronized (lock) {
                    if (!"".equals(Demo09VO.value)){
                        lock.wait();
                    }
                    String value = System.currentTimeMillis() + "_" + System.nanoTime();
                    System.out.println("Set的值是：" + value);
                    Demo09VO.value = value;
                    lock.notify();
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo09Consumer extends Thread{
    private Object lock;
    public Demo09Consumer(Object lock){
        this.lock = lock;
    }

    @Override
    public void run() {
        try{
            while(true){
                synchronized (lock){
                    if ("".equals(Demo09VO.value)){
                        lock.wait();
                    }
                    System.out.println("Get的值是：" + Demo09VO.value);
                    Demo09VO.value = "";
                    lock.notify();
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}