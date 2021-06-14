package chap3;

import java.util.ArrayList;
import java.util.List;

public class Demo13 {
    public static void main(String[] args) {
        Demo13VO vo = new Demo13VO();
        Thread[] producers = new Thread[5];
        for (int i = 0; i < producers.length; i++) {
            producers[i] = new Demo13Producer(vo);
            producers[i].setName("生产者" + (char)('A' + i));
            producers[i].start();
        }


        Thread consumer = new Demo13Consumer(vo);
        consumer.setName("消费者");
        consumer.start();
    }
}

class Demo13VO{
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

class Demo13Producer extends Thread{
    private Demo13VO vo;
    public Demo13Producer(Demo13VO vo){
        this.vo = vo;
    }

    @Override
    public void run() {
        while(true) {
            vo.push(Math.random() + "");
        }
    }
}

class Demo13Consumer extends Thread{
    private Demo13VO vo;
    public Demo13Consumer(Demo13VO vo){
        this.vo = vo;
    }

    @Override
    public void run() {
        while(true){
            vo.pop();
        }
    }
}