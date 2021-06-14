package chap2;

import java.util.ArrayList;
import java.util.List;

public class Demo17 {
    public static void main(String[] args) throws InterruptedException {
        Demo17List list = new Demo17List();
        Thread t1 = new Demo17ThreadA(list);
        t1.start();
        Thread t2 = new Demo17ThreadB(list);
        t2.start();
        Thread.sleep(5000);
        System.out.println("list size is " + list.size());
    }
}

class Demo17List{
    private List list = new ArrayList();
    synchronized public void add(Object obj){
        list.add(obj);
    }
    synchronized public int size(){
        return list.size();
    }
}

class Demo17Service{
    private Object lockObject = new Object();
    public void add(Demo17List list, Object obj){
        try {
            synchronized (list) {
                if (list.size() < 1) {
                    Thread.sleep(2000); // 模拟数据的获取
                    list.add(obj);
                }
            }
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}

class Demo17ThreadA extends Thread{
    private Demo17List list;
    public Demo17ThreadA(Demo17List list){
        this.list = list;
    }

    @Override
    public void run() {
        Demo17Service service = new Demo17Service();
        service.add(list, "a");
    }
}

class Demo17ThreadB extends Thread{
    private Demo17List list;
    public Demo17ThreadB(Demo17List list){
        this.list = list;
    }

    @Override
    public void run() {
        Demo17Service service = new Demo17Service();
        service.add(list, "b");
    }
}