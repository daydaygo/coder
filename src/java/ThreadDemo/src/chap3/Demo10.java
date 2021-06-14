package chap3;

public class Demo10 {
    public static void main(String[] args) {
        Object lock = new Object();
        int size = 2;
        Thread[] producers = new Thread[size];
        Thread[] consumers = new Thread[size];
        for (int i = 0; i < size; i++) {
            char c = (char)('A' + i);
            producers[i] = new Demo10Producer(lock);
            producers[i].setName("生产者" + c);

            consumers[i] = new Demo10Consumer(lock);
            consumers[i].setName("消费者" + c);

            producers[i].start();
            consumers[i].start();
        }
    }
}

class Demo10VO{
    public static String value = "";
}

// 生产者
class Demo10Producer extends Thread {
    private Object lock;
    public Demo10Producer(Object lock){
        this.lock = lock;
    }

    @Override
    public void run() {
        try{
            while(true) {
                synchronized (lock) {
                    if (!"".equals(Demo10VO.value)){
                        System.out.println(Thread.currentThread().getName() + "等待中……");
                        lock.wait();
                    }
                    System.out.println(Thread.currentThread().getName() + "生产中……");
                    String value = System.currentTimeMillis() + "_" + System.nanoTime();
                    Demo10VO.value = value;
                    //lock.notify();
                    lock.notifyAll();
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo10Consumer extends Thread{
    private Object lock;
    public Demo10Consumer(Object lock){
        this.lock = lock;
    }

    @Override
    public void run() {
        try{
            while(true){
                synchronized (lock){
                    if ("".equals(Demo10VO.value)){
                        System.out.println(Thread.currentThread().getName() + "等待中……");
                        lock.wait();
                    }
                    System.out.println(Thread.currentThread().getName() + "消费中……");
                    Demo10VO.value = "";
//                    lock.notify();
                    lock.notifyAll();
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}