package chap3;

import java.util.ArrayList;
import java.util.List;

public class Demo08 {
    public static void main(String[] args) throws InterruptedException {
        Demo08Service service = new Demo08Service();
        Thread t1 = new Demo08ThreadB(service);
        t1.start();
        Thread t2 = new Demo08ThreadB(service);
        t2.start();
        Thread.sleep(1000);
        Thread t3 = new Demo08ThreadA(service);
        t3.start();
    }
}

class Demo08Service{
    private List list = new ArrayList();
    private Object lock = new Object();

    public void add(){
        synchronized (lock){
            list.add("a");
            lock.notifyAll();
        }
    }
    public void subtrac(){
        try {
            synchronized (lock) {
                if (list.size() == 0) {
                    System.out.println(Thread.currentThread().getName() + "开始等待数据");
                    lock.wait();
                    System.out.println(Thread.currentThread().getName() + "结束获取数据等待");
                }
                if (list.size() > 0) {
                    list.remove(0);
                }
                System.out.println(Thread.currentThread().getName() + "：list的大小是" + list.size());
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}
class Demo08ThreadA extends Thread{
    private Demo08Service service;
    public Demo08ThreadA(Demo08Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.add();
    }
}
class Demo08ThreadB extends Thread{
    private Demo08Service service;
    public Demo08ThreadB(Demo08Service service){
        this.service = service;
    }

    @Override
    public void run() {
        service.subtrac();
    }
}