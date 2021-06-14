package chap1;

import java.util.Random;

public class Demo25 {
    public static void main(String[] args) {
        for (int i = 0; i < 5; i++) {
            Thread t1 = new Demo25Thread();
            t1.setName("A" + i);
            t1.setPriority(Thread.MAX_PRIORITY);
            t1.start();

            Thread t2 = new Demo25Thread();
            t2.setName("B" + i);
            t2.setPriority(Thread.MIN_PRIORITY);
            t2.start();
        }
    }
}

class Demo25Thread extends Thread{
    @Override
    public void run() {
        long start = System.currentTimeMillis();
        long count = 0;
        for (int i = 0; i < 10; i++) {
            for (int j = 0; j < 50000; j++) {
                Random r = new Random();
                count += i * j + r.nextInt();;
            }
        }
        long end = System.currentTimeMillis();
        System.out.println("线程" + this.getName() + "执行完成使用了" + (end  - start) + "毫秒");
    }
}
