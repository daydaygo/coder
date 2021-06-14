package chap1;

public class Demo23 {
    public static void main(String[] args) {
        Thread t1 = new Demo23Thread();
        t1.start();
    }
}

class Demo23Thread extends Thread{
    @Override
    public void run() {
        int count = 0;
        long start = System.currentTimeMillis();
        for (int i = 0; i < 500000; i++) {
            Thread.yield();
            count  += i + 1;
        }
        long end = System.currentTimeMillis();
        System.out.println("用时：" + (end - start) + "毫秒");
    }
}
