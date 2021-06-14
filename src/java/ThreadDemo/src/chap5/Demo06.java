package chap5;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;

public class Demo06 {
    public static void main(String[] args) {
        Runnable r = new Demo06Runnable();
//        ExecutorService es = Executors.newCachedThreadPool();
//        ExecutorService es = Executors.newFixedThreadPool(2);
        /*ExecutorService es = Executors.newSingleThreadExecutor();
        for (int i = 0; i < 10; i++) {
            // execute就是线程代码
            es.execute(r);
        }*/

//        ScheduledExecutorService es = Executors.newScheduledThreadPool(2);
        ScheduledExecutorService es = Executors.newSingleThreadScheduledExecutor();
        System.out.println("准备执行：" + System.currentTimeMillis());
        // 延迟执行
//        es.schedule(r, 3, TimeUnit.SECONDS);
        // 定时执行
        // scheduleAtFixedRate(需要执行的线程, 第1个需要执行的线程延迟多长时间行, 两个线程间相隔的时间, 使用到时间单位)
        es.scheduleAtFixedRate(r, 2, 5, TimeUnit.SECONDS);

        // 关闭线程池
//        es.shutdown();
    }
}

class Demo06Runnable implements Runnable{
    static long waitTime = 1000;
    @Override
    public void run() {
        try {
            System.out.println(Thread.currentThread().getName() + "开始于" + System.currentTimeMillis());
            synchronized (this) {
                waitTime += 100;
            }
            Thread.sleep(waitTime);
            System.out.println(Thread.currentThread().getName() + "结束于" + System.currentTimeMillis());
        }catch (InterruptedException e){
            e.printStackTrace();
        }
    }
}