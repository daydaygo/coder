package chap4;

import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class Demo05 {
    public static void main(String[] args) {
        Demo05Service service = new Demo05Service();
        // 一生产一消费
//        Thread producer = new Demo05ProducerThread(service);
//        producer.start();
//
//        Thread consumer = new Demo05ConsumerThread(service);
//        consumer.start();

        // 多生产多消费
        int size = 2;
        Thread[] producers = new Thread[size];
        Thread[] consumers = new Thread[size];

        for (int i = 0; i < size; i++) {
            char c = (char)('A' + i);
            producers[i] = new Demo05ProducerThread(service);
            producers[i].setName("生产者" + c);
            producers[i].start();

            consumers[i] = new Demo05ConsumerThread(service);
            consumers[i].setName("消费者" + c);
            consumers[i].start();
        }
    }
}

class Demo05Service{
    private Lock lock = new ReentrantLock();
    private Condition condition = lock.newCondition();
    private String val = "";

    public void set(){
        try{
            lock.lock();
            while(!"".equals(val)){
                System.out.println(Thread.currentThread().getName() + "开始等待");
                condition.await();
            }
            val = System.currentTimeMillis() + "-" + System.nanoTime();
            System.out.println(Thread.currentThread().getName() + "生产值：" + val);
            //condition.signal();
            condition.signalAll();
        }catch (InterruptedException e){
            e.printStackTrace();
        } finally {
            lock.unlock();
        }
    }

    public void get(){
        try{
            lock.lock();
            while("".equals(val)){
                System.out.println(Thread.currentThread().getName() + "开始等待");
                condition.await();
            }
            System.out.println(Thread.currentThread().getName() + "消费值：" + val);
            val = "";
//            condition.signal();
            condition.signalAll();
        }catch (InterruptedException e){
            e.printStackTrace();
        } finally {
            lock.unlock();
        }
    }
}

class Demo05ProducerThread extends Thread{
    private Demo05Service service;
    public Demo05ProducerThread(Demo05Service service){
        this.service = service;
    }

    @Override
    public void run() {
        while(true){
            service.set();
        }
    }
}

class Demo05ConsumerThread extends Thread{
    private Demo05Service service;
    public Demo05ConsumerThread(Demo05Service service){
        this.service = service;
    }

    @Override
    public void run() {
        while(true){
            service.get();
        }
    }
}