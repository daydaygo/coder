package chap3;

import java.util.ArrayList;
import java.util.List;

public class Demo14 {
    public static void main(String[] args) {
        Demo14VO vo = new Demo14VO();

        int size = 5;
        Thread[] producers = new Thread[size];
        Thread[] consumers = new Thread[size];

        for (int i = 0; i < size; i++) {
            char c = (char)('A' + i);
            producers[i] = new Demo14Producer(vo);
            producers[i].setName("生产者" + c);
            producers[i].start();

            consumers[i] = new Demo14Consumer(vo);
            consumers[i].setName("消费者" + c);
            consumers[i].start();
        }
    }
}

class Demo14VO{
    private List<String> list = new ArrayList<>();

    synchronized public void push(String val){
        try {
            while (list.size() == 1) {
                System.out.println(Thread.currentThread().getName() + "等待中");
                this.wait();
            }
            list.add(val);
            System.out.println(Thread.currentThread().getName() + "：添加数据" + val);
            System.out.println(Thread.currentThread().getName() + "：还有" +  list.size() + "个数据");
            this.notifyAll();
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }

    synchronized public String pop(){
        String returnValue = null;
        try {
            while (list.size() == 0) {
                System.out.println(Thread.currentThread().getName() + "等待中");
                this.wait();
            }
            returnValue = list.get(0);
            list.remove(0);
            System.out.println(Thread.currentThread().getName() + "：消费数据" + returnValue);
            System.out.println(Thread.currentThread().getName() + "：还有" + list.size() + "个数据");
            this.notifyAll();
        }catch (InterruptedException e){
            e.printStackTrace();
        }
        return returnValue;
    }
}

class Demo14Producer extends Thread{
    private Demo14VO vo;
    public Demo14Producer(Demo14VO vo){
        this.vo = vo;
    }

    @Override
    public void run() {
        while(true) {
            vo.push(Math.random() + "");
        }
    }
}

class Demo14Consumer extends Thread{
    private Demo14VO vo;
    public Demo14Consumer(Demo14VO vo){
        this.vo = vo;
    }

    @Override
    public void run() {
        while(true){
            vo.pop();
        }
    }
}