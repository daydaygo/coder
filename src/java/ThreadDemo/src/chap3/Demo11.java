package chap3;

import java.util.ArrayList;
import java.util.List;

public class Demo11 {
    public static void main(String[] args) {
        Demo11VO vo = new Demo11VO();
        Thread producer = new Demo11Producer(vo);
        producer.setName("生产者");
        producer.start();
        Thread consumer = new Demo11Consumer(vo);
        consumer.setName("消费者");
        consumer.start();
    }
}

class Demo11VO{
    private List<String> list = new ArrayList<>();

    synchronized public void push(String val){
        try {
            if (list.size() == 1) {
                System.out.println(Thread.currentThread().getName() + "等待中");
                this.wait();
            }
            list.add(val);
            System.out.println(Thread.currentThread().getName() + "：添加数据" + val);
            System.out.println(Thread.currentThread().getName() + "：还有" +  list.size() + "个数据");
            this.notify();
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }

    synchronized public String pop(){
        String returnValue = null;
        try {
            if (list.size() == 0) {
                System.out.println(Thread.currentThread().getName() + "等待中");
                this.wait();
            }
            returnValue = list.get(0);
            list.remove(0);
            System.out.println(Thread.currentThread().getName() + "：消费数据" + returnValue);
            System.out.println(Thread.currentThread().getName() + "：还有" + list.size() + "个数据");
            this.notify();
        }catch (InterruptedException e){
            e.printStackTrace();
        }
        return returnValue;
    }
}

class Demo11Producer extends Thread{
    private Demo11VO vo;
    public Demo11Producer(Demo11VO vo){
        this.vo = vo;
    }

    @Override
    public void run() {
        while(true) {
            vo.push(Math.random() + "");
        }
    }
}

class Demo11Consumer extends Thread{
    private Demo11VO vo;
    public Demo11Consumer(Demo11VO vo){
        this.vo = vo;
    }

    @Override
    public void run() {
        while(true){
            vo.pop();
        }
    }
}